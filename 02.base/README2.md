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

## 包管理

Go 语言使用包（package）这种语法元素来组织源码，所有语法可见性均定义在 package 这个级别，与 Java 、python 等语言相比，这算不上什么创新，但与 C 传统的 include 相比，则是显得“先进”了许多。

```txt
myBlog
├── conf
│   └── app.conf
├── controllers
│   ├── aboutme_controller.go
│   ├── add_article_controller.go
│   ├── album_controller.go
│   ├── base_controller.go
│   ├── default.go
│   ├── delete_article_controller.go
│   ├── exit_controller.go
│   ├── home_controller.go
│   ├── login_controller.go
│   ├── register_controller.go
│   ├── show_article_controller.go
│   ├── tags_controller.go
│   ├── update_article_controller.go
│   └── upload_controller.go
├── main.go
├── models
│   ├── album_model.go
│   ├── article_model.go
│   ├── home_model.go
│   ├── tags_model.go
│   └── user_model.go
├── myblogweb
├── routers
│   └── router.go
├── static
│   ├── css
│   │   ├── blogsheet.css
│   │   └── lib
│   │       ├── highlight.css
│   │       └── login.css
│   ├── img
│   ├── js
│   │   ├── blog.js
│   │   ├── lib
│   │   │   ├── jquery-3.3.1.min.js
│   │   │   └── jquery.url.js
│   │   └── reload.min.js
│   └── upload
│       └── img
│           └── 2018
│               └── 12
│                   └── 11
│                       ├── 1544511378-bee2.png
├── tests
│   └── default_test.go
├── utils
│   ├── myUtils.go
│   └── mysqlUtils.go
└── views
    ├── aboultme.html
    ├── album.html
    ├── block
    │   ├── home_block.html
    │   └── nav.html
    ├── home.html
    ├── index.tpl
    ├── login.html
    ├── register.html
    ├── show_article.html
    ├── tags.html
    └── write_article.html
```

Go 语言的源码复用建立在包（package）基础之上。包通过 package, import, GOPATH 操作完成。

## 1、 main 包

Go 语言的入口 main() 函数所在的包（package）叫 main，main 包想要引用别的代码，需要 import 导入！

## 2、 package

src 目录是以代码包的形式组织并保存 Go 源码文件的。每个代码包都和 src 目录下的文件夹一一对应。每个子目录都是一个代码包。

> 代码包包名和文件目录名，不要求一致。比如文件目录叫 hello，但是代码包包名可以声明为 “main”，但是同一个目录下的源码文件第一行声明的所属包，必须一致！

同一个目录下的所有.go 文件的第一行添加 包定义，以标记该文件归属的包，演示语法：

```go
package 包名
```

包需要满足：

- 一个目录下的同级文件归属一个包。也就是说，在同一个包下面的所有文件的 package 名，都是一样的。
- 在同一个包下面的文件`package`名都建议设为是该目录名，但也可以不是。也就是说，包名可以与其目录不同名。
- 包名为 main 的包为应用程序的入口包，其他包不能使用。

> 在同一个包下面的文件属于同一个工程文件，不用`import`包，可以直接使用

包可以嵌套定义，对应的就是嵌套目录，但包名应该与所在的目录一致，例如：

```go
// 文件：qf/ruby/tool.go中
package ruby
// 可以被导出的函数
func FuncPublic() {
}
// 不可以被导出的函数
func funcPrivate() {
}
```

包中，通过标识符首字母是否大写，来确定是否可以被导出。首字母大写才可以被导出，视为 public 公共的资源。

## 3、 import

要引用其他包，可以使用 import 关键字，可以单个导入或者批量导入，语法演示：

A：通常导入

```go
// 单个导入
import "package"
// 批量导入
import (
  "package1"
  "package2"
  )
```

B：点操作
我们有时候会看到如下的方式导入包

```go
import(
	. "fmt"
)
```

这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调

用的`fmt.Println("hello world")`可以省略的写成`Println("hello world")`

C：起别名

别名操作顾名思义我们可以把包命名成另一个我们用起来容易记忆的名字。导入时，可以为包定义别名，语法演示：

```go
import (
  p1 "package1"
  p2 "package2"
  )
// 使用时：别名操作，调用包函数时前缀变成了我们的前缀
p1.Method()
```

