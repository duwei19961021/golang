package main

//import (
//	"cal"
//	"fmt"
//)
//func main() {
//	fmt.Println(cal.Cal(10,20))
//}

//import "fmt"
//func test(n1 int64) {
//	if n1 > 10 {
//		n1--
//		test(n1)
//	}
//	fmt.Println(n1)
//}
//
//func main() {
//	test(64)
//}

//import (
//	"fmt"
//)
//
//func test(last int64) (sum int64) {
//	count := 0
//	count++
//	if count < 10 {
//		sum := (last + 1) * 2
//		count++
//		test(last)
//	}
//	return sum
//}
//
//func main() {
//	nice := test(1)
//	fmt.Println(nice)
//}

//func echo(n int) int {
//	return	n
//}
//func main()  {
//	n := echo(100)
//	fmt.Println(n)
//	m := echo
//	fmt.Printf("%T",m)
//}

//import (
//	"cal"
//	"fmt"
//)
//func main() {
//	sum := cal.Cal
//	fmt.Printf("%T\n",sum)
//	a := sum(100,200)
//	fmt.Printf("%T,%v",a,a)
//}

//type myint int
//func main() {
//	var sum myint  = 1
//	var num int = 1
//	var madam int
//	madam = int(sum)
//	fmt.Printf("%T,%T,%T",sum,num,madam)
//}

//func sum(n1 int, args ...int) int{
//	sum := n1
//	for i:=0;i < len(args);i++ {
//		sum += args[i] //表示取出args切片的第一个元素
//	}
//	return sum
//}
//func main() {
//	sum := sum(10,20,30)
//	fmt.Println(sum)
//}

//func swap(n1 *int, n2 *int) {
//	t := *n1
//	*n1 = *n2
//	*n2 = t
//}
//
//func main() {
//	a := 10
//	b := 20
//	swap(&a,&b)
//	fmt.Println(a,b)
//}
