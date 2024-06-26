package main

import (
	"api/user/api/user"
	"log"
	"net/http"
)

func main() {
	// register RESTful endpoint handler for '/users/'
	http.Handle("/users/", &user.UserAPI{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
