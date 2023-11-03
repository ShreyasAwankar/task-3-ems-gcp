package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Task3/models"
)

func SearchEmployees(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	defer mu.Unlock()

	collection := client.Collection("ems")

	// Get all documents in the collection
	docs, err := collection.Documents(ctx).GetAll()
	if err != nil {
		logger.Errorf("Error fetching all documents: %v", err)
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Failed to retrieve documents", http.StatusInternalServerError)
		return
	}

	var employees []models.Employee

	query := r.URL.Query()

	// Get the search criteria from query parameters
	firstName := query.Get("firstName")
	lastName := query.Get("lastName")
	email := query.Get("email")
	role := query.Get("role")

	// Iterate through the documents and retrieve their data
	for _, doc := range docs {
		var emp models.Employee
		if err := doc.DataTo(&emp); err != nil {
			logger.Errorf("Error parsing employee data: %v", err)
			continue
		}

		if (firstName == "" || emp.FirstName == firstName) &&
			(lastName == "" || emp.LastName == lastName) &&
			(email == "" || emp.Email == email) &&
			(role == "" || emp.Role == role) {
			employees = append(employees, emp)
		}

	}

	if len(employees) == 0 {
		logger.Infof("No employee found with the provided search criteria")
	} else {
		logger.Infof("Employees fetched with the provided search criteria")
	}

	// Serialize the results to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
