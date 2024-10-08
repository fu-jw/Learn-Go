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

### C

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

## 命名规范

命名是代码规范中很重要的一部分，统一的命名规则有利于提高的代码的可读性，好的命名仅仅通过命名就可以获取到足够多的信息。

Go 在命名时以字母 a 到 Z 或 a 到 Z 或下划线开头，后面跟着零或更多的字母、下划线和数字(0 到 9)。Go 不允许在命名时中使用@、$和%等标点符号。Go 是一种区分大小写的编程语言。因此，Manpower 和 manpower 是两个不同的命名。

- 当命名（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；
- 命名如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 private ）

### 包命名：package

保持 package 的名字和目录保持一致，尽量采取有意义的包名，简短，有意义，尽量和标准库不要冲突。包名应该为小写单词，不要使用下划线或者混合大小写。

```go
package demo

package main
```

### 文件命名

尽量采取有意义的文件名，简短，有意义，应该为小写单词，使用下划线分隔各个单词。

```go
my_test.go
```

### 结构体命名

- 采用驼峰命名法，首字母根据访问控制大写或者小写

- struct 申明和初始化格式采用多行，例如下面：

```go
// 多行声明
type User struct{
    Username  string
    Email     string
}

// 多行初始化
u := User{
    Username: "fredo",
    Email:    "coderfredo@gmail.com",
}

```

### 接口命名

- 命名规则基本和上面的结构体类型
- 单个函数的结构名以 “er” 作为后缀，例如 Reader , Writer 。

```go
type Reader interface {
        Read(p []byte) (n int, err error)
}

```

### 变量命名

- 和结构体类似，变量名称一般遵循驼峰法，首字母根据访问控制原则大写或者小写，但遇到特有名词时，需要遵循以下规则：
  - 如果变量为私有，且特有名词为首个单词，则使用小写，如 apiClient
  - 其它情况都应当使用该名词原有的写法，如 APIClient、repoID、UserID
  - 错误示例：UrlArray，应该写成 urlArray 或者 URLArray
- 若变量类型为 bool 类型，则名称应以 Has, Is, Can 或 Allow 开头

```go
var isExist bool
var hasConflict bool
var canManage bool
var allowGitHook bool
```

### 常量命名

常量均需使用全部大写字母组成，并使用下划线分词

```go
const APP_VER = "1.0"
```

如果是枚举类型的常量，需要先创建相应类型：

```go
type Scheme string

const (
    HTTP  Scheme = "http"
    HTTPS Scheme = "https"
)
```

## 关键字

关键字是一些特殊的用来帮助编译器理解和解析源代码的单词。

Go 中共有 25 个关键字：

```txt
break     default      func    interface  select
case      defer        go      map        struct
chan      else         goto    package    switch
const     fallthrough  if      range      type
continue  for          import  return     var
```

这些关键字可以分为四组：

- const、func、import、package、type 和 var 用来声明各种代码元素。
- chan、interface、map 和 struct 用做 一些组合类型的字面表示中。
- break、case、continue、default、 else、fallthrough、for、 goto、if、range、 return、- - select 和 switch 用在流程控制语句中。 详见基本流程控制语法。
- defer 和 go 也可以看作是流程控制关键字， 但它们有一些特殊的作用。详见协程和延迟函数调用。

## 变量

变量是为存储特定类型的值而提供给内存位置的名称。在 go 中声明变量有多种语法。

本质就是一小块内存，用于存储数据，在程序运行过程中数值可以改变

声明变量使用关键字 var

- 第一种，指定变量类型，声明后若不赋值，使用默认值

```go
var name type
name = value
```

- 第二种，根据值自行判定变量类型(类型推断 Type inference)
  如果一个变量有一个初始值，Go 将自动能够使用初始值来推断该变量的类型。因此，如果变量具有初始值，则可以省略变量声明中的类型

```go
var name = value
```

- 第三种，简短声明，省略 var, 注意 :=左侧的变量不应该是已经声明过的(多个变量同时声明时，至少保证一个是新变量)，否则会导致编译错误

```go
name := value

// 例如
sum := 100
fmt.Println(sum)
```

也可以一行多个变量

```go
// 同一类型
var m, n int = 111, 222

// 类型推断
var n1, f1, s1 = 100, 3.14, "Go"

// 一组变量
var (
stuName = "小美"
age     = 18
sex     = "女"
)
```

