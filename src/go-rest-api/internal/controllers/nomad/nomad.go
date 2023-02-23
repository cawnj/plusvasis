package nomad

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

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

func GenerateJobObject() *bytes.Buffer {

	data := []byte(`{
		"Job": {
			"ID": "nginx-test",
			"Name": "nginx-test",
			"Type": "service",
			"Datacenters": [
				"dc1"
			],
			"TaskGroups": [
				{
					"Name": "nginx-test",
					"Count": 1,
					"Tasks": [
						{
							"Name": "server",
							"Driver": "docker",
							"Config": {
								"image": "nginx",
								"ports": [
									"http"
								]
							}
						}
					],
					"Networks": [
						{
							"Mode": "bridge",
							"DynamicPorts": [
								{
									"Label": "http",
									"Value": 0,
									"To": 80
								}
							]
						}
					],
					"Services": [
						{
							"Name": "nginx-test",
							"PortLabel": "http",
							"Provider": "nomad"
						}
					]
				}
			]
		}
	}`)
	return bytes.NewBuffer(data)
}

func CreateJob(c echo.Context) error {

	resp, err := http.Post("https://nomad.local.cawnj.dev/v1/jobs", "application/json", GenerateJobObject())
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

	id := c.Param("id")
	url := fmt.Sprintf("https://nomad.local.cawnj.dev/v1/job/%s", id)
	resp, err := http.Post(url, "application/json", GenerateJobObject())
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
	respBody, err := io.ReadAll(resp.Body)
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
