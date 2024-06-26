package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func doGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	je := json.NewEncoder(w)
	je.Encode(db)
	fmt.Fprintf(w, "Got HTTP method '%v' to %v\n", r.Method, r.URL)
}
