package main

import "fmt"

//Faça um algoritmo que leia a idade de uma pessoa e mostre a sua
//classificação, de acordo com a seguinte tabela:
//até 9 anos: mirim
//10 a 13 anos: infantil
//14 a 17 anos: juvenil
//maiores de 18 anos: adulto

func main() {
	var idade int
	fmt.Print("Digite sua idade")
	fmt.Scan(&idade)
	if idade <= 9 {
		fmt.Println("Você é mirim")
	} else if idade <= 13 {
		fmt.Println("Você é infantil")
	} else if idade <= 17 {
		fmt.Println("Você é juvenil")
	} else {
		fmt.Println("Você é adulto")
	}
}
