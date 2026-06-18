import assert from "node:assert/strict";
import { spawn } from "node:child_process";
import { readdir, readFile, rm } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const frontendRoot = path.resolve(__dirname, "..");
const repoRoot = path.resolve(frontendRoot, "..");
const endpoint = "http://localhost:18081/graphql";
const siteUrl = "http://example.test";

function run(command, args, options = {}) {
  return new Promise((resolve, reject) => {
    const child = spawn(command, args, {
      stdio: options.stdio || "pipe",
      cwd: options.cwd || process.cwd(),
      env: { ...process.env, ...(options.env || {}) },
    });
    let stdout = "";
    let stderr = "";
    child.stdout?.on("data", chunk => { stdout += chunk; });
    child.stderr?.on("data", chunk => { stderr += chunk; });
    child.on("error", reject);
    child.on("exit", code => {
      if (code === 0) resolve({ stdout, stderr });
      else reject(new Error(`${command} ${args.join(" ")} exited ${code}\n${stdout}\n${stderr}`));
    });
  });
}

async function waitForHealthz(url, timeoutMs = 20000) {
  const start = Date.now();
  let lastError;
  while (Date.now() - start < timeoutMs) {
    try {
      const response = await fetch(url);
      if (response.ok && (await response.text()) === "ok") return;
    } catch (error) {
      lastError = error;
    }
    await new Promise(resolve => setTimeout(resolve, 250));
  }
  throw new Error(`fake-cms health check timed out: ${lastError?.message || "no response"}`);
}

async function gql(query, variables = {}) {
  const response = await fetch(endpoint, {
    method: "POST",
    headers: { "content-type": "application/json" },
    body: JSON.stringify({ query, variables }),
  });
  const payload = await response.json();
  if (payload.errors?.length) throw new Error(JSON.stringify(payload.errors, null, 2));
  return payload.data;
}

async function walkFiles(dir) {
  const entries = await readdir(dir, { withFileTypes: true });
  const files = [];
  for (const entry of entries) {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) files.push(...await walkFiles(full));
    else files.push(full);
  }
  return files;
}

function rel(file) {
  return "/" + path.relative(path.join(frontendRoot, "_site"), file).split(path.sep).join("/");
}

function extractJsonLd(html) {
  const match = html.match(/<script type="application\/ld\+json">([\s\S]*?)<\/script>/);
  assert(match, "expected JSON-LD script tag");
  return JSON.parse(match[1]);
}

const server = spawn("go", ["run", "./cmd/fake-cms", "serve", "--path", "testdata/cms.db", "--addr", ":18081"], {
  cwd: repoRoot,
  stdio: ["ignore", "pipe", "pipe"],
  detached: true,
});
server.stdout.on("data", chunk => process.stdout.write(`[fake-cms] ${chunk}`));
server.stderr.on("data", chunk => process.stderr.write(`[fake-cms] ${chunk}`));

async function stopServer() {
  if (!server.pid || server.killed) return;
  try {
    // Kill the whole process group so the compiled child spawned by `go run`
    // does not survive after the Node script exits.
    process.kill(-server.pid, "SIGTERM");
  } catch {
    try { server.kill("SIGTERM"); } catch {}
  }
  await new Promise(resolve => setTimeout(resolve, 500));
}

try {
  await waitForHealthz("http://localhost:18081/healthz");
  const countData = await gql(`{ articles(first: 1) { totalCount edges { node { urlSlug: slug postType seo { jsonLd } } } } }`);
  const totalCount = countData.articles.totalCount;
  assert(totalCount > 0, "seeded CMS should have articles");

  await rm(path.join(frontendRoot, "_site"), { recursive: true, force: true });
  await run("npm", ["run", "build", "--", "--quiet"], {
    cwd: frontendRoot,
    env: { CMS_ENDPOINT: endpoint, SITE_URL: siteUrl },
  });

  const files = await walkFiles(path.join(frontendRoot, "_site"));
  const htmlFiles = files.filter(file => file.endsWith("/index.html"));
  const articleFiles = htmlFiles.filter(file => {
    const parts = rel(file).split("/").filter(Boolean);
    return parts.length === 3 && parts[2] === "index.html" && !["archives", "rubrique", "author"].includes(parts[0]);
  });

  assert.equal(articleFiles.length, totalCount, "one article HTML file per CMS article");
  assert(!files.some(file => rel(file).startsWith("/tag/")), "must not generate /tag/ URLs");
  assert(files.some(file => rel(file).startsWith("/rubrique/")), "must generate /rubrique/ tag pages");

  const sitemap = await readFile(path.join(frontendRoot, "_site", "sitemap.xml"), "utf8");
  assert(sitemap.includes(`${siteUrl}/rubrique/`), "sitemap should include rubrique URLs");
  assert(sitemap.includes(`${siteUrl}/author/`), "sitemap should include author URLs");
  for (const file of htmlFiles) {
    const urlPath = rel(file).replace(/index\.html$/, "");
    if (urlPath === "/sitemap.xml") continue;
    assert(sitemap.includes(`${siteUrl}${urlPath}`), `sitemap missing generated page ${urlPath}`);
  }

  const sampleHtml = await readFile(articleFiles[0], "utf8");
  const jsonLd = extractJsonLd(sampleHtml);
  assert.equal(typeof jsonLd, "object", "JSON-LD should parse to an object");

  console.log(`integration ok: ${articleFiles.length} article pages, ${files.length} files`);
} finally {
  await stopServer();
}
