package employees

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

var mockEmployees = []Employee{
	{
		ID:         0,
		FirstName:  "John",
		LastName:   "Doe",
		Position:   "Software Engineer",
		Department: "Engineering",
		HireDate:   time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC),
		Salary:     100000,
	},
	{
		ID:         1,
		FirstName:  "Jane",
		LastName:   "Smith",
		Position:   "Product Manager",
		Department: "Product",
		HireDate:   time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		Salary:     120000,
	},
}

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
