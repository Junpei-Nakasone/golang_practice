package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./test.sql")
	defer DbConnection.Close()

	tableName := "person"
	cmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
				name STRING,
				age INT)`, tableName)
	_, err := DbConnection.Exec(cmd)

	if err != nil {
		log.Fatalln(err)
	}

	// DBテーブルpersonのname, ageカラムに値をINSERTする、
	// 値は？にしておいて後から渡せる(SQLインジェクション防止？)
	// cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	// _, err = DbConnection.Exec(cmd, "josh", 22)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // UPDATE 更新したいテーブル名 SET 更新したい値と列 WHERE 更新したい列を見つける条件
	// cmd = "UPDATE person SET age = ? WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, 25, "mike")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	cmd = "SELECT * FROM person"

	// Query()はrow(列)を返すQueryを実行する、主にSELECT文で使われる
	rows, _ := DbConnection.Query(cmd)
	fmt.Println("this is rows", rows)
	defer rows.Close()

	// 変数ppに構造体Person用のスライスを格納
	var pp []Person

	// Nextはscanメソッドで読み取ったrowの次の結果を用意する
	for rows.Next() {
		// Personを格納する変数pを宣言だけする
		var p Person

		fmt.Println("p before for statement", p)
		// &で構造体の値を指定して、Scanメソッドが構造体の値にデータを入れる
		err := rows.Scan(&p.Name, &p.Age)
		fmt.Println("this is p from for statement", p)

		if err != nil {
			log.Println(err)
		}

		// append関数でスライスppにpを入れる
		// for文が回るたびにスライスにデータが入っていく
		pp = append(pp, p)
		fmt.Println("this is pp", pp)
	}

	// 構造体ppのrangeでfor文を回し、値をaに入れる
	// ppでfor文を回すのでaも構造体になる
	for _, a := range pp {

		fmt.Println(a.Name, a.Age)

	}
}
