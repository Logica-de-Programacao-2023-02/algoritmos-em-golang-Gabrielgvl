package main

import "fmt"

// Faça um algoritmo que leia vários números inteiros e
//mostre o maior deles.
//A leitura deve ser interrompida quando for lido o valor zero.

func main() {
	var maior, x int
	x = -1
	for x != 0 {
		fmt.Print("Digite um número: ")
		fmt.Scan(&x)

		if x > maior {
			maior = x
		}
	}
	fmt.Println("O maior é", maior)
}
