package main

import "fmt"

//func test()  {
//	defer func() {
//		err := recover()
//		if err != nil{
//			fmt.Println(err)
//		}
//	}()
//	num1 := 10
//	num2 := 0
//	res := num1/num2
//	fmt.Println(res)
//}
//func main() {
//	test()
//	fmt.Println("ok")
//}

//array
//func main()  {
////var shuzu [6]float64
////shuzu[0] = 3.0
////shuzu[1] = 3.0
////shuzu[2] = 3.0
////shuzu[3] = 3.0
////shuzu[4] = 3.0
////shuzu[5] = 3.0
////var totalw float64 = 0.0
////for i := 0 ;i < len(shuzu);i++{
////	totalw += shuzu[i]
////}
////fmt.Println(totalw,len(shuzu))
////}

//func main() {
//	var shuzu [5]int
//	shuzu[0] = 10
//	fmt.Println(shuzu)
//	fmt.Printf("%p,%p,%p,%p",&shuzu,&shuzu[0],&shuzu[1],&shuzu[2])
//}

//func main() {
//	var shuzu [5]float64
//	for i:=0;i < len(shuzu);i++ {
//		fmt.Println("input\n")
//		fmt.Scanln(&shuzu[i])
//	}
//	for i:=0;i < len(shuzu);i++ {
//		fmt.Println(shuzu[i])
//	}
//}

//初始化数组的四种方式
/*
func main() {
	var shuzu1 [3]int = [3]int{1, 2, 3}
	var shuzu2 = [3]int{4,5,6}
	var shuzu3 = [...]int{7,8,9,10}
	var shuzu4 = [...]int{1:12,0:11,2:13}
	fmt.Println(shuzu1,shuzu2,shuzu3,shuzu4)
}
*/

//for-range结构遍历
func main() {
	var shuzu [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range shuzu {
		fmt.Println(index, value)
	}
}
