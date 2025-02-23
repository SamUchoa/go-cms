package models

import (
	"database/sql"
	"jsfree/cms/database"
)

type User struct{
    Name string `json:name`
    Email string `json:email`
    Password string `json:password`
    Role string `json:role`
}

func (user *User) Create() *sql.Row {
    row := database.Handler.Db.QueryRow(`INSERT INTO user (name, email, password, role)
    VALUES (?, ?, ?, ?);`, user.Name, user.Email, user.Password, user.Role)

    return row
}
