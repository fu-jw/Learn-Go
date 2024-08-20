package main

import "fmt"

func dataType() {
	var i int
	var f float64
	var b bool
	var s string

	i = 100
	f = 3.14
	b = true
	s = "Hello, World!"

	fmt.Printf("i=%d, f=%.2f, b=%t, s=%s\n", i, f, b, s)
}

// 同一个目录下只能有一个 main 函数
func main() {
	dataType()
}
