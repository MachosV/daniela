package retrieve

import (
	"encoding/json"
	"errorhandlers"
	"fmt"
	"net/http"
	"storage/retrievefromdb"
	"strconv"
)

func User(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil {
		errorhandlers.LogError(err, "Error retrieving user")
		fmt.Fprint(w, "Invalid id")
	}
	var user = retrievefromdb.RetrieveUser(int(id))
	data, err := json.MarshalIndent(user, "", "\t")
	if err != nil {
		errorhandlers.LogError(err, "Error retrieving user")
		fmt.Fprint(w, "Marshalling error")
		return
	}
	fmt.Fprint(w, string(data))
}
