package main

import "fmt"

var choice int

func main() {
	fmt.Println("---------------客户信息管理---------------")
	fmt.Println()
	fmt.Println("                1 添加客户")
	fmt.Println("                2 修改客户")
	fmt.Println("                3 删除客户")
	fmt.Println("                4 客户列表")
	fmt.Println("                5 退    出")
	fmt.Println("请选择(1-5):")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("-----------------添加客户-----------------")
	case 2:
		fmt.Println("-----------------修改客户-----------------")
	case 3:
		fmt.Println("-----------------删除客户-----------------")
	case 4:
		fmt.Println("-----------------客户列表-----------------")
	case 5:
		fmt.Println("exit")
		break
	}
}