注意：

1. 变量必须先定义才能使用
2. 变量的类型和赋值必须一致
3. 同一个作用域内，变量名不能冲突
4. 简短定义方式，左边的变量至少有一个是新的
5. 简短定义方式，不能定义全局变量
6. 变量的零值，就是默认值
   整型：默认值是 0
   浮点类型：默认是 0
   字符串：默认值""

```go
var x int //整数，默认值是0
fmt.Println(x)
var y float64 //0.0-->0
fmt.Println(y)
var str string //""
fmt.Println(str)
var s2 []int //切片[]
fmt.Println(s2)
fmt.Println(s2 == nil)
```

## 常量

在程序运行时，不会被修改的量

```go
const identifier [type] = value

// 显式类型定义：
const a string = "a"

//隐式类型定义：
const b = "b"
```

常量可以作为枚举，常量组

- 如不指定类型和初始化值，则与上一行非空常量右值相同

```go
const (
a int = 100
b
)
fmt.Printf("%T,%d\n", a, a) // int,100
fmt.Printf("%T,%d\n", b, b) // int,100
```

注意:

- 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
- 不曾使用的常量，在编译的时候，是不会报错的
- 显示指定类型的时候，必须确保常量左右值类型一致，需要时可做显示类型转换。这与变量就不一样了，变量是可以是不同的类型值

## iota

特殊常量，可以认为是一个可以被编译器修改的常量

第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 x=0, y=1, z=2 可以简写为如下形式：

```go
const (
    x = iota
    y
    z
)
```

如果中断 iota 自增，则必须显式恢复。且后续自增值按行序递增
自增默认是 int 类型，可以自行进行显示指定类型
数字常量不会分配存储空间，无须像变量那样通过内存寻址来取值，因此无法获取地址

```go
const (
a = iota //0
b        //1
c        //2
d = "hi" //独立值，iota += 1
e        //"hi"   iota += 1
f = 100  //iota +=1
g        //100  iota +=1
h = iota //7,恢复计数
i        //8
)
fmt.Println(a, b, c, d, e, f, g, h, i)
```

## 基本数据类型

数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候才需要申请大内存，就可以充分利用内存。

Go 语言按类别有以下几种数据类型：

| 序号 | 类型描述                                                                                                                                                                         |
| ---- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1    | **布尔型** <br>布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。                                                                                       |
| 2    | **数字类型** <br>整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。                                                           |
| 3    | **字符串类型:** <br>字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。                   |
| 4    | **派生类型:** <br>包括：<br> - 指针类型（Pointer）<br>- 数组类型<br>- 结构化类型 (struct)<br>- Channel 类型<br>- 函数类型<br>- 切片类型<br>- 接口类型（interface）<br>- Map 类型 |

### 数字类型

| 序号 | 类型和描述                                                                   |
| ---- | ---------------------------------------------------------------------------- |
| 1    | **uint8** <br>无符号 8 位整型 (0 到 255)                                     |
| 2    | **uint16** <br>无符号 16 位整型 (0 到 65535)                                 |
| 3    | **uint32** <br>无符号 32 位整型 (0 到 4294967295)                            |
| 4    | **uint64** <br>无符号 64 位整型 (0 到 18446744073709551615)                  |
| 5    | **int8** <br>有符号 8 位整型 (-128 到 127)                                   |
| 6    | **int16** <br>有符号 16 位整型 (-32768 到 32767)                             |
| 7    | **int32** <br>有符号 32 位整型 (-2147483648 到 2147483647)                   |
| 8    | **int64** <br>有符号 64 位整型 (-9223372036854775808 到 9223372036854775807) |

### 浮点型

| 序号 | 类型和描述                             |
| ---- | -------------------------------------- |
| 1    | **float32** <br>IEEE-754 32 位浮点型数 |
| 2    | **float64** <br>IEEE-754 64 位浮点型数 |
| 3    | **complex64** <br>32 位实数和虚数      |
| 4    | **complex128** <br>64 位实数和虚数     |

### 其他数字类型

| 序号 | 类型和描述                                   |
| ---- | -------------------------------------------- |
| 1    | **byte** <br>类似 uint8                      |
| 2    | **rune** <br>类似 int32                      |
| 3    | **uint** <br>32 或 64 位                     |
| 4    | **int** <br>与 uint 一样大小                 |
| 5    | **uintptr** <br>无符号整型，用于存放一个指针 |

