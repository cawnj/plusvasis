package auth

import (
	auth "continens/internal/authToken"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetToken(c echo.Context) error {
	var k auth.APIKey

	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&k)

	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(c.Response().Writer, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(c.Response().Writer, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
	}
	validToken, err := auth.GenerateJWT(k)
	if err != nil {
		fmt.Println("[authToken/errorResponse]: " + err.Error())
		return err
	}

	return c.JSON(http.StatusOK, validToken)
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	_, err := w.Write(jsonResp)
	if err != nil {
		log.Println("[authToken/errorResponse]", err)
	}
}
