package main

import "fmt"

func if_branch() {
	// if 分支
	/* 定义局部变量 */
	var a int = 11

	/* 使用 if 语句判断布尔表达式 */
	if a < 22 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 22\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)
	fmt.Println("--------------------------------")

	// if 变体
	if num := 100; num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}
}

func switch_branch() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C" //case 后可以由多个数值
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
	fmt.Println("--------------------------------")

	// fallthrough: 贯通后续的 case
	switch x := 5; x {
	default:
		fmt.Println(x)
	case 5:
		x += 10
		fmt.Println(x)
		fallthrough
	case 6:
		x += 20
		fmt.Println(x)
	}
	fmt.Println("--------------------------------")

	// case 变体
	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}
}

// 同一个目录下只能有一个 main 函数
// func main() {
// 	if_branch()
// 	switch_branch()
// }
