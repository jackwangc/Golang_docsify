# golang 基础

## golang 的特性

* 自动的垃圾回收
  * 因为没有 c 语言强大的指针功能
* 更丰富的内置类型, 因此不用去导包，代码简洁
  * `map`
  * `slice`
* 函数多返回值
  * 多个返回值，没有明确赋值，将返回默认的空值
  * `_,lastname := getName()` 返回需要的返回值
* 匿名函数
  * `f := func() int {return 1}`
  * go 语言中函数也是值类型，可以作为参数传递
* 接口和类型可以直接转换
  * 降低接口调整带来的代码调整工作 
  * 接口和类型没有明确的对应关系
* 并发编程
  * goroutine 
* 反射

## 变量

1. 初始化
   1. `var v1 int = 10` 
   2. `var v1 = 10`
   3. `v1 := 10` 声明的同时进行初始化，v1 类型自动判断 
2. 常量，用 `const` 修饰
   1. `ture`, `false`, `iota`
   2. `iota` 可被编译器修改，每出现一次加1，当 `const` 关键字出现重置为 0  
3. go 语言的字符串是一种基本类型
   1. 在初始化后不能被修改
   2. `x + y` 拼接字符串；`len(s)` 获取长度；`s[i]` 取字符 

## make new

1. 二者都是内存的分配（堆上），但是 `make` 只用于 `slice`、`map` 以及 `channel` 的初始化（非零值）；而 `new` 用于类型的内存分配，并且内存置为零。

2. `make` 返回的还是这三个引用类型本身；而 `new` 返回的是指向类型的指针。
3. `make`函数是无可替代的，我们在使用slice、map以及channel的时候，还是要使用make进行初始化
4. `new` 可以用短语句声明替代
## slice

### array

* golang 数组的长度是类型的一部分，因此数组长度不同，不是同一个类型
* 数组的传递是值传递，更改数组并不能更改原有的值

### slice

```go
type slice struct { 
    array unsafe.Pointer // 元素指针,数组地址
    len   int // 长度 
    cap   int // 容量
}
```

1. 创建数组切片
   1. 基于数组 `slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}`, `s1 := slice[2:5]`
   2. 直接创建 `slice := make( []int,5,10)` 创建一个初始个数为 5，容量为 10 的切片，容量可以不写，初始值为 0
   3. 切片添加元素 `slice[1] = 10`;`slice.append(slice, 3)`

2. 参数传递
   1. golang 语言的参数传递，只有值传递，没有引用传递 
   2. `slice` 作为参数传递，会修改 array 的值。

3. 扩容
   1. `apend(slice, a1, a2)` 函数添加元素超出容量会引起扩容，可以追加多个值
   2. `append` 必须有返回值
   3. 扩容会在内存中新建一个匿名数组，并复制数据到该数组
4. 扩容机制
   1. 当原 slice 容量小于 1024 的时候，新 slice 容量变成原来的 2 倍；原 slice 容量超过 1024，新 slice 容量变成原来的1.25倍。
   2. 上面的规则并不完全适用
   3. 内存对齐
      1. CPU把内存当成是一块一块的，块的大小可以是2，4，8，16字节大小，因此CPU在读取内存时是一块一块进行读取的
      2. 为了访问未对齐的内存，处理器需要作两次内存访问；而对齐的内存访问仅需要一次访问。

## map

> 键值对的未排序集合

### 使用

```go
// 1. 声明：变量名，map[键类型] 值类型
var myMap map[string] int
// 2. 初始化, 容量可选
myMap = make(map[string] int,capacity)
// 声明初始化
myMap := make(map[string] int,capacity)
// 3. 赋值
myMap["1"] = 3
// 4. 删除
delete(myMap, "1")
// 5. 查找
value,ok := myMap["1"]
if (ok) {
    fmt.print(value)
}
```

### 原理

1. 内存模型
   1. `hmap` 表示 map 的结构体 , map 的总体情况概述
   2. `buckets` 指针指向 `bmap` , `bmap` = 桶，一个桶有 8 个位置，
   3. `bmap` 中key 和 value 是各自放在一起的，这样可以省略掉 padding 字段，节省内存空间
   4. `bmap` 有第九个进入时，会在构建一个 bucket,通过 overflow 指针连接起来
2. 扩容
   1. go中的扩容和java中有很大的区别
   2. 首先会创建一个新的两倍长度的数组替换掉原来的数组，只有当访问到当前key所在的bucket的时候才会调用growWork方法进行重新hash去迁移原来的元素。
   3. 这样做的优点就是能够在扩容的时候不用因为复制整个数组而阻塞很长的时间，在redis中的map也是使用这样的方式来避免阻塞很长的时间
3. 扩容机制
   1. 装载因子超过阈值，源码里定义的阈值是 6.5。
   2. 情况1，overflow 的 bucket 数量过多，开辟一个新的 `bucket` 空间，将老元素迁移到新的 bucket 里
   3. 情况2，元素太多，bucket 太少，新建新老 bucket ,但是新的 bucket 赋值实在访问到该元素的时候
4. 插入、修改赋值过程
   1. 核心还是一个双层循环，外层遍历 bucket 和它的 overflow bucket，内层遍历整个 bucket 的各个 cell。
5. 无序
   1. 扩容时导致 key 的位置发生变化 
4. 线程不安全
   1. `sync.Map`
   2. 读写锁，分段锁

## golang 面向对象

### struct

```go
// struct
type Base struct {
    Name string
}
// 适用对象声明，会对参数进行拷贝
func(base Base) Bill() {}
// 调用指针声明，当存在一个调用指针时，所有的都必须使用指针
func (base *Base) Foo() {}

func (base *Base) Bar() {}

type Foo struct {
    // 组合
    Base
}

func (foo *Foo) Bar() {
    // 采用组合的形式实现继承
    foo.Base.Bar()
}
// 接口
type Person interface {
    Foo()
}
// 实现接口
var pp = Person(Base{Name: "jack"})
```

## 并发


---

## 参考

1. [码农桃花源](https://qcrao91.gitbook.io/go/shu-zu-he-qie-pian/qie-pian-zuo-wei-han-shu-can-shu)