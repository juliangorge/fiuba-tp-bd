package employees

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type EmployeeStorage interface {
	GetAll() ([]Employee, error)
	GetByID(id int) (*Employee, error)
	Create(e *Employee) error
	Update(e *Employee) error
	Delete(id int) error
}

type EmployeeController struct {
	storage EmployeeStorage
}

func NewEmployeeController(storage EmployeeStorage) *EmployeeController {
	return &EmployeeController{storage: storage}
}

func (c *EmployeeController) GetAll(w http.ResponseWriter, r *http.Request) {
	employees, err := c.storage.GetAll()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(employees)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (c *EmployeeController) GetByID(w http.ResponseWriter, r *http.Request) {
	stringID := r.PathValue("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	employee, err := c.storage.GetByID(id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(employee)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (c *EmployeeController) Create(w http.ResponseWriter, r *http.Request) {
	var employee Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = c.storage.Create(&employee)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *EmployeeController) Update(w http.ResponseWriter, r *http.Request) {
	stringID := r.PathValue("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var employee Employee
	err = json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	employee.ID = id
	err = c.storage.Update(&employee)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *EmployeeController) Delete(w http.ResponseWriter, r *http.Request) {
	stringID := r.PathValue("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.storage.Delete(id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
