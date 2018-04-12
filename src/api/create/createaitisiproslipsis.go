package create

import (
	"encoding/json"
	"errorhandlers"
	"fmt"
	"models"
	"net/http"
	"storage/storetodb"
)

func AitisiProslipsis(w http.ResponseWriter, r *http.Request) {
	var decoder = json.NewDecoder(r.Body)
	var aitisi = new(models.AitisiProslipsis)
	err := decoder.Decode(aitisi)
	if err != nil {
		errorhandlers.LogError(err, ": "+r.URL.Path)
		fmt.Fprintf(w, "Create user failed")
		return
	}
	err = storetodb.StoreAitisiProslipsis(aitisi)
	if err != nil {
		errorhandlers.LogError(err, ": "+r.URL.Path)
		fmt.Fprintf(w, "Store aitisi proslipsis failed")
		return
	}
	fmt.Fprint(w, "OK")
}
