package main

import(
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/* 関数のテスト */
func main(){
	var days []int
	var titles []string

	days, titles = getAllElement()

	log.Println(days)
	log.Println(titles)
}

/* MySQLの操作*/
func addElement(day int, title string, main string){
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/web")
	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	defer db.Close()

	// SQLの準備
	ins, err := db.Prepare("INSERT INTO diary VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	// SQLの実行
	res, err := ins.Exec(day, title, main)
	if err != nil {
		log.Fatal(err)
	}

	// 結果の取得
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(lastInsertID)
}

func deleteElement(day int, title string){
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/web")
	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	defer db.Close()

	// SQLの準備
	del, err := db.Prepare("DELETE FROM diary WHERE day=? AND title=?")
	if err != nil {
		log.Fatal(err)
	}
	defer del.Close()

	// SQLの実行
	res, err := del.Exec(day, title)
	if err != nil {
		log.Fatal(err)
	}

	// 結果の取得(影響を受ける行の数を取得)
	affected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(affected)
}

func getAllElement()(days []int, titles []string){
	type getInform struct {
		day int
		title string
		main string
	}
	
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/web")
	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	defer db.Close()

	// SQLの実行
	rows, err := db.Query("SELECT * FROM diary order by day")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// SQLの実行
	for rows.Next() {
		var inform getInform
		err := rows.Scan(&inform.day, &inform.title, &inform.main)

		if err != nil {
			panic(err.Error())
		}
		
		days = append(days, inform.day)
		titles = append(titles, inform.title)
	}

	return days, titles
}

/* 
参考
- [【Golang】MySQLの基本的な操作を行う](https://www.chuken-engineer.com/entry/2021/09/24/162120)
- [Go言語を真剣に勉強してみた〜データベース接続(MySQL)編〜](https://qiita.com/watataku8911/items/8c94f8f380f91aa3f3e5)
*/
