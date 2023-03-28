package templates

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type NomadJob struct {
	Name    string     `json:"containerName"`
	Image   string     `json:"dockerImage"`
	User    string     `json:"user"`
	Shell   string     `json:"shell"`
	Volumes [][]string `json:"volumes"`
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

	// output for debugging
	fmt.Printf("%+v\n", job)
	f, err := os.Create("latest-job.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	_, err = f.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}

	return buf, err
}

const JOB_TMPL = `{
	"Job": {
		"ID": "{{.User}}-{{.Name}}",
		"Name": "{{.Name}}",
		"Type": "service",
		"Datacenters": [
			"dc1"
		],
        "Meta": {
            "user": "{{.User}}",
			"shell": "{{.Shell}}",
			"volumes": "{{range $_, $v := .Volumes}}{{index $v 0}}:{{index $v 1}},{{end}}"
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
							],
							"mount": [
								{{range $_, $v := .Volumes}}
								{
									"type": "volume",
									"readonly": false,
									"source": "plusvasis-{{$.User}}-{{index $v 0}}",
									"target": "{{index $v 1}}"
								},
								{{end}}
								{
									"type": "volume",
									"readonly": false,
									"source": "plusvasis-{{.User}}",
									"target": "/userdata"
								}
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
