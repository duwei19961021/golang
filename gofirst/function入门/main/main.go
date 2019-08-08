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

import (
	"fmt"
)

func test(last int64) (sum int64) {
	count := 0
	count++
	if count < 10 {
		sum := (last + 1) * 2
		count++
		test(last)
	}
	return sum
}

func main() {
	nice := test(1)
	fmt.Println(nice)
}

//1 4  (n+1)2
