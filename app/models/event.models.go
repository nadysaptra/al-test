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

func FetchAllEvents() (Response, error) {
	var events []Event
	var res Response

	db := config.GetDBInstance()

	if result := db.Find(&events); result.Error != nil {
		fmt.Print("error FetchAllEvents")
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

func FetchEvent(event *Event) (Response, error) {
	var res Response

	db := config.GetDBInstance()

	if result := db.First(&event); result.Error != nil {
		fmt.Print("error FetchEvent")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = event

	return res, nil
}

func CreateAEvent(event *Event) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Create(&event); result.Error != nil {
		fmt.Print("error CreateAEvent")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = event

	return res, nil
}

func UpdateEvent(event *Event) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Where("id = ?", event.ID).Updates(Event{Name: event.Name, Date: event.Date, TicketPrize: event.TicketPrize}); result.Error != nil {
		fmt.Print("error UpdateEvent")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error update event record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = event

	return res, nil
}

func DeleteEvent(event *Event) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Delete(&event); result.Error != nil {
		fmt.Print("error DeleteEvent")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error update event record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"

	return res, nil
}
