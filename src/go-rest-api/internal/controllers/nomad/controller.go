package nomad

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"plusvasis/internal/templates"

	nomad "github.com/hashicorp/nomad/nomad/structs"
	"github.com/labstack/echo/v4"
)

type NomadController struct {
	Client NomadClient
}

func (n *NomadController) GetJobs(c echo.Context) error {
	data, err := n.Client.Get("/jobs?meta=true")
	if err != nil {
		log.Println("[nomad/GetJobs]", err)
		return err
	}

	var jobs []nomad.JobListStub
	err = json.Unmarshal(data, &jobs)
	if err != nil {
		log.Println("[nomad/GetJobs]", err)
		return err
	}

	var filteredJobs []nomad.JobListStub
	uid := c.Get("uid").(string)
	for _, job := range jobs {
		if job.Meta["user"] == uid {
			filteredJobs = append(filteredJobs, job)
		}
	}

	return c.JSON(http.StatusOK, filteredJobs)
}

func (n *NomadController) CreateJob(c echo.Context) error {
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

	data, err := n.Client.Post("/jobs", body)
	if err != nil {
		log.Println("[nomad/CreateJob]", err)
		return err
	}

	var resp nomad.JobRegisterResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (n *NomadController) UpdateJob(c echo.Context) error {
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

	if err := n.CheckUserAllowed(uid, jobId); err != nil {
		return err
	}

	data, err := n.Client.Post(fmt.Sprintf("/job/%s", jobId), body)
	if err != nil {
		log.Println("[nomad/UpdateJob]", err)
		return err
	}

	var resp nomad.JobRegisterResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (n *NomadController) ReadJob(c echo.Context) error {
	uid := c.Get("uid").(string)
	jobId := c.Param("id")

	data, err := n.Client.Get(fmt.Sprintf("/job/%s", jobId))
	if err != nil {
		log.Println("[nomad/ReadJob]", err)
		return err
	}

	var job nomad.Job
	err = json.Unmarshal(data, &job)
	if err != nil {
		log.Println("[nomad/ReadJob]", err)
		return err
	}

	// Doing this check here because if we use n.CheckUserAllowed, we will duplicate requests
	if job.Meta["user"] != uid {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, job)
}

func (n *NomadController) StopJob(c echo.Context) error {
	uid := c.Get("uid").(string)
	jobId := c.Param("id")
	purge := c.QueryParam("purge")

	if err := n.CheckUserAllowed(uid, jobId); err != nil {
		return err
	}

	url := fmt.Sprintf("/job/%s", jobId)
	if purge == "true" {
		url += "?purge=true"
	}
	data, err := n.Client.Delete(url)
	if err != nil {
		log.Println("[nomad/StopJob]", err)
		return err
	}

	var resp nomad.JobDeregisterResponse
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (n *NomadController) ReadJobAllocs(c echo.Context) error {
	uid := c.Get("uid").(string)
	jobId := c.Param("id")

	if err := n.CheckUserAllowed(uid, jobId); err != nil {
		return err
	}

	data, err := n.Client.Get(fmt.Sprintf("/job/%s/allocations", jobId))
	if err != nil {
		log.Println("[nomad/ReadJobAllocs]", err)
		return err
	}

	var allocs []nomad.AllocListStub
	err = json.Unmarshal(data, &allocs)
	if err != nil {
		log.Println("[nomad/ReadJobAllocs]", err)
		return err
	}

	return c.JSON(http.StatusOK, allocs)
}

func (n *NomadController) ReadJobAlloc(c echo.Context) error {
	uid := c.Get("uid").(string)
	jobId := c.Param("id")

	if err := n.CheckUserAllowed(uid, jobId); err != nil {
		return err
	}

	data, err := n.Client.Get(fmt.Sprintf("/job/%s/allocations", jobId))
	if err != nil {
		log.Println("[nomad/ReadJobAllocs]", err)
		return err
	}

	var allocs []nomad.AllocListStub
	err = json.Unmarshal(data, &allocs)
	if err != nil {
		log.Println("[nomad/ReadJobAllocs]", err)
		return err
	}
	for _, alloc := range allocs {
		if alloc.ClientStatus == "running" || alloc.ClientStatus == "pending" {
			return c.JSON(http.StatusOK, alloc)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "No running allocation found")
}

func (n *NomadController) RestartJob(c echo.Context) error {
	uid := c.Get("uid").(string)
	jobId := c.Param("id")
	allocId := c.Param("allocId")
	task := c.Param("task")

	if err := n.CheckUserAllowed(uid, jobId); err != nil {
		fmt.Println(err)
		return err
	}

	body := []byte(`{
		"TaskName": "` + task + `"
	}`)

	data, err := n.Client.Post(fmt.Sprintf("/client/allocation/%s/restart", allocId), bytes.NewBuffer(body))
	if err != nil {
		log.Println("[nomad/RestartJob]", err)
		return err
	}

	return c.JSON(http.StatusOK, bytes.NewBuffer(data))
}

func (n *NomadController) CheckUserAllowed(uid, jobId string) error {
	data, err := n.Client.Get(fmt.Sprintf("/job/%s", jobId))
	if err != nil {
		return err
	}

	var job nomad.Job
	err = json.Unmarshal(data, &job)
	if err != nil {
		return err
	}

	if job.Meta["user"] != uid {
		return echo.ErrUnauthorized
	}

	return nil
}
