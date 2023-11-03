package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Task3/models"
)

func GetEmployee(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	defer mu.Unlock()

	// Fetch employee ID from the query parameters
	empID := r.URL.Query().Get("empid")

	employeeId, err := strconv.Atoi(empID)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Invalid Employee ID", http.StatusBadRequest)
		logger.Errorf("Invalid id was employeeId was provided on GetEmployee function call.")
		return
	}

	iter := client.Collection("ems").Where("ID", "==", employeeId).Documents(ctx)

	defer iter.Stop()

	doc, err := iter.Next()
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		logger.Errorf("Employee not found in db: %v", err)
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	var employee models.Employee
	err = doc.DataTo(&employee)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		logger.Errorf("Employee not found in db: %v", err)
		http.Error(w, "Failed to retrieve employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
	logger.Infof("Employee fetched with id %v", employeeId)
}