## 运算符

表达式：(a + b) \* c

```txt
  a,b,c 叫做操作数
  +，*，叫做运算符
```

### 算术运算符

```go
+ - * / %(求余) ++ --
```

### 关系运算符

```go
== != > < >= <=
```

### 逻辑运算符

| 运算符 | 描述                                                                             |
| ------ | -------------------------------------------------------------------------------- |
| &&     | 所谓逻辑与运算符。如果两个操作数都非零，则条件变为真                             |
| \|\|   | 所谓的逻辑或操作。如果任何两个操作数是非零，则条件变为真                         |
| !      | 所谓逻辑非运算符。使用反转操作数的逻辑状态。如果条件为真，那么逻辑非操后结果为假 |

### 位运算符

| A   | B   | A&B | A\|B | A^B |
| --- | --- | --- | ---- | --- |
| 0   | 0   | 0   | 0    | 0   |
| 0   | 1   | 0   | 1    | 1   |
| 1   | 1   | 1   | 1    | 0   |
| 1   | 0   | 0   | 1    | 1   |

这里最难理解的就是^了，只要认为 AB 两者都相同的时候，为 0，其他都为 1

假设 A 为 60，B 为 13

| 运算 | 描述                                                               | 示例                            |
| ---- | ------------------------------------------------------------------ | ------------------------------- |
| &    | 二进制与操作副本位的结果，如果它存在于两个操作数                   | (A & B) = 12, 也就是 0000 1100  |
| \|   | 二进制或操作副本，如果它存在一个操作数                             | (A \| B) = 61, 也就是 0011 1101 |
| ^    | 二进制异或操作副本，如果它被设置在一个操作数就是按位取非           | (A ^ B) = 49, 也就是 0011 0001  |
| &^   | 二进制位清空&^                                                     | (A&^B)=48，也就是 110000        |
| <<   | 二进制左移位运算符。左边的操作数的值向左移动由右操作数指定的位数   | A << 2 =240 也就是 1111 0000    |
| >>   | 二进制向右移位运算符。左边的操作数的值由右操作数指定的位数向右移动 | A >> 2 = 15 也就是 0000 1111    |

### 赋值运算符

| 运算符 | 描述                                                             | 示例                            |
| ------ | ---------------------------------------------------------------- | ------------------------------- |
| =      | 简单的赋值操作符，分配值从右边的操作数左侧的操作数               | C = A + B 将分配 A + B 的值到 C |
| +=     | 相加并赋值运算符，它增加了右操作数左操作数和分配结果左操作数     | C += A 相当于 C = C + A         |
| -=     | 减和赋值运算符，它减去右操作数从左侧的操作数和分配结果左操作数   | C -= A 相当于 C = C - A         |
| \*=    | 乘法和赋值运算符，它乘以右边的操作数与左操作数和分配结果左操作数 | C _= A 相当于 C = C_ A          |
| /=     | 除法赋值运算符，它把左操作数与右操作数和分配结果左操作数         | C /= A 相当于 C = C / A         |
| %=     | 模量和赋值运算符，它需要使用两个操作数的模量和分配结果左操作数   | C %= A 相当于 C = C % A         |
| <<=    | 左移位并赋值运算符                                               | C <<= 2 相同于 C = C << 2       |
| >>=    | 向右移位并赋值运算符                                             | C >>= 2 相同于 C = C >> 2       |
| &=     | 按位与赋值运算符                                                 | C &= 2 相同于 C = C & 2         |
| ^=     | 按位异或并赋值运算符                                             | C ^= 2 相同于 C = C ^ 2         |
| \|=    | 按位或并赋值运算符                                               | C \|= 2 相同于 C = C \| 2       |

### 优先级运算符优先级

有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：

| 优先级 | 运算符            |
| ------ | ----------------- |
| 7      | ~ ! ++ --         |
| 6      | \* / % << >> & &^ |
| 5      | + - ^             |
| 4      | == != < <= >= >   |
| 3      | <-                |
| 2      | &&                |
| 1      | \|\|              |

当然，你可以通过使用括号来临时提升某个表达式的整体运算优先级。

## IO

主要由`fmt`包实现输入输出，实现了类似 C 语言 printf 和 scanf 的格式化 I/O

官网资料：<https://golang.google.cn/pkg/fmt/>

