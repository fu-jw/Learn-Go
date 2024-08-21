package main

import "fmt"

func funcSlice() {
	// 定义切片
	a := [5]int{76, 77, 78, 79, 80}
	var slice_b []int = a[1:4] //creates slice_a slice from slice_a[1] to slice_a[3]
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

// 同一个目录下只能有一个 main 函数
func main() {
	funcSlice()
}
