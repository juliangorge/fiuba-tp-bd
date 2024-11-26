package employees

import "database/sql"

type EmployeeSQLStorage struct {
	DB *sql.DB
}

func NewEmployeeSQLStorage(db *sql.DB) *EmployeeSQLStorage {
	return &EmployeeSQLStorage{DB: db}
}

func (s *EmployeeSQLStorage) GetAll() ([]Employee, error) {
	rows, err := s.DB.Query("SELECT id, first_name, last_name, position, department, hire_date, salary FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := []Employee{}
	for rows.Next() {
		var e Employee
		err := rows.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Position, &e.Department, &e.HireDate, &e.Salary)
		if err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}

	return employees, nil
}

func (s *EmployeeSQLStorage) GetByID(id int) (*Employee, error) {
	row := s.DB.QueryRow("SELECT id, first_name, last_name, position, department, hire_date, salary FROM employees WHERE id = $1", id)

	var e Employee
	err := row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Position, &e.Department, &e.HireDate, &e.Salary)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (s *EmployeeSQLStorage) Create(e *Employee) error {
	_, err := s.DB.Exec("INSERT INTO employees (first_name, last_name, position, department, hire_date, salary) VALUES ($1, $2, $3, $4, $5, $6)", e.FirstName, e.LastName, e.Position, e.Department, e.HireDate, e.Salary)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeSQLStorage) Update(e *Employee) error {
	_, err := s.DB.Exec("UPDATE employees SET first_name = $1, last_name = $2, position = $3, department = $4, hire_date = $5, salary = $6 WHERE id = $7", e.FirstName, e.LastName, e.Position, e.Department, e.HireDate, e.Salary, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmployeeSQLStorage) Delete(id int) error {
	_, err := s.DB.Exec("DELETE FROM employees WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
