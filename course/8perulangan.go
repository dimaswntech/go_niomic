package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Angka", i)
		if i == 4 {
			break
		}
	}
}