package proxy

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

type MockDialer struct {
	mock.Mock
}

func (m *MockDialer) Dial(urlStr string, requestHeader http.Header) (WsConnInterface, *http.Response, error) {
	args := m.Called(urlStr, requestHeader)
	return args.Get(0).(WsConnInterface), args.Get(1).(*http.Response), args.Error(2)
}

type MockWsConn struct {
	mock.Mock
}

func (m *MockWsConn) ReadMessage() (messageType int, p []byte, err error) {
	args := m.Called()
	return args.Int(0), args.Get(1).([]byte), args.Error(2)
}

func (m *MockWsConn) WriteMessage(messageType int, data []byte) error {
	args := m.Called(messageType, data)
	return args.Error(0)
}

func (m *MockWsConn) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockWsConn) SetReadDeadline(t time.Time) error {
	args := m.Called(t)
	return args.Error(0)
}

func setup(method string, url string) (
	*httptest.ResponseRecorder, echo.Context, *MockNomadClient, *MockDialer, NomadProxyController,
) {
	e := echo.New()
	req := httptest.NewRequest(method, url, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	nomadClient := new(MockNomadClient)
	dialer := new(MockDialer)
	nomadController := NomadProxyController{
		Client: nomadClient,
		Dialer: dialer,
	}

	return rec, c, nomadClient, dialer, nomadController
}

func TestAllocExec(t *testing.T) {
	// Setup
	jobName := "test"
	rec, c, nomadClient, dialer, nomadProxyController := setup(http.MethodGet, "/job/"+jobName)
	c.SetParamNames("id")
	c.SetParamValues(jobName)
	c.QueryParams().Set("command", "[\"/bin/bash\"]")

	// Mocks
	nomadJobAlloc := []nomad.AllocListStub{
		{
			ID:           "test",
			ClientStatus: "running",
		},
	}
	allocsJson, _ := json.Marshal(nomadJobAlloc)
	nomadClient.On("Get", "/job/"+jobName+"/allocations").Return(allocsJson, nil)

	mockWsConn := new(MockWsConn)
	dialer.On("Dial", mock.Anything, mock.Anything).Return(mockWsConn, nil, nil)

	// Assertions
	expectedCode := http.StatusOK
	if assert.NoError(t, nomadProxyController.AllocExec(c)) {
		assert.Equal(t, expectedCode, rec.Code)
	}
}
