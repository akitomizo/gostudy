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

	// varieable
	fmt.Println("main!")
	declareVariables()
	dispType()
	binaryDigits()
	testString()
	castStringToInt()

	// array slice map
	arrayAndSlice()
	sliceMake()
	mapTest()
	byteTest()

	// function
	add1, add2 := addTest(1, 2, "addStr")
	fmt.Println("add=", add1, add2)
	fmt.Println("resultTest=", resultTest())
	//
	innerF := func(x int, y int) {
		fmt.Println("innerfunction内部関数:変数へfunctionを格納", x+y)
	}
	innerF(1, 2)

	//即時関数
	func(x int, y int) {
		fmt.Println("ImmediateFunction即時関数:", x+y)
	}(1, 3)
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
	fmt.Println("------castStringToInt")
	test := "14"
	i, _ := strconv.Atoi(test) // i, err だけど使わない時は_にしておく
	fmt.Printf("%T %v", i, i)
}

//配列とスライス
func arrayAndSlice() {
	fmt.Println("------arrayAndSlice")
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println("array:", a)

	var b []int = []int{1, 2, 3, 4}
	b[0] = 100
	b = append(b, 5)
	fmt.Println("arrayb:", b)
	fmt.Println("arrayb[1:3:5]:", b[:2])
	fmt.Println("arrayb[1:3:5]:", b[2:])
	fmt.Println("arrayb[1:3:5]:", b[1:3])

	b[0] = 100
	b = append(b, 5)
	fmt.Println("arrayb:", b)
	fmt.Println("arrayb[1:3:5]:", b[:2])
	fmt.Println("arrayb[1:3:5]:", b[2:])
	fmt.Println("arrayb[1:3:5]:", b[1:3])

	var c = [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println("slice]:", c)
}

//スライスをmakeでつくる
func sliceMake() {
	fmt.Println("------sliceMake")

	a := make([]int, 0) //0のスライスをメモリーに確保
	fmt.Printf("make:len=%d cap=%d value=%v\n", len(a), cap(a), a)

	var b []int //nilのまま(メモリーに確保しない)
	fmt.Printf("var:len=%d cap=%d value=%v\n", len(b), cap(b), b)

	c := make([]int, 3) //3のスライスをメモリーに確保
	fmt.Printf("make:len=%d cap=%d value=%v\n", len(c), cap(c), c)

	test1 := make([]int, 4)
	for i := 0; i < 5; i++ {
		test1 = append(test1, i)
		fmt.Println("test1:", i, test1)
	}

	var test2 []int
	for i := 0; i < 5; i++ {
		test2 = append(test2, i)
		fmt.Println("test2:", i, test2)
	}
}

// マップ
func mapTest() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)

	m["banana"] = 300
	fmt.Println(m)

	m["newFruit"] = 500
	fmt.Println(m)

	fmt.Println(m["nothing"]) //デフォルトは0

	v, check := m["newFruit"] // 存在チェック ok=true
	fmt.Println(v, check)

	v2, check2 := m["nothing"] // 存在チェック ok=false
	fmt.Println(v2, check2)

}

func byteTest() {
	fmt.Println("------byteTest")
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}

func addTest(x, y int, str string) (int, string) {
	return x + y, str
}

func resultTest() (result string) { //戻り値を事前に設定(わかりやすい)
	result = "testStr"
	return //戻り値をあらかじめ宣言しておくと、returnの設定は不要
}

func template() {
	fmt.Println("------template")
	test := "test"
	fmt.Println(test)
}
