package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"azure-labs-test/app/models"
)

// Payment godoc
// @Summary Show payment list.
// @Description get payment list.
// @Tags Payment
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /payments [get]
func FetchAllPayments(c *fiber.Ctx) error {
	result, _ := models.FethAllPayments()
	return c.Status(result.Status).JSON(result)
}

// Payment godoc
// @Summary Show payment detail.
// @Description get payment detail.
// @Tags Payment
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /payments/{id} [get]
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

// Event godoc
// @Summary Update event.
// @Description Update event.
// @Tags Event
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param status body string true "Status"
// @Router /payments/{id} [patch]
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
