package main

import (
	"fmt"
	"log"
	"net/http"

	controllers "github.com/Task3"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Employee Management System APIs")
	// r := router.Router()
	fmt.Printf("Listening to port 4000...\n\n")
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/ems-api/v1").Subrouter()

	subRouter.HandleFunc("/employees", controllers.GetAllEmployees).Methods("GET")
	subRouter.HandleFunc("/employees/search", controllers.SearchEmployees).Methods("GET")
	subRouter.HandleFunc("/employees/employee", controllers.GetEmployee).Methods("GET")
	subRouter.HandleFunc("/employees", controllers.CreateEmployee).Methods("POST")
	subRouter.HandleFunc("/employees", controllers.UpdateEmployee).Methods("PUT")
	subRouter.HandleFunc("/employees/employee", controllers.DeleteEmployee).Methods("DELETE")

	// Create a new CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allowing all origins to make requests.(For avoiding any client CORS conflicts for API testing)
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	// Wrap your subRouter with the CORS middleware
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":4000", handler))
}
