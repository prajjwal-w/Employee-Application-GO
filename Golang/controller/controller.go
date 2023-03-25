package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"golangapi/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB

var server = "localhost"
var port = 1433
var user = "sa"
var password = "Prajjwal@18"
var database = "master"

func init() {

	//Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", server, user, password, port, database)

	var err error
	//create connection pool
	db, err = sql.Open("sqlserver", connString)

	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	err = db.PingContext(context.Background())

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected!!")
	fmt.Println()
}

// Insert a employee in Database
func addEmployee(emp model.EmployeeDetails) {

	stmt, err := db.Prepare("INSERT into master.dbo.Employee (Id,Fname,Lname,Email,Phone) VALUES(@Id,@Fname,@Lname,@Email,@Phone)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		context.Background(),
		sql.Named("Id", emp.Id),
		sql.Named("Fname", emp.Fname),
		sql.Named("Lname", emp.Lname),
		sql.Named("Email", emp.Email),
		sql.Named("Phone", emp.Phone))

	if err != nil {
		log.Fatal(err)
	}

	_, _ = result.RowsAffected()

	fmt.Println("Employee inserted Successfully")
}

// get all employees
func getAllEmployee() ([]model.EmployeeDetails, error) {

	var employeeDetails []model.EmployeeDetails

	employee, err := db.Query("SELECT * FROM master.dbo.Employee")

	if err != nil {
		return nil, fmt.Errorf("get all employee: %v", err)
	}

	defer employee.Close()
	//loop through the employee
	for employee.Next() {
		var e model.EmployeeDetails

		err := employee.Scan(&e.Id, &e.Fname, &e.Lname, &e.Email, &e.Phone)
		if err != nil {
			return nil, fmt.Errorf("get all employee: %v", err)
		}

		employeeDetails = append(employeeDetails, e)
	}
	return employeeDetails, nil

}

// Get employee by id
func getEmployeeById(id string) (model.EmployeeDetails, error) {
	var e model.EmployeeDetails
	intid, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("SELECT * FROM master.dbo.Employee WHERE Id= @Id")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(sql.Named("Id", intid))

	err = row.Scan(&e.Id, &e.Fname, &e.Lname, &e.Email, &e.Phone)
	if err != nil {
		return e, fmt.Errorf("get employee by id: %v", err)
	}
	return e, nil
}

// Update employee
func updateEmployee(emp model.EmployeeDetails) {

	result, err := db.ExecContext(
		context.Background(),
		"UPDATE master.dbo.Employee SET Fname=@Fname,Lname=@Lname,Email=@Email,Phone=@Phone Where Id=@Id",
		sql.Named("Id", emp.Id),
		sql.Named("Fname", emp.Fname),
		sql.Named("Lname", emp.Lname),
		sql.Named("Email", emp.Email),
		sql.Named("Phone", emp.Phone),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}

// Delete an employee by using Id
func deleteEmployee(id string) {
	intid, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	result, err := db.ExecContext(context.Background(), "DELETE FROM master.dbo.Employee WHERE Id = @ID", sql.Named("Id", intid))
	if err != nil {
		log.Fatal(err)
	}
	_, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Deleted ")

}

func AddaEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var emp model.EmployeeDetails
	_ = json.NewDecoder(r.Body).Decode(&emp)
	addEmployee(emp)
	json.NewEncoder(w).Encode(emp)
}

func GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allemployees, err := getAllEmployee()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(allemployees)

}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	emp, err := getEmployeeById(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(emp)
	json.NewEncoder(w).Encode(emp)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	var emp model.EmployeeDetails
	json.NewDecoder(r.Body).Decode(&emp)
	updateEmployee(emp)
	json.NewEncoder(w).Encode(emp)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteEmployee(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
