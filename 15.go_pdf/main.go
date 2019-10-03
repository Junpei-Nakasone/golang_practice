package main

import (
	"fmt"
	"log"

	// ライブラリをimport
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

// wkhtmltopdfがあるpathを定義
const path = "/usr/local/bin/wkhtmltopdf"

func main() {
	// wkhtmltopdfライブラリのNewPDFGeneratorを実行し、pdfgに格納
	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	// エラー処理
	if err != nil {
		log.Fatal(err)
	}

	// 上記で定義したpathを指定する
	wkhtmltopdf.SetPath(path)

	// PDF化したいwebページのURLを変数pに格納
	p := "https://google.com/"
	// ローカルのhtmlファイルをPDFにしたい場合は下記のように指定する
	// p := "./test.html"

	// NewPage関数を上記のpに対して行い、変数pageに格納
	page := wkhtmltopdf.NewPage(p)

	//  AddPage関数で上記pageを追加する
	pdfg.AddPage(page)

	// PDF作成、Create関数なしで実行してもエラーは発生せずPDFは作成されるが「表示する中身がありません」になる
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// 出力、WriteFile内でioutil.Writefileが実行されている
	err = pdfg.WriteFile("./google.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tada!")
}
