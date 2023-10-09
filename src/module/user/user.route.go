package user

import (
	default_routes "go-api/src/common/defaults"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get(default_routes.DEFAULT_USERS_URI, GetUser)
	app.Get(default_routes.DEFAULT_USERS_URI + "/:id", GetUserById)
	app.Put(default_routes.DEFAULT_USERS_URI + "/:id", EditUser)
	app.Delete(default_routes.DEFAULT_USERS_URI + "/:id", DeleteUser)
	app.Post(default_routes.DEFAULT_USERS_URI, CreateUser)
}