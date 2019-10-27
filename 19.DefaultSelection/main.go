package main

import (
	"fmt"
	"time"
)

func main() {
	// channelを返す
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {

		// tick channelを受信したcase, 値は何も受け取らない
		case <-tick:
			fmt.Println("tick.")
		// tick channelを変数に格納することもできる,
		// 下記のように書くと文字列だけじゃなくchannelが持つtimeのデータも表示する
		// case t := <-tick:
		// fmt.Println("tick.", t)

		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("  .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
