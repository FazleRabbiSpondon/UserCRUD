package user

import (
	"encoding/json"
	"net/http"
)

func doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jd := json.NewDecoder(r.Body)

	aUser := &User{}
	err := jd.Decode(aUser)
	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lock.Lock()
	nextUserID++
	aUser.Id = nextUserID
	db = append(db, aUser)
	lock.Unlock()

	respUser := User{Id:aUser.Id, Username: aUser.Username}
	je := json.NewEncoder(w)
	je.Encode(respUser)
}
