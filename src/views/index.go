package views

import (
	"fmt"
	"net/http"
)

/*
IndexView is the main web view of Daniela
*/
func IndexView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
