package main

import (
	"fmt"
	"time"
)

type FuncType func() string

var 好康的 FuncType

func main() {
	go func() {
		好康的 = 打电动
		fmt.Println(好康的())
	}()

	好康的 = 欢迎来我家玩
	fmt.Println(好康的())
}
func 欢迎来我家玩() string {
	// 花费 5s 前往杰哥家
	time.Sleep(5 * time.Second)
	return "登dua郎"
}

func 打电动() string {
	return "输了啦，都是你害的"
}
