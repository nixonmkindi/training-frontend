package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//Index page
func Viewer(c echo.Context) error {
	err := c.Render(http.StatusOK, "/viewer/viewer.html", echo.Map{"title": "DICOM Viewer"})
	fmt.Println(err)
	return err
}
