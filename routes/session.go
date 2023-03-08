package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func GetSession(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	attrsJson := sess.Get("attributes")
	if attrsJson == nil {
		return fiber.ErrNotFound
	}

	attrs := CasUserAttributes{}
	err = json.Unmarshal([]byte(attrsJson.(string)), &attrs)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(attrs)
}
