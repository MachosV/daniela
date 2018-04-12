package updatetodb

import (
	"errorhandlers"
	"errors"
	"storage"
)

/*StoreUser handles
saving the user object to the db
*/
func UpdateUserNickname(id int, newNickname string) error {
	var db = storage.GetDb()
	stmt, err := db.Prepare("UPDATE users " +
		"SET nickname = ? " +
		"WHERE id = ?;")
	if err != nil {
		errorhandlers.LogError(err, "Error updating user nickname")
		return errors.New("Prepare statement update user nickname failed")
	}
	_, err = stmt.Exec(
		newNickname,
		id)
	if err != nil {
		errorhandlers.LogError(err, "Error updating user nickname")
		return errors.New("Exec statement update user nickname failed")
	}
	return nil
}
