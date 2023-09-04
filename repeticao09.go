package main

import "fmt"

//Faça um algoritmo que leia vários números inteiros e mostre
//a média aritmética entre eles.
//A leitura deve ser interrompida quando for lido o valor zero.

func main() {
	var soma, quantidade, x int
	x = -1
	for x != 0 {
		fmt.Print("Digite um número: ")
		fmt.Scan(&x)

		soma += x
		quantidade++
	}
	media := soma / quantidade
	fmt.Println("A média é", media)
}
