package nomad

import (
	"bytes"
	"encoding/json"
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

func CreateJob(c echo.Context) error {

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

	resp, err := http.Post("https://nomad.local.cawnj.dev/v1/jobs", "application/json",
		bytes.NewBuffer(data))

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
