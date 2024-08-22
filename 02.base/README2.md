## 函数

### 函数是什么

函数是执行特定任务的代码块

### 函数声明

语法格式：

```go
func funcName(paramName type1, paramName type2) (output1 type1, output2 type2) {
// 这里是处理逻辑代码
。。。。。。

// 可以返回多个值
return value1, value2, ...
}
```

- func：函数由 func 开始声明
- funcName：函数名称，函数名和参数列表一起构成了函数签名。
- paramName type：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
- output1 type1, output2 type2：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
- 上面返回值声明了两个变量 output1 和 output2，如果你不想声明也可以，直接就两个类型。
- 如果只有一个返回值且不声明返回值变量，那么你可以省略包括返回值的括号（即一个返回值可以不声明返回类型）
- 函数体：函数定义的代码集合。

### 函数使用

示例代码：

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int = 200
   var ret int

   /* 调用函数并返回最大值 */
   ret = max(a, b)

   fmt.Printf( "最大值是 : %d\n", ret )
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
   /* 定义局部变量 */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}
```

### 函数的参数

形式参数：定义函数时，用于接收外部传入的数据，叫做形式参数，简称形参。

实际参数：调用函数时，传给形参的实际的数据，叫做实际参数，简称实参。

函数调用：

- 函数名称必须匹配
- 实参与形参必须一一对应：顺序，个数，类型

#### 可变参

Go 函数支持变参。接受变参的函数是有着不定数量的参数的。
为了做到这点，首先需要定义函数使其接受变参：

```go
func myFunc(arg ...int) {
    // ....
}
```

`arg ...int`告诉 Go 这个函数接受不定数量的参数。注意，这些参数的类型全部是 int。在函数体中，变量 arg 是一个 int 的 slice：

```go
for _, n := range arg {
fmt.Printf("And the number is: %d\n", n)
}
```

#### 参数传递

go 语言函数的参数也是存在**值传递**和**引用传递**

函数运用场景:

- **值传递**:

```go
package main

import (
   "fmt"
   "math"
)

func main(){
   /* 声明函数变量 */
   getSquareRoot := func(x float64) float64 {
      return math.Sqrt(x)
   }

   /* 使用函数 */
   fmt.Println(getSquareRoot(9))

}
```

- **引用传递**:

这就牵扯到了所谓的指针。
我们知道，变量在内存中是存放于一定地址上的，修改变量实际是修改变量地址处的内
存。只有 add1 函数知道 x 变量所在的地址，才能修改 x 变量的值。所以我们需要将 x 所在地址 &x 传入函数，并将函数的参数的类型由 int 改为\*int，即改为指针类型，才能在函数中修改 x 变量的值。此时参数仍然是按 copy 传递的，只是 copy 的是一个指针。

```go
package main

import "fmt"

//简单的一个函数，实现了参数+1的操作
func add1(a *int) int { // 请注意，
    *a = *a+1 // 修改了a的值
    return *a // 返回新值
}
func main() {
    x := 3
    fmt.Println("x = ", x) // 应该输出 "x = 3"
    x1 := add1(&x) // 调用 add1(&x) 传x的地址
    fmt.Println("x+1 = ", x1) // 应该输出 "x+1 = 4"
    fmt.Println("x = ", x) // 应该输出 "x = 4"
}
```

- 传指针使得多个函数能操作同一个对象。
- 传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数值传递的话, 在每次 copy 上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。
- **Go 语言中 slice，map 这三种类型的实现机制类似指针**，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变 slice 的长度，则仍需要取地址传递指针）

### 函数的返回值

一个函数被调用后，返回给调用处的执行结果，叫做函数的返回值。

调用处需要使用变量接收该结果

**注意：在 GO 语言中，一个函数可以返回多个值**。

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
   return y, x
}

func main() {
   a, b := swap("Mahesh", "Kumar")
   fmt.Println(a, b)
}
```

```go
// 注意：函数的返回值
func SumAndProduct(A, B int) (add int, Multiplied int) {
    add = A+B
    Multiplied = A*B
    return
}
```

- 空白标识符

\_是 Go 中的空白标识符。它可以代替任何类型的任何值。让我们看看这个空白标识符的用法。

