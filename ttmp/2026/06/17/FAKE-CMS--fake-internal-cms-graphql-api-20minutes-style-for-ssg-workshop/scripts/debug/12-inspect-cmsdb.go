// 12-inspect-cmsdb.go — Open the committed testdata/cms.db and print counts
// + a sample article to confirm the artifact is valid and loadable.
// Run: go run ttmp/.../scripts/debug/12-inspect-cmsdb.go
package main
import ("context";"fmt";"github.com/go-go-golems/fake-cms/internal/repo";_ "modernc.org/sqlite")
func main(){ctx:=context.Background();r,err:=repo.Open(ctx,"testdata/cms.db");if err!=nil{fmt.Println("ERR",err);return};defer r.Close()
 for _,t:=range []string{"content_node","block","author","media","term","seo","menu_item"}{var n int;r.DB.QueryRow("SELECT count(*) FROM "+t).Scan(&n);fmt.Printf("%-14s %d\n",t,n)}
 var c int;r.DB.QueryRow("SELECT count(*) FROM content_node WHERE kind='ARTICLE'").Scan(&c);fmt.Println("articles:",c)
 var c2 int;r.DB.QueryRow("SELECT count(*) FROM content_node WHERE kind='PAGE'").Scan(&c2);fmt.Println("pages:",c2)}
