package main

import "fmt"

func myFunc1() {
	/*
		函数：function
		一、概念：
			具有特定功能的代码，可以被多次调用执行。
		二、意义：
			1.可以避免重复的代码
			2.增强程序的扩展性
		三、使用
			step1：函数的定义，也叫声明
			step2：函数的调用，就是执行函数中代码的过程
		四、语法
			1.定义函数的语法
				func funcName(paramName type1, paramName type2) (output1 type1, output2 type2) {
					//这里是处理逻辑代码
					//返回多个值
					return value1, value2
				}
				A:func，定义函数的关键字
				B:funcName，函数的名字
				C:()，函数的标志
				D:参数列表：形式参数用于接收外部传入函数中的数据
				E:返回值列表：函数执行后返回给调用处的结果

			2、调用函数的语法
				函数名(实际参数)

			函数的调用处，就是函数调用的位置

			3、注意事项
				A：函数必须先定义，再调用，如果不定义：undefined: getSum
					定义了函数，没有调用，那么函数就失去了意义

				B：函数名不能冲突
				C：main()，是一个特殊的函数，作为程序的入口，由系统自动调用
					而其他函数，程序中通过函数名来调用。
	*/
	//求1-10的和
	getSum() //函数的调用处

	fmt.Println("hello world....")

	//求1-10的和
	getSum()

	fmt.Println("aaa...")

	//求1-10的和
	getSum()
}

// 定义一个函数：用于求1-10的和
func getSum() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("1-100的和是：%d\n", sum)
}

// ///////////////////////////////////////////////////////
func myFunc2() {
	/*
		函数的参数：
		形式参数：也叫形参。函数定义的时候，用于接收外部传入的数据的变量。
		函数中，某些变量的数值无法确定，需要由外部传入数据。

		实际参数：也叫实参。函数调用的时候，给形参赋值的实际的数据


		函数调用：
		1.函数名：声明的函数名和调用的函数名要统一
		2.实参必须严格匹配形参：顺序，个数，类型，一一对应的。
	*/
	// 求2个整数的和
	getAdd(20, 10)
	getAdd2(1, 2)

	fun1(1.3, 2.4, "hello")
}

func getAdd(a int, b int) {
	sum := a + b
	fmt.Printf("%d + %d = %d \n", a, b, sum)
}

func getAdd2(a, b int) { //参数的类型一致，可以简写在一起
	fmt.Printf("a:%d,b:%d\n", a, b)
}

func fun1(a, b float64, c string) {
	fmt.Printf("a:%.2f,b:%.2f,c:%s\n", a, b, c)
}

/////////////////////////////////////////////////////////
// 可变参数

