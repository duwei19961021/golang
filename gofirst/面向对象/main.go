package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//
//import "fmt"
//
////结构体创建
//type Person struct {
//	Name string
//	Age int
//}
//
//func main() {
//	p2 := Person{}
//	p2.Name = "duwei2"
//	p2.Age = 23
//	fmt.Println(p2.Age,p2.Name)
//
//	p3 := Person{"duwei3",24}
//	fmt.Println(p3)
//	var p4 *Person = new(Person)
//	(*p4).Name = "duwei4"
//	(*p4).Age = 25
//	fmt.Println(&p4)
//
//	var person *Person = &Person{} //也可以在此直接赋值
//	(*person).Name = "duwei5"
//	(*person).Age = 26
//	(*person).Age = 27
//	fmt.Println(*person)
//}

//方法

//type Person struct {
//	Name string
//}
//
//func (p Person) test() {
//	fmt.Println("test()", p.Name)
//}
//
//func (p Person) speak() {
//	fmt.Println("sepak")
//}
//
//func (p Person) jisuan(num1 int, num2 int) {
//	sum := 0
//	for i := num1; i <= num2; i++ {
//		sum += i
//	}
//	fmt.Println(sum)
//}
//
//func (p Person) getsum(a int, b int) int {
//	sum := a + b
//	return sum
//}
//
//func main() {
//	var p Person
//	p.Name = "duwei"
//	p.test()
//	p.speak()
//	p.jisuan(1, 1651)
//	res := p.getsum(10, 100)
//	fmt.Println(res)
//}

//type Methoutils struct {
//	length int
//	wide int
//}
//
//func (echo Methoutils) Echo() {
//	for i := 1;i <= 10;i++{
//		for j := 1;j <= 8;j++{
//			fmt.Print("*")
//		}
//		fmt.Println()
//	}
//}
//
//func (echo Methoutils) Echo1(m int,n int)  {
//	for i := 1;i <= m;i++{
//		for j := 1;j <= n;j++{
//			fmt.Print("*")
//		}
//		fmt.Println()
//	}
//}

//type Judge struct {
//	num int
//}
//
//func (echo Judge) jd(num1 int) bool  {
//	if num1%2 == 0{
//		return true
//	}else {
//		return false
//	}
//}

//type Calculator struct {
//	num1 float64
//	num2 float64
//}
//
//func (echo Calculator) cal(num1 float64,num2 float64,Operator uint8) (result float64)  {
//	switch Operator {
//	case '+':
//		result = num1 + num2
//	case '-':
//		result = num1 - num2
//	case '*':
//		result = num1 * num2
//	case '/':
//		result = num1 / num2
//	default:
//		fmt.Println("input error")
//	}
//	return result
//}

//type Visitor struct {
//	Name string
//	Age  int
//}

//func (vi *Visitor) Echo() {
//	fmt.Println("input your name:")
//	fmt.Scanln(&vi.Name)
//	fmt.Println("input your age:")
//	fmt.Scanln(&vi.Age)
//	if (*vi).Age > 18 {
//		fmt.Println("price=20")
//	} else {
//		fmt.Println("price=0")
//	}
//}

//func main() {
//	//var Echo Calculator
//	//var n1 float64 = 90
//	//var n2 float64  = 10
//	//var operator uint8 = '/'
//	//res := Echo.cal(float64(90),float64(10),uint8('/'))
//	//fmt.Println(res)
//	var tor Visitor
//	tor.Age = 19
//	tor.Echo()
//}

//方法、函数
//函数，函数是指针类型时只能传递地址，是值类型时只能传递值
/*
type Stu struct {
	Name string
	age int
	hoby string
}

func test01(p Stu)  {
	fmt.Println(p.Name)
}
func test02(p *Stu)  {
	fmt.Println(p.Name)
}
func (p Stu) test03() {
	fmt.Println(p.Name)
}
func (p *Stu) test04() {
	(*p).Name = "dw2"
}

func main() {
	var p Stu
	p.Name = "d2"
	test01(p)
	test02(&p)
	p.test03()
	(&p).test03()
	(&p).test04()
	fmt.Println(p.Name)
}
*/

