// Package build holds version metadata, populated at link time via -ldflags.
package build

// Version is the semantic version. Overridden at build time.
var Version = "0.0.0-dev"

// Commit is the VCS hash. Overridden at build time.
var Commit = "none"
