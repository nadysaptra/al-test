package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"azure-labs-test/app/models"
)

// User godoc
// @Summary Show user list.
// @Description get user list.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users [get]
func FetchAllUsers(c *fiber.Ctx) error {
	result, _ := models.FetchAllUsers()
	return c.Status(result.Status).JSON(result)
}

// User godoc
// @Summary Show detail user.
// @Description get detail user.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param id path int true "User Id"
// @Router /users/{id} [get]
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

// User godoc
// @Summary Show user tickets.
// @Description get user tickets.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param id path int true "User Id"
// @Router /users/{id}/tickets [get]
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

// User godoc
// @Summary Create new user.
// @Description Create new user.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param email body string true "Email"
// @Param name body string true "Name"
// @Router /users [post]
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

// User godoc
// @Summary Update user.
// @Description Update user.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param email body string true "Email"
// @Param name body string true "Name"
// @Router /users/{id} [patch]
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

// User godoc
// @Summary Delete user.
// @Description Delete user.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/{id} [delete]
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
