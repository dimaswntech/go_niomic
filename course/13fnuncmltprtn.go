package main

import "fmt"

func main() {
	var z1, z2 = calcOpt(10, 5)
	fmt.Println(z1, z2)
}
func calcOpt(z int, y int) (int, int) {
	var q1 = z + y
	var q2 = z - y
	return q1, q2
}
