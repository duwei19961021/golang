package main

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
//func main() {
//	var shuzu [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	for index, value := range shuzu {
//		fmt.Println(index, value)
//	}
//}

//冒泡排序
//func main() {
//	slice := []int{69,24,80,57,55,45}
//	for i := (len(slice) - 1); i > 0; i-- {
//		for i := 0; i < (len(slice) - 1); i++ { //长度为n的数组，两数直接比较n-1次，此层循环一次最大数会被排到最后，再次循环时最大数(排最后的数)
//		// 无需再参与比较，因此此层循环一次，上层循环减少一次，上层则使用i--
//			if slice[i] > slice[i+1] { //如果第一个数大于第二个数则交换它们的位置，否则不做交换，各自等于原来的值(以下6行)
//				t := slice[i+1]
//				slice[i+1] = slice[i]
//				slice[i] = t
//			} else {
//				slice[i] = slice[i]
//				slice[i+1] = slice[i+1]
//			}
//		}
//	}
//	fmt.Println(slice)
//}

//func mp(arr *[10]int) {
//		for j := (len(*arr) - 1); j > 0; j-- {
//			for i := 0; i < (len(*arr) - 1); i++ { //长度为n的数组，两数直接比较n-1次，此层循环一次最大数会被排到最后，再次循环时最大数(排最后的数)
//			// 无需再参与比较，因此此层循环一次，上层循环减少一次，上层则使用i--
//				if (*arr)[i] > (*arr)[i+1] { //如果第一个数大于第二个数则交换它们的位置，否则不做交换，各自等于原来的值(以下6行)
//					t := (*arr)[i+1]
//					(*arr)[i+1] = (*arr)[i]
//					(*arr)[i] = t
//				}
//			}
//		}
//	}
//
//func main() {
//	arr := [10]int{69,24,80,57,55,45,65,24,58,12}
//	mp(&arr)
//	fmt.Println(arr)
//}

//二分
//func binary(arr *[11]int,Leftindex int,Rightindex int,num int){
//	middle := (Leftindex+Rightindex)/2
//}
//func main() {
//	arr := [11]int{12,14,16,19,20,25,29,30,33,39,40}
//	binary(&arr,leftindex)
//}

//二位数组
//func main() {
//	var arr [4][6]int
//	arr [1][2]  = 1
//	for i:=0;i < 4;i++{
//		for j:=0;j < 6;j++{
//			fmt.Print(arr[i][j]," ")
//		}
//		fmt.Println()
//	}
//}

//func main() {
//var arr [4][6]int = [4][6]int{{1,2,3,4,5,6},{1,2,3,4,5,6},{1,2,3,4,5,6},{1,2,3,4,5,6}}

//fmt.Println(arr)
//for i := 0 ;i < 4;i++ {
//	for j := 0; j < 6; j++ {
//		fmt.Print(arr[i][j]," ")
//	}
//	fmt.Println()
//}

//for index,value := range arr{
//	for j,v2 := range value{
//		fmt.Printf("arr[%v][%v]=[%v]\n",index,j,v2)
//	}
//}
//}

//func main() {
//	var arr [3][5]int
//	for i,v1 := range arr{
//		for j,_ := range v1{
//			fmt.Printf("请输入第%d个班第%d位:",i+1,j+1)
//			fmt.Scanln(&arr[i][j])
//		}
//	}
//	fmt.Println(arr)
//}

//map声明
//func main() {
//	var a map[string]string //使用前必须给map分配数据空间
//	a = make(map[string]string,10) //10表示分配给map的空间，能存10个key-value
//	a["name"] = "dw" //value可重复，key不可重复(覆盖先前的)
//	a["age"] = "18"
//	a["hoby"] = "none"
//	a["home"] = "hn"
//	fmt.Println(a)
//}

//map使用
func main() {

	//var a map[string]string
	//a = make(map[string]string,10) //不写空间默认为1
	//a["name"] = "dw"
	//fmt.Println(a)

	//cities := make(map[string]string) //容量自增涨，无需定义
	//cities["city1"] = "hz"
	//cities["city2"] = "hz"
	//cities["city3"] = "hz"
	//cities["city4"] = "hz"
	//fmt.Println(cities)

	//name := map[string]string{
	//	"name1" : "aa",
	//	"name2" : "bb",
	//}
	//fmt.Println(name)

}