//
//func main() {
//	p := model.NewPerson("duwei")
//	p.SetAge(23)
//	p.SetSal(500)
//	fmt.Println(p.GetSal())
//	fmt.Println(p.GetAge())
//	fmt.Println(p.Name)
//
//	//p := model.Newaccount("duwei")
//	//p.SetAccount("duweia")
//	//p.Getbalance()
//	//fmt.Println(p.User)
//	//fmt.Println(p.Getbalance())
//}

//type Xiaoxuesheng struct {
//	Name string
//	Score float64
//	Age int
//}
//
//func (p *Xiaoxuesheng) Showinfo() {
//	fmt.Println(p.Name,p.Age,p.Score)
//}
//func (p *Xiaoxuesheng) Setscore(score float64)  {
//	p.Score = score
//}
//func (p *Xiaoxuesheng) testing() {
//	fmt.Println("xiaoxuesheng testing...")
//}
//
//type Daxuesheng struct {
//	Name string
//	Score float64
//	Age int
//}
//
//func (p *Daxuesheng) Showinfo() {
//	fmt.Println(p.Name,p.Age,p.Score)
//}
//func (p *Daxuesheng) Setscore(score float64)  {
//	p.Score = score
//}
//func (p *Daxuesheng) testing() {
//	fmt.Println("Daxuesheng testing...")
//}
//
//func main() {
//	p := Xiaoxuesheng{
//		Name:  "duwei",
//		Score: 99.5,
//		Age:   10,
//	}
//	p.Showinfo()
//	p.Setscore(98)
//	p.testing()
//
//	q := Daxuesheng{
//		Name:  "duwei2",
//		Score: 60,
//		Age:   23,
//	}
//	q.Setscore(61)
//	q.Showinfo()
//	q.testing()
//}

//继承
//语法
//type Goods struct {
//	Name string
//	Price float64
//}
//type Book struct {
//	Goods
//	Author string
//}
/*
type Student struct {
	Name string
	Age int
	Score float64
}

func (stu *Student) ShowInfo() {
	fmt.Println(stu.Name,stu.Age,stu.Score)
}
func (stu *Student) Setscore(score float64) {
	stu.Score = score
}

type Xiaoxuesheng struct {
	Student
}

func (p *Xiaoxuesheng) testing() { //此结构体特有方法
	fmt.Println("xiaoxuesheng testing...")
}

type Daxuesheng struct {
	Student
}

func (p *Daxuesheng) testing() {
	fmt.Println("Daxuesheng testing...")
}

func main()  {
	st := Xiaoxuesheng{}
	st.Student.Name = "yulushen"
	st.Student.Age = 9
	st.Student.Setscore(59)
	st.testing()
	st.Student.ShowInfo()

	dst := Daxuesheng{}
	dst.Student.Name = "duwei"
	dst.Student.Age = 23
	dst.Student.Setscore(100)
	dst.testing()
	dst.Student.ShowInfo()
}
*/
//结构体可以使用匿名结构体中的所有字段和方法

//接口
/*
type Usb interface {
	Start()
	Stop( )
}

type Phone struct {

}
func (p Phone) Start() {
	fmt.Println("Phone start...")
}
func (p Phone) Stop() {
	fmt.Println("Phone stop...")
}

type Camera struct {

}
func (c Camera) Start() {
	fmt.Println("Camera start...")
}
func (c Camera) Stop() {
	fmt.Println("Camera stop...")
}

type Computer struct {

}

func (com Computer) working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	com := Computer{}
	pho := Phone{}
	cam := Camera{}
	com.working(pho)
	com.working(cam)
}
*/

//切片排序(sort 接口)
type Student struct {
	Name  string
	Age   int
	Score int
}

type stuslice []Student

func (ss stuslice) Len() int {
	return len(ss)
}
func (ss stuslice) Less(i, j int) bool {
	//return ss[i].Age < ss[j].Age
	return ss[i].Score < ss[j].Score
}
func (ss stuslice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func main() {
	var stu1 stuslice
	for i := 0; i < 10; i++ {
		sn := Student{
			Name:  fmt.Sprintf("duwei%d", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: rand.Intn(100),
		}
		stu1 = append(stu1, sn)
	}
	for _, v := range stu1 {
		fmt.Println(v)
	}
	fmt.Println("-----------------------------------------------")
	sort.Sort(stu1)
	for _, v := range stu1 {
		fmt.Println(v)
	}
}
