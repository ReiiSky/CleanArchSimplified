package infrastructure

import (
	"github.com/Satssuki/CleanArchSimplified/pkg/system"
	"github.com/joho/godotenv"
)

// LoadDevEnvironment will load development environment variable
func LoadDevEnvironment() {
	godotenv.Load(system.CrossPlatformPath("config/dev.env"))
}
