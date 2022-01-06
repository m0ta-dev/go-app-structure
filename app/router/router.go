package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/m0taru/go-app-structure/app/controller"
	"github.com/m0taru/go-app-structure/app/handler"
	"github.com/m0taru/go-app-structure/app/middleware"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App, c *controller.UserController) fiber.Router{
	api := app.Group("/api")
	api.Get("/", handler.Hello)

	// Auth
	auth := app.Group("/auth")
	auth.Post("/signup", 	c.SignUp)
	auth.Post("/signin", 	c.SignIn)

	return api
}

// SetupRoutesForUser setup router api
func SetupRoutesForUser(api fiber.Router, c *controller.UserController) {
	
	user := api.Group("/user").Use(middleware.Auth)
	//user.Post("/", c.Create)
	user.Get("/", c.List)//.Use(middleware.Auth)
	user.Patch("/", c.Update)//.Use(middleware.Auth)
	user.Delete("/", c.Delete)//.Use(middleware.Auth)

	// users := api.Group("/users").Use(middleware.Auth)
	// users.Get("/", c.List)
}