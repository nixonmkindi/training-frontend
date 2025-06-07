package routes

import (
	"training-frontend/package/validator"
	"training-frontend/server/middlewares"
	"training-frontend/server/systems/auth"
	frontend "training-frontend/server/systems/training-frontend"

	"github.com/labstack/echo/v4"
)

// Routers function
func Routers(app *echo.Echo) {
	//Common middleware for all type of routers
	app.Use(middlewares.Cors())
	app.Use(middlewares.Gzip())
	app.Use(middlewares.Logger(true))
	app.Use(middlewares.Secure())
	app.Use(middlewares.Recover())
	app.Use(middlewares.Session()) // uncomment this to enable session
	// app.Use(middlewares.JWT(), middlewares.CheckAuth()) // uncomment this to enable authentication

	app.Validator = validator.GetValidator() //initialize custom validator

	// register static all static routes here
	app.Static("/css", "./server/public/css")
	app.Static("/adminlte", "./server/public/adminlte")
	app.Static("/images", "./server/public/images")
	app.Static("/js", "./server/public/js")
	app.Static("/dashboard", "./server/public/dashboard")

	//register subsystem routes, arrange them by alfabetically
	auth.WebRouters(app)     //add auth routes
	frontend.WebRouters(app) // add frontend routes
}
