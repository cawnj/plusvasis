package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	expectedJson := `{"Status":"OK"}`
	expectedCode := http.StatusOK

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", strings.NewReader(expectedJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, Health(c)) {
		assert.Equal(t, expectedCode, rec.Code)
		assert.JSONEq(t, expectedJson, rec.Body.String())
	}
}
