package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"poman/models"
	"poman/mydb"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var pages, _ = template.ParseGlob("templates/*.html")

	requisitos := []models.Requisito{}

	db := mydb.MyDatabase{}

	db.ConnectDatabase()

	rows, _ := db.QueryAllRequirements()

	for rows.Next() {

		var id int
		var descricao string

		rows.Scan(&id, &descricao)

		r := models.Requisito{
			Id:         id,
			Descricao:  descricao,
			Prioridade: "normal",
		}

		requisitos = append(requisitos, r)
		fmt.Println(r)

	}
	pages.ExecuteTemplate(w, "layout", requisitos)

	defer db.CloseDatabase()

}
