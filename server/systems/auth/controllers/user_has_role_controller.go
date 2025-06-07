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

const userHasRoleViewPath = "/auth/views/user_has_role/"

var UserHasRole userHasRole

type userHasRole struct{}

// Index this is a landing page
func (usp *userHasRole) List(c echo.Context) error {
	endPoint := "/user-roles/list"
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		helpers.SetErrorMessage(c, "Internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/user-permissions/list")
	}

	var userHasRoles []*models.Role
	resp, err := systems.AuthClient.Post(c, endPoint, user)

	if err != nil || resp == nil {
		helpers.SetErrorMessage(c, "error decoding user has permission entity")
		return c.Redirect(http.StatusSeeOther, "/auth/user-permissions/list")
	}
	if resp.Code == http.StatusInternalServerError {
		helpers.SetErrorMessage(c, fmt.Sprint(resp.Message))
		return c.Redirect(http.StatusSeeOther, "/auth/user-permissions/list")
	}

	helpers.Decode(resp.Data, &userHasRoles)

	userDetails, _ := User.GetUser(c, user.ID)

	data := helpers.Map{
		"data":        userHasRoles,
		"userID":      user.ID,
		"userDetails": userDetails,
	}

	err = c.Render(http.StatusOK, userHasRoleViewPath+"index", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error user has role rendering %v", err)
	}
	return nil
}

func (usp *userHasRole) ListUserRole(c echo.Context, userID int32) ([]string, error) {

	endPoint := "/user-roles/list"

	params := models.User{
		ID: userID,
	}

	var userHasRoles []models.Role

	resp, _ := systems.AuthClient.Post(c, endPoint, params)

	helpers.Decode(resp.Data, &userHasRoles)

	roles := make([]string, 0)

	for _, p := range userHasRoles {
		roles = append(roles, p.Name)
	}

	return roles, nil
}

func (usp *userHasRole) Create(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		helpers.SetErrorMessage(c, "Internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/users/list")
	}
	roles, _ := FetchAllRoles(c)

	userRoles, _ := usp.ListUserRole(c, user.ID)

	userDetails, _ := User.GetUser(c, user.ID)

	data := helpers.Map{
		"title":       "New User Role",
		"new":         true,
		"allRoles":    roles,
		"userID":      user.ID,
		"userRoles":   userRoles,
		"userDetails": userDetails,
	}
	return c.Render(http.StatusOK, userHasRoleViewPath+"create", helpers.Serve(c, data))
}

func (usp *userHasRole) Store(c echo.Context) error {
	endPoint := "/user-roles/assign"
	userHasRole := models.UserHasRole{}
	if err := c.Bind(&userHasRole); err != nil {
		helpers.SetErrorMessage(c, "internal error occured")
		return c.Redirect(http.StatusSeeOther, "/auth/user-roles/list")
	}

	params := models.UserHasRole{
		RoleID: userHasRole.RoleID,
		UserID: userHasRole.UserID,
	}

	resp, err := systems.AuthClient.Post(c, endPoint, params)

	if err != nil || resp.Code != http.StatusCreated {
		log.Errorf("error user roles entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not create user has role")
		return c.Redirect(http.StatusSeeOther, "/auth/users/list")
	}
	helpers.SetInfoMessage(c, "User Role created successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/users/list")

}

func (usp *userHasRole) Show(c echo.Context) error {
	endPoint := "/user-roles/show"
	userHasRole := models.UserHasRole{}

	if err := c.Bind(&userHasRole); err != nil {
		helpers.SetErrorMessage(c, "Internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/user-roles/list")
	}

	params := models.UserHasRole{
		UserID: userHasRole.UserID,
	}
	var userHasRoles models.UserHasRole
	resp, err := systems.AuthClient.Post(c, endPoint, params)

	if err != nil || resp == nil {
		helpers.SetErrorMessage(c, "error decoding user has role entity")
		return c.Redirect(http.StatusSeeOther, "/auth/user-roles/list")
	}
	if resp.Code == http.StatusInternalServerError {
		helpers.SetErrorMessage(c, fmt.Sprint(resp.Message))
		return c.Redirect(http.StatusSeeOther, "/auth/user-roles/list")
	}
	helpers.Decode(resp.Data, &userHasRoles)

	data := helpers.Map{
		"data": userHasRoles,
	}

	err = c.Render(http.StatusOK, userHasRoleViewPath+"show", helpers.Serve(c, data))

	if err != nil {
		log.Errorf("error rendering err-show %v", err)
	}
	return nil
}

func (usp *userHasRole) CreateDefaultRole(c echo.Context, userID int32, roleID []int32) error {
	endPoint := "/user-roles/assign"

	params := models.UserHasRole{
		RoleID: roleID,
		UserID: userID,
	}

	resp, err := systems.AuthClient.Post(c, endPoint, params)
	if err != nil || resp.Code != http.StatusCreated {
		log.Errorf("error user roles entity decoding %v", err)
	}

	return err
}

func (usp *userHasRole) ListRoleUserPermission(c echo.Context) error {
	endPoint := "/user-roles/role-user-permission"
	role := models.Role{}

	if err := c.Bind(&role); err != nil {
		helpers.SetErrorMessage(c, "Internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}

	var roleUserPermission *models.RoleUserPermission
	resp, err := systems.AuthClient.Post(c, endPoint, role)

	if err != nil || resp == nil {
		helpers.SetErrorMessage(c, "error decoding user has permission entity")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}
	if resp.Code == http.StatusInternalServerError {
		helpers.SetErrorMessage(c, fmt.Sprint(resp.Message))
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}

	helpers.Decode(resp.Data, &roleUserPermission)

	data := helpers.Map{
		"data": roleUserPermission,
	}

	err = c.Render(http.StatusOK, userHasRoleViewPath+"role_user_permission", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error role has user rendering %v", err)
	}
	return nil
}
