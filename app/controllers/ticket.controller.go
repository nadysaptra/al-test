package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"azure-labs-test/app/models"
)

// Ticket godoc
// @Summary Show ticket list.
// @Description get ticket list.
// @Tags Ticket
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /tickets [get]
func FetchAllTicket(c *fiber.Ctx) error {
	result, _ := models.FethAllTickets()
	return c.Status(result.Status).JSON(result)
}

// Ticket godoc
// @Summary Create new ticket.
// @Description Create new ticket.
// @Tags Ticket
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param status body string true "Status"
// @Param event_refer body string true "Event ID"
// @Param user_refer body string true "User ID"
// @Router /tickets [post]
func CreateTicket(c *fiber.Ctx) error {
	var ticket models.Ticket
	ticket.Status = "ordered"
	u32EventRefer, err := strconv.Atoi(c.FormValue("event_refer"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "cannot parse event_refer to uint "})
	} else {
		ticket.EventRefer = u32EventRefer
	}
	u32UserRefer, err := strconv.Atoi(c.FormValue("user_refer"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "cannot parse user_refer to uint "})
	} else {
		ticket.UserRefer = u32UserRefer
	}

	result, _ := models.CreateTicket(&ticket)
	return c.Status(result.Status).JSON(result)
}

// Ticket godoc
// @Summary Create new ticket.
// @Description Create new ticket.
// @Tags Ticket
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param status body string true "Status"
// @Param event_refer body string true "Event ID"
// @Param user_refer body string true "User ID"
// @Router /tickets/{id} [patch]
func UpdateTicket(c *fiber.Ctx) error {
	var ticket models.Ticket
	if c.Params("id") == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "id is required"})
	}
	u32Id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "cannot parse ID to uint "})
	} else {
		ticket.ID = uint(u32Id)
	}
	ticket.Status = c.FormValue("status")
	if ticket.Status == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "status is required"})
	}
	result, _ := models.UpdateTicket(&ticket)
	return c.Status(result.Status).JSON(result)
}