比如 rectProps 函数返回的结果是面积和周长，如果我们只要面积，不要周长，就可以使用空白标识符。

示例代码：

```go
func rectProps(length, width float64) (float64, float64) {
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
}
func main() {
    // 只接受一个返回值，另一个舍弃
    area, _ := rectProps(10.8, 5.6) // perimeter is discarded
    fmt.Printf("Area %f ", area)
}
```

### 函数的作用域

作用域：变量可以使用的范围。

- 局部变量

  - 一个函数内部定义的变量，就叫做局部变量
  - 变量在哪里定义，就只能在哪个范围使用，超出这个范围，我们认为变量就被销毁了。

- 全局变量

  - 一个函数外部定义的变量，就叫做全局变量
  - 所有的函数都可以使用，而且共享这一份数据

### 函数的本质

函数也是 Go 语言中的一种数据类型，可以作为另一个函数的参数，也可以作为另一个函数的返回值

### defer 函数

即延迟（defer）语句，延迟语句被用于执行一个函数调用，在这个函数之前，延迟语句返回。

#### 延迟函数

你可以在函数中添加多个 defer 语句。当函数执行到最后时，这些 defer 语句会按照逆序执行，最后该函数返回。特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题

- **如果有很多调用 defer，那么 defer 是采用`后进先出`模式**
- 在离开所在的方法时，执行（报错的时候也会执行）

```go
func ReadWrite() bool {
    file.Open("file")
    defer file.Close()
    if failureX {
          return false
    } if failureY {
          return false
    }
    return true
}
```

最后才执行`file.Close()`

示例代码：

```go
package main

import "fmt"

func main() {
    a := 1
    b := 2
    defer fmt.Println(b)
    fmt.Println(a)
}
```

运行结果：

```go
1
2
```

示例代码：

```go
package main

import (
    "fmt"
)

func finished() {
    fmt.Println("Finished finding largest")
}

func largest(nums []int) {
    defer finished()
    fmt.Println("Started finding largest")
    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    fmt.Println("Largest number in", nums, "is", max)
}

func main() {
    nums := []int{78, 109, 2, 563, 300}
    largest(nums)
}
```

运行结果：

```go
Started finding largest
Largest number in [78 109 2 563 300] is 563
Finished finding largest
```

#### 延迟方法

延迟并不仅仅局限于函数。延迟一个方法调用也是完全合法的。让我们编写一个小程序来测试这个。

示例代码：

```go
package main

import (
    "fmt"
)


type person struct {
    firstName string
    lastName string
}

func (p person) fullName() {
    fmt.Printf("%s %s",p.firstName,p.lastName)
}

func main() {
    p := person {
        firstName: "John",
        lastName: "Smith",
    }
    defer p.fullName()
    fmt.Printf("Welcome ")
}
```

运行结果：

```go
Welcome John Smith
```

#### 延迟参数

延迟函数的参数在执行延迟语句时被执行，而不是在执行实际的函数调用时执行。

让我们通过一个例子来理解这个问题。

示例代码：

```go
package main

import (
    "fmt"
)

func printA(a int) {
    fmt.Println("value of a in deferred function", a)
}
func main() {
    a := 5
    defer printA(a)
    a = 10
    fmt.Println("value of a before deferred function call", a)

}
```

运行结果：

```go
value of a before deferred function call 10
value of a in deferred function 5
```

#### 堆栈的推迟

当一个函数有多个延迟调用时，它们被添加到一个堆栈中，并在 Last In First Out（LIFO）后进先出的顺序中执行。

我们将编写一个小程序，它使用一堆 defers 打印一个字符串。示例代码：

```go
package main

import (
    "fmt"
)

func main() {
    name := "Naveen"
    fmt.Printf("Orignal String: %s\n", string(name))
    fmt.Printf("Reversed String: ")
    for _, v := range []rune(name) {
        defer fmt.Printf("%c", v)
    }
}
```

运行结果：

```go
Orignal String: Naveen
Reversed String: neevaN
```

#### defer 小结

- 当外围函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会真正的结束执行。
- 当执行外围函数中的 return 语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正返回。
- 当外围函数中的代码引发运行恐慌时，只有其中所有的延迟函数都执行完毕后，该运行时恐慌才会真正被扩展至调用函数。
