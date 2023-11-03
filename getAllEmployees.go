package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Task3/models"
	"google.golang.org/api/iterator"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var employees []models.Employee

	iter := client.Collection("ems").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			logger.Errorf("Failed to iterate through employees: %v", err)
			http.Error(w, "Failed to retrieve employees", http.StatusInternalServerError)
			return
		}

		var employee models.Employee

		err = doc.DataTo(&employee)
		if err != nil {
			logger.Errorf("Failed to parse employee data: %v", err)
			w.Header().Set("Content-Type", "text/plain")
			http.Error(w, "Failed to retrieve employees", http.StatusInternalServerError)
			return
		}

		employees = append(employees, employee)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
	logger.Info("Employees fetched successfully")
}
