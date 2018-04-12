package retrievefromdb

import (
	"errorhandlers"
	"log"
	"models"
	"storage"
)

func RetrieveUser(id int) *models.User {
	db := storage.GetDb()
	res, err := db.Query("SELECT * FROM users WHERE id = ? ;", id)
	if err != nil {
		errorhandlers.LogError(err, "Error retrieving user")
		return nil
	}
	var user models.User
	if res.Next() {
		defer res.Close()
		err = res.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Nickname,
			&user.Level,
			&user.Progress,
			&user.IsEmployer,
			&user.Stars)
		if err != nil {
			log.Println("error scanning")
			log.Println(err)
			return nil
		}
	} else {
		log.Println("No results returned")
		return nil
	}
	return &user
}
