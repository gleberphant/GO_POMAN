package cli

import (
	"fmt"
	"log"
	"poman/mydb"
)

func RunCli() {

	fmt.Println("Bem Vindo ao POMAN-GO")
	var connection mydb.MyDatabase

	if err := connection.ConnectDatabase(); err != nil {
		log.Fatal(err)
		return
	}

	defer connection.CloseDatabase()

	fmt.Println("Consultando tabelas")

	rows, _ := connection.QueryAllRequirements()

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
