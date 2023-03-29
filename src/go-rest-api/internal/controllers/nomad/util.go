package nomad

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"plusvasis/internal/templates"

	"github.com/labstack/echo/v4"
)

func decodeJobJson(job *templates.NomadJob, body io.ReadCloser) error {
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&job)

	if err != nil {
		if errors.As(err, &unmarshalErr) {
			return echo.NewHTTPError(http.StatusBadRequest, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field)
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "Bad Request "+err.Error())
		}
	}

	return nil
}
