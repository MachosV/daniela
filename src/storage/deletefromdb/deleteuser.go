package deletefromdb

import (
	"errorhandlers"
	"errors"
	"storage"
)

func DeleteUser(id int) error {
	var db = storage.GetDb()
	stmt, err := db.Prepare("DELETE from users WHERE id = ?;")
	if err != nil {
		errorhandlers.LogError(err, "Error deleting user")
		return errors.New("Prepare statement delete user failed")
	}
	_, err = stmt.Exec(id)
	stmt.Close()
	if err != nil {
		errorhandlers.LogError(err, "Error deleting user")
		return errors.New("Exec statement delete user failed")
	}
	return nil

}
