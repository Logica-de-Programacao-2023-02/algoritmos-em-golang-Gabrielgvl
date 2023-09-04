package main

import "fmt"

//Faça um algoritmo que leia três números reais e
//mostre-os em ordem crescente.

func main() {
	var x, y, z int
	fmt.Print("Digite 3 números: ")
	fmt.Scan(&x, &y, &z)

	if x > y && x > z {
		if y > z {
			fmt.Printf("%d > %d > %d", x, y, z)
		} else {
			fmt.Printf("%d > %d > %d", x, z, y)
		}
	} else if y > x && y > z {
		if x > z {
			fmt.Printf("%d > %d > %d", y, x, z)
		} else {
			fmt.Printf("%d > %d > %d", y, z, x)
		}
	} else {
		if x > y {
			fmt.Printf("%d > %d > %d", z, x, y)
		} else {
			fmt.Printf("%d > %d > %d", z, y, x)
		}
	}
}
