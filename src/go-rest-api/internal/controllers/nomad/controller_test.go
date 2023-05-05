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

func (m *MockNomadClient) ForwardRequest(c echo.Context, url string) (*http.Response, error) {
	return nil, nil
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
		Name:  "test",
		Image: "test",
		User:  "test",
		Shell: "test",
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
		Name:  "test",
		Image: "test",
		User:  "test",
		Shell: "test",
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
	nomadClient.On("Delete", "/job/"+jobName).Return(nomadDeregisterJson, nil)

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

func TestPurgeJob(t *testing.T) {
	// Setup
	jobName := "test"
	rec, c, nomadClient, nomadController := setup(http.MethodDelete, "/job/"+jobName)
	c.SetParamNames("id")
	c.SetParamValues(jobName)
	c.Set("uid", "test")
	c.QueryParams().Add("purge", "true")

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

func TestRestartJob(t *testing.T) {
	// Setup
	jobName := "test"
	rec, c, nomadClient, nomadController := setup(http.MethodPost, "/job/"+jobName+"/restart")
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
	nomadJobJson, _ := json.Marshal(nomadJob)
	nomadClient.On("Get", "/job/"+jobName).Return(nomadJobJson, nil) // CheckUserAllowed mocking

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

	restart := nomad.GenericResponse{}
	restartJson, _ := json.Marshal(restart)
	nomadClient.On("Post", "/client/allocation/test/restart", mock.Anything).Return(restartJson, nil)

	// Assertions
	expectedJson := restartJson
	expectedCode := http.StatusOK
	if assert.NoError(t, nomadController.RestartJob(c)) {
		assert.Equal(t, expectedCode, rec.Code)
		assert.JSONEq(t, string(expectedJson), rec.Body.String())
	}
}

func TestStartJob(t *testing.T) {
	// Setup
	jobName := "test"
	rec, c, nomadClient, nomadController := setup(http.MethodPost, "/job/"+jobName+"/start")
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
	nomadJobJson, _ := json.Marshal(nomadJob)
	nomadClient.On("Get", "/job/"+jobName).Return(nomadJobJson, nil) // CheckUserAllowed mocking

	nomadJobReq := templates.NomadJob{
		Name:  "test",
		Image: "test",
		User:  "test",
		Shell: "test",
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

	// Assertions
	expectedJson := nomadRegisterJson
	expectedCode := http.StatusOK
	if assert.NoError(t, nomadController.UpdateJob(c)) {
		assert.Equal(t, expectedCode, rec.Code)
		assert.JSONEq(t, string(expectedJson), rec.Body.String())
	}
}
