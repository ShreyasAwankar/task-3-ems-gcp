package controllers

import (
	"encoding/json"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/Task3/models"
	"github.com/Task3/validations"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	defer mu.Unlock()

	var employee models.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		logger.Errorf("Invalid JSON input for type Employee during controllers.CreateEmployee function call")
		w.Header().Set("Content-Type", "application/json")
		return
	}

	err1 := validations.V.Struct(employee)

	if err1 != nil {
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Invalid input for employee deatails\n-First name must contain alphabets or spaces only.\n-First name must contain alphabets or spaces only.\n-Email id must be valid eg. abc@example.com.\n-Password must be atleast 6 charecters long.\n-Phone no. should be valid.\n-Role must be either of - 'admin', 'developer', 'manager', 'tester'. (case sensetive)\n-Salary must be a number.\n-Birthdate should be in yyyy-mm-dd format.", http.StatusUnprocessableEntity)
		logger.Errorf("Invalid employee data input : occured while validating employee fields during controllers.CreateEmployee function call %v", err1)
		w.Header().Set("Content-Type", "application/json")
		return
	}

	// Query to get the maximum employee ID
	iter := client.Collection("ems").OrderBy("ID", firestore.Desc).Limit(1).Documents(ctx)

	var lastEmployee models.Employee

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		doc.DataTo(&lastEmployee)
		break
	}

	// Generate new employee ID
	EmpId := lastEmployee.ID + 1
	employee.ID = EmpId

	// Save employee data to Firestore
	_, _, err = client.Collection("ems").Add(ctx, employee)
	if err != nil {
		logger.Errorf("Failed to create employee: %v", err)
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
		return
	}

	// Setting headers and staus codes
	// newEmployeeURL := fmt.Sprintf("http://localhost:5000/employees/%v", EmpId)
	// w.Header().Set("Location", newEmployeeURL)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)
	logger.Infof("Employee created successfully with emp id %v", EmpId)
}
