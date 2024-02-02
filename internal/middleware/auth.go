package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// Protected checks if the user is logged in before allowing access to a route.
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the session from the store
		sess, err := Store.Get(c)
		if err != nil {
			// If we can't get a session, return an error.
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthenticated")
		}

		// Check if the user ID is stored in the session
		if sess.Get("user_id") == nil {
			// If a user ID is not stored, the user is not logged in.
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthenticated")
		}

		// If there is a user ID in the session, the user is logged in, so continue to the next middleware or handler
		return c.Next()
	}
}