func myFunc3() {
	/*
		可变参数：
			概念：一个函数的参数的类型确定，但是个数不确定，就可以使用可变参数。

			语法：
				参数名 ... 参数的类型

				对于函数，可变参数相当于一个切片。
				调用函数的时候，可以传入0-多个参数。

				Println(),Printf(),Print()
				append()

			注意事项：
				A：如果一个函数的参数是可变参数，同时还有其他的参数，可变参数要放在
					参数列表的最后。
				B：一个函数的参数列表中最多只能有一个可变参数。

	*/
	//1.求和
	getSum3()

	getSum3(1, 2, 3, 4, 5)

	getSum3(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	//2.切片
	s1 := []int{1, 2, 3, 4, 5}
	getSum3(s1...)
}

func getSum3(nums ...int) {
	//fmt.Printf("%T\n",nums) //[]int
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println("总和是：", sum)
}

// 可变参数必须在最后
func fun3(s1, s2 string, nums ...float64) {

}

// ////////////////////////////////////////////
// 参数传递
func myFunc4() {
	/*
		数据类型：
			一：按照数据类型来分：
					基本数据类型：
						int，float,string，bool
					复合数据类型：
						array，slice，map，struct，interface。。。。

			二：按照数据的存储特点来分：
					值类型的数据：操作的是数值本身。
						int，float64,bool，string，array
					引用类型的数据：操作的是数据的地址
						slice，map，chan

		参数传递：
			A：值传递：传递的是数据的副本。修改数据，对于原始的数据没有影响的。
				值类型的数据，默认都是值传递：基础类型，array，struct



			B：引用传递：传递的是数据的地址。导致多个变量指向同一块内存。
				引用类型的数据，默认都是引用传递：slice，map，chan

	*/
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("函数调用前，数组的数据：", arr1) //[1 2 3 4]
	fun41(arr1)
	fmt.Println("函数调用后，数组的数据：", arr1) //[1 2 3 4]

	fmt.Println("---------------------------------")

	s1 := []int{1, 2, 3, 4}
	fmt.Println("函数调用前，切片的数据：", s1) // [1 2 3 4]
	fun42(s1)
	fmt.Println("函数调用后，切片刀数据：", s1) //[100 2 3 4]
}

func fun42(s2 []int) {
	fmt.Println("函数中，切片的数据：", s2) //[1 2 3 4]
	s2[0] = 100
	fmt.Println("函数中，切片的数据更改后：", s2) //[100 2 3 4]
}
func fun41(arr2 [4]int) {
	fmt.Println("函数中，数组的数据：", arr2) //[1 2 3 4]
	arr2[0] = 100
	fmt.Println("函数中，数组的数据修改后：", arr2) //[100 2 3 4]
}

// /////////////////////////////////////////////////////////
// 返回值
func myFunc5() {
	/*
		函数的返回值：
			一个函数的执行结果，返回给函数的调用处。执行结果就叫做函数的返回值。

		return语句：
			一个函数的定义上有返回值，那么函数中必须使用return语句，将结果返回给调用处。
			函数返回的结果，必须和函数定义的一致：类型，个数，顺序。

			1.将函数的结果返回给调用处
			2.同时结束了该函数的执行

		空白标识符，专门用于舍弃数据：_
	*/

	//1.设计一个函数，用于求1-10的和，将结果在主函数中打印输出
	res := getSum5()
	fmt.Println("1-10的和：", res)

	fmt.Println(getSum52()) //5050

	res1, res2 := rectangle(5, 3)
	fmt.Println("周长：", res1, "，面积：", res2)
	res3, res4 := rectangle2(5, 3)
	fmt.Println("周长：", res3, "，面积：", res4)

	_, res5 := rectangle(5, 3)
	fmt.Println(res5)
}

// 函数，用于求矩形的周长和面积
func rectangle(len, wid float64) (float64, float64) {
	perimeter := (len + wid) * 2
	area := len * wid
	return perimeter, area
}

func rectangle2(len, wid float64) (peri float64, area float64) {
	peri = (len + wid) * 2
	area = len * wid
	return
}

func fun51() (float64, float64, string) {
	return 2.4, 5.6, "hello"
}

// 定义一个函数，带返回值
func getSum5() int {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	return sum

}

func getSum52() (sum int) { //定义函数时，指明要返回的数据是哪一个
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return
}

//////////////////////////////////////////////////////
// 作用域

// 全局变量的定义
// num3 := 1000//不支持简短定义的写法
var num3 = 1000

func myFunc6() {
	/*
		作用域：变量可以使用的范围。
			局部变量：函数内部定义的变量，就叫做局部变量。
						变量在哪里定义，就只能在哪个范围使用，超出这个范围，我们认为变量就被销毁了。

			全局变量：函数外部定义的变量，就叫做全局变量。
						所有的函数都可以使用，而且共享这一份数据

	*/
	//定义在main函数中，所以n的作用域就是main函数的范围内
	n := 10
	fmt.Println(n)

	if a := 1; a <= 10 {
		fmt.Println(a) // 1
		fmt.Println(n) // 10
	}
	//fmt.Println(a) //不能访问a，出了作用域
	fmt.Println(n)

	if b := 1; b <= 10 {
		n := 20
		fmt.Println(b) // 1
		fmt.Println(n) // 20
	}

	fun61()
	fun62()
	fmt.Println("main中访问全局变量：", num3) //2000

}

func fun61() {
	//fmt.Println(n)
	num1 := 100
	fmt.Println("fun1()函数中：num1：", num1)
	num3 = 2000
	fmt.Println("fun1()函数，访问全局变量：", num3) // 2000
}

func fun62() {
	num1 := 200
	fmt.Println(num1)
	fmt.Println("fun2()函数，访问全局变量：", num3) //2000
}

// //////////////////////////////////////////////////////
// defer 函数
func myFunc7() { //外围函数
	/*
			defer的词义："延迟","推迟"
			在go语言中，使用defer关键字来延迟一个函数或者方法的执行。

			1.defer函数或方法：一个函数或方法的执行被延迟了。

			2.defer的用法：
				A：对象.close(),临时文件的删除。。。
						文件.open()
						defer close()
						读或写

				B：go语言中关于异常的处理，使用panic()和recover()
					panic函数用于引发恐慌，导致程序中断执行
					recover函数用于恢复程序的执行，recover()语法上要求必须在defer中执行。


			3.如果多个defer函数:先延迟的后执行，后延迟的先执行。

			4.defer函数传递参数的时候：defer函数调用时，就已经传递了参数数据了，只是暂时不执行函数中的代码而已。

			5.defer函数注意点：
				defer函数：
		当外围函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会真正的结束执行。
		当执行外围函数中的return语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正返回。
		当外围函数中的代码引发运行恐慌时，只有其中所有的延迟函数都执行完毕后，该运行时恐慌才会真正被扩展至调用函数。

	*/
	//defer fun1("hello") //也被延迟了
	//fmt.Println("12345")
	//defer fun1("world") //被延迟了
	//fmt.Println("王二狗")

	a := 2
	fmt.Println(a) //2
	defer fun72(a)
	a++
	fmt.Println("main中：", a) //3

	fmt.Println(fun73())
}

func fun73() int {
	fmt.Println("fun73()函数的执行。。。")
	defer fun71("haha")
	return 0
}

func fun72(a int) { // a = 2
	fmt.Println("fun72()函数中打印a：", a) // 2
}
func fun71(s string) {
	fmt.Println(s)
}

// //////////////////////////////////////////////
// 函数类型
func myFunc8() {
	/*
		go语言的数据类型：
			基本数据类型：
					int，float，bool，string

			复合数据类型：
					array，slice，map，function，pointer，struct，interface。。。


		函数的类型：
				func(参数列表的数据类型)(返回值列表的数据类型)

	*/

	a := 10
	fmt.Printf("%T\n", a) //int
	b := [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", b) //[4]int
	/*
		[4]string
		[6]float64
	*/
	c := []int{1, 2, 3, 4}
	fmt.Printf("%T\n", c) //[]int

	d := make(map[int]string)
	fmt.Printf("%T\n", d)
	/*
		map[string]string
		map[string]map[int]string
	*/

	fmt.Printf("%T\n", fun81) //func()
	fmt.Printf("%T\n", fun82) //func(int) int
	fmt.Printf("%T\n", fun83) //func(float64, int, int) (int, int)
	fmt.Printf("%T\n", fun84) //func(string,string,int,int)(string,int ,float64)
}

func fun81() {}

func fun82(a int) int {
	return 0
}

func fun83(a float64, b, c int) (int, int) {
	return 0, 0
}

func fun84(a, b string, c, d int) (string, int, float64) {
	return "", 0, 0
}

// //////////////////////////////////////////////////////////////
// 匿名函数
func myFunc9() {
	/*
		匿名：没有名字
			匿名函数：没有名字的函数。

		定义一个匿名函数，直接进行调用。通常只能使用一次。也可以使用匿名函数赋值给某个函数变量，那么就可以调用多次了。

		匿名函数：
			Go语言是支持函数式编程：
			1.将匿名函数作为另一个函数的参数，回调函数
			2.将匿名函数作为另一个函数的返回值，可以形成闭包结构。
	*/
	fun91()
	fun91()
	fun92 := fun91
	fun92()

	//匿名函数
	func() {
		fmt.Println("我是一个匿名函数。。")
	}()

	fun93 := func() {
		fmt.Println("我也是一个匿名函数。。")
	}
	fun93()
	fun93()

	//定义带参数的匿名函数
	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

	//定义带返回值的匿名函数
	res1 := func(a, b int) int {
		return a + b
	}(10, 20) //匿名函数调用了，将执行结果给res1
	fmt.Println(res1)

	res2 := func(a, b int) int {
		return a + b
	} //将匿名函数的值，赋值给res2
	fmt.Println(res2)

	fmt.Println(res2(100, 200))
}

func fun91() {
	fmt.Println("我是fun91()函数。。")
}

// //////////////////////////////////////////////////////////////
// 高阶函数
func myFunc10() {
	/*
		高阶函数：
			根据go语言的数据类型的特点，可以将一个函数作为另一个函数的参数。

		fun1(),fun2()
		将fun1函数作为了fun2这个函数的参数。

				fun2函数：就叫高阶函数
					接收了一个函数作为参数的函数，高阶函数

				fun1函数：回调函数
					作为另一个函数的参数的函数，叫做回调函数。
	*/
	//设计一个函数，用于求两个整数的加减乘除运算
	fmt.Printf("%T\n", add)  //func(int, int) int
	fmt.Printf("%T\n", oper) //func(int, int, func(int, int) int) int

	res1 := add(1, 2)
	fmt.Println(res1)

	res2 := oper(10, 20, add)
	fmt.Println(res2)

	res3 := oper(5, 2, sub)
	fmt.Println(res3)

	fun1 := func(a, b int) int {
		return a * b
	}

	res4 := oper(10, 4, fun1)
	fmt.Println(res4)

	res5 := oper(100, 8, func(a, b int) int {
		if b == 0 {
			fmt.Println("除数不能为零")
			return 0
		}
		return a / b
	})
	fmt.Println(res5)

}
func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun) //打印3个参数
	res := fun(a, b)
	return res
}

