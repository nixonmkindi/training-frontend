package controllers

import (
	"fmt"
	"net/http"
	"training-frontend/package/log"
	"training-frontend/server/systems"

	"training-frontend/server/systems/auth/models"
	"training-frontend/server/systems/helpers"

	"github.com/labstack/echo/v4"
)

const userHasPermissionViewPath = "/auth/views/user_has_permission/"

var UserHasPermission userHasPermission

type userHasPermission struct{}

func (usp *userHasPermission) List(c echo.Context) error {
	endPoint := "/user-permissions/list"
	userID := models.ID{}

	if err := c.Bind(&userID); err != nil {
		helpers.SetErrorMessage(c, "Internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/user-permissions/list")
	}

	var userHasPermissions []*models.Permission
	resp, err := systems.AuthClient.Post(c, endPoint, userID)

	if err != nil || resp == nil {
		helpers.SetErrorMessage(c, "error decoding user has permission entity")
		return c.Redirect(http.StatusSeeOther, "/auth/users/list")
	}
	if resp.Code == http.StatusInternalServerError {
		helpers.SetErrorMessage(c, fmt.Sprint(resp.Message))
		return c.Redirect(http.StatusSeeOther, "/auth/users/list")
	}
	helpers.Decode(resp.Data, &userHasPermissions)

	user, _ := User.GetUser(c, userID.ID)

	data := helpers.Map{
		"data":   userHasPermissions,
		"userID": userID.ID,
		"user":   user,
	}

	err = c.Render(http.StatusOK, userHasPermissionViewPath+"index", helpers.Serve(c, data))

	if err != nil {
		log.Errorf("error rendering err-show %v", err)
	}
	return nil

}

func (usp *userHasPermission) ListUserPermission(c echo.Context, userID int32) ([]string, error) {

	endPoint := "/user-permissions/list"

	params := models.User{
		ID: userID,
	}

	var userHasPermissions []models.Permission

	resp, _ := systems.AuthClient.Post(c, endPoint, params)

	helpers.Decode(resp.Data, &userHasPermissions)

	permissions := make([]string, 0)

	for _, p := range userHasPermissions {
		permissions = append(permissions, p.ID)
	}

	return permissions, nil
}

// Store record
func (usp *userHasPermission) Create(c echo.Context) error {

	var user models.User

	if err := c.Bind(&user); err != nil {
		helpers.SetErrorMessage(c, "Internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/users/list")
	}

	permissions, _ := FetchAllPermissions(c)

	userPermissions, _ := usp.ListUserPermission(c, user.ID)

	userDetails, _ := User.GetUser(c, user.ID)

	data := helpers.Map{
		"title":           "New User Permission",
		"new":             true,
		"allPermissions":  permissions,
		"userID":          user.ID,
		"userPermissions": userPermissions,
		"userDetails":     userDetails,
	}
	err := c.Render(http.StatusOK, userHasPermissionViewPath+"create", helpers.Serve(c, data))

	if err != nil {
		log.Errorf("error rendering err-show %v", err)
	}
	return nil
}

func (usp *userHasPermission) Store(c echo.Context) error {
	endPoint := "/user-permissions/assign"
	userHasPermission := models.HasPermission{}
	if err := c.Bind(&userHasPermission); err != nil {
		helpers.SetErrorMessage(c, "internal error occured")
		return c.Redirect(http.StatusSeeOther, "/auth/users/list")
	}
	params := models.HasPermission{
		ID:           userHasPermission.ID,
		PermissionID: userHasPermission.PermissionID,
	}

	resp, err := systems.AuthClient.Post(c, endPoint, params)

	if err != nil || resp.Code != http.StatusCreated {
		log.Errorf("error user permission entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not create user has permission")
		return c.Redirect(http.StatusSeeOther, "/auth/users/list")
	}
	helpers.SetInfoMessage(c, "User Permission created successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/users/list")

}
