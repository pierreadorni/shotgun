package routes

import (
	"awesomeProject/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"strconv"
)

var store *session.Store

func login(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	// get the account id from the post body
	var account models.Account
	err = c.BodyParser(&account)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body: "+err.Error())
	}

	// get the account from the database from email
	db, err := models.GetDatabase()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	var dbAccount models.Account
	db = db.First(&dbAccount, "email = ?", account.Email)
	if db.Error != nil {
		return fiber.ErrNotFound
	}

	// check if the password matches
	if dbAccount.Password != account.Password {
		return fiber.ErrUnauthorized
	}

	// set the account id in the session
	sess.Set("id", dbAccount.ID)
	err = sess.Save()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}

func logout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	// delete the account id from the session
	sess.Delete("id")
	err = sess.Save()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}

func visibleByAccountId(c *fiber.Ctx) error {
	// get the account id from the sess
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.ErrUnauthorized.Code, "Failed to get session: "+err.Error())
	}
	accountId := sess.Get("id")
	if accountId == nil {
		return fiber.NewError(fiber.ErrUnauthorized.Code, "You are not logged in")
	}

	// get the account id from the url
	id := c.Params("id")
	// check if the account id from the request context matches the account id from the url
	if strconv.Itoa(int(accountId.(uint))) != id {
		return fiber.NewError(fiber.ErrForbidden.Code, "Your account id ("+strconv.Itoa(int(accountId.(uint)))+") does not match the account id in the url")
	}
	return c.Next()
}
