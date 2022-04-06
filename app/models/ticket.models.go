package models

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"azure-labs-test/config"
)

type Ticket struct {
	gorm.Model
	EventRefer int    `json:"event_refer"`
	Event      Event  `gorm:"foreignKey:EventRefer"`
	UserRefer  int    `json:"user_refer"`
	User       User   `gorm:"foreignKey:UserRefer"`
	Status     string `json:"status"` // ordered | attended | forfeit
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

func FetchTicket(ticket *Ticket) (Response, error) {
	var res Response

	db := config.GetDBInstance()

	if result := db.First(&ticket); result.Error != nil {
		fmt.Print("error FetchTicket")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = ticket

	return res, nil
}

func FetchUserTicket(user *User) (Response, error) {
	var res Response
	var tickets []Ticket

	db := config.GetDBInstance()

	if result := db.Where("user_refer = ?", user.ID).Find(&tickets); result.Error != nil {
		fmt.Print("error FetchTicket")
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

func FetchEventTicket(event *Event) (Response, error) {
	var res Response
	var tickets []Ticket

	db := config.GetDBInstance()

	if result := db.Where("event_refer = ?", event.ID).Find(&tickets); result.Error != nil {
		fmt.Print("error FetchTicket")
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

func CreateTicket(ticket *Ticket) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Create(&ticket); result.Error != nil {
		fmt.Print("error CreateTicket")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res, result.Error
	}

	var payment Payment
	payment.Status = "pending"
	payment.TicketRefer = int(ticket.ID)
	if createPayment := db.Create(&payment); createPayment.Error != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "error create payment record"
		return res, createPayment.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = ticket

	return res, nil
}

func UpdateTicket(ticket *Ticket) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Where("id = ?", ticket.ID).Updates(Ticket{Status: ticket.Status}); result.Error != nil {
		fmt.Print("error UpdateTicket")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error update record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = ticket

	return res, nil
}
