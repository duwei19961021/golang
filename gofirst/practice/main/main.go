package main

//import "fmt"
//var n1 float32 = 3.1415926542343453452137856
//
//func main() {
//	fmt.Printf("%9.f",n1)
//}

//func test(num1 int, num2 int) (shang int) {
//	defer func() {
//		err := recover()
//		if err != nil {
//			fmt.Println(err)
//		}
//	}()
//	shang = num1 / num2
//	return shang
//}
//
//func main() {
//	test(10,0)
//	fmt.Println("ok")
//}

//var shuzu [3]int = [3]int{1,2,3}
//func main() {
//	for index,value := range shuzu {
//		fmt.Println(index,value)
//	}
//}

//将[A-Z]存入数组
//func main() {
//	var shuzu [26]string
//	for i :=65;i<91;i++{
//		zimu := fmt.Sprintf("%c",i)
//		//fmt.Printf("%T,%v",zimu,zimu)
//		shuzu[i-65] = zimu
//	}
//	for _,value := range shuzu {
//		fmt.Println(value)
//	}
//}

//求数组中的最大值及其下标
//func main() {
//	var shuzu [5]int = [5]int{23,4,15,90,20}
//	maxvalue := shuzu[0]
//	maxindex := 0
//	for i := 1; i < len(shuzu); i++ {
//		if maxvalue < shuzu[i] {
//			maxvalue = shuzu[i]
//			maxindex = i
//		}
//	}
//	fmt.Println(maxvalue,maxindex)
//}

//计算数组所有元素的和和平均值
//func main() {
//	var shuzu [5]int = [5]int{10,5,38,20,28}
//	sum := 0
//	for _,value := range shuzu{
//		sum += value
//	}
//	fmt.Println(sum)
//	fmt.Println(float64(sum)/float64(len(shuzu)))
//}

//生成十个随机数，反转打印
//func main() {
//	var shuzu [10]int
//	rand.Seed(time.Now().UnixNano())
//	for i := 0 ; i < 10 ;i ++ {
//		x := rand.Intn(100)
//		shuzu[i] = x
//	}
//	fmt.Println(shuzu)
//	for i := len(shuzu)-1 ; i >= 0;i-- {
//		fmt.Println(shuzu[i])
//	}
//}

//切片
//func main() {
//	var qiepian []int = []int{1,2,2,1,4,6,3,456,3322,5}
//	slice := qiepian[1:3]
//	fmt.Println(slice,cap(slice))
//	fmt.Printf("%T",qiepian)
//}

//func main() {
//	i,j := 10 ,11
//	fmt.Println(i,j)
//	i,j = 11,10
//	fmt.Println(i,j)
//}

//切片定义
//var qiepian1 []int  = make([]int,4,10) //①
//var qiepian2 []int  = []int{1,2,3}
//
//func main() {
//	fmt.Println(qiepian1)
//	fmt.Println(qiepian2)
//}

//切片遍历
//①
//var slice [] int = []int{1,2,3,4,5,6}
//var slice [] int = make([]int,10,15)
//func main() {
//	for i := 0; i < len(slice); i++ {
//		fmt.Println(slice[i])
//	}
//	for index,value := range slice {
//		fmt.Println(index,value)
//	}
//}

//func main() {
//	var slice [] int = make([]int,4,5)
//	fmt.Printf("%p\n",&slice)
//	slice = append(slice, 5)
//	fmt.Printf("%p\n",slice)
//	slice = append(slice, 5)
//	fmt.Printf("%p",slice)
//}
/*
res:	0xc00004c420
		0xc000080030
		0xc00005c050
*/

//切片copy
//func main() {
//	var slice1 []int = []int{1,2,3,4,5,6}
//	var slice2 []int = []int{7,8,9}
//	copy(slice1,slice2)
//	fmt.Println(slice1)
//}
/*
res:	[7 8 9 4 5 6]
*/

//string切片
//var str string  = "duwei"
//func main() {
//	fmt.Println(str)
//	strb := []byte(str)
//	//fmt.Println(strb)
//	strb[0] = 'D'
//	str = string(strb)
//	fmt.Println(str)
//
//	strb1 := []rune(str)
//	strb1[0] = '杜'
//	str = string(strb1)
//	fmt.Println(strb1)
//}

//func test(n int) []uint64 {
//	slice := make([]uint64, n)
//	slice[0] = 1
//	slice[1] = 1
//	for i := 2; i < n; i++ {
//		slice[i] = slice[i-1] + slice[i-2]
//	}
//	return slice
//}
//func main() {
//	fmt.Println(test(12))
//}
