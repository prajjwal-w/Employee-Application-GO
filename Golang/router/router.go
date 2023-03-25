package router

import (
	"golangapi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/addemployee", controller.AddaEmployee).Methods("POST")
	router.HandleFunc("/api/getallemployee", controller.GetAllEmployee).Methods("GET")
	router.HandleFunc("/api/getemployeebyid/{id}", controller.GetEmployeeById).Methods("GET")
	router.HandleFunc("/api/updateemployee", controller.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/api/deleteemployeebyid/{id}", controller.DeleteEmployee).Methods("DELETE")

	return router
}
