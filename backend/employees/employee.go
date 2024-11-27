package employees

import (
	"errors"
	"time"
)

var (
	ErrEmptyFirstName  = errors.New("first name is required")
	ErrEmptyLastName   = errors.New("last name is required")
	ErrEmptyPosition   = errors.New("position is required")
	ErrEmptyDepartment = errors.New("department is required")
	ErrEmptyHireDate   = errors.New("hire date is required")
	ErrInvalidSalary   = errors.New("salary must be greater than 0")
)

type Employee struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Position   string    `json:"position"`
	Department string    `json:"department"`
	HireDate   time.Time `json:"hire_date"`
	Salary     int       `json:"salary"`
}

func (e *Employee) Validate() error {
	if e.FirstName == "" {
		return ErrEmptyFirstName
	}
	if e.LastName == "" {
		return ErrEmptyLastName
	}
	if e.Position == "" {
		return ErrEmptyPosition
	}
	if e.Department == "" {
		return ErrEmptyDepartment
	}
	if e.HireDate.IsZero() {
		return ErrEmptyHireDate
	}
	if e.Salary <= 0 {
		return ErrInvalidSalary
	}

	return nil
}
