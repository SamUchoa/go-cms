package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type DbHandler struct {
    Db *sql.DB
}

var Handler DbHandler;

func (dbHandler *DbHandler) Connect(dbUser, dbPasswd, dbName string) {
    cfg := mysql.Config{
        User: dbUser,
        Passwd: dbPasswd,
        DBName: dbName,
        Net: "tcp",
        Addr: "localhost:3306",
        AllowNativePasswords: true,
    }

    var err error
    dbHandler.Db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := dbHandler.Db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected to DB")
}
