package user

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type UserAPI struct{}

type User struct {
	Id       uint64 `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

var db = []*User{}
var nextUserID uint64
var lock sync.Mutex

func (u *UserAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		doGet(w, r)
	case http.MethodPost:
		doPost(w, r)
	case http.MethodDelete:
		doDelete(w, r)
	case http.MethodPut:
		doPut(w, r )
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsupported Method '%v' to %v\n", r.Method, r.URL)
		log.Printf("Unsupported method '%v' to %v\n", r.Method, r.URL)
	}
}
