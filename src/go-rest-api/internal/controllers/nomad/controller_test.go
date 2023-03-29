package nomad

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"plusvasis/internal/templates"

	nomad "github.com/hashicorp/nomad/nomad/structs"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNomadClient struct {
	mock.Mock
}

func (m *MockNomadClient) Get(endpoint string) ([]byte, error) {
	args := m.Called(endpoint)
	return args.Get(0).([]byte), args.Error(1)
}
func (m *MockNomadClient) Post(endpoint string, reqBody *bytes.Buffer) ([]byte, error) {
	args := m.Called(endpoint, reqBody)
	return args.Get(0).([]byte), args.Error(1)
}
func (m *MockNomadClient) Delete(endpoint string) ([]byte, error) {
	args := m.Called(endpoint)
	return args.Get(0).([]byte), args.Error(1)
}

func TestGetJobs(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/jobs", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("uid", "test")

	// Mock NomadClient
	nomadJobs := []nomad.JobListStub{
		{
			ID: "test",
			Meta: map[string]string{
				"user": "test",
			},
		},
		{
			ID: "test2",
			Meta: map[string]string{
				"user": "test2",
			},
		},
	}
	jobsJson, _ := json.Marshal(nomadJobs)

	nomadClient := new(MockNomadClient)
	nomadClient.On("Get", "/jobs?meta=true").Return(jobsJson, nil)
	nomadController := NomadController{nomadClient}

	// Assertions
	expected := []nomad.JobListStub{
		{
			ID: "test",
			Meta: map[string]string{
				"user": "test",
			},
		},
	}
	expectedJson, _ := json.Marshal(expected)
	code := http.StatusOK
	if assert.NoError(t, nomadController.GetJobs(c)) {
		assert.Equal(t, code, rec.Code)
		assert.JSONEq(t, string(expectedJson), rec.Body.String())
	}
}

func TestCreateJob(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/jobs", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Mock NomadJob POST body
	nomadJob := templates.NomadJob{
		Name: "test",
	}
	nomadJobJson, _ := json.Marshal(nomadJob)
	c.Request().Body = io.NopCloser(bytes.NewBuffer(nomadJobJson))

	// Mock NomadClient
	nomadRegister := nomad.JobRegisterResponse{
		EvalID:          "test",
		EvalCreateIndex: 1,
		JobModifyIndex:  1,
	}
	nomadRegisterJson, _ := json.Marshal(nomadRegister)
	nomadClient := new(MockNomadClient)
	nomadClient.On("Post", "/jobs", mock.Anything).Return(nomadRegisterJson, nil)
	nomadController := NomadController{nomadClient}

	// Assertions
	expectedJson := nomadRegisterJson
	code := http.StatusOK
	if assert.NoError(t, nomadController.CreateJob(c)) {
		assert.Equal(t, code, rec.Code)
		assert.JSONEq(t, string(expectedJson), rec.Body.String())
	}
}
