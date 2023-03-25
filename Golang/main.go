package main

import (
	"fmt"
	"golangapi/router"
	"net/http"
)

func main() {

	r := router.Router()
	fmt.Println("Server is getting started....")
	http.ListenAndServe(":3000", r)
	fmt.Println("Listening at port 3000....")
}
