package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	// クライアントが持っているapiKeyと仮定
	const apiKey = "Test1Key"
	const apiSecret = "Test1Secret"

	// hmac.New関数で第一引数にアルゴリズム、第二引数に暗号化する対象を指定しbyteにしてhashを作る
	h := hmac.New(sha256.New, []byte(apiKey))
	// byteの配列でhashを表示する
	fmt.Println(h)

	// go標準のWrite関数でサーバーに送りたいデータを追加する
	h.Write([]byte("TESTdata"))

	// hexパッケージのEncodeToString関数でエンコーディング
	// Sumで終わりのnilを足す？
	sign := hex.EncodeToString(h.Sum(nil))

	// 生成されたhashを表示する
	fmt.Println(sign)

}
