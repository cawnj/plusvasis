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

func setup(method string, url string) (
	*httptest.ResponseRecorder, echo.Context, *MockNomadClient, NomadController,
) {
	e := echo.New()
	req := httptest.NewRequest(method, url, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	nomadClient := new(MockNomadClient)
	nomadController := NomadController{nomadClient}

	return rec, c, nomadClient, nomadController
}

func TestGetJobs(t *testing.T) {
	// Setup
	rec, c, nomadClient, nomadController := setup(http.MethodGet, "/jobs")
	c.Set("uid", "test")

	// Mocks
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
	nomadClient.On("Get", "/jobs?meta=true").Return(jobsJson, nil)

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
	expectedCode := http.StatusOK
	if assert.NoError(t, nomadController.GetJobs(c)) {
		assert.Equal(t, expectedCode, rec.Code)
		assert.JSONEq(t, string(expectedJson), rec.Body.String())
	}
}

func TestCreateJob(t *testing.T) {
	// Setup
	rec, c, nomadClient, nomadController := setup(http.MethodPost, "/jobs")

	// Mocks
	nomadJob := templates.NomadJob{
		Name: "test",
	}
	nomadJobJson, _ := json.Marshal(nomadJob)
	c.Request().Body = io.NopCloser(bytes.NewBuffer(nomadJobJson))

	nomadRegister := nomad.JobRegisterResponse{
		EvalID:          "test",
		EvalCreateIndex: 1,
		JobModifyIndex:  1,
	}
	nomadRegisterJson, _ := json.Marshal(nomadRegister)
	nomadClient.On("Post", "/jobs", mock.Anything).Return(nomadRegisterJson, nil)

	// Assertions
	expectedJson := nomadRegisterJson
	expectedCode := http.StatusOK
	if assert.NoError(t, nomadController.CreateJob(c)) {
		assert.Equal(t, expectedCode, rec.Code)
		assert.JSONEq(t, string(expectedJson), rec.Body.String())
	}
}

func TestUpdateJob(t *testing.T) {
	// Setup
	jobName := "test"
	rec, c, nomadClient, nomadController := setup(http.MethodPost, "/job/"+jobName)
	c.SetParamNames("id")
	c.SetParamValues(jobName)
	c.Set("uid", "test")

	// Mocks
	nomadJobReq := templates.NomadJob{
		Name: "test",
		User: "test",
	}
	nomadJobReqJson, _ := json.Marshal(nomadJobReq)
	c.Request().Body = io.NopCloser(bytes.NewBuffer(nomadJobReqJson))

	nomadRegister := nomad.JobRegisterResponse{
		EvalID:          "test",
		EvalCreateIndex: 1,
		JobModifyIndex:  1,
	}
	nomadRegisterJson, _ := json.Marshal(nomadRegister)
	nomadClient.On("Post", "/job/"+jobName, mock.Anything).Return(nomadRegisterJson, nil)

	nomadJobCheckUser := nomad.Job{
		ID: "test",
		Meta: map[string]string{
			"user": "test",
		},
	}
	nomadJobCheckUserJson, _ := json.Marshal(nomadJobCheckUser)
	nomadClient.On("Get", "/job/"+jobName).Return(nomadJobCheckUserJson, nil)

	// Assertions
	expectedJson := nomadRegisterJson
	expectedCode := http.StatusOK
	if assert.NoError(t, nomadController.UpdateJob(c)) {
		assert.Equal(t, expectedCode, rec.Code)
		assert.JSONEq(t, string(expectedJson), rec.Body.String())
	}
}

func TestReadJob(t *testing.T) {
	// Setup
	jobName := "test"
	rec, c, nomadClient, nomadController := setup(http.MethodGet, "/job/"+jobName)
	c.SetParamNames("id")
	c.SetParamValues(jobName)
	c.Set("uid", "test")

	// Mocks
	nomadJob := nomad.Job{
		ID: "test",
		Meta: map[string]string{
			"user": "test",
		},
	}

	jobJson, _ := json.Marshal(nomadJob)
	nomadClient.On("Get", "/job/"+jobName).Return(jobJson, nil)

	// Assertions
	expectedCode := http.StatusOK
	if assert.NoError(t, nomadController.ReadJob(c)) {
		assert.Equal(t, expectedCode, rec.Code)
		assert.JSONEq(t, string(jobJson), rec.Body.String())
	}
}

func TestStopJob(t *testing.T) {
	// Setup
	jobName := "test"
	rec, c, nomadClient, nomadController := setup(http.MethodDelete, "/job/"+jobName)
	c.SetParamNames("id")
	c.SetParamValues(jobName)
	c.Set("uid", "test")

	// Mocks
	nomadDeregister := nomad.JobDeregisterResponse{
		EvalID:          "test",
		EvalCreateIndex: 1,
		JobModifyIndex:  1,
	}
	nomadDeregisterJson, _ := json.Marshal(nomadDeregister)
	nomadClient.On("Delete", "/job/"+jobName+"?purge=true").Return(nomadDeregisterJson, nil)

	nomadJob := nomad.Job{
		ID: "test",
		Meta: map[string]string{
			"user": "test",
		},
	}
	nomadJobJson, _ := json.Marshal(nomadJob)
	nomadClient.On("Get", "/job/"+jobName).Return(nomadJobJson, nil) // CheckUserAllowed mocking

	// Assertions
	expectedJson := nomadDeregisterJson
	expectedCode := http.StatusOK
	if assert.NoError(t, nomadController.StopJob(c)) {
		assert.Equal(t, expectedCode, rec.Code)
		assert.JSONEq(t, string(expectedJson), rec.Body.String())
	}
}

func TestReadJobAllocs(t *testing.T) {
	// Setup
	jobName := "test"
	rec, c, nomadClient, nomadController := setup(http.MethodGet, "/job/"+jobName+"/allocations")
	c.SetParamNames("id")
	c.SetParamValues(jobName)
	c.Set("uid", "test")

	// Mocks
	nomadJobAllocs := []nomad.AllocListStub{
		{
			ID: "test",
		},
	}
	allocsJson, _ := json.Marshal(nomadJobAllocs)
	nomadClient.On("Get", "/job/"+jobName+"/allocations").Return(allocsJson, nil)

	nomadJob := nomad.Job{
		ID: "test",
		Meta: map[string]string{
			"user": "test",
		},
	}
	nomadJobJson, _ := json.Marshal(nomadJob)
	nomadClient.On("Get", "/job/"+jobName).Return(nomadJobJson, nil) // CheckUserAllowed mocking

	// Assertions
	expectedCode := http.StatusOK
	if assert.NoError(t, nomadController.ReadJobAllocs(c)) {
		assert.Equal(t, expectedCode, rec.Code)
		assert.JSONEq(t, string(allocsJson), rec.Body.String())
	}
}

func TestReadJobAlloc(t *testing.T) {
	// Setup
	jobName := "test"
	rec, c, nomadClient, nomadController := setup(http.MethodGet, "/job/"+jobName+"/alloc")
	c.SetParamNames("id")
	c.SetParamValues(jobName)
	c.Set("uid", "test")

	// Mocks
	nomadJobAllocs := []nomad.AllocListStub{
		{
			ID:           "test",
			ClientStatus: "running",
		},
		{
			ID:           "test2",
			ClientStatus: "walking",
		},
	}
	allocsJson, _ := json.Marshal(nomadJobAllocs)
	nomadClient.On("Get", "/job/"+jobName+"/allocations").Return(allocsJson, nil)

	nomadJob := nomad.Job{
		ID: "test",
		Meta: map[string]string{
			"user": "test",
		},
	}
	nomadJobJson, _ := json.Marshal(nomadJob)
	nomadClient.On("Get", "/job/"+jobName).Return(nomadJobJson, nil) // CheckUserAllowed mocking

	// Assertions
	expected := []nomad.AllocListStub{
		{
			ID:           "test",
			ClientStatus: "running",
		},
	}

	expectedJson, _ := json.Marshal(expected[0])
	expectedCode := http.StatusOK
	if assert.NoError(t, nomadController.ReadJobAlloc(c)) {
		assert.Equal(t, expectedCode, rec.Code)
		assert.JSONEq(t, string(expectedJson), rec.Body.String())
	}
}
