package nomad

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"continens/internal/templates"

	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/labstack/echo/v4"
)

const NOMAD_URL = "https://nomad.local.cawnj.dev/v1"

func nomadGet(endpoint string) ([]byte, error) {
	url := NOMAD_URL + endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, echo.NewHTTPError(resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func nomadPost(endpoint string, reqBody *bytes.Buffer) (interface{}, error) {
	url := NOMAD_URL + endpoint
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func nomadDelete(endpoint string) (interface{}, error) {
	url := NOMAD_URL + endpoint
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	_, err := w.Write(jsonResp)
	if err != nil {
		log.Println("[nomad/errorResponse]", err)
	}
}

func GetJobs(c echo.Context) error {
	data, err := nomadGet("/jobs?meta=true")
	if err != nil {
		log.Println("[nomad/GetJobs]", err)
		return err
	}
	return c.JSON(http.StatusOK, data)
}

func CreateJob(c echo.Context) error {
	var job templates.NomadJob
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&job)

	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(c.Response().Writer, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(c.Response().Writer, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
	}

	body, err := templates.CreateJobJson(job)
	if err != nil {
		log.Println("[nomad/CreateJob]", err)
		return err
	}

	data, err := nomadPost("/jobs", body)
	if err != nil {
		log.Println("[nomad/CreateJob]", err)
		return err
	}
	return c.JSON(http.StatusOK, data)
}

func UpdateJob(c echo.Context) error {
	var job templates.NomadJob
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&job)

	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(c.Response().Writer, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(c.Response().Writer, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
	}

	body, err := templates.CreateJobJson(job)
	if err != nil {
		log.Println("[nomad/CreateJob]", err)
		return err
	}

	data, err := nomadPost(fmt.Sprintf("/job/%s", c.Param("id")), body)
	if err != nil {
		log.Println("[nomad/UpdateJob]", err)
		return err
	}
	return c.JSON(http.StatusOK, data)
}

func ReadJob(c echo.Context) error {
	data, err := nomadGet(fmt.Sprintf("/job/%s", c.Param("id")))
	if err != nil {
		log.Println("[nomad/ReadJob]", err)
		return err
	}
	if data == nil {
		return echo.ErrNotFound
	}

	var job structs.Job
	err = json.Unmarshal(data, &job)
	if err != nil {
		log.Println("[nomad/ReadJob]", err)
		return err
	}
	if job.Meta["user"] != c.Get("uid") {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, job)
}

func StopJob(c echo.Context) error {
	data, err := nomadDelete(fmt.Sprintf("/job/%s?purge=true", c.Param("id")))
	if err != nil {
		log.Println("[nomad/StopJob]", err)
		return err
	}
	return c.JSON(http.StatusOK, data)
}

func ReadJobAllocs(c echo.Context) error {

	data, err := nomadGet(fmt.Sprintf("/job/%s/allocations", c.Param("id")))
	if err != nil {
		log.Println("[nomad/ReadJobAllocs]", err)
		return err
	}

	if data == nil {
		return c.JSONBlob(http.StatusBadRequest, []byte(`{ "Response" : "Job Not Found" }`))
	}
	return c.JSON(http.StatusOK, data)
}
