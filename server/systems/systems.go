package systems

import (
	"training-frontend/package/client"
	"training-frontend/package/config"
	"training-frontend/package/log"
	"training-frontend/package/util"

	"github.com/k0kubun/pp"
)

// TODO: move these into config.yml configuration file
const (
	AuthBaseUrl    = "http://localhost:4320/auth/api/v1"
	BackendBaseUrl = "http://localhost:4317/training-backend/api/v1"
	AuditBaseUrl   = "http://localhost:4331/audit/api/v1"
)

var AuthClient *client.Client
var BackendClient *client.Client

func Init() {
	pp.Printf("Initialising Training Frontend Client...\n")
	cfg, _ := config.New()

	//FRONTEND
	frontendSystem := "training_frontend"
	frontendKey, err := cfg.GetSystemPrivateKey(frontendSystem)
	if util.IsError(err) {
		log.Errorf("error getting keys: %v", err)
		panic("system is not well started")
	}

	AuthClient, err = client.New(AuthBaseUrl, frontendKey, frontendSystem)
	if util.IsError(err) {
		log.Errorf("error initializing auth client: %v", err)
		panic("auth system client could not be initiated")
	}

	BackendClient, err = client.New(BackendBaseUrl, frontendKey, frontendSystem)
	if util.IsError(err) {
		log.Errorf("error initializing backend client: %v", err)
		panic("backend client could not be initiated")
	}

	pp.Printf("Training Frontend Client Initialised...\n")
}