// 加法运算
func add(a, b int) int {
	return a + b
}

// 减法
func sub(a, b int) int {
	return a - b
}

// //////////////////////////////////////////////////////////////
// 闭包
func myFunc11() {
	/*
		go语言支持函数式编程：
			支持将一个函数作为另一个函数的参数，
			也支持将一个函数作为另一个函数的返回值。

		闭包(closure)：
			一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量(外层函数中的参数，或者外层函数中直接定义的变量)，并且该外层函数的返回值就是这个内层函数。

			这个内层函数和外层函数的局部变量，统称为闭包结构。


			局部变量的生命周期会发生改变，正常的局部变量随着函数调用而创建，随着函数的结束而销毁。
			但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还要继续使用。


	*/
	res1 := increment()      //res1 = fun
	fmt.Printf("%T\n", res1) //func() int
	fmt.Println(res1)
	v1 := res1()
	fmt.Println(v1) //1
	v2 := res1()
	fmt.Println(v2) //2
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())

	res2 := increment()
	fmt.Println(res2)
	v3 := res2()
	fmt.Println(v3) //1
	fmt.Println(res2())

	fmt.Println(res1())
}

func increment() func() int { //外层函数
	//1.定义了一个局部变量
	i := 0
	//2.定义了一个匿名函数，给变量自增并返回
	fun := func() int { //内层函数
		i++
		return i
	}
	//3.返回该匿名函数
	return fun
}

// func main() {
// 	myFunc1()
// 	myFunc2()
// 	// 可变参数
// 	myFunc3()
// 	// 参数传递
// 	myFunc4()
// 	// 返回值
// 	myFunc5()
// 	// 作用域
// 	myFunc6()
// 	// defer 函数
// 	myFunc7()
// 	// 函数类型
// 	myFunc8()
// 	// 匿名函数
// 	myFunc9()
// 	// 高阶函数
// 	myFunc10()
// 	// 闭包
// 	myFunc11()
// }
