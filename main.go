package main

import (
	"fmt"
	"lesson/pkg/statement"

	// "lesson/pkg/definition"
	"time"
)

// 初期化処理
func init() {
	fmt.Println("main--init")
}

// メイン関数
func main() {

	fmt.Println("**********main**********")
	fmt.Println("実行時間：", time.Now())
	statement.Statement()
	// definition.Definition()
}

func template() {
	fmt.Println("------template")
	test := "test"
	fmt.Println(test)
}
