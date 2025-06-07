package paperless

import (
	"training-frontend/server/systems/training-frontend/controllers"

	"github.com/labstack/echo/v4"
)

// WebRouters initialises web routes
func WebRouters(app *echo.Echo) {
	landing := app.Group("/")
	{
		landing.GET("", controllers.Index)
	}

	frontend := app.Group("/training-frontend")
	{
		frontend.GET("", controllers.Index)
		frontend.GET("/home", controllers.Index)
		frontend.GET("/default-error", controllers.DefaultErrorPage)
		frontend.GET("/no-exam-data", controllers.NoExaminationDataDefaultPage)
	}
}
