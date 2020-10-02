package application

import (
	"os"

	"github.com/Satssuki/CleanArchSimplified/cmd/application/handler"
	"github.com/Satssuki/CleanArchSimplified/internal/infrastructure"
	"github.com/labstack/echo/v4"
)

// RunServer run server with specified port
func RunServer(port string) {
	infrastructure.SetupMongoConnection(os.Getenv("DBNAME"), os.Getenv("MONGOURL"))
	client, err := infrastructure.CreateMongoConnection()

	if err != nil {
		panic(err)
	}

	infrastructure.SetDefaultClient(client)
	server := echo.New()
	defer server.Start(port)
	registerRoute(server)

}

func registerRoute(root *echo.Echo) {
	root.POST("/user", handler.InsertUser)
}
