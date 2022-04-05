package models

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"azure-labs-test/config"
)

type Payment struct {
	gorm.Model
	Status      string `json:"status"` // pending | payed | cancelled
	TicketRefer int
	Ticket      Ticket `gorm:"foreignKey:TicketRefer"`
}

func FethAllPayments() (Response, error) {
	var payments []Payment
	var res Response

	db := config.GetDBInstance()

	if result := db.Find(&payments); result.Error != nil {
		fmt.Print("error FethAllPayments")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = payments

	return res, nil
}

func CreatePayment(payment *Payment) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Create(&payment); result.Error != nil {
		fmt.Print("error CreatePayment")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = payment

	return res, nil
}

func UpdatePayment(payment *Payment) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Save(&payment); result.Error != nil {
		fmt.Print("error UpdatePayment")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error update payment record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = payment

	return res, nil
}
