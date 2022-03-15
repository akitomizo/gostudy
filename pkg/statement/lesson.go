package statement

import (
	"fmt"
	"log"
	"os"
)

func Statement() {
	fmt.Println("------statement")

	testIf()
	testFor()
	testRange()
	testSwitch()
	testDefer()
	testLog()
	paniAndRecover()
	fmt.Println("------end")
}

func testIf() {

	fmt.Println("------testIf")
	num := 9
	if num%2 == 0 {
		fmt.Println("if")
	} else if num%3 == 0 {
		fmt.Println("elseif")
	} else {
		fmt.Println("else")
	}

	if result := 1; result == 2 {
		fmt.Println("result-if")
	} else {
		fmt.Println("result-else")
	}
}

func testFor() {
	fmt.Println("------testFor")
	// continue
	for i := 0; i < 4; i++ {
		if i == 2 {
			fmt.Println("continue:i=", i)
			continue
		}
		fmt.Println("i=", i)
	}

	// break
	for i := 0; i < 4; i++ {
		if i == 2 {
			fmt.Println("break:i=", i)
			break
		}
		fmt.Println("i=", i)
	}

	// 省略形
	j := 1
	for j < 10 {
		j++
		fmt.Println("j=", j)
	}
}
func testRange() {

	l := []string{"test1", "test2", "test3"}
	// これと同じ
	// for i := 0; i < len(l); i++ {
	// 	fmt.Println("i=", i, "l[i]=", l[i])
	// }

	for i, v := range l {
		fmt.Println("i=", i, "l[i]=", v)
	}

	for _, v := range l {
		fmt.Println("l[i]=", v)
	}

	m := map[string]int{"aplle": 100, "banana": 200}

	for k, v := range m {
		fmt.Println("k=", k, "v=", v)
	}

	for v := range m {
		fmt.Println("v=", v)
	}
}

func testSwitch() {
	fmt.Println("------switch")
	os := "windows"
	switch os {
	case "mac":
		fmt.Println("mac")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("Defaut!")
	}

	switch os2 := "mac"; os2 {
	case "mac":
		fmt.Println("mac")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("Defaut!")
	}
}

func testDefer() {
	fmt.Println("------defer")
	fmt.Println("1")
	defer fmt.Println("5")
	defer fmt.Println("4")
	defer fmt.Println("3")
	fmt.Println("2")

	//使用例
	file, _ := os.Open("pkg/statement/lesson.go")
	defer file.Close() //　※Openと一緒にCloseしておけば忘れにくい
	data := make([]byte, 10)
	file.Read(data)
	fmt.Println(string(data))
}

func testLog() {
	fmt.Println("------log")
	log.Println("Logging!")
	log.Printf("%T %v", "test", "test")
	// log.Fatalf("%T %v", "error", "error") // fatalで出力すると、その時点で実行終了する。
	// log.Fatalln("error")
}

// Panicは使用の推奨されない
func paniAndRecover() {

	fmt.Println("------PanicAndRecover")
	// panic("panicTest")　// ※recover前に実行するとPanic発生
	defer func() {
		s := recover() // recoverでPanicの復旧
		fmt.Println(s)
		fmt.Println("------PanicAndRecover1")
	}()
	panic("panicTest") // panic発生
	fmt.Println("------PanicAndRecover2")
}

func template() {
	fmt.Println("------template")
	test := "test"
	fmt.Println(test)
}
