package application

import (
	"os"

	"github.com/Satssuki/CleanArchSimplified/cmd/application/handler"
	"github.com/Satssuki/CleanArchSimplified/internal/infrastructure"
	"github.com/labstack/echo/v4"
)

// RunServer run server with specified port
func RunServer(port string) error {
	infrastructure.LoadDevEnvironment()
	infrastructure.SetupMongoConnection(os.Getenv("DBNAME"), os.Getenv("MONGOURL"))
	client, err := infrastructure.CreateMongoConnection()
	err = infrastructure.MongoClientPing(client)

	if err == nil {
		infrastructure.SetDefaultClient(client)
		server := echo.New()
		registerRoute(server)

		server.Start(port)
	}
	return err
}

func registerRoute(root *echo.Echo) {
	root.POST("/user", handler.InsertUser)
}
