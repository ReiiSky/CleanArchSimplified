package handler

import (
	"io/ioutil"

	"github.com/Satssuki/CleanArchSimplified/internal/domain"
	"github.com/Satssuki/CleanArchSimplified/internal/interfaces"
	"github.com/labstack/echo/v4"
)

// InsertUser route for inserting new user
func InsertUser(c echo.Context) error {
	logic := domain.CreateUserLogic()
	bytes, err := ioutil.ReadAll(c.Request().Body)
	var statusCode int = 400
	if err == nil {
		err = interfaces.ComposeToConcreteObject(logic, &bytes)
		if err == nil {
			var status string
			status, err = logic.Register()

			if err == nil {
				statusCode, message := ParseDomainMessage(status)
				if statusCode >= 200 {
					return c.String(statusCode, message)
				}
			}
		}
	}
	return c.String(statusCode, err.Error())
}
