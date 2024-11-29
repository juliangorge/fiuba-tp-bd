package employee_tags

type EmployeeTag struct {
	EmployeeID int      `bson:"employee_id" json:"employee_id"`
	Tags       []string `bson:"tags" json:"tags"`
}
