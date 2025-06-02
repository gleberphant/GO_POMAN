package mydb

import (
	"database/sql"

	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type MyDatabase struct {
	con *sql.DB
}

func (db *MyDatabase) ConnectDatabase() error {

	var err error

	db.con, err = sql.Open("sqlite3", "file:database.sqlite")

	if err != nil {
		fmt.Printf("Error ao abrir conexão %s", err.Error())
		return err
	}

	fmt.Println("Conectado a base de dados")

	return nil
}

func (db *MyDatabase) QueryAllRequirements() (*sql.Rows, error) {

	rows, err := db.con.Query(`SELECT "id", "descricao" FROM "requisitos" `)

	if err != nil {
		return nil, err
	}

	return rows, nil

}

func (db *MyDatabase) CloseDatabase() {
	db.con.Close()
}
