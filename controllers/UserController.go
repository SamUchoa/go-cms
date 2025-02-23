package controllers

import (
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
//    if len(newUser.Password) < 8

fmt.Fprintf(res, "created: %s", plainBody)
}
