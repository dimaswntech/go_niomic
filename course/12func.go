package main

import "fmt"

func main() {
	showMessage()
	fmt.Println(showItem())
	fmt.Println(calcOpt(12, 12))
}
func showMessage() {
	fmt.Println("yes")
}
func showItem() string {
	return "table"
}
func calcOpt(z int, y int) int {
	var q = z + y
	return q
}
