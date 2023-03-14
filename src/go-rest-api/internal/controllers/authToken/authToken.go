package auth

import (
	auth "continens/internal/authToken"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetToken(c echo.Context) error {
	validToken, err := auth.GenerateJWT()
	if err != nil {
		fmt.Println("Error Occured")
		return err
	}

	return c.JSON(http.StatusOK, validToken)

}
