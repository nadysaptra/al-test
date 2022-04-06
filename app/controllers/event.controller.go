package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"azure-labs-test/app/models"
)

func FetchAllEvents(c *fiber.Ctx) error {
	result, _ := models.FetchAllEvents()
	return c.Status(result.Status).JSON(result)
}

func FetchEvent(c *fiber.Ctx) error {
	var event models.Event
	convertToUint32, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "cannot parse ID to uint "})
	} else {
		event.ID = uint(convertToUint32)
	}
	result, _ := models.FetchEvent(&event)
	return c.Status(result.Status).JSON(result)
}

func FetchEventTickets(c *fiber.Ctx) error {
	var event models.Event
	convertToUint32, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "cannot parse ID to uint "})
	} else {
		event.ID = uint(convertToUint32)
	}
	result, _ := models.FetchEventTicket(&event)
	return c.Status(result.Status).JSON(result)
}

func CreateEvent(c *fiber.Ctx) error {
	var event models.Event

	event.Date = c.FormValue("date")
	event.Name = c.FormValue("name")
	intVar, err := strconv.Atoi(c.FormValue("ticket_prize"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "ticket prize should be integer"})
	}
	event.TicketPrize = intVar
	if event.Date == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "date is required"})
	}
	if event.Name == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "name is required"})
	}
	if event.TicketPrize == 0 {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "ticket_prize is required"})
	}
	result, _ := models.CreateAEvent(&event)
	return c.Status(result.Status).JSON(result)
}

func UpdateEvent(c *fiber.Ctx) error {
	var event models.Event
	if c.Params("id") == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "id is required"})
	}
	convertToUint32, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "cannot parse ID to uint "})
	} else {
		event.ID = uint(convertToUint32)
	}
	if c.FormValue("date") != "" {
		event.Date = c.FormValue("date")
	}
	if c.FormValue("name") != "" {
		event.Name = c.FormValue("name")
	}
	if c.FormValue("ticket_prize") != "" {
		intVar, err := strconv.Atoi(c.FormValue("ticket_prize"))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "ticket prize should be integer"})
		}
		event.TicketPrize = intVar
	}

	result, _ := models.UpdateEvent(&event)
	return c.Status(result.Status).JSON(result)
}

func DeleteEvent(c *fiber.Ctx) error {
	var event models.Event
	if c.Params("id") == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "id is required"})
	}
	convertToUint32, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "cannot parse ID to uint "})
	} else {
		event.ID = uint(convertToUint32)
	}

	result, _ := models.DeleteEvent(&event)
	return c.Status(result.Status).JSON(result)
}
