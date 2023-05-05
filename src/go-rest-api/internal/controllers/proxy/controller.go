package proxy

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	nomadController "plusvasis/internal/controllers/nomad"

	"github.com/gorilla/websocket"
	nomad "github.com/hashicorp/nomad/nomad/structs"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type NomadProxyController struct {
	Client nomadController.NomadClient
	Dialer DialerInterface
}

type DialerInterface interface {
	Dial(urlStr string, requestHeader http.Header) (WsConnInterface, *http.Response, error)
}

type DefaultDialer struct{}

func (d *DefaultDialer) Dial(urlStr string, requestHeader http.Header) (WsConnInterface, *http.Response, error) {
	return websocket.DefaultDialer.Dial(urlStr, requestHeader)
}

type WsConnInterface interface {
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
	Close() error
	SetReadDeadline(t time.Time) error
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const idleTimeout = 30 * time.Second

// AllocExec godoc
//
//	@Summary		AllocExec
//	@Description	Execute commands in a Nomad job allocation via WebSockets
//	@Tags			proxy
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string	true	"Allocation ID"
//	@Param			command	query	string	true	"Command to execute"
//	@Success		200
//	@Failure		401	{object}	echo.HTTPError
//	@Failure		404
//	@Failure		500
//	@Failure		502
//	@Security		BearerAuth
//	@Router			/job/{id}/exec [get]
func (n *NomadProxyController) AllocExec(c echo.Context) error {
	jobId := c.Param("id")
	command := c.QueryParam("command")
	if command == "" {
		return fmt.Errorf("missing query parameters")
	}

	uid := c.Get("uid").(string)
	if err := n.checkUserAllowed(uid, jobId); err != nil {
		return err
	}

	alloc, err := n.parseRunningAlloc(jobId)
	if err != nil {
		return err
	}

	baseURL := "wss://nomad.local.cawnj.dev/"
	path := "v1/client/allocation/" + alloc.ID + "/exec"
	queryParams := url.Values{}
	queryParams.Add("command", command)
	queryParams.Add("task", alloc.TaskGroup)
	queryParams.Add("tty", "true")
	queryParams.Add("ws_handshake", "true")

	url := baseURL + path + "?" + queryParams.Encode()

	nomadConn, _, err := n.Dialer.Dial(url, nil)
	if err != nil {
		return echo.ErrBadGateway
	}
	defer nomadConn.Close()

	clientConn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return errors.Wrap(err, "failed to upgrade connection")
	}
	defer clientConn.Close()

	err = nomadConn.SetReadDeadline(time.Now().Add(idleTimeout))
	if err != nil {
		return echo.ErrInternalServerError
	}
	err = clientConn.SetReadDeadline(time.Now().Add(idleTimeout))
	if err != nil {
		return echo.ErrInternalServerError
	}

	fmt.Printf("Started terminal session for job %s\n", jobId)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		n.forwardMessages(clientConn, nomadConn)
		wg.Done()
	}()
	go func() {
		n.forwardMessages(nomadConn, clientConn)
		wg.Done()
	}()
	wg.Wait()

	fmt.Printf("Stopped terminal session for job %s\n", jobId)
	return nil
}

func (n *NomadProxyController) forwardMessages(srcConn, dstConn WsConnInterface) {
	for {
		msgType, msg, err := srcConn.ReadMessage()
		if err != nil {
			break
		}
		err = dstConn.WriteMessage(msgType, msg)
		if err != nil {
			break
		}
		err = srcConn.SetReadDeadline(time.Now().Add(idleTimeout))
		if err != nil {
			break
		}
	}
}

func (n *NomadProxyController) parseRunningAlloc(jobId string) (*nomad.AllocListStub, error) {
	data, err := n.Client.Get(fmt.Sprintf("/job/%s/allocations", jobId))
	if err != nil {
		return nil, err
	}

	var allocs []nomad.AllocListStub
	err = json.Unmarshal(data, &allocs)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}
	for _, alloc := range allocs {
		if alloc.ClientStatus == "running" || alloc.ClientStatus == "pending" {
			return &alloc, nil
		}
	}

	return nil, echo.ErrNotFound
}

func (n *NomadProxyController) checkUserAllowed(uid, jobId string) error {
	data, err := n.Client.Get(fmt.Sprintf("/job/%s", jobId))
	if err != nil {
		return err
	}

	var job nomad.Job
	err = json.Unmarshal(data, &job)
	if err != nil {
		return echo.ErrInternalServerError
	}

	if job.Meta["user"] != uid {
		return echo.ErrUnauthorized
	}

	return nil
}

// StreamLogs godoc
//
//	@Summary		StreamLogs
//	@Description	Stream logs from a Nomad job allocation
//	@Tags			proxy
//	@Produce		json
//	@Param			id		path	string	true	"Job ID"
//	@Param			task	query	string	true	"Task name (same as job name)"
//	@Param			type	query	string	true	"Log type (stdout or stderr)"
//	@Success		200
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Failure		500
//	@Security		BearerAuth
//	@Router			/job/{id}/logs [get]
func (n *NomadProxyController) StreamLogs(c echo.Context) error {
	jobId := c.Param("id")
	logType := c.QueryParam("type")
	if logType == "" {
		return echo.ErrBadRequest
	}

	uid := c.Get("uid").(string)
	if err := n.checkUserAllowed(uid, jobId); err != nil {
		return err
	}

	alloc, err := n.parseRunningAlloc(jobId)
	if err != nil {
		return err
	}

	baseURL := "https://nomad.local.cawnj.dev/"
	path := "v1/client/fs/logs/" + alloc.ID
	queryParams := url.Values{}
	queryParams.Add("task", alloc.TaskGroup)
	queryParams.Add("type", logType)
	queryParams.Add("follow", "true")
	queryParams.Add("offset", "50000")
	queryParams.Add("origin", "end")

	url := baseURL + path + "?" + queryParams.Encode()
	resp, err := n.Client.ForwardRequest(c, url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("Started streaming %s for job %s\n", logType, jobId)

	err = streamResponse(c, resp)
	if err != nil {
		return err
	}

	fmt.Printf("Stopped streaming %s for job %s\n", logType, jobId)
	return nil
}

func streamResponse(c echo.Context, resp *http.Response) error {
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return echo.ErrInternalServerError
		}
		defer gzipReader.Close()
		resp.Body = gzipReader
	}

	buf := make([]byte, 8192)
	for {
		bytesRead, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return echo.ErrInternalServerError
		}

		select {
		case <-c.Request().Context().Done():
			return nil
		default:
		}

		_, err = c.Response().Writer.Write(buf[:bytesRead])
		if err != nil {
			return echo.ErrInternalServerError
		}

		if flusher, ok := c.Response().Writer.(http.Flusher); ok {
			flusher.Flush()
		}
	}

	return nil
}
