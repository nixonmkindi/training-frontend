package controllers

import (
	"training-frontend/server/systems"
	"training-frontend/server/systems/auth/models"
	"training-frontend/server/systems/helpers"

	"github.com/labstack/echo/v4"
)

const userViewPath = "/auth/views/user/"

var User users

type users struct{}

func (usr *users) GetUser(c echo.Context, userID int32) (models.User, error) {
	endPoint := "/users/show"

	params := models.User{
		ID: userID,
	}

	var users models.User
	resp, _ := systems.AuthClient.Post(c, endPoint, params)

	helpers.Decode(resp.Data, &users)

	return users, nil
}
