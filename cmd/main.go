package main

import (
	"log"

	"github.com/RamiroCuenca/go-jwt-notes/auth"
	"github.com/RamiroCuenca/go-jwt-notes/common/logger"
)

func main() {
	// Parse the certificates/keys
	err := auth.LoadCertificates("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("Could not load the certificates/keys. Error: %v", err)
	}

	// Init Zap logger so that we can use it all over the app
	logger.InitZapLogger()

	// Init postgres database
	// connection.NewPostgresClient()

	// Get routes
	mux := Routes()

	// Init server
	sv := NewServer(mux)

	// Run server
	logger.Log().Info("Server running over port :8000 ...")
	sv.Run()
}
