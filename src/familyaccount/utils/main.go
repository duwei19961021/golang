package utils

import "fmt"

type FamilyAccount struct {
	income   float64 //收入
	reason   string  //说明
	spending float64 //支出
	balance  float64 //余额
	record   string  //记录
	yn       string
	flg      bool
	choice   int
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		balance: 1000,
		record:  "",
		yn:      "",
		flg:     false,
	}

}

func (this *FamilyAccount) info() {
	fmt.Println("------------------当前收支明细记录---------------------")
	fmt.Println("收支     账户金额     收支金额          说    明        ")
	if len(this.record) == 0 {
		fmt.Println("当前无任何收支明细！")
	} else {
		fmt.Println(this.record)
	}

}
func (this *FamilyAccount) come() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.income)
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.reason)
	this.balance += this.income
	this.record += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.income, this.reason)
}
func (this *FamilyAccount) spend() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.spending)
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.reason)
	if this.spending > this.balance {
		fmt.Println("余额不足")
	} else {
		this.balance -= this.spending
		this.record += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.spending, this.reason)
	}
}

func (this *FamilyAccount) quit() {
	fmt.Println("确定退出？(y/n)")
	for {
		fmt.Scanln(&this.yn)
		if this.yn == "y" || this.yn == "n" {
			break
		} else {
			fmt.Println("输入有误，重新输入")
		}
	}
	if this.flg {
		return
	}
}

func (this *FamilyAccount) Mainmenu() {
	for {
		fmt.Println("---------------家庭收支记账软件-----------------------")
		fmt.Println()
		fmt.Println("                 1 收支明细                           ")
		fmt.Println("                 2 登记收入                           ")
		fmt.Println("                 3 登记支出                           ")
		fmt.Println("                 4 退   出")
		fmt.Println()
		fmt.Println("请选择(1-4):")
		fmt.Scanln(&this.choice)
		switch this.choice {
		case 1:
			this.info()
		case 2:
			this.come()
		case 3:
			this.spend()
		case 4:
			this.quit()
		}
	}
}
