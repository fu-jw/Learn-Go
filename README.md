# Learn-Go

## Go 简介

Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。

Go 是从 2007 年由 Google 开发，并最终于 2009 年 11 月开源，在 2012 年早些时候发布了 Go 稳定版本，它是基于编译、垃圾收集和并发的编程语言。

Go 语言专门针对多处理器系统应用程序的编程进行了优化，使用 Go 编译的程序可以媲美 C / C++代码的速度，而且更加安全、支持并行进程。
作为出现在 21 世纪的语言，其近 C 的执行性能和近解析型语言的开发效率，以及近乎于完美的编译速度，已经风靡全球。

Golang 的哲学理念：“Less is more or less is less”。

- 学习曲线容易
- 效率：快速的编译时间，开发效率和运行效率高
- "出身名门、血统纯正"
- 自由高效：组合的思想、无侵入式的接口
- 强大的标准库
- 部署方便：二进制文件，Copy 部署
- 并行和异步编程几乎无痛点
  目前几个比较流行的领域，Go 都有用武之地。

- 云计算基础设施领域
  代表项目：docker、kubernetes、etcd、consul、cloudflare CDN、七牛云存储等。

- 基础软件
  代表项目 tidb、influxdb、cockroachdb 等。

- 微服务
  代表项目：go-kit、micro、monzo bank 的 typhon、bilibili 等。

- 互联网基础设施
  代表项目：以太坊、hyperledger 等。

- 分布式系统、数据库代理器、中间件等，例如 Etcd
  DevOps - Go / Python / Shell / Ruby

作为一名 Go 语言开发者，主要的就业领域包括：

- Golang 开发工程师 / Golang 研发工程师
- Golang 服务器后台开发 / 游戏服务器开发
- 云计算平台(golang)开发工程师
- 区块链开发(golang)工程师
- Golang 架构师

## 核心特性

Go 语言之所以厉害，是因为它在服务端的开发中，总能抓住程序员的痛点，以最直接、简单、高效、稳定的方式来解决问题。这里并不会深入讨论 GO 语言的具体语法，只介绍语言中关键的、对简化编程具有重要意义的方面。

### 并发编程

Go 语言在并发编程方面比绝大多数语言要简洁不少，这一点是其最大亮点之一，也是其在未来进入高并发高性能场景的重要筹码。

不同于传统的多进程或多线程，golang 的并发执行单元是一种称为 goroutine 的协程。

由于在共享数据场景中会用到锁，再加上 GC，其并发性能有时不如异步复用 IO 模型，因此相对于大多数语言来说，golang 的并发编程简单比并发性能更具卖点。

在当今这个多核时代，并发编程的意义不言而喻。当然，很多语言都支持多线程、多进程编程，但遗憾的是，实现和控制起来并不是那么令人感觉轻松和愉悦。Golang 不同的是，语言级别支持协程(goroutine)并发（协程又称微线程，比线程更轻量、开销更小，性能更高），操作起来非常简单，语言级别提供关键字（go）用于启动协程，并且在同一台机器上可以启动成千上万个协程。协程经常被理解为轻量级线程，一个线程可以包含多个协程，共享堆不共享栈。协程间一般由应用程序显式实现调度，上下文切换无需下到内核层，高效不少。协程间一般不做同步通讯，而 golang 中实现协程间通讯有两种：1）共享内存型，即使用全局变量+mutex 锁来实现数据共享；2）消息传递型，即使用一种独有的 channel 机制进行异步通讯。

对比 JAVA 的多线程和 GO 的协程实现，明显更直接、简单。这就是 GO 的魅力所在，以简单、高效的方式解决问题，关键字 go，或许就是 GO 语言最重要的标志。

高并发是 Golang 语言最大的亮点

### 内存回收(GC)

从 C 到 C++，从程序性能的角度来考虑，这两种语言允许程序员自己管理内存，包括内存的申请和释放等。因为没有垃圾回收机制所以 C/C++运行起来速度很快，但是随着而来的是程序员对内存使用上的很谨小慎微的考虑。因为哪怕一点不小心就可能会导致“内存泄露”使得资源浪费或者“野指针”使得程序崩溃等，尽管 C++11 后来使用了智能指针的概念，但是程序员仍然需要很小心的使用。后来为了提高程序开发的速度以及程序的健壮性，java 和 C#等高级语言引入了 GC 机制，即程序员不需要再考虑内存的回收等，而是由语言特性提供垃圾回收器来回收内存。但是随之而来的可能是程序运行效率的降低。

