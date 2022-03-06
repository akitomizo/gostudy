package main

import (
	"fmt"
	"strconv"
)

// 関数外で実行できる変数の宣言
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "str"
	t, f bool    = true, false
)

const pi = 3.14

// 初期化処理
func init() {
	fmt.Println("init")
	fmt.Println(i, f64, s, t, f)
	fmt.Println(pi)
}

// メイン関数
func main() {
	fmt.Println("main!")
	declareVariables()
	dispType()
	binaryDigits()
	testString()
	castStringToInt()

}

//関数内でしか実行できない変数の宣言
func declareVariables() {
	fmt.Println("------declareVariables")
	xi := 1
	xf64 := 1.2
	xs := "str"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	// fmt.Println("testfunc", time.Now())
}

//関数内でしか実行できない変数の宣言
func dispType() {
	fmt.Println("------dispType")
	fmt.Printf("%T", f64)
}

//二進数の操作
func binaryDigits() {
	fmt.Println("------binaryDigits")
	fmt.Println(1 << 1)
	fmt.Println(1 << 2)
	fmt.Println(1 << 3)
}

//関数内でしか実行できない変数の宣言
func testString() {
	fmt.Println("------testString")
	fmt.Println("Hello World")
	fmt.Println("Hello World"[0])
	fmt.Println(string("Hello World"[0]))
}

func castStringToInt() {
	test := "14"
	i, _ := strconv.Atoi(test) // i, err だけど使わない時は_にしておく
	fmt.Printf("%T%v", i, i)
}
