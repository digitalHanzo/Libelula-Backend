package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"../models"
)

//temporary data
var families []models.Family

//DB connection
//(declared on bugController)

//GetBugsEndPoint - Get all data from PG
func GetFamiliesEndPoint(res http.ResponseWriter, req *http.Request) {
	queryStmt, err := database.Prepare("SELECT * FROM families")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := queryStmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		families = append(families, models.Family{ID: id, Name: name})
	}
	json.NewEncoder(res).Encode(families)

	families = nil
}
