package main

import (
	"os"

	"github.com/Satssuki/CleanArchSimplified/cmd/application"
	config "github.com/Satssuki/CleanArchSimplified/pkg/system"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(config.CrossPlatformPath("config/dev.env"))
	application.RunServer(os.Getenv("PORT"))
}
