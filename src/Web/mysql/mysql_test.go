package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func insertDB(db *sql.DB) {
	stmt, _ := db.Prepare("INSERT user_info SET username=?,create_time=?")
	res, _ := stmt.Exec("皮皮落", "2018-12-09")
	id, _ := res.LastInsertId()
	fmt.Println("id=", id)
}

func updateDB(db *sql.DB) {
	stmt, _ := db.Prepare("UPDATE user_info SET username=? WHERE id=?")
	res, _ := stmt.Exec("ppzl", "2")
	cnt, _ := res.RowsAffected()
	fmt.Println("affected=", cnt)
}

func selectDB(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM user_info")
	for rows.Next() {
		var id int
		var username string
		var create_time string
		rows.Scan(&id, &username, &create_time)
		fmt.Printf("id=%d, username=%s, create_time=%s\n", id, username, create_time)
	}
}

func TestSql(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang?charset=utf8")
	if err != nil {
		panic(err)
	}
	//insertDB(db)
	//updateDB(db)
	selectDB(db)
}
