package main

/*func main()  {
	fmt.Println("hello word","duwei")
}
go build helloworld.go 生成helloworld.exe二进制文件
常用转义字符：\t,\n,\\,\",\r回车*/

/*func main()  {
	fmt.Println("hello\tworld")
	fmt.Println("hello\nworld")
	fmt.Println("hello\\world")
	fmt.Println("hello\"world")
	fmt.Println("helloddddd\rworld")
}*/

/*func main()  {
	fmt.Println("hello world")
}*/

/*
var (
	a1 = 300
	a2 = 600
	a3 = "duwei"
)
var  ii int  = 55
func main()  {
	//使用变量的三种方式：
	//1、指定变量类型，声明后若不赋值，则使用默认值，int默认值为0
	//定义变量
	var i int
	//i = 20
	//使用变量
	fmt.Println(i)
	//2、根据值自行判断变量类型(类型推导)
	var num  = 10.11
	fmt.Println(num)
	//3、省略var，注意:=左侧的变量不应该是已经声明过的，否则会编译错误
	number := 100
	name := "tom"
	fmt.Println(number,name)
	//4、多变量声明
	var n1, n2, n3 int
	var n4, name1 ,n5  = 100, "duwei", 200
	//n4, name1 ,n5  := 100, "duwei", 200
	fmt.Println(n4, name1, n5)
	fmt.Println(n1,n2,n3)
	//5、一次声明多个全局变量，函数外部声明。
	fmt.Println(a1,a2,a3)
	fmt.Println(ii)
}
*/

//+号的使用
/*func main() {
	var i int = 1
	var j int = 2
	var r int = i + j
	fmt.Println(r) //加法运算

	var str1 = "hello"
	var str2 = "world"
	var hw = str1 + str2
	fmt.Println(hw) //字符串拼接
	fmt.Println(true,false)
}*/

import "fmt"

func main() {
	//var i int8  = 10
	//fmt.Println(i)
	//var n1 = 1000000 //n1 是啥类型
	//fmt.Printf("数据类型为：",n1)
	//fmt.Printf("占用字节数为：",unsafe.Sizeof(n1))
	//fmt.Printf("n1的数据类型为%T，占用字节数为%b",n1,unsafe.Sizeof(n1))
	//编程中遵守保小不保大的原则。

	//浮点型float,默认是float64。
	//var n1 float32  = 2.55
	//fmt.Printf("n1类型为%T",n1)

	var n2 float32 = 3.14159165416516516516516516516516544544
	var n3 float64 = 3.51654651616161
	fmt.Println(n2, n3)
}
