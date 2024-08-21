package main

import "fmt"

// break
func funcArray() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	//  为数组 n 初始化元素
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	// 输出每个数组元素的值
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	// 数组的长度：也可以不指定，由编译器决定
	a := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is", len(a))

	// 使用 range 遍历数组
	arr := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range arr { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)

	// 数组是值类型
	aa := [...]string{"USA", "China", "India", "Germany", "France"}
	bb := aa // aa copy of aa is assigned to bb
	bb[0] = "Singapore"
	fmt.Println("aa is ", aa)
	fmt.Println("bb is ", bb)
}

// 同一个目录下只能有一个 main 函数
func main() {
	funcArray()
}
