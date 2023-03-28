package templates

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"text/template"
)

type NomadJob struct {
	Name    string     `json:"containerName"`
	Image   string     `json:"dockerImage"`
	User    string     `json:"user"`
	Shell   string     `json:"shell"`
	Volumes [][]string `json:"volumes"`
	Env     [][]string `json:"env"`
	Port    int        `json:"port"`
	Expose  bool       `json:"expose"`
}

func last(i int, slice interface{}) bool {
	v := reflect.ValueOf(slice)
	return i == v.Len()-1
}

func CreateJobJson(job NomadJob) (*bytes.Buffer, error) {
	t, err := template.New("").Funcs(template.FuncMap{
		"last": last,
	}).Parse(JOB_TMPL)
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
			"volumes": "{{range $i, $v := .Volumes}}{{index $v 0}}:{{index $v 1}}{{if not (last $i $.Volumes)}},{{end}}{{end}}",
			"env": "{{range $i, $v := .Env}}{{index $v 0}}={{index $v 1}}{{if not (last $i $.Env)}},{{end}}{{end}}",
			"port": "{{.Port}}"
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
						},
						"Env": {
							{{range $i, $v := .Env}}
							"{{index $v 0}}": "{{index $v 1}}"{{if not (last $i $.Env)}},{{end}}
							{{end}}
						}
					}
				],
				"Networks": [
					{
						"Mode": "host",
						"DynamicPorts": [
							{
								"Label": "http",
								"To": {{.Port}}
							}
						]
					}
				],
				"Services": [
					{
						"Name": "{{.Name}}",
						"PortLabel": "http",
						{{if .Expose}}
						"Tags": [
							"traefik.enable=true",
							"traefik.http.routers.{{.User}}-{{.Name}}.entrypoints=https",
							"traefik.http.routers.{{.User}}-{{.Name}}.rule=Host(` + "`" + `{{.User}}-{{.Name}}.local.plusvasis.xyz` + "`" + `)",
							"traefik.port=${NOMAD_PORT_http}"
						],
						{{end}}
						"Provider": "nomad"
					}
				]
			}
		]
	}
}`
