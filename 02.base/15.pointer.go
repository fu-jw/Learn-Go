package main

import "fmt"

func pointer() {
	/*
		指针：pointer
			存储了另一个变量的内存地址的变量。
	*/
	//1.定义一个int类型的变量
	a := 10
	fmt.Println("a的数值是：", a)     //10
	fmt.Printf("%T\n", a)        //int
	fmt.Printf("a的地址是：%p\n", &a) //a的地址是：0x1400000e0f8

	//2.创建一个指针变量，用于存储变量a的地址
	var p1 *int
	fmt.Println(p1)           //<nil>，空指针
	p1 = &a                   //p1指向了a的内存地址
	fmt.Println("p1的数值：", p1) //p1的数值： 0x1400000e0f8, p1中存储的是a的地址
	fmt.Printf("p1自己的地址：%p\n", &p1)
	fmt.Println("p1的数值，是a的地址，该地址存储的数据：", *p1) //10, 获取指针指向的变量的数值

	//3.操作变量，更改数值 ，并不会改变地址
	a = 100
	fmt.Println(a)
	fmt.Printf("%p\n", &a) // 0x1400000e0f8

	//4.通过指针，改变变量的数值
	*p1 = 200
	fmt.Println(a)

	//5.指针的指针
	var p2 **int
	fmt.Println(p2) // <nil>
	p2 = &p1
	fmt.Printf("%T,%T,%T\n", a, p1, p2)               //int, *int, **int
	fmt.Println("p2的数值：", p2)                         //p1的地址, 0x1400005a038
	fmt.Printf("p2自己的地址：%p\n", &p2)                   // p2自己的地址：0x1400005a040
	fmt.Println("p2中存储的地址，对应的数值，就是p1的地址，对应的数据：", *p2) // 0x1400000e0f8
	fmt.Println("p2中存储的地址，对应的数值，再获取对应的数值：", **p2)     // 200

}

// ----------------------------------------------------------------
// 指针数组
func pointerArr() {
	/*
		数组指针：首先是一个指针，一个数组的地址。
			*[4]Type

		指针数组：首先是一个数组，存储的数据类型是指针
			[4]*Type


			*[5]float64,指针，一个存储了5个浮点类型数据的数组的指针
			*[3]string，指针，数组的指针，存储了3个字符串
			[3]*string，数组，存储了3个字符串的指针地址的数组
			[5]*float64，数组，存储了5个浮点数据的地址的数组
			*[5]*float64，指针，一个数组的指针，存储了5个float类型的数据的指针地址的数组的指针
			*[3]*string，指针，存储了3个字符串的指针地址的数组的指针
			**[4]string，指针，存储了4个字符串数据的数组的指针的指针
			**[4]*string，指针，存储了4个字符串的指针地址的数组，的指针的指针
	*/
	//1.创建一个普通的数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)

	//2.创建一个指针，存储该数组的地址--->数组指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)         //&[1 2 3 4]
	fmt.Printf("%p\n", p1)  //数组arr1的地址
	fmt.Printf("%p\n", &p1) //p1指针自己的地址

	//3.根据数组指针，操作数组
	(*p1)[0] = 100
	fmt.Println(arr1)

	p1[0] = 200 //简化写法
	fmt.Println(arr1)

	//4.指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr2) //[1 2 3 4]
	fmt.Println(arr3)
	arr2[0] = 100
	fmt.Println(arr2)
	fmt.Println(a)
	*arr3[0] = 200
	fmt.Println(arr3)
	fmt.Println(a)

	b = 1000
	fmt.Println(arr2)
	fmt.Println(arr3)
	for i := 0; i < len(arr3); i++ {
		fmt.Println(*arr3[i])
	}
}

// ----------------------------------------------------------------
// 指针函数
func pointerFunc() {
	/*
		函数指针：一个指针，指向了一个函数的指针。
			因为go语言中，function，默认看作一个指针，没有*。


			slice,map,function

		指针函数：一个函数，该函数的返回值是一个指针。

	*/
	var a func()
	a = fun31
	a()

	arr1 := fun32()
	fmt.Printf("arr1的类型：%T，地址：%p，数值：%v\n", arr1, &arr1, arr1)

	arr2 := fun33()
	fmt.Printf("arr32的类型：%T，地址：%p，数值：%v\n", arr2, &arr2, arr2)
	fmt.Printf("arr32指针中存储的数组的地址：%p\n", arr2)
}
func fun33() *[4]int {
	arr := [4]int{5, 6, 7, 8}
	fmt.Printf("函数中arr的地址：%p\n", &arr)
	return &arr
}

func fun32() [4]int { //普通函数
	arr := [4]int{1, 2, 3, 4}
	return arr
}

func fun31() {
	fmt.Println("fun31().....")
}

// ----------------------------------------------------------------
// 指针参数
func pointerParam() {
	/*
		指针作为参数：

		参数的传递：值传递，引用传递
	*/
	a := 10
	fmt.Println("fun41()函数调用前，a:", a)
	funP1(a)
	fmt.Println("fun41()函数调用后，a：", a)

	funP2(&a)
	fmt.Println("fun42()函数调用后，a:", a)

	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("fun43()函数调用前：", arr1)
	funP3(arr1)
	fmt.Println("fun43()函数调用后：", arr1)

	funP4(&arr1)
	fmt.Println("fun44()函数调用后：", arr1)

}
func funP4(p2 *[4]int) { // 引用传递
	fmt.Println("funP4()函数中的数组指针：", p2)
	p2[0] = 200
	fmt.Println("funP4()函数中的数组指针：", p2)
}

func funP3(arr2 [4]int) { // 值传递
	fmt.Println("funP3()函数中数组的：", arr2)
	arr2[0] = 100
	fmt.Println("funP3()函数中修改数组：", arr2)
}

func funP1(num int) { // 值传递：num = a = 10
	fmt.Println("funP1()函数中，num的值：", num)
	num = 100
	fmt.Println("funP1()函数中修改num：", num)
}

func funP2(p1 *int) { //传递的是a的地址，就是引用传递
	fmt.Println("funP2()函数中，p1：", *p1)
	*p1 = 200
	fmt.Println("funP2()函数中，修改p1：", *p1)
}

// func main() {
// 	// pointer()
// 	// pointerArr()
// 	// pointerFunc()
// 	pointerParam()
// }
