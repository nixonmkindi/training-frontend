package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	} else {
		code = 500
	}
	if code == http.StatusUnauthorized {
		c.Redirect(http.StatusSeeOther, "/auth/login")
	}
	errorPage := fmt.Sprintf("/views/error/%d.html", code)
	c.Render(code, errorPage, nil)
}
