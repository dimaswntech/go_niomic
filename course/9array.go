package main

import "fmt"

func main() {
	var buah = [4]string{"apel", "jeruk", "anggur", "melon"}
	fmt.Println(len(buah))
	fmt.Println(buah)
	buah[0] = "ace"
	fmt.Println(buah)
}