GC 过程是：先 stop the world，扫描所有对象判活，把可回收对象在一段 bitmap 区中标记下来，接着立即 start the world，恢复服务，同时起一个专门 gorountine 回收内存到空闲 list 中以备复用，不物理释放。物理释放由专门线程定期来执行。

GC 瓶颈在于每次都要扫描所有对象来判活，待收集的对象数目越多，速度越慢。一个经验值是扫描 10w 个对象需要花费 1ms，所以尽量使用对象少的方案，比如我们同时考虑链表、map、slice、数组来进行存储，链表和 map 每个元素都是一个对象，而 slice 或数组是一个对象，因此 slice 或数组有利于 GC。

GC 性能可能随着版本不断更新会不断优化，这块没仔细调研，团队中有 HotSpot 开发者，应该会借鉴 jvm gc 的设计思想，比如分代回收、safepoint 等。

- 内存自动回收，再也不需要开发人员管理内存
- 开发人员专注业务实现，降低了心智负担
- 只需要 new 分配内存，不需要释放

### 内存分配

初始化阶段直接分配一块大内存区域，大内存被切分成各个大小等级的块，放入不同的空闲 list 中，对象分配空间时从空闲 list 中取出大小合适的内存块。内存回收时，会把不用的内存重放回空闲 list。空闲内存会按照一定策略合并，以减少碎片。

### 编译

编译涉及到两个问题：编译速度和依赖管理

目前 Golang 具有两种编译器，一种是建立在 GCC 基础上的 Gccgo，另外一种是分别针对 64 位 x64 和 32 位 x86 计算机的一套编译器(6g 和 8g)。

依赖管理方面，由于 golang 绝大多数第三方开源库都在 github 上，在代码的 import 中加上对应的 github 路径就可以使用了，库会默认下载到工程的 pkg 目录下。

另外，编译时会默认检查代码中所有实体的使用情况，凡是没使用到的 package 或变量，都会编译不通过。这是 golang 挺严谨的一面。

### 网络编程

由于 golang 诞生在互联网时代，因此它天生具备了去中心化、分布式等特性，具体表现之一就是提供了丰富便捷的网络编程接口，比如 socket 用 net.Dial(基于 tcp/udp，封装了传统的 connect、listen、accept 等接口)、http 用 http.Get/Post()、rpc 用 client.Call('class_name.method_name', args, &reply)，等等。

> 高性能 HTTP Server

### 函数多返回值

在 C，C++ 中，包括其他的一些高级语言是不支持多个函数返回值的。但是这项功能又确实是需要的，所以在 C 语言中一般通过将返回值定义成一个结构体，或者通过函数的参数引用的形式进行返回。而在 Go 语言中，作为一种新型的语言，目标定位为强大的语言当然不能放弃对这一需求的满足，所以支持函数多返回值是必须的。

函数定义时可以在入参后面再加(a,b,c)，表示将有 3 个返回值 a、b、c。这个特性在很多语言都有，比如 python。

这个语法糖特性是有现实意义的，比如我们经常会要求接口返回一个三元组（errno,errmsg,data），在大多数只允许一个返回值的语言中，我们只能将三元组放入一个 map 或数组中返回，接收方还要写代码来检查返回值中包含了三元组，如果允许多返回值，则直接在函数定义层面上就做了强制，使代码更简洁安全。

### 语言交互性

语言交互性指的是本语言是否能和其他语言交互，比如可以调用其他语言编译的库。

在 Go 语言中直接重用了大部份的 C 模块，这里称为 Cgo.Cgo 允许开发者混合编写 C 语言代码，然后 Cgo 工具可以将这些混合的 C 代码提取并生成对于 C 功能的调用包装代码。开发者基本上可以完全忽略这个 Go 语言和 C 语言的边界是如何跨越的。

