package main

import (
	"fmt"
	"log"
)

func cliApp() {

	fmt.Println("Bem Vindo ao POMAN-GO")
	var connection MyDatabase

	if err := connection.connectDatabase(); err != nil {
		log.Fatal(err)
		return
	}

	defer connection.closeDatabase()

	fmt.Println("Consultando tabelas")

	rows, _ := connection.queryAllRequirements()

	defer rows.Close()

	fmt.Println("Resultados: ")

	for rows.Next() {
		var id, descricao string
		erro := rows.Scan(&id, &descricao)
		if erro != nil {
			fmt.Printf("erro leitura da linha : %s\n", erro)
		} else {
			fmt.Printf("| ID: %s | DESCRIÇÃO: %s |\n", id, descricao)
		}
	}

}
