package main

import "fmt"

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("Hello, World!")
	testfunc()
}

func testfunc() {
	fmt.Println("testfunc")
}
