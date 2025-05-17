package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/marcboeker/go-duckdb/v2"
	"fmt"
)

type MyDatabase struct {
	con *sql.DB
}

func (db *MyDatabase) connectDatabase() error {

	var err error

	db.con, err = sql.Open("sqlite3", "file:database.sqlite")
	//db.con, err = sql.Open("duckdb", "database.duckdb")

	if err != nil {
		fmt.Printf("Error ao abrir conexão %s", err.Error())
		return err
	}

	fmt.Println("Conectado a base de dados")

	return nil
}

func (db *MyDatabase) queryAllRequirements() (*sql.Rows, error) {

	rows, err := db.con.Query(`SELECT "id", "descricao" FROM "requisitos" `)

	if err != nil {
		return nil, err
	}

	return rows, nil

}

func (db *MyDatabase) closeDatabase() {
	db.con.Close()
}