golang 可以和 C 程序交互，但不能和 C++交互。可以有两种替代方案：1）先将 c++编译成动态库，再由 go 调用一段 c 代码，c 代码通过 dlfcn 库动态调用动态库（记得 export LD_LIBRARY_PATH）；2）使用 swig(没玩过)

### 异常处理

golang 不支持 try...catch 这样的结构化的异常解决方式，因为觉得会增加代码量，且会被滥用，不管多小的异常都抛出。golang 提倡的异常处理方式是：

- 普通异常：被调用方返回 error 对象，调用方判断 error 对象。
- 严重异常：指的是中断性 panic（比如除 0），使用 defer...recover...panic 机制来捕获处理。严重异常一般由 golang 内部自动抛出，不需要用户主动抛出，避免传统 try...catch 写得到处都是的情况。当然，用户也可以使用 panic('xxxx')主动抛出，只是这样就使这一套机制退化成结构化异常机制了。

### 其他一些有趣的特性

- 类型推导：类型定义：支持 var abc = 10 这样的语法，让 golang 看上去有点像动态类型语言，但 golang 实际上时强类型的，前面的定义会被自动推导出是 int 类型。

> 作为强类型语言，隐式的类型转换是不被允许的，记住一条原则：让所有的东西都是显式的。

> 简单来说，Go 是一门写起来像动态语言，有着动态语言开发效率的静态语言。

- 一个类型只要实现了某个 interface 的所有方法，即可实现该 interface，无需显式去继承。

> Go 编程规范推荐每个 Interface 只提供一到两个的方法。这样使得每个接口的目的非常清晰。另外 Go 的隐式推导也使得我们组织程序架构的时候更加灵活。在写 JAVA／C++程序的时候，我们一开始就需要把父类／子类／接口设计好，因为一旦后面有变更，修改起来会非常痛苦。而 Go 不一样，当你在实现的过程中发现某些方法可以抽象成接口的时候，你直接定义好这个接口就 OK 了，其他代码不需要做任何修改，编译器的自动推导会帮你做好一切。

- 不能循环引用：即如果 a.go 中 import 了 b，则 b.go 要是 import a 会报 import cycle not allowed。好处是可以避免一些潜在的编程危险，比如 a 中的 func1()调用了 b 中的 func2()，如果 func2()也能调用 func1()，将会导致无限循环调用下去。

- defer 机制：在 Go 语言中，提供关键字 defer，可以通过该关键字指定需要延迟执行的逻辑体，即在函数体 return 前或出现 panic 时执行。这种机制非常适合善后逻辑处理，比如可以尽早避免可能出现的资源泄漏问题。

> 可以说，defer 是继 goroutine 和 channel 之后的另一个非常重要、实用的语言特性，对 defer 的引入，在很大程度上可以简化编程，并且在语言描述上显得更为自然，极大的增强了代码的可读性。

- “包”的概念：和 python 一样，把相同功能的代码放到一个目录，称之为包。包可以被其他包引用。main 包是用来生成可执行文件，每个程序只有一个 main 包。包的主要用途是提高代码的可复用性。通过 package 可以引入其他包。

- 编程规范：GO 语言的编程规范强制集成在语言中，比如明确规定花括号摆放位置，强制要求一行一句，不允许导入没有使用的包，不允许定义没有使用的变量，提供 gofmt 工具强制格式化代码等等。奇怪的是，这些也引起了很多程序员的不满，有人发表 GO 语言的 XX 条罪状，里面就不乏对编程规范的指责。要知道，从工程管理的角度，任何一个开发团队都会对特定语言制定特定的编程规范，特别像 Google 这样的公司，更是如此。GO 的设计者们认为，与其将规范写在文档里，还不如强制集成在语言里，这样更直接，更有利用团队协作和工程管理。

- 交叉编译：比如说你可以在运行 Linux 系统的计算机上开发运行 Windows 下运行的应用程序。这是第一门完全支持 UTF-8 的编程语言，这不仅体现在它可以处理使用 UTF-8 编码的字符串，就连它的源码文件格式都是使用的 UTF-8 编码。Go 语言做到了真正的国际化！

## 一个小故事

