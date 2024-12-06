package main

import "fmt"

func brojac(broj int) (int, int, int, int) {
	a := broj / 100
	b := (broj / 10) % 10
	c := broj % 10
	zbr := a + b + c
	return a, b, c, zbr
}
func main() {
	fmt.Println("Unesi troznamenkasti broj: ")
	var broj int
	fmt.Scanln(&broj)
	fmt.Println(brojac(broj))
}
