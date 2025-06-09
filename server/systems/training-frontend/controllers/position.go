package controllers

import (
	"net/http"
	"strconv"
	"training-frontend/package/log"
	"training-frontend/package/util"
	"training-frontend/server/systems"
	"training-frontend/server/systems/helpers"
	"training-frontend/server/systems/training-frontend/models"

	"github.com/labstack/echo/v4"
)

const positionViewPath = "/training-frontend/views/position/"

func ShowPosition(c echo.Context) error {
	endPoint := "/position/show"
	positionData := &models.GetID{}

	positionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Errorf("error parsing position id: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, Please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}
	positionData.ID = int32(positionID)

	resp, err := systems.BackendClient.Post(c, endPoint, positionData)

	if err != nil || resp == nil {
		log.Errorf("error getting position data: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	if resp.Code == http.StatusInternalServerError {
		log.Errorf("training backend: error getting position: %v", resp.Error)
		helpers.SetErrorMessage(c, "Internal system error, please try again later!")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	var position models.Position
	helpers.Decode(resp.Data, &position)
	data := helpers.Map{
		"data": position,
	}
	err = c.Render(http.StatusOK, positionViewPath+"show", helpers.Serve(c, data))

	if err != nil {
		log.Errorf("error position list page rendering: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, Please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	return nil
}

func ListPosition(c echo.Context) error {
	endPoint := "/position/list"

	//1. frontend error
	resp, err := systems.BackendClient.Post(c, endPoint, nil)
	if util.IsError(err) || resp == nil {
		log.Errorf("error getting position list: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/home")
	}

	//backend error
	if resp.Code == http.StatusInternalServerError {
		log.Errorf("training backend: error getting position list: %v", resp.Error)
		helpers.SetErrorMessage(c, "Internal system error, please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/home")
	}

	//no data
	if resp.Code == http.StatusAccepted {
		log.Info("no position record")
		helpers.SetErrorMessage(c, "No records of positions")
	}

	//data available
	var positions []*models.Position
	helpers.Decode(resp.Data, &positions)

	data := helpers.Map{
		"data": positions,
	}

	err = c.Render(http.StatusOK, positionViewPath+"index", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error position list page rendering: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, Please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-backend/home")
	}

	return nil
}

func CreatePosition(c echo.Context) error {
	data := helpers.Map{
		"title": "New Position",
	}

	err := c.Render(http.StatusOK, positionViewPath+"create", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error position create page rendering %v", err)
		helpers.SetErrorMessage(c, "Internal system error, Please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	return nil
}

func StorePosition(c echo.Context) error {
	endPoint := "/position/create"

	position := models.Position{}

	if err := c.Bind(&position); util.IsError(err) {
		log.Errorf("error binding position: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	if err := c.Validate(position); util.IsError(err) {
		log.Errorf("error validating position: %v", err)
		helpers.SetErrorMessage(c, "Could not validate submitted data")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	//userID, _, _ := auth.GetUserFromContext(c)
	position.CreatedBy = 1

	resp, err := systems.BackendClient.Post(c, endPoint, position)

	if err != nil || resp == nil {
		log.Errorf("error saving position data: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	if resp.Code == http.StatusInternalServerError {
		log.Errorf("training backend: error storing position: %v", resp.Error)
		helpers.SetErrorMessage(c, "Internal system error, please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	helpers.SetInfoMessage(c, "Position created successfully")

	err = c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	if err != nil {
		log.Errorf("error position list page rendering: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, Please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	return nil
}

func EditPosition(c echo.Context) error {
	endPoint := "/position/show"

	id := &models.GetID{}

	positionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Errorf("error parsing position id: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, Please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}
	id.ID = int32(positionID)

	resp, err := systems.BackendClient.Post(c, endPoint, id)
	if util.IsError(err) || resp == nil {
		log.Errorf("error getting position data: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	if resp.Code == http.StatusInternalServerError {
		log.Errorf("training backend: error getting position: %v", resp.Error)
		helpers.SetErrorMessage(c, "Internal system error, please try again later!")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	var positionData models.Position

	helpers.Decode(resp.Data, &positionData)

	data := helpers.Map{
		"data": positionData,
	}

	err = c.Render(http.StatusOK, positionViewPath+"edit", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error position list page rendering: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, Please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	return nil
}

func UpdatePosition(c echo.Context) error {
	endPoint := "/position/update"

	position := models.Position{}

	if err := c.Bind(&position); err != nil {
		log.Errorf("error binding position data: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, please try again later!")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	if err := c.Validate(position); util.IsError(err) {
		log.Errorf("error validating position: %v", err)
		helpers.SetErrorMessage(c, "Could not validate submitted data")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	//userID, _, _ := auth.GetUserFromContext(c)
	position.UpdatedBy = 1
	resp, err := systems.BackendClient.Post(c, endPoint, position)
	if util.IsError(err) || resp == nil {
		log.Errorf("error getting position: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	if resp.Code == http.StatusInternalServerError {
		log.Errorf("training backend: error updating position: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, please try again later!")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	helpers.SetInfoMessage(c, "Position updated successfully")
	err = c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")

	if err != nil {
		log.Errorf("error position list page rendering: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, Please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	return nil
}

func DeletePosition(c echo.Context) error {
	endPoint := "/position/delete"

	id := &models.DeleteIDs{}

	positionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Errorf("error parsing position id: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, Please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	id.ID = int32(positionID)

	//userID, _, _ := auth.GetUserFromContext(c)
	id.DeletedBy = 1
	if err := c.Validate(id); util.IsError(err) {
		log.Errorf("error validating position: %v", err)
		helpers.SetErrorMessage(c, "Could not validate submitted data")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	resp, err := systems.BackendClient.Post(c, endPoint, id)
	if util.IsError(err) || resp == nil {
		log.Errorf("error deleting position: %v", err)
		helpers.SetErrorMessage(c, "Internal system error, please try again later!")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	if resp.Code == http.StatusInternalServerError {
		log.Errorf("training backend: error deleting position: %v", resp.Error)
		helpers.SetErrorMessage(c, "Internal system error, please try again later!")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	helpers.SetInfoMessage(c, "Position is deleted successfully")
	err = c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	if err != nil {
		log.Errorf("error position list page rendering %v", err)
		helpers.SetErrorMessage(c, "Internal system error, Please try again later")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/position/list")
	}

	return nil
}
