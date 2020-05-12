package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"../db"
	"../models"
)

//TEMPORARY DATA
var bugs []models.Bug

//DB connection
var database = db.Connection()

//GetBugsEndPoint - Get all data from PG
func GetBugsEndPoint(res http.ResponseWriter, req *http.Request) {
	queryStmt, err := database.Prepare("SELECT * FROM bugs")

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
		var description string
		var familyId int

		if err := rows.Scan(&id, &name, &description, &familyId); err != nil {
			log.Fatal(err)
		}
		bugs = append(bugs, models.Bug{ID: id, Name: name, Description: description, FamilyID: familyId})
	}
	json.NewEncoder(res).Encode(bugs)

	bugs = nil
}

//GetBugEndPoint - Get specific bug from PG
func GetBugEndPoint(res http.ResponseWriter, req *http.Request) {

	//obtener parametro por ruta
	params := mux.Vars(req)

	id := params["id"]

	queryStmt, err := database.Prepare(fmt.Sprintf("SELECT * FROM bugs WHERE id = '%s';", id))

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
		var description string
		var familyId int

		if err := rows.Scan(&id, &name, &description, &familyId); err != nil {
			log.Fatal(err)
		}
		bugs = append(bugs, models.Bug{ID: id, Name: name, Description: description, FamilyID: familyId})
	}
	json.NewEncoder(res).Encode(bugs)

	bugs = nil
}

//CreateBugEndPoint - Insert on DB
func CreateBugEndPoint(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	req.ParseForm()

	var insertBug models.Bug

	//Get data from body and storing in new object
	insertBug.Name = req.FormValue("name")
	insertBug.Description = req.FormValue("description")
	id, err := strconv.Atoi(req.FormValue("family_id"))
	insertBug.FamilyID = id

	queryStmt, err := database.Prepare(fmt.Sprintf("INSERT INTO bugs (name, description, family_id) VALUES ('%s', '%s', %d);", insertBug.Name, insertBug.Description, insertBug.FamilyID))

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

//EditBugEndPoint - Editing from db
func EditBugEndPoint(res http.ResponseWriter, req *http.Request) {

	//obtener parametro por ruta
	params := mux.Vars(req)
	id := params["id"]

	req.ParseForm()
	//Get data from body and storing in new object
	var editBug models.Bug
	editBug.Name = req.FormValue("name")
	editBug.Description = req.FormValue("description")
	familyid, err := strconv.Atoi(req.FormValue("family_id"))
	editBug.FamilyID = familyid

	queryStmt, err := database.Prepare(fmt.Sprintf("UPDATE bugs SET name = '%s', description = '%s' , family_id = %d WHERE id = %s;", editBug.Name, editBug.Description, editBug.FamilyID, id))

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

//DeleteBugEndPoint - Deleting from db
func DeleteBugEndPoint(res http.ResponseWriter, req *http.Request) {

	//obtener parametro por ruta
	params := mux.Vars(req)

	id := params["id"]

	queryStmt, err := database.Prepare(fmt.Sprintf("DELETE FROM bugs WHERE id = '%s';", id))

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
