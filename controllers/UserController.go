package controllers

import (
	"fmt"
	"jsfree/cms/database/models"
	"net/http"
)


func Create(res http.ResponseWriter, req *http.Request){
    newUser := models.User{
        Name: "samuel",
        Email: "amuel@exemplo.com",
        Password: "senha123",
        Role: "otário",
    }
    result := newUser.Create()
    fmt.Println(result)
    fmt.Fprintf(res, "funcionou")
}
