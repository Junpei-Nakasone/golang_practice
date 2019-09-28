package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// サーバーサイドのDBと仮定
var DB = map[string]string{
	"TEST1Key": "TEST1Secret",
}

// サーバーサイドでクライアントから受け取った情報が正しいものか判定する関数
// 引数はクライアントから送られてきたデータというイメージ
func Server(apiKey, sign string, data []byte) {
	// TEST1のSecretはDBから参照する
	apiSecret := DB[apiKey]
	// hmacのNew関数でapiSecretmを暗号化
	h := hmac.New(sha256.New, []byte(apiSecret))
	// 上記のhにクライアントから送られてきたdataを追加(dataはServer関数の第３引数)
	h.Write(data)
	// hをEncodeしてexpectedHMACに格納
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	// クライアントから送られてきたsign(Server関数の第二引数)とexpextedHMACが同じかチェックする、この場合trueになる
	fmt.Println(sign == expectedHMAC)
}

func main() {
	const apiKey = "TEST1Key"
	const apiSecret = "TEST1Secret"

	// byteの何らかのデータを送ると仮定
	data := []byte("data")
	h := hmac.New(sha256.New, []byte(apiSecret))

	h.Write(data)
	sign := hex.EncodeToString(h.Sum(nil))

	fmt.Println(sign)

	Server(apiKey, sign, data)
}
