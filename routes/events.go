package routes

import (
	"awesomeProject/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func getEvents(c *fiber.Ctx) error {
	db, err := models.GetDatabase()
	if err != nil {
		return err
	}

	var events []models.Event
	// get all events
	db.Find(&events)
	if db.Error != nil {
		return db.Error
	}
	return c.JSON(events)
}

func getEvent(c *fiber.Ctx) error {
	db, err := models.GetDatabase()
	if err != nil {
		return err
	}

	var event models.Event

	// join Subscriptions table with Events table
	db.Preload("Subscribers").Find(&event, c.Params("id"))

	if db.Error != nil {
		return db.Error
	}
	return c.JSON(event)
}

func createEvent(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}
	attrsJson := sess.Get("attributes")
	if attrsJson == nil {
		return fiber.ErrUnauthorized
	}
	attrs := CasUserAttributes{}
	err = json.Unmarshal([]byte(attrsJson.(string)), &attrs)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	db, err := models.GetDatabase()
	if err != nil {
		return err
	}

	event := models.Event{}
	if err := c.BodyParser(&event); err != nil {
		return err
	}

	if event.Title == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Title is required")
	}
	if event.Description == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Description is required")
	}
	if event.StartTime == "" {
		return fiber.NewError(fiber.StatusBadRequest, "StartTime is required")
	}
	if event.EndTime == "" {
		return fiber.NewError(fiber.StatusBadRequest, "EndTime is required")
	}

	event.Owner = attrs.Uid

	if err := db.Create(&event).Error; err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(event)
}

func updateEvent(c *fiber.Ctx) error {
	// only the owner can update the event
	sess, err := store.Get(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	attrsJson := sess.Get("attributes")
	if attrsJson == nil {
		return fiber.ErrUnauthorized
	}
	attrs := CasUserAttributes{}
	err = json.Unmarshal([]byte(attrsJson.(string)), &attrs)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if attrs.Uid != c.Params("id") {
		return fiber.ErrForbidden
	}

	// update the event
	db, err := models.GetDatabase()
	if err != nil {
		return err
	}

	event := models.Event{}
	if err := c.BodyParser(&event); err != nil {
		return err
	}

	if err := db.Save(&event).Error; err != nil {
		return err
	}
	return c.JSON(event)
}

func deleteEvent(c *fiber.Ctx) error {
	// only the owner can delete the event
	sess, err := store.Get(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	attrsJson := sess.Get("attributes")
	if attrsJson == nil {
		return fiber.ErrUnauthorized
	}
	attrs := CasUserAttributes{}
	err = json.Unmarshal([]byte(attrsJson.(string)), &attrs)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	// get the event
	db, err := models.GetDatabase()
	if err != nil {
		return err
	}

	event := models.Event{}
	db.Find(&event, c.Params("id"))
	if db.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "Event not found")
	}

	if attrs.Uid != event.Owner {
		return fiber.ErrForbidden
	}

	// delete the event
	if err := db.Delete(&event, c.Params("id")).Error; err != nil {
		return err
	}
	return c.JSON(event)
}

func subscribe(c *fiber.Ctx) error {
	db, err := models.GetDatabase()
	if err != nil {
		return err
	}

	event := models.Event{}
	if err := db.First(&event, c.Params("id")).Error; err != nil {
		return err
	}

	sess, err := store.Get(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	attrsJson := sess.Get("attributes")
	if attrsJson == nil {
		return fiber.ErrUnauthorized
	}

	attrs := CasUserAttributes{}
	err = json.Unmarshal([]byte(attrsJson.(string)), &attrs)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	err = db.Model(&event).Association("Subscribers").Append(&models.User{Uid: attrs.Uid})
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}
