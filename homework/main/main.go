package main

//p173
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		var num int = rand.Intn(100)
		fmt.Println(num)
		time.Sleep(100 * time.Millisecond)
		arr[i] = num + 1
	}
	fmt.Println(arr)
	for index, value := range arr {
		if value == 55 {
			fmt.Println("yes", index)
			break
		}
	}
}

//	fmt.Println(shuzu)
//	for i := len(shuzu)-1 ; i >= 0;i-- {
//		fmt.Println(shuzu[i])

//for k:=len(arr)-1;k>0 ;k--  {
//	for j:=0;j<len(arr)-1;j++{
//		if (arr)[j] < (arr)[j+1]{
//			var t int
//			t = (arr)[j+1]
//			(arr)[j+1] = (arr)[j]
//			(arr)[j] = t
//		}
//	}
//}

//p197