很久以前，有一个 IT 公司，这公司有个传统，允许员工拥有 20%自由时间来开发实验性项目。在 2007 的某一天，公司的几个大牛，正在用 c++ 开发一些比较繁琐但是核心的工作，主要包括庞大的分布式集群，大牛觉得很闹心，后来 c++ 委员会来他们公司演讲，说 c++ 将要添加大概 35 种新特性。这几个大牛的其中一个人，名为：Rob Pike，听后心中一万个 xxx 飘过，“c++ 特性还不够多吗？简化 c++ 应该更有成就感吧”。于是乎，Rob Pike 和其他几个大牛讨论了一下，怎么解决这个问题，过了一会，Rob Pike 说要不我们自己搞个语言吧，名字叫“go”，非常简短，容易拼写。其他几位大牛就说好啊，然后他们找了块白板，在上面写下希望能有哪些功能。接下来的时间里，大牛们开心的讨论设计这门语言的特性，经过漫长的岁月，他们决定，以 c 语言为原型，以及借鉴其他语言的一些特性，来解放程序员，解放自己，然后在 2009 年，go 语言诞生。

以下就是这些大牛所罗列出的 Go 要有的功能：

- 规范的语法（不需要符号表来解析）
- 垃圾回收（独有）
- 无头文件
- 明确的依赖
- 无循环依赖
- 常量只能是数字
- int 和 int32 是两种类型
- 字母大小写设置可见性（letter case sets visibility）
- 任何类型（type）都有方法（不是类型）
- 没有子类型继承（不是子类）
- 包级别初始化以及明确的初始化顺序
- 文件被编译到一个包里
- 包 package-level globals presented in any order
- 没有数值类型转换（常量起辅助作用）
- 接口隐式实现（没有“implement”声明）
- 嵌入（不会提升到超类）
- 方法按照函数声明（没有特别的位置要求）
- 方法即函数
- 接口只有方法（没有数据）
- 方法通过名字匹配（而非类型）
- 没有构造函数和析构函数
- postincrement（如++i）是状态，不是表达式
- 没有 preincrement(i++)和 predecrement
- 赋值不是表达式
- 明确赋值和函数调用中的计算顺序（没有“sequence point”）
- 没有指针运算
- 内存一直以零值初始化
- 局部变量取值合法
- 方法中没有“this”
- 分段的堆栈
- 没有静态和其它类型的注释
- 没有模板
- 内建 string、slice 和 map
- 数组边界检查

## Go 和其他语言比较

Go 的很多语言特性借鉴与它的三个祖先：C，Pascal 和 CSP。Go 的语法、数据类型、控制流等继承于 C，Go 的包、面对对象等思想来源于 Pascal 分支，而 Go 最大的语言特色，基于管道通信的协程并发模型，则借鉴于 CSP 分支。

### Java

编译语言，速度适中，目前的大型网站都是拿 java 写的，比如淘宝、京东等。主要特点是稳定，开源性好，具有自己的一套编写规范，开发效率适中，目前最主流的语言。

> 作为编程语言中的大腕。具有最大的知名度和用户群。无论风起云涌，我自巍然不动。他强任他强，清风拂山岗；他横由他横，明月照大江。

### C#

执行速度快（4.28），学习难度适中，开发速度适中。但是由于 c#存在很多缺点，京东、携程等大型网站前身都是用 c#开发的，但是现在都迁移到了 java 上。

### C/C++

现存编程语言中的老祖，其他语言皆由此而生。执行速度最快无人能及。但是写起来最为复杂，开发难度大。

### Javascript

编程语言中特立独行的傲娇美女。前端处理能力是其它语言无法比拟。发展中的 js 后端处理能力也是卓越不凡。前后端通吃，舍我其谁？

### Python

脚本语言，速度最慢（258s），代码简洁、学习进度短，开发速度快。豆瓣就是拿 python 写的。Python 著名的服务器框架有 django，flask。但是 python 在大型项目上不太稳定，因此有些用 python 的企业后来迁移到了 java 上。

### Scala

编译语言，比 python 快十倍，和 java 差不多，但是学习进度慢，而且在实际编程中，如果对语言不够精通，很容易造成性能严重下降。，后来比如 Yammer 就从 scala 迁移到了 java 上。微服务框架有 lagom 等。

### Go

