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

const userCategoryViewPath = "/auth/views/user_category/"

var UserCategory userCategory

type userCategory struct{}

// Index this is a landing page
func (rhp *userCategory) List(c echo.Context) error {
	endPoint := "/user-categories/list"

	resp, err := systems.AuthClient.Post(c, endPoint, nil)

	if err != nil || resp.Code == http.StatusNoContent {
		log.Errorf("error occurred while posting:%v\n", err)
		helpers.SetErrorMessage(c, "An error has occurred..")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/home")
	}

	var userCategories []*models.UserCategory

	helpers.Decode(resp.Data, &userCategories)

	data := helpers.Map{
		"data": userCategories,
	}

	err = c.Render(http.StatusOK, userCategoryViewPath+"index", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error user category rendering %v", err)
	}
	return nil
}

// Store record
func (rhp *userCategory) Create(c echo.Context) error {
	data := helpers.Map{
		"title": "New User Category",
		"new":   true,
	}
	return c.Render(http.StatusOK, userCategoryViewPath+"create", helpers.Serve(c, data))
}

func (rhp *userCategory) Store(c echo.Context) error {
	endPoint := "/user-categories/create"
	userCategory := models.UserCategory{}
	if err := c.Bind(&userCategory); err != nil {
		helpers.SetErrorMessage(c, "internal error occured")
		return c.Redirect(http.StatusSeeOther, "/auth/user-categories/list")
	}

	userID, _, _ := auth.GetUserFromContext(c)
	userCategory.CreatedBy = userID

	resp, err := systems.AuthClient.Post(c, endPoint, userCategory)

	if err != nil || resp.Code != http.StatusCreated {
		log.Errorf("error user categories entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not create user category")
		return c.Redirect(http.StatusSeeOther, "/auth/user-categories/list")
	}
	helpers.SetInfoMessage(c, "User Category created successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/user-categories/list")

}

func (rhp *userCategory) Show(c echo.Context) error {
	endPoint := "/user-categories/show"
	userCategory := models.UserCategory{}

	if err := c.Bind(&userCategory); err != nil {
		helpers.SetErrorMessage(c, "Internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/user-categories/list")
	}

	var userCategories models.UserCategory
	resp, err := systems.AuthClient.Post(c, endPoint, userCategory)

	if err != nil || resp == nil {
		helpers.SetErrorMessage(c, "error decoding user category entity")
		return c.Redirect(http.StatusSeeOther, "/auth/user-categories/list")
	}
	if resp.Code != http.StatusOK {
		helpers.SetErrorMessage(c, fmt.Sprint(resp.Message))
		return c.Redirect(http.StatusSeeOther, "/auth/user-categories/list")
	}
	helpers.Decode(resp.Data, &userCategories)

	data := helpers.Map{
		"data": userCategories,
	}

	err = c.Render(http.StatusOK, userCategoryViewPath+"show", helpers.Serve(c, data))

	if err != nil {
		log.Errorf("error rendering err-show %v", err)
	}
	return nil

}

func (rhp *userCategory) Edit(c echo.Context) error {
	endPoint := "/user-categories/show"
	userCategory := models.UserCategory{}
	if err := c.Bind(&userCategory); err != nil {
		helpers.SetErrorMessage(c, "internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/user-categories/list")
	}

	var userCategories models.UserCategory
	resp, err := systems.AuthClient.Post(c, endPoint, userCategory)

	if err != nil || resp == nil {
		helpers.SetErrorMessage(c, "error decoding user category entity")
		return c.Redirect(http.StatusSeeOther, "/auth/user-categories/list")
	}
	if resp.Code == http.StatusInternalServerError {
		helpers.SetErrorMessage(c, fmt.Sprint(resp.Message))
		return c.Redirect(http.StatusSeeOther, "/auth/user-categories/list")
	}
	helpers.Decode(resp.Data, &userCategories)

	data := helpers.Map{
		"data": userCategories,
	}

	err = c.Render(http.StatusOK, userCategoryViewPath+"edit", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error rendering err-edit %v", err)
	}
	return nil
}

func (rhp *userCategory) Update(c echo.Context) error {
	endPoint := "/user-categories/update"

	userCategory := models.UserCategory{}
	if err := c.Bind(&userCategory); err != nil {
		helpers.SetErrorMessage(c, "internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/user-categories/list")
	}

	userID, _, _ := auth.GetUserFromContext(c)
	userCategory.UpdatedBy = userID

	resp, err := systems.AuthClient.Post(c, endPoint, userCategory)

	if err != nil || resp.Code != http.StatusAccepted {
		log.Errorf("error user categories entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not update user category")
		return c.Redirect(http.StatusSeeOther, "/auth/user-categories/list")
	}
	helpers.SetInfoMessage(c, "User Category updated successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/user-categories/list")

}

func (rhp *userCategory) Delete(c echo.Context) error {
	endPoint := "/user-categories/delete"

	userCategory := models.UserCategory{}
	if err := c.Bind(&userCategory); err != nil {
		helpers.SetErrorMessage(c, "internal error occured!")
		return c.Redirect(http.StatusSeeOther, "/auth/user-categories/list")
	}

	userID, _, _ := auth.GetUserFromContext(c)
	userCategory.DeletedBy = userID

	resp, err := systems.AuthClient.Post(c, endPoint, userCategory)

	if err != nil || resp.Code != http.StatusAccepted {
		log.Errorf("error user category entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not delete user category")
		return c.Redirect(http.StatusSeeOther, "/auth/user-categories/list")
	}
	helpers.SetInfoMessage(c, "User category deleted successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/user-categories/list")
}
