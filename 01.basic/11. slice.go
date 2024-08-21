package main

import "fmt"

func funcSlice() {
	// 定义切片
	a := [5]int{76, 77, 78, 79, 80}
	var slice_b []int = a[1:4] //creates slice_b slice from a[1] to a[3]
	fmt.Println(slice_b)       // [77 78 79]

	// 修改切片
	slice_c := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	slice_d := slice_c[2:5]
	fmt.Println("array before", slice_c)
	// array before [57 89 90 82 100 78 67 69 59]
	for i := range slice_d {
		slice_d[i]++
	}
	fmt.Println("array after", slice_c)
	// array after [57 89 91 83 101 78 67 69 59]

	// 当多个片共享相同的底层数组时，每个元素所做的更改将在数组中反映出来。
	num := [3]int{78, 79, 80}
	slice_e := num[:] //creates a slice which contains all elements of the array
	slice_f := num[:]
	fmt.Println("array before change 1", num)
	// array before change 1 [78 79 80]
	slice_e[0] = 100
	fmt.Println("array after modification to slice slice_e", num)
	// array after modification to slice slice_e [100 79 80]
	slice_f[1] = 101
	fmt.Println("array after modification to slice slice_f", num)
	// array after modification to slice slice_f [100 101 80]

	// len 和 cap
	var number1 = make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(number1), cap(number1), number1)
	// len=3 cap=5 slice=[0 0 0]

	// 一个切片在未初始化之前默认为 nil，长度为 0
	var number2 []int
	fmt.Printf("len=%d cap=%d slice=%v\n", len(number2), cap(number2), number2)
	// len=0 cap=0 slice=[]
	if number2 == nil {
		fmt.Printf("切片是空的")
	}
	// 切片是空的

	// append 和 copy
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func funcSlice2() {
	/*
		slice := arr[start:end]
		 	切片中的数据：[start,end)
		 	arr[:end],从头到end
		 	arr[start:]从start到末尾

		 从已有的数组上，直接创建切片，该切片的底层数组就是当前的数组。
		 	长度是从start到end切割的数据量。
			但是容量从start到数组的末尾。
	*/
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("----------1.已有数组直接创建切片--------------------")
	s1 := a[:5]  //1-5
	s2 := a[3:8] //4-8
	s3 := a[5:]  // 6-10
	s4 := a[:]   // 1-10
	fmt.Println("a:", a)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", s1)

	fmt.Println("----------2.长度和容量--------------------")
	fmt.Printf("s1	len:%d,cap:%d\n", len(s1), cap(s1)) //s1	len:5,cap:10
	fmt.Printf("s2	len:%d,cap:%d\n", len(s2), cap(s2)) //s2	len:5,cap:7
	fmt.Printf("s3	len:%d,cap:%d\n", len(s3), cap(s3)) //s3	len:5,cap:5
	fmt.Printf("s4	len:%d,cap:%d\n", len(s4), cap(s4)) //s4	len:10,cap:10

	fmt.Println("----------3.更改数组的内容--------------------")
	a[4] = 100
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println("----------4.更改切片的内容--------------------")
	s2[2] = 200
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println("----------4.更改切片的内容--------------------")
	s1 = append(s1, 1, 1, 1, 1)
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println("----------5.添加元素切片扩容--------------------")
	fmt.Println(len(s1), cap(s1))

	s1 = append(s1, 2, 2, 2, 2, 2)
	fmt.Println(a)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(len(s1), cap(s1))
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", &a)
}

func funcSlice3() {
	/*
		按照类型来分：
			基本类型：int，float，string，bool
			复合类型：array，slice，map，struct，pointer，function，chan

		按照特点来分：
			值类型：int，float，string，bool，array
				传递的是数据副本
			引用类型：Slice
				传递的地址,多个变量指向了同一块内存地址，

		所以：切片是引用类型的数据，存储了底层数组的引用
	*/

	//1.数组：值类型
	a1 := [4]int{1, 2, 3, 4}
	a2 := a1 //值传递：传递的是数据
	fmt.Println(a1, a1)
	a1[0] = 100
	fmt.Println(a1, a2)

	//2.切片：引用类型
	s1 := []int{1, 2, 3, 4}
	s2 := s1
	fmt.Println(s1, s2)
	s1[0] = 100
	fmt.Println(s1, s2)

	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p\n", &s2)
}

func funcSlice4() {
	/*
		深拷贝：拷贝的是数据本身。
			值类型的数据，默认都是深拷贝：array，int，float，string，bool，struct

		浅拷贝：拷贝的是数据 地址。
			导致多个变量指向同一块内存
			引用类型的数据，默认都是浅拷贝：slice，map，

			因为切片是引用类型的数据，直接拷贝的是地址。

		func copy(dst, src []Type) int
			可以实现切片的拷贝
	*/

	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0) //len:0,cap:0
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1)
	fmt.Println(s2)

	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)

	//copy()
	s3 := []int{7, 8, 9}
	fmt.Println(s2)
	fmt.Println(s3)

	//copy(s2,s3) //将s3中的元素，拷贝到s2中
	//copy(s3,s2) //将s2中的元素，拷贝到s3中
	copy(s3[1:], s2[2:])
	fmt.Println(s2)
	fmt.Println(s3)
}

// 同一个目录下只能有一个 main 函数
// func main() {
// 	funcSlice()
// 	funcSlice2()
// 	funcSlice3()
// 	funcSlice4()
// }
