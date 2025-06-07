package server

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"training-frontend/package/config"
	"training-frontend/package/log"
	"training-frontend/package/util"
	"training-frontend/server/routes"
	"training-frontend/server/systems"
	"training-frontend/server/systems/helpers"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	// Echo instance
	e := echo.New()
	//Define renderer
	e.Renderer = Renderer()

	//Disable echo banner
	e.HideBanner = true

	//handle all errors
	//e.HTTPErrorHandler = controllers.CustomHTTPErrorHandler

	// Routes
	routes.Routers(e)
	cfg, err := config.New()
	if err != nil {
		log.Errorf("error creating a new config")
		return
	}

	e.Server.IdleTimeout = 5 * time.Minute

	generateRoutes(e)
	helpers.Init()
	systems.Init()
	address := fmt.Sprintf("%s:%d", cfg.WebServer.Host, cfg.WebServer.Port)
	e.Logger.Fatal(e.Start(address))
}

// this function generates all routes and store them in json format
func generateRoutes(e *echo.Echo) {
	//generate json route and store it into /cmd/permission
	path, _ := os.Getwd() //get working directory
	path = path + "/.storage/routes/routes.json"

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if !util.IsError(err) {
		os.WriteFile(path, data, 0644)
	}
}

// GOOS=linux GOARCH=amd64 go build -o bin/paperless
