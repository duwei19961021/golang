package model

import "fmt"

type person struct {
	Name string
	age  int
	sal  float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	p.age = age
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	p.sal = sal
}

func (p *person) GetSal() float64 {
	return p.sal
}

type account struct {
	User    string
	pwd     string
	balance float64
}

func Newaccount(name string) *account {
	return &account{
		User: name,
	}
}
func (p *account) SetAccount(pwd string) {
	if len(pwd) != 6 {
		fmt.Println("length error")
	} else {
		p.pwd = pwd
	}
}
func (p *account) Getbalance() float64 {
	return p.balance
}
