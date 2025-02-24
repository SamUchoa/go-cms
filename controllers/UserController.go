package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"jsfree/cms/database/models"
	"log"
	"net/http"
)


func Create(res http.ResponseWriter, req *http.Request){
    plainBody, err := io.ReadAll(req.Body)

    if err != nil {
        log.Fatal(err)
    }

    newUser := models.User{}

    fmt.Println(json.Unmarshal(plainBody, &newUser))

    if newUser.Name == "" || newUser.Email == "" || newUser.Password == "" || newUser.Role == "" { 
        res.WriteHeader(400)
        fmt.Fprintf(res, "preencha todos os campos")
    }

    hash := sha256.Sum256([]byte(newUser.Password))
    newUser.Password = fmt.Sprintf("%x", hash)
    fmt.Println(newUser.Password)

    if err != nil {
        log.Fatal(err)
    }
//    if len(newUser.Password) < 8

    row := newUser.Create()
    
    if err := row.Err(); err != nil {
        log.Fatal(err)
    }
    fmt.Fprintf(res, "created: %+v", newUser)
}
