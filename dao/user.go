package dao

import (
	"gocode/models"
	"log"
)

func GetUserNameById(userId int) string {
	var userName string
	row := DB.QueryRow("select  user_name from blog_user where uid=?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	_ = row.Scan(&userName)
	return userName
}

func GetUser(username, passwd string) *models.User {
	row := DB.QueryRow("select *  from blog_user where user_name=? and passwd= ?", username, passwd)
	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}
	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}

func UserIsexit(username, passwd string) *models.User {
	row := DB.QueryRow("select * from blog_user where user_name=?", username)
	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd)
	if err == nil {
		return nil
	}
	_, err = DB.Exec("insert into blog_user(user_name,passwd) values(?,?) ", username, passwd)
	err = row.Scan(&user.Uid, &user.UserName, &user.Passwd)
	if err != nil {
		log.Println(err)
	}
	return user
}
