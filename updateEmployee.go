package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Task3/models"
	"github.com/Task3/validations"
)

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
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

	var employee models.Employee
	err = json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/plain")
		logger.Errorf("Invalid JSON input for employee: %v", err)
		return
	}

	employee.ID = employeeId

	err = validations.V.Struct(employee)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Invalid input for employee details\n-First name must contain alphabets or spaces only.\n-First name must contain alphabets or spaces only.\n-Email id must be valid eg. abc@example.com.\n-Password must be atleast 6 charecters long.\n-Phone no. should be valid.\n-Role must be either of - 'admin', 'developer', 'manager', 'tester'. (case sensetive)\n-Salary must be a number.\n-Birthdate should be in yyyy-mm-dd format.", http.StatusUnprocessableEntity)
		logger.Errorf("Invalid employee data input: %v", err)
		return
	}

	// Find the document based on a field other than the document ID
	query := client.Collection("ems").Where("ID", "==", employeeId)
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		logger.Errorf("Error finding document: %v", err)
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Failed to retrieve employee", http.StatusInternalServerError)
		return
	}

	if len(docs) == 0 {
		logger.Infof("Employee not found")
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Employee not found ", http.StatusNotFound)
		return
	}

	docRef := docs[0].Ref

	// Perform the update
	_, updateErr := docRef.Set(ctx, employee)
	if updateErr != nil {
		logger.Errorf("Failed to update employee: %v", updateErr)
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Failed to update employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
	logger.Infof("Employee with employee ID %d updated successfully", employeeId)
}
