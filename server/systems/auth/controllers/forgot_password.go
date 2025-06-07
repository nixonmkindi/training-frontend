package controllers

import (
	"net/http"
	"training-frontend/package/log"
	"training-frontend/server/systems"
	"training-frontend/server/systems/auth/auth"
	"training-frontend/server/systems/auth/models"
	"training-frontend/server/systems/helpers"

	"github.com/labstack/echo/v4"
)

const forgotPasswordViewPath = "/auth/views/auth/password/"

var ForgotPassword forgotPassword

type forgotPassword struct{}

// Show Forgot~ Password Form
func (fp *forgotPassword) ShowForgotPasswordForm(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		helpers.SetErrorMessage(c, "internal error occured")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	data := helpers.Map{
		"title":  "DIT Paperless | Reset Password",
		"userID": user.ID,
	}

	err := c.Render(http.StatusOK, forgotPasswordViewPath+"forgot_password.html", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error role rendering %v", err)
	}
	return nil
}

func (fp *forgotPassword) PasswordReset(c echo.Context) error {
	user := models.ID{}
	if err := c.Bind(&user); err != nil {
		helpers.SetErrorMessage(c, "internal error occured")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	data := helpers.Map{
		"title":  "DIT Paperless | Reset Password",
		"userID": user.ID,
	}

	err := c.Render(http.StatusOK, forgotPasswordViewPath+"change_password.html", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error role rendering %v", err)
	}
	return nil
}

func (fp *forgotPassword) ResetPassword(c echo.Context) error {
	endPoint := "/users/forgot-password"
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		helpers.SetErrorMessage(c, "internal error occured")
		return c.Redirect(http.StatusSeeOther, "/training-frontend/home")
	}

	user.IsPasswordReset = true
	userID, _, _ := auth.GetUserFromContext(c)
	user.UpdatedBy = userID

	resp, err := systems.AuthClient.Post(c, endPoint, user)
	if err != nil || resp.Code != http.StatusAccepted {
		log.Errorf("error user entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not update password")
		if user.IsHOD {
			return c.Redirect(http.StatusSeeOther, "/soma/staff/list-per-department")
		} else {
			return c.Redirect(http.StatusSeeOther, "/auth/users/list")
		}
	}

	helpers.SetInfoMessage(c, "Password updated successfully")
	if user.IsHOD {
		return c.Redirect(http.StatusSeeOther, "/soma/staff/list-per-department")
	} else {
		return c.Redirect(http.StatusSeeOther, "/auth/users/list")
	}
}

func (fp *forgotPassword) UserResetPassword(c echo.Context) error {
	endPoint := "/users/forgot-password"
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		helpers.SetErrorMessage(c, "internal error occured")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	userID, campusID, _ := auth.GetUserFromContext(c)
	params := models.User{
		ID:       userID,
		CampusID: campusID,
		Password: user.Password,
	}

	resp, err := systems.AuthClient.Post(c, endPoint, params)

	if err != nil || resp.Code != http.StatusAccepted {
		log.Errorf("error user entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not update password")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}
	helpers.SetInfoMessage(c, "Password updated successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/login")
}

func (fp *forgotPassword) ResetStudentPassword(c echo.Context) error {

	students := FetchUserByCategory(c, 1)
	data := helpers.Map{
		"students": students,
	}

	err := c.Render(http.StatusOK, forgotPasswordViewPath+"reset_student_password", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error role rendering %v", err)
	}
	return nil
}

func (fp *forgotPassword) ResetPasswordForStudent(c echo.Context) error {
	endPoint := "/users/forgot-password"
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		helpers.SetErrorMessage(c, "internal error occured")
		return c.Redirect(http.StatusSeeOther, "/auth/reset-student-password")
	}

	user.IsPasswordReset = true
	userID, _, _ := auth.GetUserFromContext(c)
	user.UpdatedBy = userID

	resp, err := systems.AuthClient.Post(c, endPoint, user)
	if err != nil || resp.Code != http.StatusAccepted {
		log.Errorf("error user entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not update password")
		return c.Redirect(http.StatusSeeOther, "/auth/reset-student-password")
	}

	helpers.SetInfoMessage(c, "Password updated successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/reset-student-password")
}

func (fp *forgotPassword) ResetStudentEmail(c echo.Context) error {

	students := FetchUserByCategory(c, 1)
	data := helpers.Map{
		"students": students,
	}

	err := c.Render(http.StatusOK, forgotPasswordViewPath+"reset_student_email", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error role rendering %v", err)
	}
	return nil
}

func (fp *forgotPassword) UpdateEmailForStudent(c echo.Context) error {
	endPoint := "/users/update-email"
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		helpers.SetErrorMessage(c, "internal error occured")
		return c.Redirect(http.StatusSeeOther, "/auth/reset-email-for-student")
	}

	userID, _, _ := auth.GetUserFromContext(c)
	user.UpdatedBy = userID

	resp, err := systems.AuthClient.Post(c, endPoint, user)
	if err != nil || resp.Code != http.StatusAccepted {
		log.Errorf("error user entity decoding %v", err)
		helpers.SetErrorMessage(c, "Error! Could not update password")
		return c.Redirect(http.StatusSeeOther, "/auth/reset-email-for-student")
	}

	helpers.SetInfoMessage(c, "Password updated successfully")
	return c.Redirect(http.StatusSeeOther, "/auth/reset-email-for-student")
}

func (fp *forgotPassword) ResetStaffPassword(c echo.Context) error {
	user := models.ID{}
	if err := c.Bind(&user); err != nil {
		helpers.SetErrorMessage(c, "internal error occured")
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	data := helpers.Map{
		"title":  "DIT Paperless | Reset Password",
		"userID": user.ID,
		"isHOD":  true,
	}

	err := c.Render(http.StatusOK, forgotPasswordViewPath+"change_password.html", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error role rendering %v", err)
	}
	return nil
}
