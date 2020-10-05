package handler

import (
	"io/ioutil"

	"github.com/Satssuki/CleanArchSimplified/internal/domain"
	"github.com/Satssuki/CleanArchSimplified/internal/interfaces"
	"github.com/Satssuki/CleanArchSimplified/pkg/structs"
	"github.com/labstack/echo/v4"
)

// InsertUser route for inserting new user
func InsertUser(c echo.Context) error {
	logic := domain.CreateUserLogic()
	bytes, err := ioutil.ReadAll(c.Request().Body)
	var statusCode int = 400
	var responseMessage structs.Any

	if err == nil {
		err = interfaces.ComposeToConcreteObject(logic, &bytes)
		if err == nil {
			var message string
			message, err = logic.Register()

			if err == nil {
				statusCode, message = ParseDomainMessage(message)
				if statusCode >= 200 {
					responseMessage["message"] = message
				}
			}
		}
	}
	return c.JSON(statusCode, responseMessage)
}
