package dao

import(
	"model"
	"utils"
)

func CheckUser(username string, password string) (*model.User, err){
		sql := "select * from users where"
	  row :=utils.Db.QueryRow(sql, username, password)
	  user := &model.User{}
	  row.Scan(&user.ID,&user.Username, &user.Password)
} 


func SaveUser(){
	
}