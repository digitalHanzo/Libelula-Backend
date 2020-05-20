package main

import (
	"./controllers"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//ENRUTADOR
	router := mux.NewRouter()

	//endpoints
	router.HandleFunc("/bugs", controllers.GetBugsEndPoint).Methods("GET")
	router.HandleFunc("/bug/{id}", controllers.GetBugEndPoint).Methods("GET")
	//add
	router.HandleFunc("/bug", controllers.CreateBugEndPoint).Methods("POST")
	//del
	router.HandleFunc("/bug/{id}", controllers.DeleteBugEndPoint).Methods("DELETE")
	//edit
	router.HandleFunc("/bug/{id}", controllers.EditBugEndPoint).Methods("PUT")

	router.HandleFunc("/families", controllers.GetFamiliesEndPoint).Methods("GET")
	router.HandleFunc("/family/{id}", controllers.GetFamilyEndPoint).Methods("GET")

	//Start server
	println("Server mounted on http://localhost:420")
	log.Fatal(http.ListenAndServe(":0420", router))

}
