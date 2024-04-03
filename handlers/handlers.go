package handlers

import (
	"go_fiber/database"
	"go_fiber/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []models.User
	db.Find(&users)
	return c.JSON(fiber.Map{
		"success": true,
		"users":   users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	/*user := &models.User{
		// Note: when writing to external database,
		// we can simply use - Name: c.FormValue("user")
		Name: utils.CopyString(c.FormValue("user")),
	}*/

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(503).SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"success": false,
			"user":    user,
		})
	}

	database.Insert(user)

	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}
