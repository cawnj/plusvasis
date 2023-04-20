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

func (m *MockNomadClient) Post(endpoint string, body *bytes.Buffer) ([]byte, error) { return nil, nil }
func (m *MockNomadClient) Delete(endpoint string) ([]byte, error)                   { return nil, nil }

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

func setup(method, url string) (
	*HijackableResponseWriter, echo.Context, *MockNomadClient, *MockDialer, NomadProxyController,
) {
	e := echo.New()
	req := createMockWsRequest(method, url)
	rec := &HijackableResponseWriter{
		ResponseWriter: httptest.NewRecorder(),
		Conn:           &MockConn{},
	}
	c := e.NewContext(req, rec)

	nomadClient := new(MockNomadClient)
	dialer := new(MockDialer)
	nomadController := NomadProxyController{
		Client: nomadClient,
		Dialer: dialer,
	}

	return rec, c, nomadClient, dialer, nomadController
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
	rec, c, nomadClient, dialer, nomadProxyController := setup(http.MethodGet, "/job/"+jobName+"/exec")
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

	httpResponse := createMockHttpResponse(http.StatusOK)
	dialer.On("Dial", mock.Anything, mock.Anything).Return(rec.Conn, httpResponse, nil)

	// Assertions
	assert.NoError(t, nomadProxyController.AllocExec(c))
}
