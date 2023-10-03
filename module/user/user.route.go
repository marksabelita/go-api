package user

import (
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	BASE_URI := "users"

	
	app.Get(BASE_URI, GetUser)
	
	app.Get(BASE_URI + "/:id", GetUserById)
	app.Put(BASE_URI + "/:id", EditUser)
	app.Delete(BASE_URI + "/:id", DeleteUser)
	app.Post(BASE_URI, CreateUser)
}