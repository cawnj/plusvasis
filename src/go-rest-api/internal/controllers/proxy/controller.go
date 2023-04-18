package proxy

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"plusvasis/internal/controllers/nomad"

	"github.com/gorilla/websocket"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/labstack/echo/v4"
)

type NomadProxyController struct {
	Client nomad.NomadClient
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

const idleTimeout = 30 * time.Second

func (n *NomadProxyController) AllocExec(c echo.Context) error {
	id := c.Param("id")
	rawQuery := c.Request().URL.RawQuery
	if rawQuery == "" {
		return c.String(http.StatusBadRequest, "Missing params")
	}

	alloc, err := n.parseRunningAlloc(id)
	if err != nil {
		return c.String(http.StatusNotFound, "Allocation not found")
	}

	nomadURL := "wss://nomad.local.cawnj.dev/v1/client/allocation/" + alloc.ID + "/exec"
	nomadURL += "?" + rawQuery

	nomadConn, _, err := websocket.DefaultDialer.Dial(nomadURL, nil)
	if err != nil {
		return c.String(http.StatusBadGateway, "Error connecting to Nomad")
	}
	nomadConn.SetReadDeadline(time.Now().Add(idleTimeout))
	defer nomadConn.Close()

	clientConn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error upgrading request to WebSocket")
	}
	clientConn.SetReadDeadline(time.Now().Add(idleTimeout))
	defer clientConn.Close()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		log.Printf("Forwarding messages from client to nomad for job %s", id)
		n.forwardMessages(clientConn, nomadConn, "client")
		wg.Done()
	}()
	go func() {
		log.Printf("Forwarding messages from nomad to client for job %s", id)
		n.forwardMessages(nomadConn, clientConn, "nomad")
		wg.Done()
	}()
	wg.Wait()

	log.Printf("Closing connections for job %s", id)
	return nil
}

func (n *NomadProxyController) forwardMessages(srcConn, dstConn *websocket.Conn, name string) {
	for {
		msgType, msg, err := srcConn.ReadMessage()
		if err != nil {
			break
		}
		err = dstConn.WriteMessage(msgType, msg)
		if err != nil {
			break
		}
		log.Printf("%s: %s\n", name, msg)
		srcConn.SetReadDeadline(time.Now().Add(idleTimeout))
	}
}

func (n *NomadProxyController) parseRunningAlloc(jobId string) (*structs.AllocListStub, error) {
	data, err := n.Client.Get(fmt.Sprintf("/job/%s/allocations", jobId))
	if err != nil {
		return nil, err
	}

	var allocs []structs.AllocListStub
	err = json.Unmarshal(data, &allocs)
	if err != nil {
		return nil, err
	}
	for _, alloc := range allocs {
		if alloc.ClientStatus == "running" || alloc.ClientStatus == "pending" {
			return &alloc, nil
		}
	}

	return nil, fmt.Errorf("no running alloc found for job %s", jobId)
}
