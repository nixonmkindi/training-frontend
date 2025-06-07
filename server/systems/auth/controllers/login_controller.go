package controllers

import (
	"fmt"
	"net/http"
	"training-frontend/package/log"
	"training-frontend/package/util"
	"training-frontend/server/systems"
	"training-frontend/server/systems/auth/auth"
	"training-frontend/server/systems/auth/models"
	"training-frontend/server/systems/helpers"

	"github.com/labstack/echo/v4"
)

const loginViewPath = "auth/views/auth/"

var Login login

type login struct{}

// Show Login Form
func (lg *login) ShowLoginForm(c echo.Context) error {
	data := helpers.Map{
		"title": "DIT Paperless | Login",
	}
	err := c.Render(http.StatusOK, loginViewPath+"login.html", helpers.Serve(c, data))
	if err != nil {
		fmt.Print(err)
	}
	return nil
}

// LoginUser
func (lg *login) Login(c echo.Context) error {

	user := models.Login{}
	if err := c.Bind(&user); util.IsError(err) {
		helpers.SetErrorMessage(c, "Something went wrong, please try again later or contact the system admin")
		return c.Redirect(http.StatusSeeOther, loginViewPath+"login.html")
	}

	if err := c.Validate(&user); util.IsError(err) {
		helpers.SetErrorMessage(c, "Error occured while validating your data!")
		log.Errorf("error occurred while validating login data: %v", err)
		return c.Redirect(http.StatusSeeOther, loginViewPath+"login.html")
	}

	endPoint := "/auth/login"
	resp, err := systems.AuthClient.Post(c, endPoint, user)

	if err != nil || resp.Code != http.StatusOK {
		helpers.SetErrorMessage(c, resp.Error)
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	var userAuth models.UserAuth
	helpers.Decode(resp.Data, &userAuth)

	if !hasAny(userAuth.UserACL.Roles, []string{"Action Officer", "Staff"}) {
		log.Errorf("user with user category %v attempted login", err)
		helpers.SetErrorMessage(c, "PAPERLESS is for staff only, please login into SOMA")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	auth.SetTokensAndSetCookies(userAuth.AuthToken, c)

	//Cache user acl
	key := helpers.GetACLKey(user.Email)
	err = helpers.StoreCache(key, userAuth)
	if util.IsError(err) {
		helpers.SetErrorMessage(c, "Authentication error has occured")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	if userAuth.User.CampusID == 0 {
		helpers.SetInfoMessage(c, "Campus settings")
		return c.Redirect(http.StatusSeeOther, "/auth/users/select-campus")
	}

	if userAuth.User.IsPasswordReset {
		helpers.SetInfoMessage(c, "Please Reset Your Password")
		return c.Redirect(http.StatusSeeOther, "/auth/forgot-password")
	}

	helpers.SetInfoMessage(c, "Welcome, you have successfully logged into your account")
	return c.Redirect(http.StatusSeeOther, "/training-frontend/home")
}

// Logout
func (lg *login) Logout(c echo.Context) error {
	_, _, email := auth.GetUserFromContext(c)
	aclKey := helpers.GetACLKey(email)
	helpers.ClearCache(aclKey) //clear permission cache
	auth.ClearSession(c)
	helpers.SetInfoMessage(c, "You have successfully logged out of your account")
	return c.Redirect(http.StatusSeeOther, "/auth/login")
}

func hasAny(s1 []string, s2 []string) bool {
	for _, a := range s1 {
		for _, b := range s2 {
			if a == b {
				return true
			}
		}
	}
	return false
}
