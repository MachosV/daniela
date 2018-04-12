package delete

import (
	"errorhandlers"
	"fmt"
	"net/http"
	"storage/deletefromdb"
	"strconv"
)

func User(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	if err != nil {
		errorhandlers.LogError(err, "Error deleting user")
		fmt.Fprint(w, "Error deleting user")
		return
	}
	var error = deletefromdb.DeleteUser(int(id))
	if error != nil {
		fmt.Fprint(w, "Error deleting user")
		return
	}
	fmt.Fprint(w, "OK")
}
