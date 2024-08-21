package main

import "fmt"

func operator() {

	// 1.算数运算符：+，-，*，/，%
	var a, b int = 10, 5
	fmt.Println(a+b, a-b, a*b, a/b, a%b)

	// 2. 关系运算符：==，!=，<，<=，>，>=
	fmt.Println(a == b, a != b, a < b, a <= b, a > b, a >= b)

	// 3. 逻辑运算符：&&，||，!
	fmt.Println(true && false, true || false, !true)

	// 4. 位运算符：&，|，^，<<，>>
	fmt.Println(a&b, a|b, a^b, a<<1, a>>1)

	// 5. 赋值运算符：=，+=，-=，*=，/=，%=
	var c int = 1
	c += 2
	fmt.Println(c)
	c -= 1
	fmt.Println(c)
	c *= 2
	fmt.Println(c)
	c /= 2
	fmt.Println(c)
	c %= 2
	fmt.Println(c)

}

// 同一个目录下只能有一个 main 函数
// func main() {
// 	operator()
// }
