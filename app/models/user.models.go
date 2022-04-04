package models

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"azure-labs-test/config"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FethAllUsers() (Response, error) {
	var users []User
	var res Response

	db := config.GetDBInstance()

	if result := db.Find(&users); result.Error != nil {
		fmt.Print("error FethAllUsers")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = users

	return res, nil
}

func CreateAUser(user *User) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Create(&user); result.Error != nil {
		fmt.Print("error CreateAUser")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = user

	return res, nil
}
