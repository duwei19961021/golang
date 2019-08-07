package main

import "fmt"

func cal(n1 float64, n2 float64) (sum float64) {
	sum = n2 + n1
	return sum
}
func main() {
	result := cal(10, 20)
	fmt.Println(result)
}
