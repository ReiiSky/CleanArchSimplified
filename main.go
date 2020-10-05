package main

import (
	"os"

	"github.com/Satssuki/CleanArchSimplified/cmd/application"
)

func main() {
	application.RunServer(os.Getenv("PORT"))
}
