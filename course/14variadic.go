package main

import "fmt"

func main() {
	var rata_rata = hitung(1, 3, 2, 4, 12, 4)
	var pesan = fmt.Sprintln("Rata - rata: ", rata_rata)
	fmt.Println(pesan)
}
func hitung(angka ...int) float64 {
	var total int = 0
	for i := 0; i < len(angka); i++ {
		total = total + angka[i]
	}
	var avg = float64(total) / float64(len(angka))
	return avg
}
