package handler

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"teams-secure-notes/internal/model"
	"teams-secure-notes/internal/repository"
	"teams-secure-notes/internal/middleware"
	"log"
)


// RegisterUser creates a new user
func RegisterUser(c *fiber.Ctx) error {
	// Placeholder for user struct
	user := new(model.User)

	// Parse body into user struct
	if err := c.BodyParser(user); err != nil {
		return fiber.ErrBadRequest
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	user.Password = string(hashedPassword)

	// Save the user to the database
	
	if err := repository.CreateUser(user); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// LoginUser authenticates a user
func LoginUser(c *fiber.Ctx) error {
	// Placeholder for login credentials
	login := new(struct{ Username, Password string })

	// Parse body into login struct
	if err := c.BodyParser(login); err != nil {
		return fiber.ErrBadRequest
	}

	// Retrieve user by username
	user, err := repository.GetUserByUsername(login.Username)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		log.Println(login.Password)
		log.Println(user.Password)
		newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(login.Password), bcrypt.DefaultCost)
		if err != nil {
			return fiber.ErrInternalServerError
		}
		log.Println(string(newHashedPassword))
		log.Println(bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)))
		return (fiber.NewError(fiber.StatusUnauthorized, "Invalid password")) 
	}


	// Start new session
	sess, err := middleware.Store.Get(c)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	
	// Save the session
	sess.Set("user_id", user.ID)
	err = sess.Save()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	 
	// Return JSON or a success message
	response := struct {
		Message string `json:"message"`
	}{
		Message: "User logged in successfully",
	}

	return c.JSON(response)
}

