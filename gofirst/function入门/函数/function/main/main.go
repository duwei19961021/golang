package main

//import (
//	"fmt"
//	"utils"
//)
////init函数
////func init()  {
////	fmt.Println("init first")
////}
////func main() {
////	fmt.Println("main first")
////}
//func main() {
//	age := utils.Age
//	fmt.Println(age)
//}

//匿名函数
//import "fmt"
//func main() {
//	//在定义匿名函数时就直接调用，只能调用一次
//	res1 := func (n1 int,n2 int) int{
//		return n1 + n2
//	}(10,20)
//	fmt.Println(res1)
//}

//func main()  { //将函数赋值给一个变量，能多次调用
//	a := func (n1 int, n2 int) int {
//		return n1 + n2
//	}
//	res2 := a(10,20)
//	fmt.Println(res2)
//	res3 := a(10,30)
//	fmt.Println(res3)
//}

//var Fun1  = func(n1 int,n2 int) int{ //全局匿名函数
//	return n1 + n2
//}
//func main() {
//	res := Fun1(100,200)
//	fmt.Println(res)
//}

//闭包

//defer,在函数执行完毕后及时释放资源
