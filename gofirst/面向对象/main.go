package main

import "fmt"

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

type Visitor struct {
	Name string
	Age  int
}

func (vi *Visitor) Echo() {
	fmt.Println("input your name:")
	fmt.Scanln(&vi.Name)
	fmt.Println("input your age:")
	fmt.Scanln(&vi.Age)
	if (*vi).Age > 18 {
		fmt.Println("price=20")
	} else {
		fmt.Println("price=0")
	}
}

func main() {
	//var Echo Calculator
	//var n1 float64 = 90
	//var n2 float64  = 10
	//var operator uint8 = '/'
	//res := Echo.cal(float64(90),float64(10),uint8('/'))
	//fmt.Println(res)
	var tor Visitor
	tor.Age = 19
	tor.Echo()
}
