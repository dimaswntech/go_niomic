package main

import "fmt"

func main() {
	var buah = []string{"apel", "jeruk", "anggur", "melon"}
	buah = append(buah, "papaya")
	for i := 0; i < len(buah); i++ {
		fmt.Println("Elemen ke - :", i, buah[i])
	}
}
