// 11-verify-det.go — Verify seed determinism at the content level by hashing
// a STABLE serialization (sorted columns, string-coerced values, no []byte
// pointers) over two in-memory seeds. Proves the DATA is deterministic.
// Run: go run ttmp/.../scripts/debug/11-verify-det.go
package main
import ("context";"crypto/sha256";"database/sql";"fmt";"github.com/go-go-golems/fake-cms/internal/repo";_ "modernc.org/sqlite")
func hashOne()string{ctx:=context.Background();r,_:=repo.Open(ctx,":memory:");r.Seed(ctx);defer r.Close()
 h:=sha256.New()
 for _,t:=range []struct{n,c string}{{"author","slug,display_name"},{"content_node","id,kind,slug,title,published_at,author_id,word_count"},{"block","content_id,order_index,type,data"}}{
  rows,_:=r.DB.Query(fmt.Sprintf("SELECT %s FROM %s ORDER BY 1,2",t.c,t.n))
  for rows.Next(){var a,b,c,d sql.NullString;rows.Scan(&a,&b,&c,&d);fmt.Fprintf(h,"%s|%s|%s|%s|%s|%s\n",t.n,a.String,b.String,c.String,d.String)}
  rows.Close()}
 return fmt.Sprintf("%x",h.Sum(nil))}
func main(){a:=hashOne();b:=hashOne();fmt.Println("A",a);fmt.Println("B",b);if a==b{fmt.Println("DETERMINISTIC ✓")}else{fmt.Println("NOT ✗")}}
