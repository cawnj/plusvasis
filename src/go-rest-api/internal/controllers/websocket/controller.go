package websocket

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func AllocExec(c echo.Context) error {
	id := c.Param("id")
	rawQuery := c.Request().URL.RawQuery
	if rawQuery == "" {
		return c.String(http.StatusBadRequest, "Missing params")
	}

	nomadURL := "wss://nomad.local.cawnj.dev/v1/client/allocation/" + id + "/exec"
	nomadURL += "?" + rawQuery

	nomadConn, _, err := websocket.DefaultDialer.Dial(nomadURL, nil)
	if err != nil {
		return c.String(http.StatusBadGateway, "Error connecting to Nomad")
	}
	defer nomadConn.Close()

	clientConn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error upgrading request to WebSocket")
	}
	defer clientConn.Close()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		forwardMessages(clientConn, nomadConn)
		wg.Done()
	}()
	go func() {
		forwardMessages(nomadConn, clientConn)
		wg.Done()
	}()
	wg.Wait()

	return nil
}

func forwardMessages(srcConn, dstConn *websocket.Conn) {
	for {
		msgType, msg, err := srcConn.ReadMessage()
		if err != nil {
			break
		}
		err = dstConn.WriteMessage(msgType, msg)
		if err != nil {
			break
		}
	}
}
