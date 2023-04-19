package nomad

import (
	"bytes"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

const NOMAD_URL = "https://nomad.local.cawnj.dev/v1"

type NomadClient interface {
	Get(endpoint string) ([]byte, error)
	Post(endpoint string, reqBody *bytes.Buffer) ([]byte, error)
	Delete(endpoint string) ([]byte, error)
}

type DefaultNomadClient struct{}

func (n *DefaultNomadClient) Get(endpoint string) ([]byte, error) {
	url := NOMAD_URL + endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}
	if resp.StatusCode != 200 {
		return nil, echo.NewHTTPError(resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}
	return body, nil
}

func (n *DefaultNomadClient) Post(endpoint string, reqBody *bytes.Buffer) ([]byte, error) {
	url := NOMAD_URL + endpoint
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}
	if resp.StatusCode != 200 {
		return nil, echo.NewHTTPError(resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}
	return body, nil
}

func (n *DefaultNomadClient) Delete(endpoint string) ([]byte, error) {
	url := NOMAD_URL + endpoint
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}
	if resp.StatusCode != 200 {
		return nil, echo.NewHTTPError(resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}
	return body, nil
}
