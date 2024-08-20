package main

import "fmt"

func funcConst() {
	/*
		常量：
		1.概念：同变量类似，程序执行过程中数值不能改变
		2.语法：
			显式类型定义
			隐式类型定义
		3.常数：
			固定的数值：11，"hi"
	*/
	fmt.Println(11)
	fmt.Println("hi")

	//1.定义常量
	const PATH string = "https:www.baidu.com"
	const PI = 3.14
	fmt.Println(PATH)
	fmt.Println(PI)

	//2.尝试修改常量的数值
	//PATH = "http://www.sina.com" //cannot assign to PATH

	//3.定义一组常量
	const C1, C2, C3 = 22, 3.14, "哈哈"
	const (
		MALE    = 0
		FEMALE  = 1
		UNKNOWN = 3
	)
	//4.一组常量中，如果某个常量没有初始值，默认和上一行一致
	const (
		a int = 100
		b
		c string = "fredo"
		d
		e
	)
	fmt.Printf("%T,%d\n", a, a) // int,100
	fmt.Printf("%T,%d\n", b, b) // int,100
	fmt.Printf("%T,%s\n", c, c) // string,fredo
	fmt.Printf("%T,%s\n", d, d) // string,fredo
	fmt.Printf("%T,%s\n", e, e) // string,fredo

	//5. 枚举类型：使用常量组作为枚举类型。一组相关数值的数据
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
}

// 同一个目录下只能有一个 main 函数
func main() {
	funcConst()
}