### 打印输出

**打印：**

[func Print(a ...interface{}) (n int, err error)](https://golang.google.cn/pkg/fmt/#Print)

**格式化打印：**

[func Printf(format string, a ...interface{}) (n int, err error)](https://golang.google.cn/pkg/fmt/#Printf)

**打印后换行**

[func Println(a ...interface{}) (n int, err error)](https://golang.google.cn/pkg/fmt/#Println)

格式化打印中的常用占位符：

```txt
格式化打印占位符：
    %v,原样输出
    %T，打印类型
    %t,bool类型
    %s，字符串
    %f，浮点
    %d，10进制的整数
    %b，2进制的整数
    %o，8进制
    %x，%X，16进制
      %x：0-9，a-f
      %X：0-9，A-F
    %c，打印字符
    %p，打印地址
    。。。
```

示例代码：

```go
func out() {
  a := 100           //int
  b := 3.14          //float64
  c := true          // bool
  d := "Hello World" //string
  e := `Fredo`        //string
  f := 'A'
  fmt.Printf("%T,%b\n", a, a)
  fmt.Printf("%T,%f\n", b, b)
  fmt.Printf("%T,%t\n", c, c)
  fmt.Printf("%T,%s\n", d, d)
  fmt.Printf("%T,%s\n", e, e)
  fmt.Printf("%T,%d,%c\n", f, f, f)
  fmt.Println("-----------------------")
  fmt.Printf("%v\n", a)
  fmt.Printf("%v\n", b)
  fmt.Printf("%v\n", c)
  fmt.Printf("%v\n", d)
  fmt.Printf("%v\n", e)
  fmt.Printf("%v\n", f)
}
```

### 键盘输入

常用方法：

[func Scan(a ...interface{}) (n int, err error)](https://golang.google.cn/pkg/fmt/#Scan)

[func Scanf(format string, a ...interface{}) (n int, err error)](https://golang.google.cn/pkg/fmt/#Scanf)

[func Scanln(a ...interface{}) (n int, err error)](https://golang.google.cn/pkg/fmt/#Scanln)

```go
func in() {
  var x int
  var y float64
  fmt.Println("请输入一个整数，一个浮点类型：")
  fmt.Scanln(&x, &y) //读取键盘的输入，通过操作地址，赋值给x和y   阻塞式
  fmt.Printf("x的数值：%d，y的数值：%f\n", x, y)

  fmt.Scanf("%d,%f", &x, &y)
  fmt.Printf("x:%d,y:%f\n", x, y)
}
```

## 程序的流程结构

程序的流程控制结构一共有三种：顺序结构，选择结构，循环结构。

顺序结构：从上向下，逐行执行。

选择结构：条件满足，某些代码才会执行。0-1 次

- 分支语句：if，switch，select

循环结构：条件满足，某些代码会被反复的执行多次。0-N 次

- ​ 循环语句：for

### 分支语句

#### if 语句

语法格式：

```go
if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
}
```

```go
if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
} else {
  /* 在布尔表达式为 false 时执行 */
}
```

```go
if 布尔表达式1 {
   /* 在布尔表达式1为 true 时执行 */
} else if 布尔表达式2{
   /* 在布尔表达式1为 false ,布尔表达式2为true时执行 */
} else{
   /* 在上面两个布尔表达式都为false时，执行*/
}
```

示例代码：

```go
func branch() {
   /* 定义局部变量 */
   var a int = 10

   /* 使用 if 语句判断布尔表达式 */
   if a < 20 {
       /* 如果条件为 true 则执行以下语句 */
       fmt.Printf("a 小于 20\n" )
   }
   fmt.Printf("a 的值为 : %d\n", a)
}
```

#### if 变体

如果其中包含一个可选的语句组件(在评估条件之前执行)，则还有一个变体。它的语法是

```go
if statement; condition {
}

if condition{


}
```

示例代码：

```go
func branch() {
    if num := 10; num % 2 == 0 { //checks if number is even
        fmt.Println(num,"is even")
    }  else {
        fmt.Println(num,"is odd")
    }
}
```

> 需要注意的是，num 的定义在 if 里，那么只能够在该 if..else 语句块中使用，否则编译器会报错的。

#### switch

switch 是一个条件语句，它计算表达式并将其与可能匹配的列表进行比较，并根据匹配执行代码块。它可以被认为是一种惯用的方式来写多个 if else 子句。

switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。
switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。

而如果 switch 没有表达式，它会匹配 true

Go 里面 switch 默认相当于每个 case 最后带有 break，匹配成功后不会自动向下执行其他 case，而是跳出整个 switch, 但是可以使用 fallthrough 强制执行后面的 case 代码。

变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。
您可以**同时测试多个可能符合条件的值，使用逗号分割它们**，例如：case val1, val2, val3。

```go
switch var1 {
  case val1:
    ...
  case val2:
    ...
  default:
    ...
}
```

示例代码：

```go
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
}
```

#### fallthrough

如需贯通后续的 case，就添加 fallthrough

```go
func main() {
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
}
```

#### case 变体

case 中的表达式是可选的，可以省略。如果该表达式被省略，则被认为是 switch true，并且每个 case 表达式都被计算为 true，并执行相应的代码块。

```go
func switch_branch() {
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
```

> 注意事项:
>
> 1. case 后的常量值不能重复
> 2. case 后可以有多个常量值
> 3. fallthrough 应该是某个 case 的最后一行。如果它出现在中间的某个地方，编译器就会抛出错误。

### 循环语句

循环语句表示条件满足，可以反复的执行某段代码。

for 是唯一的循环语句。(Go 没有 while 循环)

#### for 语句

语法结构：

```go
for init; condition; post { }
```

> 初始化语句只执行一次。在初始化循环之后，将检查该条件。如果条件计算为 true，那么{}中的循环体将被执行，然后是 post 语句。post 语句将在循环的每次成功迭代之后执行。在执行 post 语句之后，该条件将被重新检查。如果它是正确的，循环将继续执行，否则循环终止。

示例代码：

```go
func cyclic() {
  for i := 0; i < 10; i++ {
    fmt.Printf(" %d", i)
  }
}
```

> 在 for 循环中声明的变量仅在循环范围内可用。因此，i 不能在外部访问循环。

#### for 变体

**所有的三个组成部分，即初始化、条件和 post 都是可选的。**

```go
for condition { }
```

效果与 while 相似

```go
for { }
```

效果与 for(;;) 一样

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环

```go
for key, value := range oldMap {
    newMap[key] = value
}
```

### break、continue、goto

#### break

跳出循环体。break 语句用于在结束其正常执行之前突然终止 for 循环

```go
func funcBreak() {
  for i := 1; i <= 10; i++ {
    if i > 5 {
      break //loop is terminated if i > 5
    }
    fmt.Printf("%d ", i)
  }
  fmt.Printf("\nline after for loop")
}
```

#### continue

跳出一次循环。continue 语句用于跳过 for 循环的当前迭代。在 continue 语句后面的 for 循环中的所有代码将不会在当前迭代中执行。循环将继续到下一个迭代。

```go
func funcContinue() {
  for i := 1; i <= 10; i++ {
    if i%2 == 0 {
      continue
    }
    fmt.Printf("%d ", i)
  }
}
```

#### goto

可以无条件地转移到过程中指定的行

```go
func funcGoto() {
  /* 定义局部变量 */
  var a int = 10

/* 循环 */
LOOP:
  for a < 20 {
    if a == 15 {
      /* 跳过迭代 */
      a = a + 1
      goto LOOP
    }
    fmt.Printf("a的值为 : %d\n", a)
    a++
  }
}
```

统一错误处理
多处错误处理存在代码重复时是非常棘手的，例如：

```go
  err := firstCheckError()
  if err != nil {
      goto onExit
  }
  err = secondCheckError()
  if err != nil {
      goto onExit
  }
  fmt.Println("done")
  return
onExit:
    fmt.Println(err)
    exitProcess()
```

## 数组

### 数组概述

Go 语言提供了数组类型的数据结构。 数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。

数组元素可以通过索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。数组的下标取值范围是从 0 开始，到长度减 1。

数组一旦定义后，大小不能更改。

### 数组语法

**声明和初始化数组**：

- 需要指明数组的大小和存储的数据类型。

```go
var variable_name [SIZE] variable_type
```

示例代码：

```go
var balance [10] float32
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```

初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：

```go
var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```

```go
balance[4] = 50.0
```

数组的其他创建方式：

```go
  var a [4] float32 // 等价于：var arr2 = [4]float32{}
  fmt.Println(a) // [0 0 0 0]
  var b = [5] string{"ruby", "王二狗", "rose"}
  fmt.Println(b) // [ruby 王二狗 rose  ]
  var c = [5] int{'A', 'B', 'C', 'D', 'E'} // byte
  fmt.Println(c) // [65 66 67 68 69]
  d := [...] int{1,2,3,4,5}// 根据元素的个数，设置数组的大小
  fmt.Println(d)//[1 2 3 4 5]
  e := [5] int{4: 100} // [0 0 0 0 100]
  fmt.Println(e)
  f := [...] int{0: 1, 4: 1, 9: 1} // [1 0 0 0 1 0 0 0 0 1]
  fmt.Println(f)
```

**访问数组元素**：

```go
float32 salary = balance[9]
```

示例代码：

```go
func funcArray() {
   var n [10]int /* n 是一个长度为 10 的数组 */
   var i,j int

   // 为数组 n 初始化元素
   for i = 0; i < 10; i++ {
      n[i] = i + 100 /* 设置元素为 i + 100 */
   }

   // 输出每个数组元素的值
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }
}
```

**数组的长度**：

- 通过将数组作为参数传递给 len 函数，可以获得数组的长度。
- 也可以让编译器决定

```go
a := [...]float64{67.7, 89.8, 21, 78}
fmt.Println("length of a is",len(a))
```

使用 range 遍历数组：

```go
  arr := [...]float64{67.7, 89.8, 21, 78}
  sum := float64(0)
  for i, v := range arr { //range returns both the index and value
    fmt.Printf("%d the element of a is %.2f\n", i, v)
    sum += v
  }
  fmt.Println("\nsum of all elements of a", sum)
```

如果您只需要值并希望忽略索引，那么可以通过使用\_ blank 标识符替换索引来实现这一点。

```go
for _, v := range a { //ignores index
}
```

### 多维数组

Go 语言支持多维数组，以下为常用的多维数组声明语法方式：

```go
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
```

```go
var threedim [5][10][4]int
```

三维数组

```go
a = [3][4]int{
 {0, 1, 2, 3} ,   /*  第一行索引为 0 */
 {4, 5, 6, 7} ,   /*  第二行索引为 1 */
 {8, 9, 10, 11}   /*  第三行索引为 2 */
}
```

### 数组是值类型

Go 中的数组是值类型，而不是引用类型。这意味着当它们被分配给一个新变量时，将把原始数组的副本分配给新变量。如果对新变量进行了更改，则不会在原始数组中反映。

```go
a := [...]string{"USA", "China", "India", "Germany", "France"}
b := a // a copy of a is assigned to b
b[0] = "Singapore"
fmt.Println("a is ", a)
fmt.Println("b is ", b)

```

数组的大小是类型的一部分。因此[5]int 和[25]int 是不同的类型。因此，数组不能被调整大小。不要担心这个限制，因为切片的存在是为了解决这个问题。

```go
a := [3]int{5, 78, 8}
var b [5]int
b = a //not possible since [3]int and [5]int are distinct types
```

## 切片

### 切片概述

Go 语言切片是对数组的抽象。
Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大

切片是一种方便、灵活且强大的包装器。切片本身没有任何数据。它们只是对现有数组的引用。

切片与数组相比，不需要设定长度，在[]中不用设定值，相对来说比较自由

从概念上面来说 slice 像一个结构体，这个结构体包含了三个元素：

1. 指针，指向数组中 slice 指定的开始位置
2. 长度，即 slice 的长度
3. 最大长度，也就是 slice 开始位置到数组的最后位置的长度

### 切片语法

**定义切片**：

```go
var identifier []type
```

切片不需要说明长度。
或使用 make()函数来创建切片:

```go
var slice1 []type = make([]type, len)
// 也可以简写为
slice1 := make([]type, len)
```

```go
make([]T, length, capacity)
```

**初始化**：

```go
s[0] = 1
s[1] = 2
s[2] = 3
```

```go
s :=[] int {1,2,3 }
```

```go
s := arr[startIndex:endIndex]
```

将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片（**前闭后开**），长度为 endIndex-startIndex

```go
s := arr[startIndex:]
```

缺省 endIndex 时将表示一直到 arr 的最后一个元素

```go
s := arr[:endIndex]
```

缺省 startIndex 时将表示从 arr 的第一个元素开始

```go
func main() {
  a := [5]int{76, 77, 78, 79, 80}
  var slice_b []int = slice_a[1:4] //creates b slice from a[1] to a[3]
  fmt.Println(slice_b)
}
```

### 修改切片

slice 没有自己的任何数据。它只是底层数组的一个表示。对 slice 所做的任何修改都将反映在底层数组中。

示例代码：

```go
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
```

当多个片共享相同的底层数组时，每个元素所做的更改将在数组中反映出来。

```go
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
```

### len 和 cap

切片的长度是切片中元素的数量。
切片的容量是从创建切片的索引开始的底层数组中元素的数量。

切片是可索引的，并且可以由 len() 方法获取长度
切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少

```go
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
```

### append 和 copy

append 向 slice 里面追加一个或者多个元素，然后返回一个和 slice 一样类型的 slice
copy 函数 copy 从源 slice 的 src 中复制元素到目标 dst，并且返回复制的元素的个数

append 函数会改变 slice 所引用的数组的内容，从而影响到引用同一数组的其它 slice。 但当 slice 中没有剩
余空间（即(cap-len) == 0）时，此时将动态分配新的数组空间（2 倍扩容）。返回的 slice 数组指针将指向这个空间，而原
数组的内容将保持不变；其它引用此数组的 slice 则不受影响

下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法

```go
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
```

## map

### map 概述

map 是 Go 中的内置类型，它将一个值与一个键关联起来。可以使用相应的键检索值。

Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值
Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的，也是引用类型

使用 map 过程中需要注意的几点：

- map 是无序的，每次打印出来的 map 都会不一样，它不能通过 index 获取，而必须通过 key 获取
- map 的长度是不固定的，也就是和 slice 一样，也是一种引用类型
- 内置的 len 函数同样适用于 map，返回 map 拥有的 key 的数量
- map 的 key 可以是所有可比较的类型，如布尔型、整数型、浮点型、复杂型、字符串型……也可以键。

### map 使用

可以使用内建函数 make 也可以使用 map 关键字来定义 Map:

```go
/* 声明变量，默认 map 是 nil */
var map_variable map[key_data_type]value_data_type

/* 使用 make 函数 */
map_variable = make(map[key_data_type]value_data_type)
```

```go
rating := map[string]float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2 }
```

如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对

### 常用操作

- delete

delete(map, key) 函数用于删除集合的元素, 参数为 map 和其对应的 key。删除函数不返回任何值。

```go
/* 删除元素 */
delete(countryCapitalMap,"France");
fmt.Println("Entry for France is deleted")

fmt.Println("删除元素后 map")

/* 打印 map */
for country := range countryCapitalMap {
  fmt.Println("Capital of",country,"is",countryCapitalMap[country])
}
```

- ok-idiom
  我们可以通过 key 获取 map 中对应的 value 值。语法为：

```go
map[key]
```

但是当 key 如果不存在的时候，我们会得到该 value 值类型的默认值，比如 string 类型得到空字符串，int 类型得到 0。但是程序不会报错。

所以我们可以使用 ok-idiom 获取值，可知道 key/value 是否存在

```go
value, ok := map[key]
```

```go
/* 查看元素在集合中是否存在 */
capital, ok := countryCapitalMap["United States"]
/* 如果 ok 是 true, 则存在，否则不存在 */
if ok {
  fmt.Println("Capital of United States is", capital)
} else {
  fmt.Println("Capital of United States is not present")
}
```

- len
  使用 len 函数可以确定 map 的长度。

```go
fmt.Println("Length of map is:", len(countryCapitalMap))
```

### map 是引用类型的

与切片相似，映射是引用类型。当将映射分配给一个新变量时，它们都指向相同的内部数据结构。因此，一个的变化会反映另一个。

```go
var newMap map[string]string = countryCapitalMap
newMap["United Kingdom"] = "London"
fmt.Println("New map after adding UK:", newMap)
fmt.Println("Original map:", countryCapitalMap)
/* 注意：对 newMap 的操作并不影响 countryCapitalMap */
fmt.Println("Length of new map is:", len(newMap))
fmt.Println("Is newMap the same as countryCapitalMap?", &newMap == &countryCapitalMap)
/* 注意：newMap 和 countryCapitalMap 并不是同一个 map */
fmt.Println("Country capital of United Kingdom is:", newMap["United Kingdom"])
fmt.Println("Country capital of United Kingdom is:", countryCapitalMap["United Kingdom"])
```

> 注意：
> map 不能使用==操作符进行比较。
> ==只能用来检查 map 是否为空。否则会报错：
> invalid operation: map1 == map2 (map can only be compared to nil)

## string

### string 概述

Go 中的字符串是一个字节的切片。可以通过将其内容封装在“”中来创建字符串。Go 中的字符串是 Unicode 兼容的，并且是 UTF-8 编码的。

```go
name := "Hello World"
fmt.Println(name)
```

### string 使用

- 访问字符串中的单个字节

```go
func funcString() {
  name := "Hello World"
  for i := 0; i < len(name); i++ {
    // 打印每个字符的字节码
    fmt.Printf("十六进制：%x ", name[i])
    fmt.Printf("十进制：%d ", name[i])
  }
  fmt.Printf("\n")
  for i := 0; i < len(name); i++ {
    // 打印每个字符
    fmt.Printf("%c ", name[i])
  }
}
```

### strings 包

访问 strings 包，可以有很多操作 string 的函数。

```go
/*
  strings包下的关于字符串的函数

*/
s1 := "helloworld"
//1.是否包含指定的内容-->bool
fmt.Println(strings.Contains(s1, "abc"))
//2.是否包含chars中任意的一个字符即可
fmt.Println(strings.ContainsAny(s1, "abcd"))
//3.统计substr在s中出现的次数
fmt.Println(strings.Count(s1, "lloo"))

//4.以xxx前缀开头，以xxx后缀结尾
s2 := "20240828.txt"
if strings.HasPrefix(s2, "202408") {
  fmt.Println("24年8月的文件。。")
}
if strings.HasSuffix(s2, ".txt") {
  fmt.Println("文本文档。。")
}

//索引
//helloworld
fmt.Println(strings.Index(s1, "lloo"))     //查找substr在s中的位置，如果不存在就返回-1
fmt.Println(strings.IndexAny(s1, "abcdh")) //查找chars中任意的一个字符，出现在s中的位置
fmt.Println(strings.LastIndex(s1, "l"))    //查找substr在s中最后一次出现的位置

//字符串的拼接
ss1 := []string{"abc", "world", "hello"}
s3 := strings.Join(ss1, "-")
fmt.Println(s3)

//切割
s4 := "123,4563,aaa,49595,45"
ss2 := strings.Split(s4, ",")
//fmt.Println(ss2)
for i := 0; i < len(ss2); i++ {
  fmt.Println(ss2[i])
}

//重复，自己拼接自己count次
s5 := strings.Repeat("hello", 5)
fmt.Println(s5)

//替换
//helloworld
s6 := strings.Replace(s1, "l", "*", -1)
fmt.Println(s6)
//fmt.Println(strings.Repeat("hello",5))

s7 := "heLLo WOrlD**123.."
fmt.Println(strings.ToLower(s7))
fmt.Println(strings.ToUpper(s7))

/*
  截取子串：
  substring(start,end)-->substr
  str[start:end]-->substr
    包含start，不包含end下标
*/
fmt.Println(s1)
s8 := s1[:5]
fmt.Println(s8)
fmt.Println(s1[5:])
```

### strconv 包

访问 strconv 包，可以实现 string 和其他数值类型之间的转换。

```go
/*
  strconv包：字符串和基本类型之前的转换
    string convert
*/
//fmt.Println("aa"+100)
//1.bool类型
s1 := "true"
b1, err := strconv.ParseBool(s1)
if err != nil {
  fmt.Println(err)
  return
}
fmt.Printf("%T,%t\n", b1, b1)

ss1 := strconv.FormatBool(b1)
fmt.Printf("%T,%s\n", ss1, ss1)

//2.整数
s2 := "100"
i2, err := strconv.ParseInt(s2, 2, 64)
if err != nil {
  fmt.Println(err)
  return
}
fmt.Printf("%T,%d\n", i2, i2)

ss2 := strconv.FormatInt(i2, 10)
fmt.Printf("%T,%s\n", ss2, ss2)

//itoa(),atoi()
i3, err := strconv.Atoi("-42") //转为int类型
fmt.Printf("%T,%d\n", i3, i3)
ss3 := strconv.Itoa(-42)
fmt.Printf("%T,%s\n", ss3, ss3)
```
