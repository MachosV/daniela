package main

import (
	"api/create"
	"api/delete"
	"api/retrieve"
	"api/update"
	"log"
	"net/http"
	_ "storage"
	"views"
)

func main() {
	http.HandleFunc("/api/createiaitisiproslipsis", create.AitisiProslipsis)
	http.HandleFunc("/api/updateusernickname", update.UpdateUserNickname)
	http.HandleFunc("/api/deleteuser", delete.User)
	http.HandleFunc("/api/retrieveuser", retrieve.User)
	http.HandleFunc("/api/createuser", create.User)
	http.HandleFunc("/", views.IndexView)
	log.Println("Server is up and running.")
	http.ListenAndServe("localhost:8000", nil)
}
