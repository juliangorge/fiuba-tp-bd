package employees

import (
	"bdd-back/utils"
	"encoding/json"
	"log"
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
	ctx := r.Context()
	employees, err := c.storage.GetAll()
	if err != nil {
		utils.HandleHttpError(ctx, w, "Failed to Get All Employees", http.StatusInternalServerError, err)
		return
	}

	bytes, err := json.Marshal(employees)
	if err != nil {
		utils.HandleHttpError(ctx, w, "Failed to Marshal Employees", http.StatusInternalServerError, err)
		return
	}

	w.Write(bytes)
}

func (c *EmployeeController) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	stringID := r.PathValue("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		utils.HandleHttpError(ctx, w, "Invalid ID", http.StatusBadRequest, err)
		return
	}

	employee, err := c.storage.GetByID(id)
	if err != nil {
		log.Printf("ERROR: Failed to Get Employee by ID: %v", err)
		if err == ErrEmployeeNotFound {
			utils.HandleHttpError(ctx, w, "Employee not found", http.StatusNotFound, err)
		} else {
			utils.HandleHttpError(ctx, w, "Failed to Get Employee by ID", http.StatusInternalServerError, err)
		}
		return
	}

	utils.HttpJsonResponse(ctx, w, http.StatusOK, employee)
}

func (c *EmployeeController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var employee Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = employee.Validate()
	if err != nil {
		utils.HandleHttpError(ctx, w, "Invalid Employee", http.StatusBadRequest, err)
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
	ctx := r.Context()
	stringID := r.PathValue("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		utils.HandleHttpError(ctx, w, "Invalid ID", http.StatusBadRequest, err)
		return
	}

	var employee Employee
	err = json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		utils.HandleHttpError(ctx, w, "Invalid request body", http.StatusBadRequest, err)
		return
	}

	err = employee.Validate()
	if err != nil {
		utils.HandleHttpError(ctx, w, "Invalid Employee", http.StatusBadRequest, err)
		return
	}

	_, err = c.storage.GetByID(id)
	if err != nil {
		utils.HandleHttpError(ctx, w, "Employee not found", http.StatusNotFound, err)
		return
	}

	employee.ID = id
	err = c.storage.Update(&employee)
	if err != nil {
		utils.HandleHttpError(ctx, w, "Failed to Update Employee", http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *EmployeeController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	stringID := r.PathValue("id")
	id, err := strconv.Atoi(stringID)
	if err != nil {
		utils.HandleHttpError(ctx, w, "Invalid ID", http.StatusBadRequest, err)
		return
	}

	err = c.storage.Delete(id)
	if err != nil {
		utils.HandleHttpError(ctx, w, "Failed to Delete Employee", http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
