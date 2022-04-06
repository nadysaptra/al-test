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

func FetchAllUsers() (Response, error) {
	var users []User
	var res Response

	db := config.GetDBInstance()

	if result := db.Find(&users); result.Error != nil {
		fmt.Print("error FetchAllUsers")
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

func FetchUser(user *User) (Response, error) {
	var res Response

	db := config.GetDBInstance()

	if result := db.First(&user); result.Error != nil {
		fmt.Print("error FetchUser")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = user

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

func UpdateUser(user *User) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Where("id = ?", user.ID).Updates(User{Name: user.Name, Email: user.Email}); result.Error != nil {
		fmt.Print("error UpdateUser")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error update user record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = user

	return res, nil
}

func DeleteUser(user *User) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Delete(&user); result.Error != nil {
		fmt.Print("error DeleteUser")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error update user record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"

	return res, nil
}
