package main

import "fmt"

// 1.定义结构体
type Person struct {
	name    string
	age     int
	sex     string
	address string
}

func funcStruct() {
	/*
		结构体：是由一系列具有相同类型或不同类型的数据构成的数据集合
			结构体成员是由一系列的成员变量构成，这些成员变量也被称为“字段”
	*/
	//1.方法一
	var p1 Person
	fmt.Println(p1)
	p1.name = "小明"
	p1.age = 30
	p1.sex = "男"
	p1.address = "北京市"
	fmt.Printf("姓名：%s,年龄：%d,性别：%s,地址：%s\n", p1.name, p1.age, p1.sex, p1.address)

	//2.方法二
	p2 := Person{}
	p2.name = "小红"
	p2.age = 28
	p2.sex = "女"
	p2.address = "上海市"
	fmt.Printf("姓名：%s,年龄：%d,性别：%s,地址：%s\n", p2.name, p2.age, p2.sex, p2.address)

	//3.方法三
	p3 := Person{name: "如花", age: 20, sex: "女", address: "杭州市"}
	fmt.Println(p3)
	p4 := Person{
		name:    "老王",
		age:     40,
		sex:     "男",
		address: "武汉市",
	}
	fmt.Println(p4)

	//4.方法四
	p5 := Person{"李小花", 25, "女", "成都"}
	fmt.Println(p5)
}

// =================================================================
// 结构体指针
func structPointer() {
	/*
		数据类型：
			值类型：int，float，bool，string，array，struct

			引用类型：slice，map，function，pointer


		通过指针：
			new()，不是nil，空指针
				指向了新分配的类型的内存空间，里面存储的零值。
	*/
	//1.结构体是值类型
	p1 := Person{"小明", 30, "男", "北京市"}
	fmt.Println(p1)
	fmt.Printf("%p,%T\n", &p1, p1)

	p2 := p1
	fmt.Println(p2)
	fmt.Printf("%p,%T\n", &p2, p2)

	p2.name = "小红"
	fmt.Println(p2)
	fmt.Println(p1)

	//2.定义结构体指针
	var pp1 *Person
	pp1 = &p1
	fmt.Println(pp1)
	fmt.Printf("%p,%T\n", pp1, pp1)
	fmt.Println(*pp1)

	//(*pp1).name = "李四"
	pp1.name = "王五"
	fmt.Println(pp1)
	fmt.Println(p1)

	//使用内置函数new(),go语言中专门用于创建某种类型的指针的函数
	pp2 := new(Person)
	fmt.Printf("%T\n", pp2)
	fmt.Println(pp2)
	//(*pp2).name
	pp2.name = "Jerry"
	pp2.age = 20
	pp2.sex = "男"
	pp2.address = "上海市"
	fmt.Println(pp2)

	pp3 := new(int)
	fmt.Println(pp3)
	fmt.Println(*pp3)
}

// ==================================================================
// 匿名结构体
func anonStruct() {
	/*
		匿名结构体和匿名字段：

		匿名结构体：没有名字的结构体，
			在创建匿名结构体时，同时创建对象
			变量名 := struct{
				定义字段Field
			}{
				字段进行赋值
			}

		匿名字段：一个结构体的字段没有字段名


		匿名函数：

	*/
	s1 := Student{name: "张三", age: 18}
	fmt.Println(s1.name, s1.age)

	func() {
		fmt.Println("hello world...")
	}()

	s2 := struct {
		name string
		age  int
	}{
		name: "李四",
		age:  19,
	}
	fmt.Println(s2.name, s2.age)

	//w1 := Worker{name:"小明",age:30}
	//fmt.Println(w1.name,w1.age)

	w2 := Worker{"小红", 32}
	fmt.Println(w2)
	fmt.Println(w2.string)
	fmt.Println(w2.int)
}

