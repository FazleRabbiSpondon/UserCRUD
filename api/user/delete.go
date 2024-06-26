package user

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func doDelete(w http.ResponseWriter, r *http.Request) {

	fields := strings.Split(r.URL.String(), "/")
	id, err := strconv.ParseUint(fields[len(fields)-1], 10, 64)
	fmt.Println(fields)
	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Request to delete user %v", id)

	lock.Lock()
	var tmp = []*User{}
	for _, u := range db {
		if id == u.Id {
			continue
		}
		tmp = append(tmp, u)
	}

	db = tmp
	lock.Unlock()

}
