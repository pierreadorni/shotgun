package routes

import (
	"awesomeProject/models"
	"github.com/gofiber/fiber/v2"
)

func getAccounts(c *fiber.Ctx) error {
	db, err := models.GetDatabase()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	var accounts []models.Account
	// find accounts but remove the password field
	db = db.Select("id", "email", "created_at", "updated_at").Find(&accounts)
	return c.JSON(accounts)
}

func getAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	db, err := models.GetDatabase()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	var account models.Account
	db = db.First(&account, id)
	if db.Error != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(account)
}

func createAccount(c *fiber.Ctx) error {
	db, err := models.GetDatabase()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	var account models.Account
	err = c.BodyParser(&account)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body: "+err.Error())
	}

	if account.Email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Email is required")
	}
	if account.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Password is required")
	}

	db.Create(&account)
	if db.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create account: "+db.Error.Error())
	}

	return c.JSON(account)
}

func updateAccount(c *fiber.Ctx) error {
	db, err := models.GetDatabase()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	var account models.Account
	db = db.First(&account, c.Params("id"))
	if db.Error != nil {
		return fiber.ErrNotFound
	}

	err = c.BodyParser(&account)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body: "+err.Error())
	}

	db.Save(&account)
	if db.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update account: "+db.Error.Error())
	}

	return c.JSON(account)
}

func deleteAccount(c *fiber.Ctx) error {
	db, err := models.GetDatabase()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	var account models.Account
	db = db.First(&account, c.Params("id"))
	if db.Error != nil {
		return fiber.ErrNotFound
	}

	db.Delete(&account)
	if db.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete account: "+db.Error.Error())
	}

	return c.SendString("Account deleted")
}
