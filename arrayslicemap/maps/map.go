package main

import "fmt"

func main() {
	aprovados := make(map[int]string)
	aprovados[1234678978] = "Zoro"
	aprovados[1234678900] = "Sanji"
	aprovados[1234678901] = "Luffy"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
}
