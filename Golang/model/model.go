package model

type EmployeeDetails struct {
	Id    int    `json:"empId"`
	Fname string `json:"firstname"`
	Lname string `json:"lastname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
