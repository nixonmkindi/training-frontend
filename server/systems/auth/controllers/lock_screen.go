package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const lockScreenViewPath = "/auth/"

var LockScreen lockScreen

type lockScreen struct{}

// Show Lock Screen
func (ls *lockScreen) ShowLockScreen(c echo.Context) error {

	err := c.Render(http.StatusOK, lockScreenViewPath+"lockscreen.html", echo.Map{
		"title": "DIT Paperless | Lock Screen",
	})
	return err
}
