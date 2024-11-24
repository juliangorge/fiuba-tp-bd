package employees

import "time"

type Employee struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Position   string    `json:"position"`
	Department string    `json:"department"`
	HireDate   time.Time `json:"hire_date"`
	Salary     int       `json:"salary"`
}
