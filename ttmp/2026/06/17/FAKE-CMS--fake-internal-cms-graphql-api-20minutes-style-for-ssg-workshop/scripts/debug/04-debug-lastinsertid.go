// 04-debug-lastinsertid.go — Probe: does modernc.org/sqlite return a usable
// LastInsertId() for INTEGER PRIMARY KEY columns?
//
// Run from the repo root:
//   go run ttmp/2026/06/17/FAKE-CMS--*/scripts/debug/04-debug-lastinsertid.go
//
// Why this exists:
//   The deterministic seed relies on LastInsertId() to capture author/media/
//   content_node row IDs and feed them as foreign keys into dependent rows.
//   We needed to confirm modernc populates it (it does: 1,2,3,...) and that
//   the IDs survive into subsequent inserts.
package main

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

func main() {
	db, _ := sql.Open("sqlite", ":memory:")
	_, _ = db.Exec("CREATE TABLE t(id INTEGER PRIMARY KEY, v TEXT)")
	for i, v := range []string{"a", "b", "c"} {
		res, err := db.Exec("INSERT INTO t(v) VALUES(?)", v)
		id, e := res.LastInsertId()
		fmt.Printf("i=%d v=%s err=%v lastid=%d liderr=%v\n", i, v, err, id, e)
	}
	rows, _ := db.Query("SELECT id,v FROM t")
	for rows.Next() {
		var id int
		var v string
		rows.Scan(&id, &v)
		fmt.Println("row", id, v)
	}
	rows.Close()
}
