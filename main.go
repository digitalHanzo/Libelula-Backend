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

	// bugs = append(bugs, models.Bug{ID: "12", Name: "Libelula", Description: "Best bug ever", Family: &models.Family{ID: "3", Name: "Example", SubFamily: "Example"}})
	// bugs = append(bugs, models.Bug{ID: "10", Name: "Escorpion", Description: "Best bug ever"})

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

	//Start server
	println("Server mounted on http://localhost:420")
	log.Fatal(http.ListenAndServe(":0420", router))

}
