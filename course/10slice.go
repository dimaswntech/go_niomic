package main

import "fmt"

func main() {
	var buah = []string{"apel", "anggur"}
	fmt.Println(buah)
	buah = append(buah, "mangga")
	fmt.Print(len(buah))
}
