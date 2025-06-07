package getdata

import (
	"net/http"
	"training-frontend/package/log"
	"training-frontend/server/systems"
	authh "training-frontend/server/systems/auth/models"
	"training-frontend/server/systems/helpers"

	"github.com/labstack/echo/v4"
	"github.com/yudai/pp"
)

var GetData getData

type getData struct{}

func (at *getData) GetAllUsers(c echo.Context) ([]authh.User, error) {

	endPoint := "/users/list"

	resp, err := systems.AuthClient.Post(c, endPoint, nil)

	if err != nil || resp.Code == http.StatusNoContent {
		pp.Printf("error occurred while posting:%v\n", err)
		helpers.SetErrorMessage(c, "An error has occurred..")
		return nil, c.Redirect(http.StatusSeeOther, "/training-frontend/home")
	}

	var users []authh.User

	helpers.Decode(resp.Data, &users)

	return users, nil
}

func (at *getData) FetchUserByCategory(c echo.Context, categoryID int32) []authh.User {

	endPoint := "/users/list-by-category"
	user := authh.User{
		UserCategoryID: categoryID,
	}

	resp, err := systems.AuthClient.Post(c, endPoint, user)
	var users []authh.User

	if err != nil || resp.Code == http.StatusNoContent {
		log.Errorf("error occurred while posting: %v", err)
		return users
	}

	helpers.Decode(resp.Data, &users)

	return users
}

func (usr *getData) GetUser(c echo.Context, userID int32) (authh.User, error) {
	endPoint := "/users/show"

	params := authh.User{
		ID: userID,
	}

	var users authh.User
	resp, _ := systems.AuthClient.Post(c, endPoint, params)

	helpers.Decode(resp.Data, &users)

	return users, nil
}
