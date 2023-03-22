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

func checkUserAllowed(uid, jobId string) error {
	data, err := nomadGet(fmt.Sprintf("/job/%s", jobId))
	if err != nil {
		return err
	}

	var job structs.Job
	err = json.Unmarshal(data, &job)
	if err != nil {
		return err
	}

	if job.Meta["user"] != uid {
		return echo.ErrUnauthorized
	}

	return nil
}

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

func nomadPost(endpoint string, reqBody *bytes.Buffer) ([]byte, error) {
	url := NOMAD_URL + endpoint
	resp, err := http.Post(url, "application/json", reqBody)
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

func nomadDelete(endpoint string) ([]byte, error) {
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
	if resp.StatusCode != 200 {
		return nil, echo.NewHTTPError(resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func decodeJobJson(job *templates.NomadJob, body io.ReadCloser) error {
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&job)

	if err != nil {
		if errors.As(err, &unmarshalErr) {
			return echo.NewHTTPError(http.StatusBadRequest, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field)
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "Bad Request "+err.Error())
		}
	}

	return nil
}

func GetJobs(c echo.Context) error {
	data, err := nomadGet("/jobs?meta=true")
	if err != nil {
		log.Println("[nomad/GetJobs]", err)
		return err
	}

	var jobs []structs.JobListStub
	err = json.Unmarshal(data, &jobs)
	if err != nil {
		log.Println("[nomad/GetJobs]", err)
		return err
	}

	var filteredJobs []structs.JobListStub
	uid := c.Get("uid").(string)
	for _, job := range jobs {
		if job.Meta["user"] == uid {
			filteredJobs = append(filteredJobs, job)
		}
	}

	return c.JSON(http.StatusOK, filteredJobs)
}

func CreateJob(c echo.Context) error {
	var job templates.NomadJob
	err := decodeJobJson(&job, c.Request().Body)
	if err != nil {
		return err
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

	var resp structs.JobRegisterResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func UpdateJob(c echo.Context) error {
	var job templates.NomadJob
	err := decodeJobJson(&job, c.Request().Body)
	if err != nil {
		return err
	}

	body, err := templates.CreateJobJson(job)
	if err != nil {
		log.Println("[nomad/UpdateJob]", err)
		return err
	}

	uid := c.Get("uid").(string)
	jobId := c.Param("id")

	if err := checkUserAllowed(uid, jobId); err != nil {
		return err
	}

	data, err := nomadPost(fmt.Sprintf("/job/%s", jobId), body)
	if err != nil {
		log.Println("[nomad/UpdateJob]", err)
		return err
	}

	var resp structs.JobRegisterResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func ReadJob(c echo.Context) error {
	uid := c.Get("uid").(string)
	jobId := c.Param("id")

	data, err := nomadGet(fmt.Sprintf("/job/%s", jobId))
	if err != nil {
		log.Println("[nomad/ReadJob]", err)
		return err
	}

	var job structs.Job
	err = json.Unmarshal(data, &job)
	if err != nil {
		log.Println("[nomad/ReadJob]", err)
		return err
	}

	// Doing this check here because if we use checkUserAllowed, we will duplicate requests
	if job.Meta["user"] != uid {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, job)
}

func StopJob(c echo.Context) error {
	uid := c.Get("uid").(string)
	jobId := c.Param("id")

	if err := checkUserAllowed(uid, jobId); err != nil {
		return err
	}

	data, err := nomadDelete(fmt.Sprintf("/job/%s?purge=true", jobId))
	if err != nil {
		log.Println("[nomad/StopJob]", err)
		return err
	}

	var resp structs.JobRegisterResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func ReadJobAllocs(c echo.Context) error {
	uid := c.Get("uid").(string)
	jobId := c.Param("id")

	if err := checkUserAllowed(uid, jobId); err != nil {
		return err
	}

	data, err := nomadGet(fmt.Sprintf("/job/%s/allocations", jobId))
	if err != nil {
		log.Println("[nomad/ReadJobAllocs]", err)
		return err
	}

	var allocs []structs.AllocListStub
	err = json.Unmarshal(data, &allocs)
	if err != nil {
		log.Println("[nomad/ReadJobAllocs]", err)
		return err
	}

	return c.JSON(http.StatusOK, allocs)
}
