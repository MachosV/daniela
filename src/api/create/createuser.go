package create

import (
	"encoding/json"
	"errorhandlers"
	"fmt"
	"models"
	"net/http"
	"storage/storetodb"
)

/*User reads a json object
and sends it over to storage for saving
*/
func User(w http.ResponseWriter, r *http.Request) {
	var decoder = json.NewDecoder(r.Body)
	var user = new(models.User)
	err := decoder.Decode(user)
	if err != nil {
		errorhandlers.LogError(err, ": "+r.URL.Path)
		fmt.Fprintf(w, "Create user failed")
		return
	}
	err = storetodb.StoreUser(user)
	if err != nil {
		errorhandlers.LogError(err, ": "+r.URL.Path)
		fmt.Fprintf(w, "Store user failed")
		return
	}
	fmt.Fprint(w, "OK")
}
