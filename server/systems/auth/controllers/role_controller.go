package controllers

import (
	"fmt"
	"net/http"
	"training-frontend/package/log"
	"training-frontend/server/systems"

	"training-frontend/server/systems/auth/auth"
	"training-frontend/server/systems/auth/models"
	"training-frontend/server/systems/helpers"

	"github.com/labstack/echo/v4"
)

const roleViewPath = "/auth/views/role/"

var Role role

type role struct{}

// Index this is a landing page
func (rhp *role) List(c echo.Context) error {
	endPoint := "/roles/list"

	resp, err := systems.AuthClient.Post(c, endPoint, nil)

	if err != nil || resp.Code == http.StatusNoContent {
		log.Errorf("error occurred while posting:%v\n", err)
		helpers.SetErrorMessage(c, "An error has occurred..")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/home")
	}

	var roles []*models.Role

	helpers.Decode(resp.Data, &roles)

	data := helpers.Map{
		"data": roles,
	}

	err = c.Render(http.StatusOK, roleViewPath+"index", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error role rendering %v", err)
	}
	return nil
}

// Store record
func (rhp *role) Create(c echo.Context) error {
	data := helpers.Map{
		"title": "New Role",
		"new":   true,
	}
	return c.Render(http.StatusOK, roleViewPath+"create", helpers.Serve(c, data))
}

func (rhp *role) Store(c echo.Context) error {
	endPoint := "/roles/create"
	role := models.Role{}
	if err := c.Bind(&role); err != nil {
		helpers.SetErrorMessage(c, "internal error occured")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}

	// userID, _, _ := auth.GetUserFromContext(c)

	params := models.Role{
		Name:        role.Name,
		Description: role.Description,
		CreatedBy:   1,
	}
	resp, err := systems.AuthClient.Post(c, endPoint, params)

	if err != nil || resp.Code != http.StatusCreated {
		log.Errorf("error role entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not create role")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}
	helpers.SetInfoMessage(c, "Role created successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/roles/list")

}

func (rhp *role) Show(c echo.Context) error {
	endPoint := "/roles/show"
	role := models.Role{}

	if err := c.Bind(&role); err != nil {
		helpers.SetErrorMessage(c, "Internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}

	params := models.ID{
		ID: role.ID,
	}
	var roles models.Role
	resp, err := systems.AuthClient.Post(c, endPoint, params)

	if err != nil || resp == nil {
		helpers.SetErrorMessage(c, "error decoding role entity")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}
	if resp.Code == http.StatusInternalServerError {
		helpers.SetErrorMessage(c, fmt.Sprint(resp.Message))
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}
	helpers.Decode(resp.Data, &roles)

	data := helpers.Map{
		"data": roles,
	}

	err = c.Render(http.StatusOK, roleViewPath+"show", helpers.Serve(c, data))

	if err != nil {
		log.Errorf("error rendering err-show %v", err)
	}
	return nil

}

func (rhp *role) Edit(c echo.Context) error {
	endPoint := "/roles/show"

	role := models.Role{}
	if err := c.Bind(&role); err != nil {
		helpers.SetErrorMessage(c, "internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}

	resp, err := systems.AuthClient.Post(c, endPoint, role)

	var roles models.Role
	if err != nil || resp == nil {
		helpers.SetErrorMessage(c, "error decoding role entity")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}
	if resp.Code == http.StatusInternalServerError {
		helpers.SetErrorMessage(c, fmt.Sprint(resp.Message))
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}
	helpers.Decode(resp.Data, &roles)

	data := helpers.Map{
		"data": roles,
	}

	err = c.Render(http.StatusOK, roleViewPath+"edit", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error rendering err-edit %v", err)
	}
	return nil
}

func (rhp *role) Update(c echo.Context) error {
	endPoint := "/roles/update"
	role := models.Role{}
	if err := c.Bind(&role); err != nil {
		helpers.SetErrorMessage(c, "internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}

	user, _, _ := auth.GetUserFromContext(c)
	role.UpdatedBy = user

	resp, err := systems.AuthClient.Post(c, endPoint, role)

	if err != nil || resp.Code != http.StatusAccepted {
		log.Errorf("error role entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not update role")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}
	helpers.SetInfoMessage(c, "Role updated successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/roles/list")

}

func (rhp *role) Delete(c echo.Context) error {
	endPoint := "/roles/delete"

	roles := models.Role{}
	if err := c.Bind(&roles); err != nil {
		helpers.SetErrorMessage(c, "internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}

	user, _, _ := auth.GetUserFromContext(c)
	roles.DeletedBy = user

	resp, err := systems.AuthClient.Post(c, endPoint, roles)

	if err != nil || resp.Code != http.StatusAccepted {
		log.Errorf("error role entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not delete role")
		return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
	}
	helpers.SetInfoMessage(c, "Role deleted successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/roles/list")
}

func GetRole(c echo.Context, roleID int32) (*models.Role, error) {
	endPoint := "/roles/show"

	id := models.Role{}

	id.ID = roleID

	var roles *models.Role

	resp, err := systems.AuthClient.Post(c, endPoint, id)

	if err != nil {
		helpers.SetErrorMessage(c, "error decoding role entity")
		return roles, err
	}

	helpers.Decode(resp.Data, &roles)

	return roles, nil

}
