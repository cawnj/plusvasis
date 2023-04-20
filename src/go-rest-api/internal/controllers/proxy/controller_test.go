package proxy

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net"
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
	return nil, nil
}
func (m *MockNomadClient) Delete(endpoint string) ([]byte, error) {
	return nil, nil
}
func (m *MockNomadClient) ForwardRequest(c echo.Context, url string) (*http.Response, error) {
	args := m.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}

type MockDialer struct {
	mock.Mock
}

func (m *MockDialer) Dial(urlStr string, requestHeader http.Header) (WsConnInterface, *http.Response, error) {
	args := m.Called(urlStr, requestHeader)
	return args.Get(0).(WsConnInterface), args.Get(1).(*http.Response), args.Error(2)
}

type MockConn struct {
	net.Conn
}

func (m *MockConn) ReadMessage() (messageType int, p []byte, err error) { return 0, nil, nil }
func (m *MockConn) WriteMessage(messageType int, data []byte) error     { return nil }
func (m *MockConn) Read(p []byte) (n int, err error)                    { return 0, nil }
func (m *MockConn) Write(p []byte) (n int, err error)                   { return 0, nil }
func (m *MockConn) Close() error                                        { return nil }
func (m *MockConn) SetDeadline(t time.Time) error                       { return nil }
func (m *MockConn) SetReadDeadline(t time.Time) error                   { return nil }
func (m *MockConn) SetWriteDeadline(t time.Time) error                  { return nil }
func (m *MockConn) LocalAddr() net.Addr                                 { return nil }
func (m *MockConn) RemoteAddr() net.Addr                                { return nil }

type HijackableResponseWriter struct {
	http.ResponseWriter
	Conn *MockConn
}

func (h *HijackableResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	rw := bufio.NewReadWriter(
		bufio.NewReader(
			bytes.NewBuffer(nil)), bufio.NewWriter(bytes.NewBuffer(nil)),
	)
	return h.Conn, rw, nil
}

func setup(method string, url string) (
	*httptest.ResponseRecorder, echo.Context, *MockNomadClient, NomadProxyController,
) {
	e := echo.New()
	req := httptest.NewRequest(method, url, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	nomadClient := new(MockNomadClient)
	dialer := new(MockDialer)
	controller := NomadProxyController{
		Client: nomadClient,
		Dialer: dialer,
	}

	return rec, c, nomadClient, controller
}

func setupWithHrw(method, url string) (
	*HijackableResponseWriter, echo.Context, *MockNomadClient, *MockDialer, NomadProxyController,
) {
	e := echo.New()
	req := createMockWsRequest(method, url)
	hrw := &HijackableResponseWriter{
		ResponseWriter: httptest.NewRecorder(),
		Conn:           &MockConn{},
	}
	c := e.NewContext(req, hrw)

	nomadClient := new(MockNomadClient)
	dialer := new(MockDialer)
	controller := NomadProxyController{
		Client: nomadClient,
		Dialer: dialer,
	}

	return hrw, c, nomadClient, dialer, controller
}

func createMockHttpResponse(statusCode int) *http.Response {
	recorder := httptest.NewRecorder()
	recorder.WriteHeader(statusCode)
	response := recorder.Result()
	return response
}

func createMockWsRequest(method, url string) *http.Request {
	req := httptest.NewRequest(method, url, nil)
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Upgrade", "websocket")
	req.Header.Set("Sec-Websocket-Version", "13")
	req.Header.Set("Sec-Websocket-Key", "test")
	return req
}

func TestAllocExec(t *testing.T) {
	// Setup
	jobName := "test"
	hrw, c, client, dialer, controller := setupWithHrw(http.MethodGet, "/job/"+jobName+"/exec")
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
	client.On("Get", "/job/"+jobName+"/allocations").Return(allocsJson, nil)

	httpResponse := createMockHttpResponse(http.StatusOK)
	dialer.On("Dial", mock.Anything, mock.Anything).Return(hrw.Conn, httpResponse, nil)

	// Assertions
	assert.NoError(t, controller.AllocExec(c))
}

func TestStreamLogs(t *testing.T) {
	// Setup
	jobName := "test"
	rec, c, client, controller := setup(http.MethodGet, "/job/"+jobName+"/logs")
	c.SetParamNames("id")
	c.SetParamValues(jobName)
	c.QueryParams().Set("task", "test")
	c.QueryParams().Set("type", "test")

	// Mocks
	nomadJobAlloc := []nomad.AllocListStub{
		{
			ID:           "test",
			ClientStatus: "running",
		},
	}
	allocsJson, _ := json.Marshal(nomadJobAlloc)
	client.On("Get", "/job/"+jobName+"/allocations").Return(allocsJson, nil)

	httpResponse := createMockHttpResponse(http.StatusOK)
	client.On("ForwardRequest", mock.Anything).Return(httpResponse, nil)

	// Assertions
	expectedCode := http.StatusOK
	if assert.NoError(t, controller.StreamLogs(c)) {
		assert.Equal(t, expectedCode, rec.Code)
	}
}
