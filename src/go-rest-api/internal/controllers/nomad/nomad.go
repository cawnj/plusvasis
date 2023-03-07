package nomad

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"continens/internal/templates"

	"github.com/labstack/echo/v4"
)

func GetJobs(c echo.Context) error {

	resp, err := http.Get("https://nomad.local.cawnj.dev/v1/jobs")
	if err != nil {
		log.Println("[nomad/GetJobs]", err)
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[nomad/GetJobs]", err)
		return err
	}
	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("[nomad/GetJobs]", err)
		return err
	}
	return c.JSON(http.StatusOK, data)
}

func CreateJob(c echo.Context) error {
	job := templates.CreateJobObject("nginx-test", "nginx", "cawnj")
	data, err := templates.CreateJobJson(job)
	if err != nil {
		log.Println("[nomad/CreateJob]", err)
		return err
	}

	resp, err := http.Post("https://nomad.local.cawnj.dev/v1/jobs", "application/json", data)
	if err != nil {
		log.Println("[nomad/CreateJob]", err)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[nomad/CreateJob]", err)
		return err
	}
	var res interface{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println("[nomad/CreateJob]", err)
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func UpdateJob(c echo.Context) error {
	job := templates.CreateJobObject("nginx-test", "nginx", "cawnj")
	data, err := templates.CreateJobJson(job)
	if err != nil {
		log.Println("[nomad/CreateJob]", err)
		return err
	}

	id := c.Param("id")
	url := fmt.Sprintf("https://nomad.local.cawnj.dev/v1/job/%s", id)
	resp, err := http.Post(url, "application/json", data)
	if err != nil {
		log.Println("[nomad/UpdateJob]", err)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[nomad/UpdateJob]", err)
		return err
	}
	var res interface{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println("[nomad/UpdateJob]", err)
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func ReadJob(c echo.Context) error {
	id := c.Param("id")
	url := fmt.Sprintf("https://nomad.local.cawnj.dev/v1/job/%s", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("[nomad/ReadJob]", err)
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[nomad/ReadJob]", err)
		return err
	}
	var data interface{}
	err = json.Unmarshal(body, &data)
	if data != nil {
		if err != nil {
			log.Println("[nomad/ReadJob]", err)
			return err
		}
	}
	if data == nil {
		encodedJSON := []byte(`{
			"Response" : "Job Not Found"
		}`)

		return c.JSONBlob(http.StatusBadRequest, encodedJSON)
	}
	return c.JSON(http.StatusOK, data)
}

func StopJob(c echo.Context) error {
	id := c.Param("id")
	url := fmt.Sprintf("https://nomad.local.cawnj.dev/v1/job/%s", id)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Println("[nomad/StopJob]", err)
		return err
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[nomad/StopJob]", err)
		return err
	}
	defer resp.Body.Close()

	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[nomad/StopJob]", err)
		panic(err.Error())
	}
	var data interface{}
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		log.Println("[nomad/StopJob]", err)
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func ReadJobAllocs(c echo.Context) error {
	id := c.Param("id")
	url := fmt.Sprintf("https://nomad.local.cawnj.dev/v1/job/%s/allocations", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("[nomad/ReadJobAllocs]", err)
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[nomad/ReadJobAllocs]", err)
		return err
	}
	var data interface{}
	err = json.Unmarshal(body, &data)
	if data != nil {
		if err != nil {
			log.Println("[nomad/ReadJobAllocs]", err)
			return err
		}
	}
	if data == nil {
		encodedJSON := []byte(`{
			"Response" : "Job Not Found"
		}`)

		return c.JSONBlob(http.StatusBadRequest, encodedJSON)
	}
	return c.JSON(http.StatusOK, data)
}
