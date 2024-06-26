package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func doPut(w http.ResponseWriter, r *http.Request) {

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

	

	w.Header().Set("Content-Type", "application/json")
	jd := json.NewDecoder(r.Body)

	aUser := &User{}
	err2 := jd.Decode(aUser)
	if nil != err2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lock.Lock()
//	nextUserID++
	aUser.Id = id
	db = append(db, aUser)
	lock.Unlock()

	respUser := User{Id:aUser.Id, Username: aUser.Username}
	je := json.NewEncoder(w)
	je.Encode(respUser)

}
