package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store = session.New() // default is a memory store

// AuthRequired is a middleware to check if a user is logged in
func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if a session is associated with the current request
		sess, err := Store.Get(c)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		// Check if auth session exists
		if sess.Get("user_id") == nil {
			return fiber.ErrUnauthorized
		}

		return c.Next()
	}
}

