package controllers

import (
	"net/http"
	"time"
	"training-frontend/package/log"
	"training-frontend/server/systems/helpers"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	data := helpers.Map{
		"title": "Training Frontend|Index",
		"test":  "1234",
		"today": time.Now().Format("02/01/2006"),
	}

	err := c.Render(http.StatusOK, "training-frontend/views/index", helpers.Serve(c, data))
	if err != nil {
		log.Errorf("error rendering paperless page %v", err)
	}
	return err
}

func DefaultErrorPage(c echo.Context) error {
	data := helpers.Map{}

	return c.Render(http.StatusOK, "training-frontend/views/default_error.html", helpers.Serve(c, data))
}

func NoExaminationDataDefaultPage(c echo.Context) error {
	data := helpers.Map{}

	return c.Render(http.StatusOK, "training-frontend/views/no_exam_data.html", helpers.Serve(c, data))
}
