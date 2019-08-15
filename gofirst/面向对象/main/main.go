package main

import (
	"fmt"
)

//import "fmt"

////结构体创建
//type Person struct {
//	Name string
//	Age int
//}

//func main() {
//	p2 := Person{}
//	p2.Name = "duwei2"
//	p2.Age = 23
//	fmt.Println(p2.Age,p2.Name)

//	p3 := Person{"duwei3",24}
//	fmt.Println(p3)
//	var p4 *Person = new(Person)
//	(*p4).Name = "duwei4"
//	(*p4).Age = 25
//	fmt.Println(p4)

//	var person *Person = &Person{} //也可以在此直接赋值
//	(*person).Name = "duwei5"
//	(*person).Age = 26
//	(*person).Age = 27
//	fmt.Println(*person)
//}

//方法

type Person struct {
	Name string
}

func (p Person) test() {
	fmt.Println("test()", p.Name)
}

func (p Person) speak() {
	fmt.Println("sepak")
}

func (p Person) jisuan(num1 int, num2 int) {
	sum := 0
	for i := num1; i <= num2; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func (p Person) getsum(a int, b int) int {
	sum := a + b
	return sum
}

func main() {
	var p Person
	p.Name = "duwei"
	p.test()
	p.speak()
	p.jisuan(1, 1651)
	res := p.getsum(10, 100)
	fmt.Println(res)
}
