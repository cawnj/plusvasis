package test

import (
	"continens/internal/controllers/nomad"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	mockJob = `'test': 'job'`
)

func health(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}

func TestHealth(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/health")

	// Assertions
	if assert.NoError(t, health(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetJobs(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/jobs")
	c.Set("uid", "test")

	// Assertions
	if assert.NoError(t, nomad.GetJobs(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