type Worker struct {
	//name string
	//age int
	string //匿名字段
	int    //匿名字段，默认使用数据类型作为名字，那么匿名字段的类型就不能重复，否则会冲突
	//string
}

type Student struct {
	name string
	age  int
}

// //////////////////////////////////////////////////////////////
// 嵌套结构体
func nestStruct() {
	/*
		结构体嵌套：一个结构体中的字段，是另一个结构体类型。
			has a
	*/

	b1 := Book{}
	b1.bookName = "西游记"
	b1.price = 45.8

	s1 := Student1{}
	s1.name = "小明"
	s1.age = 18
	s1.book = b1 //值传递
	fmt.Println(b1)
	fmt.Println(s1)
	fmt.Printf("学生姓名：%s，学生年龄：%d，看的书是：《%s》，书的价格是:%.2f\n", s1.name, s1.age, s1.book.bookName, s1.book.price)

	s1.book.bookName = "红楼梦"
	fmt.Println(s1)
	fmt.Println(b1)

	b4 := Book{bookName: "呼啸山庄", price: 76.9}
	s4 := Student2{name: "Ruby", age: 18, book: &b4}
	fmt.Println(b4)
	fmt.Println(s4)
	fmt.Println("\t", s4.book)

	s4.book.bookName = "雾都孤儿"
	fmt.Println(b4)
	fmt.Println(s4)
	fmt.Println("\t", s4.book)

	s2 := Student1{name: "小红", age: 19, book: Book{bookName: "Go语言是怎样炼成的", price: 89.7}}
	fmt.Println(s2.name, s2.age)
	fmt.Println("\t", s2.book.bookName, s2.book.price)

	s3 := Student1{
		name: "Jerry",
		age:  17,
		book: Book{
			bookName: "十万个为啥",
			price:    55.9,
		},
	}
	fmt.Println(s3.name, s3.age)
	fmt.Println("\t", s3.book.bookName, s3.book.price)
}

// 1.定义一个书的结构体
type Book struct {
	bookName string
	price    float64
}

// 2.定义学生的结构体
type Student1 struct {
	name string
	age  int
	book Book
}

type Student2 struct {
	name string
	age  int
	book *Book // book的地址
}

// =================================================================
// OOP

// 1.定义父类
type Person3 struct {
	name string
	age  int
}

// 2.定义子类
type Student3 struct {
	Person3        //模拟继承结构
	school  string //子类的新增属性
}

func OOP() {
	/*
		面向对象：OOP

		Go语言的结构体嵌套：
			1.模拟继承性：is - a
				type A struct{
					field
				}
				type B struct{
					A //匿名字段
				}

			2.模拟聚合关系：has - a
				type C struct{
					field
				}
				type D struct{
					c C //聚合关系
				}
	*/

	//1.创建父类的对象
	p1 := Person3{name: "张三", age: 30}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age)

	//2.创建子类的对象
	s1 := Student3{Person3{"李四", 17}, "清华大学"}
	fmt.Println(s1)

	s2 := Student3{Person3: Person3{name: "rose", age: 19}, school: "北京大学"}
	fmt.Println(s2)

	var s3 Student3
	s3.Person3.name = "王五"
	s3.Person3.age = 19
	s3.school = "南开大学"
	fmt.Println(s3)

	s3.name = "Jerry"
	s3.age = 16
	fmt.Println(s3)

	fmt.Println(s1.name, s1.age, s1.school)
	fmt.Println(s2.name, s2.age, s2.school)
	fmt.Println(s3.name, s3.age, s3.school)
	/*
	   s3.Person3.name---->s3.name
	   Student3结构体将Person3结构体作为一个匿名字段了
	   那么Person3中的字段，对于Student3来讲，就是提升字段
	   Student3对象直接访问Person3中的字段
	*/

}

// func main() {
// 	// funcStruct()
// 	// structPointer()
// 	// anonStruct()
// 	// nestStruct()
// 	OOP()
// }
