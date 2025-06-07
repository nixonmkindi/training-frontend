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

const roleHasPermissionViewPath = "/auth/views/role_has_permission/"

var RoleHasPermission roleHasPermission

type roleHasPermission struct{}

// Index this is a landing page
func (rol *roleHasPermission) List(c echo.Context) error {

	endPoint := "/role-permissions/list"

	var roleHasPermission models.HasPermission

	if err := c.Bind(&roleHasPermission); err != nil {
		helpers.SetErrorMessage(c, "Internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}

	var roleHasPermissions []*models.Permission

	resp, err := systems.AuthClient.Post(c, endPoint, roleHasPermission)
	if err != nil || resp == nil {
		helpers.SetErrorMessage(c, "error decoding role has permission entity")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}
	if resp.Code == http.StatusInternalServerError {
		helpers.SetErrorMessage(c, fmt.Sprint(resp.Message))
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}

	helpers.Decode(resp.Data, &roleHasPermissions)

	role, _ := GetRole(c, roleHasPermission.ID)

	data := helpers.Map{
		"data":   roleHasPermissions,
		"roleID": roleHasPermission.ID,
		"role":   role,
	}

	err = c.Render(http.StatusOK, roleHasPermissionViewPath+"index", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error role has permission rendering %v", err)
	}
	return nil
}

// Index this is a landing page
func (rol *roleHasPermission) ListRolePermission(c echo.Context, roleID int32) ([]string, error) {

	endPoint := "/role-permissions/list"

	params := models.Role{
		ID: roleID,
	}

	var roleHasPermissions []models.Permission

	resp, _ := systems.AuthClient.Post(c, endPoint, params)

	helpers.Decode(resp.Data, &roleHasPermissions)

	permissions := make([]string, 0)

	for _, p := range roleHasPermissions {
		permissions = append(permissions, p.ID)
	}

	return permissions, nil
}

// Store record
func (rol *roleHasPermission) Create(c echo.Context) error {

	var roleHasPermission models.Role

	if err := c.Bind(&roleHasPermission); err != nil {
		helpers.SetErrorMessage(c, "Internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}

	allPermissions, _ := FetchAllPermissions(c)

	rolePermissions, _ := rol.ListRolePermission(c, roleHasPermission.ID)
	subsystems, _ := GetAllSubsystems(c)

	role, _ := GetRole(c, roleHasPermission.ID)

	data := helpers.Map{
		"title":           "New Role Permission",
		"new":             true,
		"allPermissions":  allPermissions,
		"roleID":          roleHasPermission.ID,
		"rolePermissions": rolePermissions,
		"subsystems":      subsystems,
		"role":            role,
	}
	err := c.Render(http.StatusOK, roleHasPermissionViewPath+"create", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error rendering err-show %v", err)
	}
	return nil
}

func (rol *roleHasPermission) Store(c echo.Context) error {
	endPoint := "/role-permissions/assign"
	roleHasPermission := models.HasPermission{}
	if err := c.Bind(&roleHasPermission); err != nil {
		helpers.SetErrorMessage(c, "internal error occured")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}
	params := models.HasPermission{
		ID:           roleHasPermission.ID,
		PermissionID: roleHasPermission.PermissionID,
	}
	resp, err := systems.AuthClient.Post(c, endPoint, params)

	if err != nil || resp.Code != http.StatusCreated {
		log.Errorf("error role has permission entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not create role has permission")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}
	helpers.SetInfoMessage(c, "Role Permission created successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
}