编程界的小鲜肉。高并发能力无人能及。即具有像 Python 一样的简洁代码、开发速度，又具有 C 语言一样的执行效率，优势突出。

## Go 源码文件

![GO源码文件分类](https://image.fu-jw.com/img/2024/08/19/66c29479dc97a.png)

### 命令源码文件

声明自己属于 main 代码包、包含无参数声明和结果声明的 main 函数。

命令源码文件被安装以后，GOPATH 如果只有一个工作区，那么相应的可执行文件会被存放当前工作区的 bin 文件夹下；如果有多个工作区，就会安装到 GOBIN 指向的目录下。

命令源码文件是 Go 程序的入口。

同一个代码包中最好也不要放多个命令源码文件。多个命令源码文件虽然可以分开单独 go run 运行起来，但是无法通过 go build 和 go install。

```go
package main
import "fmt"
func main(){
	fmt.Print("Hello World !!!")
}
```

> 注意：文件夹中放了两个命令源码文件，同时都声明自己属于 main 代码包时
>
> - 两个文件都可执行命令：`go run`
> - 不可以使用命令 `go build`和`go install`
> - 所以命令源码文件应该是被单独放在一个代码包中

### 库源码文件

库源码文件就是不具备命令源码文件上述两个特征的源码文件。存在于某个代码包中的普通的源码文件。

库源码文件被安装后，相应的归档文件（.a 文件）会被存放到当前工作区的 pkg 的平台相关目录下。

### 测试源码文件

名称以 \_test.go 为后缀的代码文件，并且必须包含 Test 或者 Benchmark 名称前缀的函数：

```go
func TestXXX( t *testing.T) {

}
```

名称以 Test 为名称前缀的函数，只能接受 \*testing.T 的参数，这种测试函数是功能测试函数。

```go
func BenchmarkXXX( b *testing.B) {

}
```

名称以 Benchmark 为名称前缀的函数，只能接受 \*testing.B 的参数，这种测试函数是性能测试函数。

## Go 的命令

目前 Go 的最新版 1.23 里面基本命令有以下 19 个

```txt
The commands are:

	bug         start a bug report
	build       compile packages and dependencies
	clean       remove object files and cached files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         update packages to use new APIs
	fmt         gofmt (reformat) package sources
	generate    generate Go files by processing source
	get         add dependencies to current module and install them
	install     compile and install packages and dependencies
	list        list packages or modules
	mod         module maintenance
	work        workspace maintenance
	run         compile and run Go program
	telemetry   manage telemetry data and settings
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         report likely mistakes in packages
```

### go run

专门用来运行命令源码文件的命令，注意，这个命令不是用来运行所有 Go 的源码文件的！

go run 命令只能接受一个命令源码文件以及若干个库源码文件（必须同属于 main 包）作为文件参数，且不能接受测试源码文件。它在执行时会检查源码文件的类型。如果参数中有多个或者没有命令源码文件，那么 go run 命令就只会打印错误提示信息并退出，而不会继续执行。

> 使用 `go run -n test.go` 可以查看执行过程
>
> - 先执行了 compile 命令
> - 然后 link，生成了归档文件.a 和 最终可执行文件
> - 最终的可执行文件放在 exe 文件夹里面
> - 命令的最后一步就是执行了可执行文件

![go run 执行过程](https://image.fu-jw.com/img/2024/08/19/66c29db642183.png)

### go build

主要是用于测试编译。在包的编译过程中，若有必要，会同时编译与之相关联的包。

1. 如果是普通包，当你执行 go build 命令后，不会产生任何文件。
2. 如果是 main 包，当只执行 go build 命令后，会在当前目录下生成一个可执行文件。如果需要在$GOPATH/bin 目录下生成相应的 exe 文件，需要执行 go install 或者使用 go build -o 路径/可执行文件。
3. 如果某个文件夹下有多个文件，而你只想编译其中某一个文件，可以在 go build 之后加上文件名，例如 go build a.go；go build 命令默认会编译当前目录下的所有 go 文件。
4. 你也可以指定编译输出的文件名。比如，我们可以指定 go build -o 可执行文件名，默认情况是你的 package 名(非 main 包)，或者是第一个源文件的文件名(main 包)。
5. go build 会忽略目录下以”\_”或者”.”开头的 go 文件。
6. 如果你的源代码针对不同的操作系统需要不同的处理，那么你可以根据不同的操作系统后缀来命名文件。

当代码包中有且仅有一个命令源码文件的时候，在文件夹所在目录中执行 go build 命令，会在该目录下生成一个与**目录同名的可执行文件**，

- 因此不可在统一目录下存在两个命令源码文件！
- 此时执行命令`go install`就会将可执行文件移动到 bin 目录下

go build 和 go install

- go build 编译命令源码文件，则会在该命令的执行目录中生成一个可执行文件
- go build 后面不追加目录路径的话，它就把当前目录作为代码包并进行编译。go build 命令后面如果跟了代码包导入路径作为参数，那么该代码包及其依赖都会被编译。

> 注意如果 go build 用来编译非命令源码文件，即库源码文件，go build 执行完是不会产生任何结果的。
> 这种情况下，go build 命令只是检查库源码文件的有效性，只会做检查性的编译，而不会输出任何结果文件。

![go build 执行过程](https://image.fu-jw.com/img/2024/08/19/66c2a24140199.png)

### go install

用来编译并安装代码包或者源码文件

go install 命令在内部实际上分成了两步操作：第一步是生成结果文件(可执行文件或者.a 包)，第二步会把编译好的结果移到$GOPATH/pkg或者​$GOPATH/bin。

- 可执行文件： 一般是 go install 带 main 函数的 go 文件产生的，有函数入口，所有可以直接运行。
- .a 应用包： 一般是 go install 不包含 main 函数的 go 文件产生的，没有函数入口，只能被调用。

go install 用于编译并安装指定的代码包及它们的依赖包。当指定的代码包的依赖包还没有被编译和安装时，该命令会先去处理依赖包。与 go build 命令一样，传给 go install 命令的代码包参数应该以导入路径的形式提供。并且，go build 命令的绝大多数标记也都可以用于 实际上，go install 命令只比 go build 命令多做了一件事，即：安装编译后的结果文件到指定目录。

安装代码包会在当前工作区的 pkg 的平台相关目录下生成归档文件（即 .a 文件）。 安装命令源码文件会在当前工作区的 bin 目录（如果 GOPATH 下有多个工作区，就会放在 GOBIN 目录下）生成可执行文件。

- go install 命令如果后面不追加任何参数，它会把当前目录作为代码包并安装。这和 go build 命令是完全一样的。
- go install 命令后面如果跟了代码包导入路径作为参数，那么该代码包及其依赖都会被安装。
- go install 命令后面如果跟了命令源码文件以及相关库源码文件作为参数的话，只有这些文件会被编译并安装。

![go install 执行过程](https://image.fu-jw.com/img/2024/08/19/66c2a4de808d0.png)

### go get

用于从远程代码仓库（比如 Github ）上下载并安装代码包。

> 注意，go get 命令会把当前的代码包下载到 $GOPATH 中的第一个工作区的 src 目录中，并安装。

> 智能下载: 在使用它检出或更新代码包之后，它会寻找与本地已安装 Go 语言的版本号相对应的标签（tag）或分支（branch）。比如，本机安装 Go 语言的版本是 1.x，那么 go get 命令会在该代码包的远程仓库中寻找名为 “go1” 的标签或者分支。如果找到指定的标签或者分支，则将本地代码包的版本切换到此标签或者分支。如果没有找到指定的标签或者分支，则将本地代码包的版本切换到主干的最新版本。

![go get 执行过程](https://image.fu-jw.com/img/2024/08/19/66c2a6da898e1.png)

### 其他命令

- go clean 命令是用来移除当前源码包里面编译生成的文件
- go fmt 命令主要是用来帮你格式化所写好的代码文件
- go test 命令，会自动读取源码目录下面名为\*\_test.go 的文件，生成并运行测试用的可执行文件
- go doc 命令其实就是一个很强大的文档工具
- go fix 用来修复以前老版本的代码到新版本，例如 go1 之前老版本的代码转化到 go1
- go version 查看 go 当前的版本
- go env 查看当前 go 的环境变量
- go list 列出当前全部安装的 package
