package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../models"
	"github.com/gorilla/mux"
)

//temporary data
var families []models.Family

//DB connection
//(declared on bugController)

//GetFamiliesEndPoint - Get all data from PG
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

//GetFamilyEndPoint - Get specific bug from PG
func GetFamilyEndPoint(res http.ResponseWriter, req *http.Request) {

	//obtener parametro por ruta
	params := mux.Vars(req)

	id := params["id"]

	queryStmt, err := database.Prepare(fmt.Sprintf("SELECT * FROM families WHERE id = '%s';", id))

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

//CreateFamilyEndPoint - Insert on DB
func CreateFamilyEndPoint(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	req.ParseForm()

	var insertFamily models.Family

	//Get data from body and storing in new object
	insertFamily.Name = req.FormValue("name")

	queryStmt, err := database.Prepare(fmt.Sprintf("INSERT INTO families (name) VALUES ('%s');", insertFamily.Name))

	if err != nil {
		json.NewEncoder(res).Encode(err)
	} else {

		inserted, err := queryStmt.Query()
		if err != nil {
			json.NewEncoder(res).Encode(err)

		} else {
			json.NewEncoder(res).Encode(inserted)
		}
	}

}

//EditFamilyEndPoint - Editing from db
func EditFamilyEndPoint(res http.ResponseWriter, req *http.Request) {

	//obtener parametro por ruta
	params := mux.Vars(req)
	id := params["id"]

	req.ParseForm()
	//Get data from body and storing in new object
	var editFamily models.Family
	editFamily.Name = req.FormValue("name")

	queryStmt, err := database.Prepare(fmt.Sprintf("UPDATE families SET name = '%s' WHERE id = %s;", editFamily.Name, id))

	if err != nil {
		json.NewEncoder(res).Encode(err)
	} else {

		updated, err := queryStmt.Query()
		if err != nil {
			json.NewEncoder(res).Encode(err)

		} else {
			json.NewEncoder(res).Encode(updated)
		}
	}
}

//DeleteFamilyEndPoint - Deleting from db
func DeleteFamilyEndPoint(res http.ResponseWriter, req *http.Request) {

	//obtener parametro por ruta
	params := mux.Vars(req)

	id := params["id"]

	queryStmt, err := database.Prepare(fmt.Sprintf("DELETE FROM families WHERE id = '%s';", id))

	if err != nil {
		log.Fatal(err)
	}

	rows, err := queryStmt.Query()
	if err != nil {
		panic(err)
	} else {

		response := map[string]string{
			"status":  "200",
			"message": "row deleted",
		}

		json.NewEncoder(res).Encode(response)
	}
	defer rows.Close()

}
