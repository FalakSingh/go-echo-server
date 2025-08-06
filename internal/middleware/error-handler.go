package appMiddleware

import (
	appError "echo-server/internal/helper/errors"
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
)

func CustomErrorHandler(err error, c echo.Context) {
	var (
		code    = http.StatusInternalServerError
		message string
	)

	if ae, ok := err.(*appError.AppError); ok {
		code = ae.Status
		message = ae.Message
	} else if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		// he.Message is of type interface{}, so we convert it to string
		if msgStr, ok := he.Message.(string); ok {
			message = msgStr
		} else {
			message = http.StatusText(code)
		}
	} else {
		message = err.Error()
	}

	log.Printf("Error: %v", err)

	customErr := &appError.AppError{
		Status:  code,
		Message: message, // now it's type-safe
	}

	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			_ = c.NoContent(code)
		} else {
			if err := c.JSON(code, customErr); err != nil {
				log.Printf("Failed to send error response: %v", err)
			}
		}
	}
}

