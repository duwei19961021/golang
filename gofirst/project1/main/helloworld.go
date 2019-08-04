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

//var n2 float32 = 3.14159165416516516516516516516516544544
//var n3 float64 = 3.51654651616161
//fmt.Println(n2, n3)

//var m = "duwei"
//fmt.Printf("m数据类型为%T",m)

//var n1 byte = 'a'
//var n2 byte = '0'
//fmt.Println(n1,n2) //输出对应的ascii码
//fmt.Printf("n1=%c n2=%d",n1,n2,"\n")
//var n3 int = '杜'
//fmt.Printf("n3=%c n3对应的码值=%d",n3,n3)
//var n4 int  = 26460
//fmt.Printf("n4=%c",n4)
//
//字符类型是可以进行运算的，相当于一个整数，运算时按照码值运行
//var n5  = 10 + 'a'
//fmt.Println("n5=",n5)

/*import (
	"fmt"
)

func main()  {
	//var b = false
	//fmt.Println(unsafe.Sizeof(b))

	var address string = "湖南" //string,go中的字符串一旦定义之后不可变
	fmt.Println(address)
	var name = `\n` //反引号表示原生输出
	fmt.Println(name)
}
*/

/*基本数据类型
整形：0
浮点型：0
字符串：""
bool：false
*/

/*
基本数据类型转换

//import "fmt"
//func main()  {
//	var i int32  = 100
//	var ii float32 = float32(i)
//	fmt.Printf("i=%v ii=%v",i,ii)
//}


import "fmt"
func main()  {
	var n1 int32  = 12
	var n2 int64
	var n3 int8
	//n2 = n1 +20
	//n3 = n1 + 20	//n1是int32类型，加上20之后整体还是int32类型，int32类型不能赋值给int8类型
	n2 = int64(n1 +20)
	n3 = int8(n1 + 20)
	fmt.Println(n2,n3)
}
*/

/*
基本数据类型和string的转换
import (
	"fmt"
	"strconv"
	_ "strconv"
)

func main()  {
	var num1 int = 99
	//var num2 float64  = 99.999
	//var b bool = true
	//var Byte byte  = 'a'
	var str string //空字符串
	//使用fmt.Sprintf方法
	//str = fmt.Sprintf("%d",num1)
	//fmt.Printf("type %T\n",str)
	//
	//str = fmt.Sprintf("%f",num2)
	//fmt.Printf("type %T\n",str)
	//
	//str = fmt.Sprintf("%t",b)
	//fmt.Printf("type %T\n",str)
	//
	//str = fmt.Sprintf("%s",Byte)
	//fmt.Sprintf("%T",str)

	//使用strconv
	str = strconv.FormatInt(int64(num1),10)
	fmt.Printf("type %T str=%q",str,str)
}
*/

/*
string转基本数据类型
import (
	"fmt"
	"strconv"
)

func main() {
	var s1 string = "true"
	var s2	bool
	s2 , _=strconv.ParseBool(s1)
	fmt.Printf("type %T",s2)
}
*/

/*
指针
import (
	"fmt"
)

func main()  {
	var i int64 = 10
	var i1 int64 = 10
	var i2  int64
	fmt.Println("i地址=",&i,i,"i1地址=",&i1,i1,"i2地址=",&i2,i2)
}

import (
	"fmt"
)

func main() {
	var num int  = 64
	fmt.Println("num地址=",&num)
	var ptr *int = &num
	fmt.Println("ptr地址=",ptr)
}
*/

/*
import "fmt"
func main()  {
	var _num int = 10
	fmt.Println("value:",_num)
}
*/

/*
算数运算符

import (
	"fmt"
)
var num1 int = 100
var num2 int = 40

func main() {
	var sum int = num1 + num2
	fmt.Println("和=",sum)
	var cha int = num1 - num2
	fmt.Println("差=",cha)
	var cheng int = num1 * num2
	fmt.Println("积=",cheng)

	var chu int = num1 / num2  //只保留整数部分
	fmt.Println("商=",chu)
	fmt.Println("保留小数=",100/40.0)

	var quyu int = num1 % num2
	fmt.Println("余数",quyu) // a%b=a-a/b*b

	var i int = 100
	i++ //i = i+1
	fmt.Println(i)
	i-- //i = i-1
	fmt.Println(i)
}
*/

//func main() {
//	var days int = 97
//	var week int = days/7
//	var day int = days%7
//	fmt.Printf("周=%v 天=%v",week,day)
//}
//import "fmt"
//func main() {
//	var a int = 8
//	var b int = 7
//	fmt.Println(a > b)
//}

//import "fmt"
//func main() {
//	var age int = 40
//	if age != 50 {
//		fmt.Println("ok")
//	}
//}

//import "fmt"
//
//func main() {
//	var a int = 10
//	a /= a
//	fmt.Println(a)
//}

//func main() {
//	var num1 int = 11
//	var num2 int = 10
//	var t int
//	t = num1
//	num1 = 10
//	num2 = t
//	fmt.Println(num1,num2)
//}

//import "fmt"
//func main() {
//	var num1 int = 11
//	var num2 int = 10
//	var ptr1 *int = &num1
//	var ptr2 *int = &num2
//	num2 = *ptr1
//	*ptr2 = *ptr1
//	fmt.Println(num1,num2)
//}

//import "fmt"
//
//func main() {
//	var a int = 10
//	var b int = 9
//	if a > b {
//		a = a
//	}	else {
//		a = b
//	}
//	fmt.Println(a)
//}

//键盘输入
//import "fmt"
//func main() {
//	var name string
//	var age byte
//	var isPass bool
//	//fmt.Println("请输入姓名：")
//	//fmt.Scanln(&name)
//	//fmt.Println(name)
//	fmt.Println("input:")
//	fmt.Scanf("%s %d %t",&name,&age,&isPass)
//	fmt.Println(name,age,isPass)
//}

//import "fmt"
//func main() {
//	var a int = 1 << 2
//	fmt.Println(a)
//}
