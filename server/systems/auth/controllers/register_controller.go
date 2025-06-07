package controllers

import (
	"net/http"
	"training-frontend/package/log"
	"training-frontend/server/systems/helpers"

	"github.com/labstack/echo/v4"
)

const registerViewPath = "auth/views/auth/"

var Register register

type register struct{}

// Show Registration Form
func (r *register) ShowRegistrationForm(c echo.Context) error {

	data := helpers.Map{
		"title": "Training Frontend | Register",
	}
	return c.Render(http.StatusOK, registerViewPath+"register.html", helpers.Serve(c, data))
}

// Create UserAccount
func (r *register) CreateAccount(c echo.Context) error {

	endPoint := "/users/create"

	log.Info(endPoint)
	return c.Redirect(http.StatusSeeOther, "/auth/login")
}
