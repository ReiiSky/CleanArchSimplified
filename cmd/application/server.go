package application

import (
	"os"

	"github.com/Satssuki/CleanArchSimplified/cmd/application/handler"
	"github.com/Satssuki/CleanArchSimplified/internal/infrastructure"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// RunServer run server with specified port
func RunServer(port string) {
	infrastructure.SetupMongoConnection(os.Getenv("DBNAME"), os.Getenv("MONGOURL"))
	client, err := infrastructure.CreateMongoConnection()
	err = infrastructure.MongoClientPing(client)

	if err != nil {
		log.
			Panic().
			Msgf("error connect to database with message '%v'", err)

		panic(err)
	} else {
		log.
			Info().
			Msg("mongo successfully connected")
	}

	infrastructure.SetDefaultClient(client)
	server := echo.New()
	registerRoute(server)

	log.
		Info().
		Msgf("server running on port %v", port)
	server.Start(port)
}

func registerRoute(root *echo.Echo) {
	root.POST("/user", handler.InsertUser)
}
