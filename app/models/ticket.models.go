package models

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"azure-labs-test/config"
)

type Ticket struct {
	gorm.Model
	EventRefer int
	Event      Event `gorm:"foreignKey:EventRefer"`
	UserRefer  int
	User       User   `gorm:"foreignKey:UserRefer"`
	Status     string `json:"status"` // attended | forfeit
}

func FethAllTickets() (Response, error) {
	var tickets []Ticket
	var res Response

	db := config.GetDBInstance()

	if result := db.Find(&tickets); result.Error != nil {
		fmt.Print("error FethAllTickets")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = tickets

	return res, nil
}

func CreateATicket(ticket *Ticket) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Create(&ticket); result.Error != nil {
		fmt.Print("error CreateATicket")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = ticket

	return res, nil
}
