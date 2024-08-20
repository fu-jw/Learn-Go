package main

import "fmt"

func funcIOTA() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "hi" //独立值，iota += 1
		e        //"hi"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

// 同一个目录下只能有一个 main 函数
func main() {
	funcIOTA()
}
