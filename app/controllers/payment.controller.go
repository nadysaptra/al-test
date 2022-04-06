package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"azure-labs-test/app/models"
)

func FetchAllPayments(c *fiber.Ctx) error {
	result, _ := models.FethAllPayments()
	return c.Status(result.Status).JSON(result)
}

func FetchPayment(c *fiber.Ctx) error {
	var payment models.Payment
	convertToUint32, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "cannot parse ID to uint "})
	} else {
		payment.ID = uint(convertToUint32)
	}
	result, _ := models.FetchPayment(&payment)
	return c.Status(result.Status).JSON(result)
}

func UpdatePayment(c *fiber.Ctx) error {
	var payment models.Payment
	payment.Status = c.FormValue("status")
	if payment.Status == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "status is required"})
	}
	if c.Params("id") == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "id is required"})
	}
	convertToUint32, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "cannot parse ID to uint "})
	} else {
		payment.ID = uint(convertToUint32)
	}

	result, _ := models.UpdatePayment(&payment)
	return c.Status(result.Status).JSON(result)
}
