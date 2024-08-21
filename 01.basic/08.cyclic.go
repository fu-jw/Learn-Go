package main

import "fmt"

func cyclic() {
	for i := 0; i < 10; i++ {
		fmt.Printf(" %d", i)
	}
	fmt.Println()
	fmt.Println("--------------------")

	// for 变体

	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

// 同一个目录下只能有一个 main 函数
// func main() {
// 	cyclic()
// }