D：\_操作
如果仅仅需要导入包时执行初始化操作，并不需要使用包内的其他函数，常量等资源。则可以在导入包时，匿名导入。

这个操作经常是让很多人费解的一个操作符，请看下面这个 import：

```go
import (
   "database/sql"
   _ "github.com/ziutek/mymysql/godrv"
 )
```

\_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的 init 函数。也就是说，使用下划线作为包的别名，会仅仅执行 init()。

> 导入的包的路径名，可以是相对路径也可以是绝对路径，推荐使用绝对路径（起始于工程根目录）。

## 4、GOPATH 环境变量

import 导入时，会从 GO 的安装目录（也就是 GOROOT 环境变量设置的目录）和 GOPATH 环境变量设置的目录中，检索 src/package 来导入包。如果不存在，则导入失败。
GOROOT，就是 GO 内置的包所在的位置。
GOPATH，就是我们自己定义的包的位置。

通常我们在开发 Go 项目时，调试或者编译构建时，需要设置 GOPATH 指向我们的项目目录，目录中的 src 目录中的包就可以被导入了。

## 5、init() 包初始化

下面我们详细的来介绍一下这两个函数：init()、main() 是 go 语言中的保留函数。我们可以在源码中，定义 init() 函数。此函数会在包被导入时执行，例如如果是在 main 中导入包，包中存在 init()，那么 init() 中的代码会在 main() 函数执行前执行，用于初始化包所需要的特定资料。例如：
包源码：

```go
src/userPackage/tool.go

package userPackage
import "fmt"
func init() {
  fmt.Println("tool init")
}
```

主函数源码：

```go
src/main.go

package main
import (
  "userPackage"
  )
func main() {
  fmt.Println("main run")
  // 使用userPackage
  userPackage.SomeFunc()
}
```

执行时，会先输出 "tool init"，再输出 "main run"。

下面我们详细的来介绍一下 init()、main() 这两个函数。在 go 语言中的区别如下：
相同点：

两个函数在定义时不能有任何的参数和返回值。
该函数只能由 go 程序自动调用，不可以被引用。

不同点：

init 可以应用于任意包中，且可以重复定义多个。
main 函数只能用于 main 包中，且只能定义一个。

两个函数的执行顺序：

在 main 包中的 go 文件默认总是会被执行。

对同一个 go 文件的 init( ) 调用顺序是从上到下的。

对同一个 package 中的不同文件，将文件名按字符串进行“从小到大”排序，之后顺序调用各文件中的 init()函数。

对于不同的 package，如果不相互依赖的话，按照 main 包中 import 的顺序调用其包中的 init() 函数。

如果 package 存在依赖，调用顺序为最后被依赖的最先被初始化，例如：导入顺序 main –> A –> B –> C，则初始化顺序为 C –> B –> A –> main，一次执行对应的 init 方法。main 包总是被最后一个初始化，因为它总是依赖别的包

避免出现循环 import，例如：A –> B –> C –> A。

一个包被其它多个包 import，但只能被初始化一次

## 6、管理外部包

go 允许 import 不同代码库的代码。对于 import 要导入的外部的包，可以使用 go get 命令取下来放到 GOPATH 对应的目录中去。

举个例子，比如说我们想通过 go 语言连接 mysql 数据库，那么需要先下载 mysql 的数据包，打开终端并输入以下命令：

```shell
localhost:~ ruby$ go get github.com/go-sql-driver/mysql
```

安装之后，就可以在 gopath 目录的 src 下，看到对应的文件包目录

> 也就是说，对于 go 语言来讲，其实并不关心你的代码是内部还是外部的，总之都在 GOPATH 里，任何 import 包的路径都是从 GOPATH 开始的；唯一的区别，就是内部依赖的包是开发者自己写的，外部依赖的包是 go get 下来的。

## 扩展

我们可以通过 go install 来编译包文件。

我们知道一个非 main 包在编译后会生成一个.a 文件（在临时目录下生成，除非使用 go install 安装到`$GOROOT`或 ​`$GOPATH`下，否则你看不到.a），用于后续可执行程序链接使用。

比如 Go 标准库中的包对应的源码部分路径在：`$GOROOT/src`，而标准库中包编译后的.a 文件路径在`$GOROOT/pkg/darwin_amd64`下。
