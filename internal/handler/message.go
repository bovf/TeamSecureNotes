package handler

import "github.com/gofiber/fiber/v2"

// GetMessages returns a list of messages
func GetMessages(c *fiber.Ctx) error {
    // Logic to retrieve messages from the database
     return c.SendString("GetMessages called")
}
// CreateMessage creates a new message
func CreateMessage(c *fiber.Ctx) error {
    // Logic to save a new message to the database
    return c.SendString("CreateMessage called")
}

