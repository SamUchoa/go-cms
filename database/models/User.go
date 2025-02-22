package models

import (
	"database/sql"
	"jsfree/cms/database"
)

type User struct{
    Name string
    Email string
    Password string
    Role string
}

func (user *User) Create() *sql.Row {
    row := database.Handler.Db.QueryRow(`INSERT INTO user (name, email, password, role)
    VALUES (?, ?, ?, ?);`, user.Name, user.Email, user.Password, user.Role)

    return row
}
