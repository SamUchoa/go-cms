package main

import (
	"fmt"
	"jsfree/cms/database"
	"jsfree/cms/controllers"
	"log"
	"net/http"
	"os"
)

func main() {
    dbUser := os.Getenv("DB_USER")
    dbPwd := os.Getenv("DB_PASSWD")
    dbName := os.Getenv("DB_NAME")

    database.Handler.Connect(dbUser, dbPwd, dbName)
    
    http.HandleFunc("/", controllers.Create)    

    fmt.Println("listening on http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
