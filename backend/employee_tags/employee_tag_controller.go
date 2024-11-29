package employee_tags

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"bdd-back/utils"
)

type EmployeeTagStorage interface {
	GetAllTagsByID(ctx context.Context, employeeID int) (EmployeeTag, error)
	InsertTag(ctx context.Context, employeeID int, tagToAdd string) error
	RemoveTag(ctx context.Context, employeeID int, tagToRemove string) error
}

type EmployeeTagController struct {
	storage EmployeeTagStorage
}

func NewEmployeeTagController(storage EmployeeTagStorage) *EmployeeTagController {
	return &EmployeeTagController{storage: storage}
}

func (c *EmployeeTagController) GetAllTagsByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	stringID := r.PathValue("employee_id")
	employeeID, err := strconv.Atoi(stringID)
	if err != nil {
		utils.HandleHttpError(ctx, w, "Invalid ID", http.StatusBadRequest, err)
		return
	}

	employee, err := c.storage.GetAllTagsByID(ctx, employeeID)
	if err != nil {
		log.Printf("ERROR: Failed to Get Employee by ID: %v", err)
		utils.HandleHttpError(ctx, w, "Failed to Get Employee Tags by EmployeeID", http.StatusInternalServerError, err)
		return
	}

	utils.HttpJsonResponse(ctx, w, http.StatusOK, employee)
}

func (c *EmployeeTagController) InsertTag(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	stringID := r.PathValue("employee_id")
	employeeID, err := strconv.Atoi(stringID)
	if err != nil {
		utils.HandleHttpError(ctx, w, "Invalid ID", http.StatusBadRequest, err)
		return
	}
	tagToAdd := r.PathValue("tag_name")
	if len(tagToAdd) == 0 {
		utils.HandleHttpError(ctx, w, "Invalid Tag name", http.StatusBadRequest, err)
		return
	}

	err = c.storage.InsertTag(ctx, employeeID, tagToAdd)
	if err != nil {
		utils.HandleHttpError(ctx, w, "Failed to Update Employee Tag", http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *EmployeeTagController) RemoveTag(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	stringID := r.PathValue("employee_id")
	employeeID, err := strconv.Atoi(stringID)
	if err != nil {
		utils.HandleHttpError(ctx, w, "Invalid ID", http.StatusBadRequest, err)
		return
	}
	tagToRemove := r.PathValue("tag_name")
	if len(tagToRemove) == 0 {
		utils.HandleHttpError(ctx, w, "Invalid Tag name", http.StatusBadRequest, err)
		return
	}

	err = c.storage.RemoveTag(ctx, employeeID, tagToRemove)
	if err != nil {
		utils.HandleHttpError(ctx, w, "Failed to Remove Employee Tag", http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
