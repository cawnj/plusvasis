package templates

import (
	"bytes"
	"fmt"
	"text/template"
)

type NomadJob struct {
	ID    string
	Name  string
	Image string
	User  string
}

func CreateJobJson(job NomadJob) (*bytes.Buffer, error) {
	t, err := template.New("").Parse(JOB_TMPL)
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	err = t.Execute(buf, job)
	if err != nil {
		return nil, err
	}
	return buf, err
}

func CreateJobObject(name, image, user string) NomadJob {
	return NomadJob{
		ID:    fmt.Sprintf("%s-%s", user, name),
		Name:  name,
		Image: image,
		User:  user,
	}
}

const JOB_TMPL = `{
	"Job": {
		"ID": "{{.ID}}",
		"Name": "{{.ID}}",
		"Type": "service",
		"Datacenters": [
			"dc1"
		],
        "Meta": {
            "user": "{{.User}}"
        },
		"TaskGroups": [
			{
				"Name": "{{.Name}}",
				"Count": 1,
				"Tasks": [
					{
						"Name": "{{.Name}}",
						"Driver": "docker",
						"Config": {
							"image": "{{.Image}}",
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
								"To": 80
							}
						]
					}
				],
				"Services": [
					{
						"Name": "{{.Name}}",
						"PortLabel": "http",
						"Provider": "nomad"
					}
				]
			}
		]
	}
}`
