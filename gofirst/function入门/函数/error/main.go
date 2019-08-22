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
//func main() {

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
/*
	studentMap := make(map[string]map[string]string)
	studentMap["stu1"] = make(map[string]string,3)
	studentMap["stu1"]["name"] = "tom"
	studentMap["stu1"]["sex"] = "man"

	studentMap["stu2"] = make(map[string]string,3)
	studentMap["stu2"]["name"] = "alice"
	studentMap["stu2"]["sex"] = "woman"
	fmt.Println(studentMap) //统计key，value
	//fmt.Println(len(studentMap))
	fmt.Println(studentMap["stu1"])
	fmt.Println(studentMap["stu1"]["name"])
	fmt.Println(studentMap["stu1"]["sex"])
*/
//map增删改查
/*
	studentMap["stu2"]["name"] = "alice~" //若key已存在则是修改
	fmt.Println(studentMap)
	studentMap["stu2"]["address"] = "hz" //若key不存在则是增加
	fmt.Println(studentMap)
	delete(studentMap["stu2"],"address") //删除key，如删除的key也不会报错=不操作
	fmt.Println(studentMap)
	val,ok := studentMap["stu1"]
	if ok {
		fmt.Println("y",val)
	}else {
		fmt.Println("n")
	}
*/
//删除map
//①遍历，逐个删除
//②make一个新的空间，原来的内存会被gc回收
/*
	studentMap = make(map[string]map[string]string)
	fmt.Println(studentMap)
*/
//遍历map
/*
	for _,value := range studentMap{
		for key1,value1 := range value{
			fmt.Println(key1,value1)
		}
	}
*/

//map切片
/*
	var monster []map[string]string
	monster = make([]map[string]string, 2) //容量为2
	if monster[0] == nil {
		monster[0] = make(map[string]string, 2)
		monster[0]["name"] = "m1"
		monster[0]["age"] = "18"
	}
	if monster[1] == nil {
		monster[1] = make(map[string]string, 2)
		monster[1]["name"] = "m2"
		monster[1]["age"] = "19"
	}
*/

//if monster[2] == nil {	//越界了，不合法，报错
//	monster[2] = make(map[string]string,2)
//	monster[2]["name"] = "m3"
//	monster[2]["age"] = "20"
//}

//使用切片的append函数，动态增加。
/*
	newmonster := map[string]string{
		"name" : "m3",
		"age" : "21",
	}
	monster = append(monster,newmonster)
	fmt.Println(monster)
*/

//}

//func main() {
//	file,err := os.Open("C:\\Users\\x1c\\GolandProjects\\golang\\gofirst\\function入门\\函数\\error\\a.txt")
//	fmt.Println(*file,err)
//	file.Close()
//}

//读文件1
/*
func main() {
	file,err :=os.Open("C:\\Users\\x1c\\GolandProjects\\golang\\gofirst\\function入门\\函数\\error\\a.txt")
	if err != nil{
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		fmt.Print(str)
	}
}
*/

/*
func main() {
	filePath := "d:/golang.txt"
	file,err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE|os.O_APPEND,0777)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	str := "hello world!\n"
	writer := bufio.NewWriter(file)
	for i:=0;i<5;i++{
		writer.WriteString(str)
	}
	writer.Flush()
}
*/

/*
func main()  {
	oldfilePath := "d:/send.txt"
	newfilePath := "d:/accept.txt"
	file, err := os.OpenFile(oldfilePath,os.O_RDONLY,0777)
	if err != nil{
		fmt.Println(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		fmt.Print(str)
		file,err := os.OpenFile(newfilePath,os.O_RDWR|os.O_APPEND,0777)
		if err != nil {
			fmt.Println(err)
		}
		writer := bufio.NewWriter(file)
		writer.WriteString(str)
		writer.Flush()
	}
}
*/

/*
var count struct{
	zimucount int
	shuzicount int
	spacecount int
	othercount int
}

func main() {
		filePath := "d:\\send.txt"
		file, err := os.Open(filePath)
		if err != nil{
			fmt.Println(err)
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		for {
			str,err := reader.ReadString('\n')
			if err == io.EOF{
				break
			}
			for _,v := range str{
				switch  {
				case v >= 'a' && v <='z' :
					count.zimucount++
				case v >= 'A' && v <= 'Z':
					count.zimucount++
				case v == ' ' || v == '\t':
					count.spacecount++
				case v >='0' && v <= '9':
					count.shuzicount++
				default:
					count.othercount++
				}
			}
		}
		fmt.Printf("zimu:%v\n",count.zimucount)
		fmt.Printf("shuzi:%v\n",count.shuzicount)
		fmt.Printf("space:%v\n",count.spacecount)
		fmt.Printf("other:%v\n",count.othercount)
}
*/

/*
func main() {
	fmt.Println(len(os.Args))
	for i,v := range os.Args{
		fmt.Printf("args[%v]=%v\n",i,v)
	}
}
*/

//命令行参数
/*
func main() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user,"u","","用户名")
	flag.StringVar(&pwd,"pwd","","密码")
	flag.StringVar(&host,"h","","主机")
	flag.IntVar(&port,"p",3306,"用户名")
	flag.Parse()
	fmt.Printf("user:%v,pwd:%v,host:%v,port:%v",user,pwd,host,port)
}
*/

//序列化
import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name    string
	Age     int
	Sal     float64
	Address string
	Skill   string
}

func test() {
	mon := Monster{
		Name:    "duwei",
		Age:     23,
		Sal:     1000,
		Address: "hunan",
		Skill:   "wu",
	}
	data, err := json.Marshal(&mon)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
func main() {
	test()
}
