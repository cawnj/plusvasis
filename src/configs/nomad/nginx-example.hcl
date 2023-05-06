# a simple nomad job that runs an nginx web server
# this job will will be exposed by traefik at "https://nginx.plusvasis.xyz"

job "nginx" {
  datacenters = ["dc1"]
  type = "service"
  group "nginx" {
    count = 1
    network {
      mode = "bridge" 
      port "http" {
        to = 80
      }
    }
    service {
      name = "nginx"
      port = "http"
      provider = "nomad"
      tags = [
        "traefik.enable=true",
        "traefik.http.routers.nginx.entrypoints=https",
        "traefik.http.routers.nginx.rule=Host(`nginx.plusvasis.xyz`)",
      ]
    }
    task "server" {
      driver = "docker"
      config {
        image = "nginx"
        ports = ["http"]
      }
    }
  }
}
