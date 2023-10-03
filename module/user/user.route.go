package user

import (
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	BASE_URI := "users"

	// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
	app.Get(BASE_URI, GetUser)
	
	app.Get(BASE_URI + "/:id", GetUserById)
	app.Put(BASE_URI + "/:id", EditUser)
	app.Delete(BASE_URI + "/:id", DeleteUser)
	app.Post(BASE_URI, CreateUser)
}