package main

func main() {
	//str := "hello你"
	//fmt.Println(len(str)) //按字节返回，一个汉字三个字节

	//str1 := []rune(str) //字符串遍历，处理带有中文的字符串
	//for i:=0;i<len(str1);i++{
	//	fmt.Printf("%c\n",str1[i])
	//}

	//str := "123"
	//n, err := strconv.Atoi(str) //字符串转整数。利用此特性进行数据校验：数字字符串判断
	//if err != nil {
	//	fmt.Println("转换错误:",err)
	//}else {
	//	fmt.Println("转换成功:",n)
	//}

	//var str int = 123456
	//str1 := strconv.Itoa(str) //整数转字符串
	//fmt.Printf("%T",str1)

	//var str string = "duwei"
	//var bytes = []byte(str) //字符串转byte
	//fmt.Printf("%T,%v",bytes,bytes)

	//str := string([]byte{100,117,119,101,105}) //byte转字符串
	//fmt.Println(str)
	//
	//str = strconv.FormatInt(123,16) //十进制转其他进制：16
	//fmt.Println(str)

	//res := strings.Contains("dsafsduweiduwsdfeiduwesdiduweidd","duwei") //判断是否包含，返回bool类型
	//fmt.Printf("%v",res)

	//res := strings.Count("dsafsduweiduwsdfeiduwesdiduweidd","duwei") //统计指定子串的次数
	//fmt.Println(res)

	//res := strings.EqualFold("abc","ABC") //判断字符串相等，忽略大小写
	//fmt.Println(res)

	//res := strings.Index("dddduweidsafsdfuweiduwsdfeiduwesdiduweidd","duwei") //返回最第一个出现的索引
	//	//fmt.Println(res)

	//	//res2 := strings.LastIndex("dsafsdfuweiduwsdfeiduwesdiduweidd","duwei") //返回最最后一个出现的索引
	//	//fmt.Println(res2)

	//res := strings.Replace("dsafs duwei duwsdfei duwei sdi duwei dd","duwei","lisi",-1)
	//fmt.Println(res)
	//res := strings.Replace("duwei#duwei","#","",-1) // 删除某字符
	//fmt.Println(res)

	//res := strings.Split("a,b,c,d,e,f,g",",") //分割
	//fmt.Println(res)

	//res := strings.ToLower("DUWEI")
	//res2 := strings.ToUpper("duwei")
	//fmt.Println(res,"\n",res2)

	//res := strings.TrimSpace("   sdfsdf  sdfsdf  ") //去掉两端空格
	//fmt.Println(res)

	//res := strings.Trim("!!!!@hello!!!","!") //去掉两端指定的字符串
	//fmt.Println(res)
	//res := strings.Trim("!#duwei#@!#$#!","!#")
	//fmt.Println(res)
	//strings.TrimLeft() //去掉左端字符
	//strings.TrimRight() //去掉右端字符

	//res := strings.HasPrefix("ftp://10.0.0.1","ftp") //判断是否以某某开头
	//fmt.Println(res)

	//res := strings.HasSuffix("ftp://10.0.0.1duwei","duwei") //判断是否以某某结尾
	//fmt.Println(res)

}
