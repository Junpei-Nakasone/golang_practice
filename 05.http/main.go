package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// http.パッケージのGet関数で指定したURLの情報を取得し、respに格納、Getは*Responseとerrorを返す
	resp, err := http.Get("http://example.com")
	if err != nil {
		return
	}
	// deferでrespのBody.Close()を指定しておく
	defer resp.Body.Close()

	// respのBodyをioutil.ReadAllしてbodyに格納
	body, _ := ioutil.ReadAll(resp.Body)
	// Printlnでbodyを表示, ReadAllは[]byteで返すのでstringを指定する
	fmt.Println(string(body))

	// urlパッケージのParse関数を使ってURLが正しいか判定し、baseに格納
	base, err := url.Parse("http://example.com")
	// baseを表示し、エラーがあればエラーも表示する
	fmt.Println(base, err)

	// 上のbaseの後に加えたいURLをurl.Parseで指定し、referenceに格納
	reference, _ := url.Parse("/test?name=user1")
	// baseにResolveReference関数を使ってreferenceを追加する, stringはなくてもTERMINALでは出力される、WEB上では変な挙動になる？
	endpoint := base.ResolveReference(reference).String()
	// baseの後ろにreferenceが追加されて表示される
	fmt.Println(endpoint)

	////////////Advanced//////////
	// GETの場合のHTTPリクエストの作り方, http.NewRequestの第３引数はhttpリクエストのbody
	req, err := http.NewRequest("GET", endpoint, nil)
	// POSTの場合のHTTPリクエストの作り方, URLで見えないようにbodyに入れて送る
	// req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("PASSWORD")))

	// ヘッダー情報を書き込める
	req.Header.Add("HeaderSample", `W`)

	// URLのクエリーをqに格納
	q := req.URL.Query()

	// qを表示すると上記のreferenceで指定したqueryがmapで表示される
	fmt.Println(q)

	// qにクエリーを追加できる
	q.Add("age", "&31")

	// age &31が追加されて表示される
	fmt.Println(q)

	// web上で使うにはEncodeが必要？
	req.URL.RawQuery = q.Encode()
	// qの中の＆がEncodeされてひょうじされる(&は区切り文字なので、データに&がある場合は%26にEncodeされる
	fmt.Println(q.Encode())

	// 変数clientにhttp.Clientのポインタを指定
	var client *http.Client = &http.Client{}

	// ClientのDo関数はhttpリクエストを送りhttpレスポンスを受け取るので、上記のreqを引数に入れて、respに格納
	resp, _ = client.Do(req)

	// 上記のreqなしと同様にbodyに格納して表示
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
