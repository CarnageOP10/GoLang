package main

import (
	"fmt"
	"log"
	"mongoAPI/router"
	"net/http"
)

func main() {
	fmt.Println("Server is getting started")

	r:= router.Router()
	log.Fatal(http.ListenAndServe(":4500", r))
	fmt.Println("listening......")
}
