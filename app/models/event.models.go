package models

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"azure-labs-test/config"
)

type Event struct {
	gorm.Model
	Name        string `json:"name"`
	TicketPrize int    `json:"ticket_prize"`
	Date        string `json:"date"`
}

func FethAllEvents() (Response, error) {
	var events []Event
	var res Response

	db := config.GetDBInstance()

	if result := db.Find(&events); result.Error != nil {
		fmt.Print("error FethAllEvents")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = events

	return res, nil
}

func CreateEvent(user *Event) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Create(&user); result.Error != nil {
		fmt.Print("error CreateEvent")
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
