package storage

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	//initializing the driver
	_ "github.com/go-sql-driver/mysql"
)

var dbConnection *sql.DB

type databaseStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Database string `json:"database"`
}

func init() {
	var databaseObject = loadCredentials()
	_ = databaseObject
	db, err := sql.Open("mysql",
		databaseObject.Username+":"+databaseObject.Password+"@tcp("+databaseObject.Host+")/"+databaseObject.Database)
	if err != nil {
		log.Println(err)
		log.Fatal("Database connection failed")
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		log.Fatal("Database ping failed")
	}
	dbConnection = db
}

/*
GetDb returns a db connection
*/
func GetDb() *sql.DB {
	return dbConnection
}

func loadCredentials() *databaseStruct {
	jsonFile, err := os.Open("db.settings") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var database = new(databaseStruct)
	json.Unmarshal(byteValue, database)
	jsonFile.Close()
	return database
}
