package controller

import (
	"net/http"
	"encoding/json"
	"book-store/dao"
)

func Login(w http.ResponseWriter, r *http.Request){
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user, _ := dao.CheckUser()
	if user.ID > 0{

	} else {
       
	}
}