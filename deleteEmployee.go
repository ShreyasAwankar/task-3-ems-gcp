package controllers

import (
	"net/http"
	"strconv"
)

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Fetch employee ID from the query parameters
	empID := r.URL.Query().Get("empid")

	employeeId, err := strconv.Atoi(empID)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Invalid Employee ID", http.StatusBadRequest)
		logger.Errorf("Invalid ID provided for employee: %v", err)
		return
	}

	// Find the document to delete based on the employee ID
	query := client.Collection("ems").Where("ID", "==", employeeId)
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		logger.Errorf("Error finding document: %v", err)
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Failed to retrieve employee", http.StatusInternalServerError)
		return
	}

	if len(docs) != 1 {
		logger.Errorf("Employee not found")
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	// Get the reference to the document to be deleted
	docRef := docs[0].Ref

	// Perform the deletion
	_, deleteErr := docRef.Delete(ctx)
	if deleteErr != nil {
		logger.Errorf("Failed to delete employee: %v", deleteErr)
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Failed to delete employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	logger.Infof("Employee with employee ID %d deleted successfully", employeeId)
}
