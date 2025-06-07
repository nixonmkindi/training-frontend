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

const subSystemViewPath = "/auth/views/sub_system/"

var SubSystems subSystem

type subSystem struct{}

// Index this is a landing page
func (rhp *subSystem) List(c echo.Context) error {
	endPoint := "/subsystems/list"

	resp, err := systems.AuthClient.Post(c, endPoint, nil)

	if err != nil || resp.Code == http.StatusNoContent {
		log.Errorf("error occurred while posting:%v\n", err)
		helpers.SetErrorMessage(c, "An error has occurred..")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/home")
	}

	var subSystems []*models.SubSystems

	helpers.Decode(resp.Data, &subSystems)

	data := helpers.Map{
		"data": subSystems,
	}

	err = c.Render(http.StatusOK, subSystemViewPath+"index", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error sub system rendering %v", err)
	}
	return nil
}

// Store record
func (rhp *subSystem) Create(c echo.Context) error {
	data := helpers.Map{
		"title": "New Sub system",
		"new":   true,
	}
	return c.Render(http.StatusOK, subSystemViewPath+"create", helpers.Serve(c, data))
}

func (rhp *subSystem) Store(c echo.Context) error {
	endPoint := "/subsystems/create"

	subSystem := models.SubSystems{}
	if err := c.Bind(&subSystem); err != nil {
		helpers.SetErrorMessage(c, "internal error occured")
		return c.Redirect(http.StatusSeeOther, "/auth/subsystems/list")
	}

	userID, _, _ := auth.GetUserFromContext(c)
	subSystem.CreatedBy = userID

	resp, err := systems.AuthClient.Post(c, endPoint, subSystem)

	if err != nil || resp.Code != http.StatusCreated {
		log.Errorf("error sub system entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not create sub system")
		return c.Redirect(http.StatusSeeOther, "/auth/subsystems/list")
	}
	helpers.SetInfoMessage(c, "Subsystem created successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/subsystems/list")

}

func (rhp *subSystem) Show(c echo.Context) error {
	endPoint := "/subsystems/show"
	subSystem := models.SubSystems{}

	if err := c.Bind(&subSystem); err != nil {
		helpers.SetErrorMessage(c, "Internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/subsystems/list")
	}

	var subSystems models.SubSystems
	resp, err := systems.AuthClient.Post(c, endPoint, subSystem)

	if err != nil || resp == nil {
		helpers.SetErrorMessage(c, "error decoding sub system entity")
		return c.Redirect(http.StatusSeeOther, "/auth/subsystems/list")
	}
	if resp.Code != http.StatusOK {
		helpers.SetErrorMessage(c, fmt.Sprint(resp.Message))
		return c.Redirect(http.StatusSeeOther, "/auth/subsystems/list")
	}
	helpers.Decode(resp.Data, &subSystems)

	data := helpers.Map{
		"data": subSystems,
	}

	err = c.Render(http.StatusOK, subSystemViewPath+"show", helpers.Serve(c, data))

	if err != nil {
		log.Errorf("error rendering err-show %v", err)
	}
	return nil

}

func (rhp *subSystem) Edit(c echo.Context) error {
	endPoint := "/subsystems/show"
	subSystem := models.SubSystems{}
	if err := c.Bind(&subSystem); err != nil {
		helpers.SetErrorMessage(c, "internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/subsystems/list")
	}

	var subSystems models.SubSystems
	resp, err := systems.AuthClient.Post(c, endPoint, subSystem)

	if err != nil || resp == nil {
		helpers.SetErrorMessage(c, "error decoding sub system entity")
		return c.Redirect(http.StatusSeeOther, "/auth/subsystems/list")
	}
	if resp.Code != http.StatusOK {
		helpers.SetErrorMessage(c, fmt.Sprint(resp.Message))
		return c.Redirect(http.StatusSeeOther, "/auth/subsystems/list")
	}
	helpers.Decode(resp.Data, &subSystems)

	data := helpers.Map{
		"data": subSystems,
	}

	err = c.Render(http.StatusOK, subSystemViewPath+"edit", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error rendering err-edit %v", err)
	}
	return nil
}

func (rhp *subSystem) Update(c echo.Context) error {
	endPoint := "/subsystems/update"

	subSystem := models.SubSystems{}
	if err := c.Bind(&subSystem); err != nil {
		helpers.SetErrorMessage(c, "internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/subsystems/list")
	}

	userID, _, _ := auth.GetUserFromContext(c)
	subSystem.UpdatedBy = userID

	resp, err := systems.AuthClient.Post(c, endPoint, subSystem)

	if err != nil || resp.Code != http.StatusAccepted {
		log.Errorf("error sub system entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not update sub system")
		return c.Redirect(http.StatusSeeOther, "/auth/subsystems/list")
	}
	helpers.SetInfoMessage(c, "Subsystem updated successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/subsystems/list")

}

func (rhp *subSystem) Delete(c echo.Context) error {
	endPoint := "/subsystems/delete"
	subSystems := models.SubSystems{}
	if err := c.Bind(&subSystems); err != nil {
		helpers.SetErrorMessage(c, "internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/subsystems/list")
	}
	params := models.SubSystems{
		ID:        subSystems.ID,
		DeletedBy: 1,
	}
	resp, err := systems.AuthClient.Post(c, endPoint, params)

	if err != nil || resp.Code != http.StatusAccepted {
		log.Errorf("error sub system entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not delete sub system")
		return c.Redirect(http.StatusSeeOther, "/auth/subsystems/list")
	}
	helpers.SetInfoMessage(c, "User deleted successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/subsystems/list")
}

func GetAllSubsystems(c echo.Context) ([]*models.SubSystems, error) {
	endPoint := "/subsystems/list"

	resp, err := systems.AuthClient.Post(c, endPoint, nil)
	var subSystems []*models.SubSystems

	if err != nil || resp.Code == http.StatusNoContent {
		log.Errorf("error occurred while posting:%v\n", err)
		return subSystems, nil
	}

	helpers.Decode(resp.Data, &subSystems)

	return subSystems, nil
}
