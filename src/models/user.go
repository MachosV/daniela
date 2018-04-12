package models

/*
User struct represents (obviously)
a user
*/
type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Nickname   string `json:"nickname"`
	Level      int    `json:"level"`
	Progress   int    `json:"progress"`
	IsEmployer bool   `json:"isEmployer"`
	Stars      uint   `json:"stars"`
}
