package main

/*
var principal float64 //本金
var income float64	//收入
var reason string	//说明
var	spending float64	//支出
var balance  float64	//余额
var record	string	//记录

func main() {
	for {
		var choice int
		var flg bool
		fmt.Println("---------------家庭收支记账软件-----------------------")
		fmt.Println()
		fmt.Println("                 1 收支明细                           ")
		fmt.Println("                 2 登记收入                           ")
		fmt.Println("                 3 登记支出                           ")
		fmt.Println("                 4 退   出")
		fmt.Println()
		fmt.Println("请选择(1-4):")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("------------------当前收支明细记录---------------------")
			fmt.Println("收支     账户金额     收支金额          说    明        ")
			if len(record) == 0{
				fmt.Println("当前无任何收支明细！")
			}else {
				fmt.Println(record)
			}
		case 2:
			fmt.Println("本次收入金额:")
			fmt.Scanln(&income)
			fmt.Println("本次收入说明:")
			fmt.Scanln(&reason)
			balance += income
			record += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, income, reason)
		case 3:
			fmt.Println("本次支出金额:")
			fmt.Scanln(&spending)
			fmt.Println("本次支出说明:")
			fmt.Scanln(&reason)
			if spending > balance{
				fmt.Println("余额不足")
			}else {
				balance -= spending
				record += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, spending, reason)
			}

		case 4:
			var yn string
			fmt.Println("确定退出？(y/n)")
			for {
				fmt.Scanln(&yn)
				if yn == "y" || yn == "n" {
						break
				} else {
					fmt.Println("输入有误，重新输入")
				}
			}
			if yn == "y"{
				flg = true
			}
		}
		if flg{
			break
		}
	}
}
*/

import "familyaccount/utils"

func main() {
	utils.NewFamilyAccount().Mainmenu()
}
