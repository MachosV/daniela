package storetodb

import (
	"errorhandlers"
	"errors"
	"models"
	"storage"
)

/*StoreUser handles
saving the user object to the db
*/
func StoreUser(user *models.User) error {
	var db = storage.GetDb()
	stmt, err := db.Prepare("INSERT INTO users (" +
		"name," +
		"surname," +
		"nickname," +
		"isEmployer" +
		") VALUES (?,?,?,?)")
	if err != nil {
		errorhandlers.LogError(err, "Error storing user")
		return errors.New("Prepare statement store user failed")
	}
	_, err = stmt.Exec(
		user.Name,
		user.Surname,
		user.Nickname,
		user.IsEmployer)
	stmt.Close()
	if err != nil {
		errorhandlers.LogError(err, "Error storing user")
		return errors.New("Exec statement store user failed")
	}
	return nil
}
