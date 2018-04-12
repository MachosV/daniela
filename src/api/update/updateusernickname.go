package update

import (
	"errorhandlers"
	"fmt"
	"net/http"
	"storage/updatetodb"
	"strconv"
)

func UpdateUserNickname(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
	newNickname := r.PostFormValue("nickname")
	if err != nil {
		errorhandlers.LogError(err, "Error retrieving user")
		fmt.Fprint(w, "Invalid id")
		return
	}
	err = updatetodb.UpdateUserNickname(int(id), newNickname)
	if err != nil {
		errorhandlers.LogError(err, "Error retrieving user")
		fmt.Fprint(w, "Invalid id")
		return
	}
	fmt.Fprint(w, "OK")
}
