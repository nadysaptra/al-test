package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"azure-labs-test/app/models"
)

func FetchAllUsers(c *fiber.Ctx) error {
	result, _ := models.FetchAllUsers()
	return c.Status(result.Status).JSON(result)
}

func FetchUser(c *fiber.Ctx) error {
	var user models.User
	convertToUint32, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "cannot parse ID to uint "})
	} else {
		user.ID = uint(convertToUint32)
	}
	result, _ := models.FetchUser(&user)
	return c.Status(result.Status).JSON(result)
}

func FetchUserTickets(c *fiber.Ctx) error {
	var user models.User
	convertToUint32, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "cannot parse ID to uint "})
	} else {
		user.ID = uint(convertToUint32)
	}
	result, _ := models.FetchUserTicket(&user)
	return c.Status(result.Status).JSON(result)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	user.Email = c.FormValue("email")
	user.Name = c.FormValue("name")
	if user.Email == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "email is required"})
	}
	if user.Name == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "name is required"})
	}
	result, _ := models.CreateAUser(&user)
	return c.Status(result.Status).JSON(result)
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	if c.Params("id") == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "id is required"})
	}
	convertToUint32, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "cannot parse ID to uint "})
	} else {
		user.ID = uint(convertToUint32)
	}
	if c.FormValue("email") != "" {
		user.Email = c.FormValue("email")
	}
	if c.FormValue("name") != "" {
		user.Name = c.FormValue("name")
	}

	result, _ := models.UpdateUser(&user)
	return c.Status(result.Status).JSON(result)
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User
	if c.Params("id") == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "id is required"})
	}
	convertToUint32, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "cannot parse ID to uint "})
	} else {
		user.ID = uint(convertToUint32)
	}

	result, _ := models.DeleteUser(&user)
	return c.Status(result.Status).JSON(result)
}
