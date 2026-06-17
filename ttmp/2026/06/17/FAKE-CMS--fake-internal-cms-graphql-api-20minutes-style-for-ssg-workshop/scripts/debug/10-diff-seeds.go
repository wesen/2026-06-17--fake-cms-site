// 10-diff-seeds.go — Dump the first article row + first block rows of two
// independent seeds to find WHERE determinism breaks (values vs ids).
// Run from repo root: go run ttmp/.../scripts/debug/10-diff-seeds.go
package main
import ("context";"database/sql";"fmt";"os";"github.com/go-go-golems/fake-cms/internal/repo";_ "modernc.org/sqlite")
func dump(label, path string){ctx:=context.Background();r,err:=repo.Open(ctx,path);if err!=nil{fmt.Println(label,"open",err);return};defer r.Close()
 if err:=r.Seed(ctx);err!=nil{fmt.Println(label,"seed",err);return}
 fmt.Println("===",label,"===")
 var slug,pub,mod string;var wc sql.NullInt64
 r.DB.QueryRow("SELECT slug,published_at,modified_at,word_count FROM content_node WHERE kind='ARTICLE' ORDER BY id LIMIT 1").Scan(&slug,&pub,&mod,&wc)
 fmt.Printf("article1: slug=%s pub=%s mod=%s wc=%d\n",slug,pub,mod,wc.Int64)
 rows,_:=r.DB.Query("SELECT id,type,data FROM block WHERE content_id=(SELECT id FROM content_node WHERE kind='ARTICLE' ORDER BY id LIMIT 1) ORDER BY order_index LIMIT 4")
 for rows.Next(){var id int;var ty,da string;rows.Scan(&id,&ty,&da);fmt.Printf("  block id=%d type=%s data=%s\n",id,ty,da)};rows.Close()}
func main(){os.Remove("/tmp/a.db");os.Remove("/tmp/b.db");dump("A","/tmp/a.db");dump("B","/tmp/b.db")}
