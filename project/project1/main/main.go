package main

type record struct {
	Income   float64
	Spending float64
	Balance  float64
	Reason   string
}

func (Ic *record) income(ic float64, rs string) {
	(*Ic).Income = ic
	Ic.Reason = rs
}
func (Sd *record) spending(sd float64) float64 {
	//return Sd.Balance -= sd
}

func main() {

}

//func main() {
//	var choice int
//	fmt.Println("---------------家庭收支记账软件-----------------------")
//	fmt.Println()
//	fmt.Println("                 1 收支明细                           ")
//	fmt.Println("                 2 登记收入                           ")
//	fmt.Println("                 3 登记支出                           ")
//	fmt.Println("                 4 退   出")
//	fmt.Println()
//	fmt.Println("                 请选择(1-4):")
//	fmt.Scanln(&choice)
//	switch choice {
//	case 1:
//		fmt.Println("------------------当前收支明细记录---------------------")
//		fmt.Println("收支     账户金额     收支金额          说    明        ")
//	}
//
//
//
//
//
//
//
//
//
//
//
//}
