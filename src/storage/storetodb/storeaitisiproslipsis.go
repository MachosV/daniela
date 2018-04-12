package storetodb

import (
	"errorhandlers"
	"errors"
	"models"
	"storage"
)

func StoreAitisiProslipsis(aitisi *models.AitisiProslipsis) error {
	var db = storage.GetDb()
	stmt, err := db.Prepare("INSERT INTO aitiseis_proslipsis (" +
		"idEmployee," +
		"idEmployer," +
		") VALUES (?,?)")
	if err != nil {
		errorhandlers.LogError(err, "Error storing aitisi proslipsis")
		return errors.New("Prepare statement store aitisi proslipsis failed")
	}
	_, err = stmt.Exec(
		aitisi.IDEmployee,
		aitisi.IDEmployer)
	stmt.Close()
	if err != nil {
		errorhandlers.LogError(err, "Error storing aitisi proslipsis")
		return errors.New("Exec statement store aitisi proslipsis failed")
	}
	return nil
}
