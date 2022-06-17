package main

import "fmt"

func main() {
	var tinggiBadan = map[string]int{"aldo": 182, "yosep": 178}
	tampil_mahasiswa(tinggiBadan)
}
func tampil_mahasiswa(tinggiBadan map[string]int) {
	for k, v := range tinggiBadan {
		fmt.Println(k, "\t:", v, "cm")
	}
}
