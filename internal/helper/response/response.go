package response

import "github.com/labstack/echo/v4"

type AppResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(c echo.Context, status int, message string, data any) error {
	return c.JSON(status, echo.Map{"message": message, "data": data})
}
