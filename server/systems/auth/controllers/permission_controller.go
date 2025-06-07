package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"training-frontend/package/log"
	"training-frontend/server/systems"

	"training-frontend/server/systems/auth/auth"
	"training-frontend/server/systems/auth/models"
	"training-frontend/server/systems/helpers"

	"github.com/labstack/echo/v4"
)

const permissionViewPath = "/auth/views/permission/"

var Permission permissions

type permissions struct{}

// Index this is a landing page
func (per *permissions) List(c echo.Context) error {
	endPoint := "/permissions/list"

	resp, err := systems.AuthClient.Post(c, endPoint, nil)

	if err != nil || resp.Code == http.StatusNoContent {
		log.Errorf("error occurred while posting:%v\n", err)
		helpers.SetErrorMessage(c, "An error has occurred..")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/home")
	}

	var permissions []*models.Permission

	helpers.Decode(resp.Data, &permissions)

	data := helpers.Map{
		"data": permissions,
	}

	err = c.Render(http.StatusOK, permissionViewPath+"index", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error permission rendering %v", err)
	}
	return nil
}

func (per *permissions) GeneratePermission(c echo.Context) error {
	endPoint := "/permissions/create"

	//Getting working directory path
	path, err := os.Getwd()
	if err != nil {
		log.Errorf("an error has occured while getting path%v", err)
	}
	path += "/.storage/routes/routes.json"

	//Reading json file
	file, _ := os.ReadFile(path)

	var routes []*models.Route
	json.Unmarshal(file, &routes)

	resp, err1 := systems.AuthClient.Post(c, endPoint, routes)

	if err1 != nil || resp.Code != http.StatusCreated {
		log.Errorf("error occurred while posting:%v\n", err)
		helpers.SetErrorMessage(c, "An error has occurred..")
		return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
	}

	helpers.SetInfoMessage(c, "Permissions created successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
}

func (per *permissions) Show(c echo.Context) error {
	endPoint := "/permissions/show"
	permission := models.StringID{}

	if err := c.Bind(&permission); err != nil {
		helpers.SetErrorMessage(c, "Internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
	}

	var permissions models.Permission
	resp, err := systems.AuthClient.Post(c, endPoint, permission)

	if err != nil || resp == nil {
		helpers.SetErrorMessage(c, "error decoding permission entity")
		return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
	}
	if resp.Code == http.StatusInternalServerError {
		helpers.SetErrorMessage(c, fmt.Sprint(resp.Message))
		return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
	}

	helpers.Decode(resp.Data, &permissions)

	data := helpers.Map{
		"data": permissions,
	}

	err = c.Render(http.StatusOK, permissionViewPath+"show", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error rendering err-show %v", err)
	}
	return nil

}

func (rhp *permissions) Edit(c echo.Context) error {
	endPoint := "/permissions/show"

	permission := models.Permission{}
	if err := c.Bind(&permission); err != nil {
		helpers.SetErrorMessage(c, "internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
	}

	var permissions models.Permission
	resp, err := systems.AuthClient.Post(c, endPoint, permission)

	if err != nil || resp == nil {
		helpers.SetErrorMessage(c, "error decoding permission entity")
		return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
	}
	if resp.Code == http.StatusInternalServerError {
		helpers.SetErrorMessage(c, fmt.Sprint(resp.Message))
		return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
	}
	helpers.Decode(resp.Data, &permissions)

	data := helpers.Map{
		"data": permissions,
	}

	err = c.Render(http.StatusOK, permissionViewPath+"edit", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error rendering err-edit %v", err)
	}
	return nil
}

func (rhp *permissions) Update(c echo.Context) error {
	endPoint := "/permissions/update"
	permission := models.Permission{}
	if err := c.Bind(&permission); err != nil {
		helpers.SetErrorMessage(c, "internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
	}

	resp, err := systems.AuthClient.Post(c, endPoint, permission)

	if err != nil || resp.Code != http.StatusAccepted {
		log.Errorf("error permission entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not update permission")
		return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
	}
	helpers.SetInfoMessage(c, "Permission updated successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")

}

func (rhp *permissions) Delete(c echo.Context) error {
	endPoint := "/permissions/soft-delete"
	permissions := models.Permission{}
	if err := c.Bind(&permissions); err != nil {
		helpers.SetErrorMessage(c, "internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
	}

	user, _, _ := auth.GetUserFromContext(c)
	permissions.DeletedBy = user

	resp, err := systems.AuthClient.Post(c, endPoint, permissions)

	if err != nil || resp.Code != http.StatusAccepted {
		log.Errorf("error permission entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not delete permission")
		return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
	}
	helpers.SetInfoMessage(c, "Permission deleted successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
}

func (rhp *permissions) HardDelete(c echo.Context) error {
	endPoint := "/permissions/hard-delete"
	permissions := models.Permission{}
	if err := c.Bind(&permissions); err != nil {
		helpers.SetErrorMessage(c, "internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
	}

	resp, err := systems.AuthClient.Post(c, endPoint, permissions)

	if err != nil || resp.Code != http.StatusAccepted {
		log.Errorf("error permission entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not delete permission")
		return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
	}
	helpers.SetInfoMessage(c, "Permission deleted successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
}

func (rhp *permissions) ForceDelete(c echo.Context) error {
	endPoint := "/permissions/force-delete"

	resp, err := systems.AuthClient.Post(c, endPoint, nil)

	if err != nil || resp.Code != http.StatusAccepted {
		log.Errorf("error permission entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not delete permission")
		return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
	}
	helpers.SetInfoMessage(c, "Permission deleted successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/permissions/list")
}
