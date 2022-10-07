# interface

more information please see [go 语言学习-接口](https://www.cnblogs.com/itogo/p/8645486.html#:~:text=go语言学习-接口%201%20接口的定义%20定义方式如下%3A%20type%20Namer%20interface%20{,可以使用下面的方法%3A%20...%204%20标准库中的常用接口%20io.Reader%20和%20io.Writer%20)

## 1. 接口的定义

```go
type Namer interface {
    method1(param_list) return_list
    method2(param_list) return_list
    ...
}
```

这里的 `Namer` 就是接口的名字，只要符合标识符的规则即可。不过，通常约定的接口的名字最好以 `er`,　 `r`, `able` 结尾(视情况而定)，这样一眼就知道它是接口。

## 2. 实现接口

在 `go` 中实现接口很简单，不需要显式的声明实现了某个接口，想要实现某个接口，只需要实现接口中的所有方法即可。

```go {.line-numbers}
package main

import "fmt"
import "math"

type Shaper interface {
    Area() float32
    Circumference() float32
}

type Rect struct {
    Width float32
    Height float32
}

type Circle struct {
    Radius float32
}

func (r Rect) Area() int {
    return r.Width * r.Height
}

func (r Rect) Circumference() int {
    return 2 * (r.Width + r.Height)
}

func (c Circle) Area() int {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() int {
    return math.Pi * 2 * c.Radius
}

func main() {
    r := Rect{10, 20}
    fmt.Printf("Rect w: %f, d: %f, Area: %f, Circumference: %f", r.Width, r.Height, r.Area(), r.Circumference())

    c := Circle{5}
    fmt.Printf("Circle r: %f, Area: %f, Circumference: %f", c.Radius, c.Area(), c.Circumference())
}
```

上面我们定义了一个 `Shaper` 的接口，其中包含两个方法 `Area` 和　 `Circumference`，分别用来计算面积和周长。然后我们定义了两个结构体 `Rect`, `Circle` 并分别实现了这两个方法。但是上面的程序似乎并没有体现出接口和两个实现类型的关系，下面我们将他们关联起来使用:

```go
func showInfo(s Shaper) {
    fmt.Printf("Area: %f, Circumference: %f", s.Area(), s.Circumference())
}
```

注意，这里方法的参数是一个接口类型的，因此我们可以将实现接口的类型的实例传递进去，像下面这样:

```go
r := Rect{10, 20}
showInfo(r)

c := Circle{5}
showInfo(c)
```

## 3. 获取实现接口的实际类型

在上面的 `showInfo` 中我们传入了接口类型的对象，如果将实现了接口的类型传递进去，那么会将实际类型的其他特性掩盖住，因此通常我们会想要获取其真正的类型, 可以使用下面的方法:

```go
func showInfo(s Shaper) {
    switch s.(type) {
    case Rect:
        fmt.Println("This is Rect")
    case Circle:
        fmt.Println("This is Circle")
    }
    fmt.Printf("Area: %f, Circumference: %f\n", s.Area(), s.Circumference())
}
```

另外可以使用类型断言，来判断某一时刻接口是否是某个具体类型

```go
v, ok := s.(Rect)   // s 是一个接口类型
```

## 4. 标准库中的常用接口

### 4.1. io.Reader 和 io.Writer

这两个接口定义了实现 `io` 功能的基本操作，因此某种类型只要实现了这两个接口就可以进行 `io` 操作。

`Reader` 接口的定义为:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

仅仅只有这一个方法，`Read`方法将从数据流中读取 `len(p)` 个字节的数据到字节数组 `p`　中，并且返回读取的字节数(即使发生了错误，`n`也会返回已经读取的字节数)。

我们可能会经常用到的实现了 `Reader` 接口的对象有: `os.Stdin`(标准输入), `os.File`的实例(文件对象)等等，　我们可以对其调用 `Read` 方法来读取数据。

`Writer` 接口的定义:

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

`Write` 将 `len(p)` 个字节的数据从 `p` 中写入到基本数据流中。写入的字节数 `n（0 <= n <= len(p)）`以及任何遇到的引起写入提前停止的错误。

类似的实现了 `Writer` 接口的对象有: `os.Stdout`, `os.Stderr`, `os.File` 等等。可以使用 `Write` 方法向其中写入数据。

### 4.2. sort

排序操作和字符串格式化一样是很多程序经常使用的操作。尽管一个最短的快排程序只要 15 行就可以搞定，但是一个健壮的实现需要更多的代码，并且我们不希望每次我们需要的时候都重写或者拷贝这些代码。

幸运的是，`sort` 包内置的提供了根据一些排序函数来对任何序列排序的功能。它的设计非常独到。在很多语言中，排序算法都是和序列数据类型关联，同时排序函数和具体类型元素关联。

相比之下，`Go`语言的 `sort.Sort` 函数不会对具体的序列和它的元素做任何假设。相反，它使用了一个接口类型 `sort.Interface` 来指定通用的排序算法和可能被排序到的序列类型之间的约定。这个接口的实现由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切片。

一个内置的排序算法需要知道三个东西：序列的长度，表示两个元素比较的结果，一种交换两个元素的方式；这就是 `sort.Interface` 的三个方法：

```go
package sort
type Interface interface {
    Len() int            // 获取元素数量
    Less(i, j int) bool // i，j是序列元素的指数。
    Swap(i, j int)        // 交换元素
}
```

为了对序列进行排序，我们需要定义一个实现了这三个方法的类型，然后对这个类型的一个实例应用 `sort.Sort` 函数。思考对一个字符串切片进行排序，这可能是最简单的例子了。下面是这个新的类型 `MyStringList` 和它的 `Len`，`Less` 和 `Swap` 方法

```go
type MyStringList  []string
func (p MyStringList ) Len() int { return len(m) }
func (p MyStringList ) Less(i, j int) bool { return m[i] < m[j] }
func (p MyStringList ) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
```

### 4.3. 使用 sort.Interface 接口进行排序

对一系列字符串进行排序时，使用字符串切片（`[]string`）承载多个字符串。使用 `type` 关键字，将字符串切片（`[]string`）定义为自定义类型 `MyStringList`。为了让 `sort` 包能识别 `MyStringList`，能够对 `MyStringList` 进行排序，就必须让 `MyStringList` 实现 `sort.Interface` 接口。

下面是对字符串排序的详细代码（代码 1）：

```go {.line-numbers}
package main

import (
    "fmt"
    "sort"
)

// 将[]string定义为MyStringList类型
type MyStringList []string

// 实现sort.Interface接口的获取元素数量方法
func (m MyStringList) Len() int {
    return len(m)
}

// 实现sort.Interface接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
    return m[i] < m[j]
}

// 实现sort.Interface接口的交换元素方法
func (m MyStringList) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}

func main() {

    // 准备一个内容被打乱顺序的字符串切片
    names := MyStringList{
        "3. Triple Kill",
        "5. Penta Kill",
        "2. Double Kill",
        "4. Quadra Kill",
        "1. First Blood",
    }

    // 使用sort包进行排序
    sort.Sort(names)

    // 遍历打印结果
    for _, v := range names {
            fmt.Printf("%s\n", v)
    }

}
```

代码输出结果：

```text
1. First Blood
2. Double Kill
3. Triple Kill
4. Quadra Kill
5. Penta Kill
```

代码说明如下：

- 第 9 行，接口实现不受限于结构体，任何类型都可以实现接口。要排序的字符串切片 `[]string` 是系统定制好的类型，无法让这个类型去实现 `sort.Interface` 排序接口。因此，需要将 `[]string` 定义为自定义的类型。
- 第 12 行，实现获取元素数量的 `Len()` 方法，返回字符串切片的元素数量。
- 第 17 行，实现比较元素的 `Less()` 方法，直接取 `m` 切片的 `i` 和 `j` 元素值进行小于比较，并返回比较结果。
- 第 22 行，实现交换元素的 `Swap()` 方法，这里使用 `Go` 语言的多变量赋值特性实现元素交换。
- 第 29 行，由于将 `[]string` 定义成 `MyStringList` 类型，字符串切片初始化的过程等效于下面的写法：

```go
names := []string {
    "3. Triple Kill",
    "5. Penta Kill",
    "2. Double Kill",
    "4. Quadra Kill",
    "1. First Blood",
}
```

- 第 38 行，使用 `sort` 包的 `Sort()` 函数，将 `names`（`MyStringList` 类型）进行排序。排序时，`sort` 包会通过 `MyStringList` 实现的 `Len()`、`Less()`、`Swap()` 这 3 个方法进行数据获取和修改。
- 第 41 行，遍历排序好的字符串切片，并打印结果。

### 4.4. 常见类型的便捷排序

通过实现 `sort.Interface` 接口的排序过程具有很强的可定制性，可以根据被排序对象比较复杂的特性进行定制。例如，需要多种排序逻辑的需求就适合使用 `sort`.`Interface` 接口进行排序。但大部分情况中，只需要对字符串、整型等进行快速排序。Go 语言中提供了一些固定模式的封装以方便开发者迅速对内容进行排序。

### 4.5. 字符串切片的便捷排序

`sort` 包中有一个 `StringSlice` 类型，定义如下：

```go
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p StringSlice) Sort() { Sort(p) }
```

`sort` 包中的 `StringSlice` 的代码与 `MyStringList` 的实现代码几乎一样。因此，只需要使用 `sort` 包的 `StringSlice` 就可以更简单快速地进行字符串排序。将代码 1 中的排序代码简化后如下所示：

```go
names := sort.StringSlice{
    "3. Triple Kill",
    "5. Penta Kill",
    "2. Double Kill",
    "4. Quadra Kill",
    "1. First Blood",
}

sort.Sort(names)
```

简化后，只要两句代码就实现了字符串排序的功能。

### 4.6. 对整型切片进行排序

除了字符串可以使用 `sort` 包进行便捷排序外，还可以使用 `sort.IntSlice` 进行整型切片的排序。`sort.IntSlice` 的定义如下：除了字符串可以使用 `sort` 包进行便捷排序外，还可以使用 `sort.IntSlice` 进行整型切片的排序。`sort.IntSlice` 的定义如下：

```go
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p IntSlice) Sort() { Sort(p)
```

`sort` 包在 `sort.Interface` 对各类型的封装上还有更进一步的简化，下面使用 `sort.Strings` 继续对代码 1 进行简化，代码如下：

### 4.7. 对结构体数据进行排序

除了基本类型的排序，也可以对结构体进行排序。结构体比基本类型更为复杂，排序时不能像数值和字符串一样拥有一些固定的单一原则。结构体的多个字段在排序中可能会存在多种排序的规则，例如，结构体中的名字按字母升序排列，数值按从小到大的顺序排序。一般在多种规则同时存在时，需要确定规则的优先度，如先按名字排序，再按年龄排序等。

#### 4.7.1. 完整实现 sort.Interface 进行结构体排序

将一批英雄名单使用结构体定义，英雄名单的结构体中定义了英雄的名字和分类。排序时要求按照英雄的分类进行排序，相同分类的情况下按名字进行排序，详细代码实现过程如下。

结构体排序代码（代码 2）：

```go {.line-numbers}
package main

import (
    "fmt"
    "sort"
)

// 声明英雄的分类
type HeroKind int

// 定义HeroKind常量, 类似于枚举
const (
    None HeroKind = iota
    Tank
    Assassin
    Mage
)

// 定义英雄名单的结构
type Hero struct {
    Name string  // 英雄的名字
    Kind HeroKind  // 英雄的种类
}

// 将英雄指针的切片定义为Heros类型
type Heros []*Hero

// 实现sort.Interface接口取元素数量方法
func (s Heros) Len() int {
    return len(s)
}

// 实现sort.Interface接口比较元素方法
func (s Heros) Less(i, j int) bool {

    // 如果英雄的分类不一致时, 优先对分类进行排序
    if s[i].Kind != s[j].Kind {
        return s[i].Kind < s[j].Kind
    }

    // 默认按英雄名字字符升序排列
    return s[i].Name < s[j].Name
}

// 实现sort.Interface接口交换元素方法
func (s Heros) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func main() {

    // 准备英雄列表
    heros := Heros{
        &Hero{"吕布", Tank},
        &Hero{"李白", Assassin},
        &Hero{"妲己", Mage},
        &Hero{"貂蝉", Assassin},
        &Hero{"关羽", Tank},
        &Hero{"诸葛亮", Mage},
    }

    // 使用sort包进行排序
    sort.Sort(heros)

    // 遍历英雄列表打印排序结果
    for _, v := range heros {
        fmt.Printf("%+v\n", v)
    }
}
```

代码输出如下：

```text
&{Name:关羽 Kind:1}
&{Name:吕布 Kind:1}
&{Name:李白 Kind:2}
&{Name:貂蝉 Kind:2}
&{Name:妲己 Kind:3}
&{Name:诸葛亮 Kind:3}
```

代码说明如下：

- 第 9 行，将 `int` 声明为 `HeroKind` 英雄类型，后面会将这个类型当做枚举来使用。
- 第 13 行，定义一些英雄类型常量，可以理解为枚举的值。
- 第 26 行，为了方便实现 `sort.Interface` 接口，将 `[]*Hero` 定义为 `Heros` 类型。
- 第 29 行，`Heros` 类型实现了 `sort.Interface` 的 `Len()` 方法，返回英雄的数量。
- 第 34 行， `Heros` 类型实现了 `sort.Interface` 的 `Less()` 方法，根据英雄字段的比较结果决定如何排序。
- 第 37 行，当英雄的分类不一致时，优先按分类的枚举数值从小到大排序。
- 第 42 行，英雄分类相等的情况下，默认根据英雄的名字字符升序排序。
- 第 46 `行，Heros` 类型实现了 `sort.Interface` 的 `Swap()` 方法，交换英雄元素的位置。
- 第 53 ～ 60 行，准备一系列英雄数据。
- 第 63 行，使用 `sort` 包进行排序。
- 第 66 行，遍历所有排序完成的英雄数据。

#### 4.7.2. 使用 sort.Slice 进行切片元素排序

从 `Go 1.8` 开始，`Go` 语言在 `sort` 包中提供了 `sort.Slice()` 函数进行更为简便的排序方法。`sort.Slice()` 函数只要求传入需要排序的数据，以及一个排序时对元素的回调函数，类型为 `func(i,j int)bool`，`sort.Slice()` 函数的定义如下：

```go
func Slice(slice interface{}, less func(i, j int) bool)
```

使用 `sort.Slice()` 函数，对代码 2 重新优化的完整代码如下：

```go {.line-numbers}
package main

import (
    "fmt"
    "sort"
)

type HeroKind int

const (
    None = iota
    Tank
    Assassin
    Mage
)

type Hero struct {
    Name string
    Kind HeroKind
}

func main() {

    heros := []*Hero{
        {"吕布", Tank},
        {"李白", Assassin},
        {"妲己", Mage},
        {"貂蝉", Assassin},
        {"关羽", Tank},
        {"诸葛亮", Mage},
    }

    sort.Slice(heros, func(i, j int) bool {
        if heros[i].Kind != heros[j].Kind {
            return heros[i].Kind < heros[j].Kind
        }

        return heros[i].Name < heros[j].Name
    })

    for _, v := range heros {
        fmt.Printf("%+v\n", v)
    }
}
```

第 33 行到第 39 行加粗部分是新添加的 sort.Slice() 及回调函数部分。对比前面的代码，这里去掉了 Heros 及接口实现部分的代码。

使用 `sort.Slice()` 不仅可以完成结构体切片排序，还可以对各种切片类型进行自定义排序。
