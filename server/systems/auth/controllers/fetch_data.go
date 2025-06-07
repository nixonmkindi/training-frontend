package controllers

import (
	"net/http"
	"training-frontend/package/log"
	"training-frontend/server/systems"
	"training-frontend/server/systems/auth/models"
	"training-frontend/server/systems/helpers"

	"github.com/labstack/echo/v4"
)

func FetchAllPermissions(c echo.Context) ([]models.Permission, error) {

	endPoint := "/permissions/list"

	resp, err := systems.AuthClient.Post(c, endPoint, nil)

	var permissions []models.Permission

	if err != nil || resp.Code == http.StatusNoContent {
		log.Errorf("error occurred while posting: %v", err)
		return permissions, err
	}

	helpers.Decode(resp.Data, &permissions)

	return permissions, nil
}

func FetchAllUsers(c echo.Context) ([]models.User, error) {

	endPoint := "/users/list"

	resp, err := systems.AuthClient.Post(c, endPoint, nil)

	var users []models.User

	if err != nil || resp.Code == http.StatusNoContent {
		log.Errorf("error occurred while posting: %v", err)
		return users, err
	}

	helpers.Decode(resp.Data, &users)

	return users, nil
}

func FetchUserByCategory(c echo.Context, categoryID int32) []models.User {

	endPoint := "/users/list-by-category"
	user := models.User{
		UserCategoryID: categoryID,
	}

	resp, err := systems.AuthClient.Post(c, endPoint, user)
	var users []models.User

	if err != nil || resp.Code == http.StatusNoContent {
		log.Errorf("error occurred while posting: %v", err)
		return users
	}

	helpers.Decode(resp.Data, &users)

	return users
}

func GetRoleByName(c echo.Context, name string) (models.Role, error) {
	endPoint := "/roles/get-role-by-name"

	role := models.Role{}
	role.Name = name

	resp, err := systems.AuthClient.Post(c, endPoint, role)

	var roles models.Role

	if err != nil || resp.Code == http.StatusNoContent {
		log.Errorf("error occurred while posting: %v", err)
		return roles, err
	}

	helpers.Decode(resp.Data, &roles)

	return roles, nil
}

func FetchAllRoles(c echo.Context) ([]*models.Role, error) {
	endPoint := "/roles/list"

	resp, err := systems.AuthClient.Post(c, endPoint, nil)

	var roles []*models.Role

	if err != nil || resp.Code == http.StatusNoContent {
		log.Errorf("error occurred while posting: %v", err)
		return roles, err
	}

	helpers.Decode(resp.Data, &roles)

	return roles, nil
}

func FetchAllUserCategories(c echo.Context) ([]*models.UserCategory, error) {
	endPoint := "/user-categories/list"

	resp, err := systems.AuthClient.Post(c, endPoint, nil)

	var userCategories []*models.UserCategory

	if err != nil || resp.Code == http.StatusNoContent {
		log.Errorf("error occurred while posting: %v", err)
		return userCategories, err
	}

	helpers.Decode(resp.Data, &userCategories)

	return userCategories, nil
}

func GetUserByEmail(c echo.Context, email string) (models.User, error) {
	endPoint := "/users/get-user-by-email"

	params := models.User{
		Email: email,
	}

	var users models.User
	resp, _ := systems.AuthClient.Post(c, endPoint, params)

	helpers.Decode(resp.Data, &users)

	return users, nil

}
