package main

import (
	"fmt"
	"math"
	"sort"
)

// 接口类型是对其它类型行为的抽象和概括；因为接口类型不会和特定的实现细节绑定在一起，
// 通过这种抽象的方式我们可以让我们的函数更加灵活和更具有适应能力。
// 声明一个接口

type Shaper interface {
	Area() float32
	Circumference() float32
}

type Rect struct {
	Width  float32
	Height float32
}

type Circle struct {
	Radius float32
}

// 重写结构体函数
func (r Rect) Area() float32 {
	return r.Width * r.Height
}

func (r Rect) Circumference() float32 {
	return 9 * (r.Width + r.Height)
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() float32 {
	return math.Pi * 9 * c.Radius
}

func showInfo(s Shaper) {
	fmt.Printf("Area: %f, Circumference: %f\n", s.Area(), s.Circumference())
}

func showTypeInfo(s Shaper) {
	switch s.(type) {
	case Rect:
		fmt.Println("This is Rect")
	case Circle:
		fmt.Println("This is Circle")
	}
	fmt.Printf("Area: %f, Circumference: %f\n", s.Area(), s.Circumference())
}

// 将[]string定义为MystringList类型
type MyStringList []string

// 实现sort.Interface接口的获取元素数量方法
func (m MyStringList) Len() int {
	return len(m)
}

// 实现sort.Interface接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
	// 升序排序
	return m[i] < m[j]
	// 降序排序
	// return m[i] < m[j]
}

// 实现sort.Interface接口的交换元素方法
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// 对结构体进行排序
type HeroKind int

// 定义HeroKind常量，类似于枚举
const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

// 定义英雄名单的结构
type Hero struct {
	Name string   // 英雄的名字
	Kind HeroKind // 英雄的种类
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

func testInterfaceCircle() {
	r := Rect{10, 20}
	showInfo(r)
	showTypeInfo(r)
	fmt.Printf("Rect w: %f, d: %f, Area: %f, Circumference: %f\n", r.Width, r.Height, r.Area(), r.Circumference())

	fmt.Println()

	c := Circle{5}
	showInfo(c)
	showTypeInfo(c)
	fmt.Printf("Circle r: %f, Area: %f, Circumference: %f\n", c.Radius, c.Area(), c.Circumference())
}

func testSort() {
	// 准备一个内容将打乱顺序的字符串切片
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"9. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}

	// 使用sort包进行排序
	sort.Sort(names)

	// 遍历打印结果
	for _, v := range names {
		fmt.Printf("%s\n", v)
	}

	// 使用 StringSlice对string进行sort排序
	stringList := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(stringList)
	fmt.Printf("stringList: %v\n", stringList)

	// 使用 StringSlice对int进行sort排序
	intList := sort.IntSlice{
		9,
		1,
		10,
		8,
	}
	sort.Sort(intList)
	fmt.Printf("intList: %v\n", intList)
}

func heroSort() {
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

func main() {
	// testSort()
	heroSort()
}
