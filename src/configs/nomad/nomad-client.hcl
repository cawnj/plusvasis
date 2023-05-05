data_dir = "/opt/nomad/data"
bind_addr = "0.0.0.0"

name = "sonic"

advertise {
  http = "192.168.1.202"
  rpc  = "192.168.1.202"
  serf = "192.168.1.202"
}

client {
  enabled = true
  servers = ["192.168.1.201"]
}

plugin "docker" {
  config {
    volumes {
      enabled = true
    }
  }
}