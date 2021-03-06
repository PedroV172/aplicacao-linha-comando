package main

import (
	"fmt"
	"linha-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("-----------------")

	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
		fmt.Println("Caiu no Erro!")
	}
	fmt.Println("-----------------")
}
