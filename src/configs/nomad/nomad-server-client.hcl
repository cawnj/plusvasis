data_dir = "/opt/nomad/data"
bind_addr = "0.0.0.0"

name = "leo"

advertise {
  http = "192.168.1.201"
  rpc  = "192.168.1.201"
  serf = "192.168.1.201"
}

server {
  enabled = true
  bootstrap_expect = 1
}

client {
  enabled = true
  servers = ["127.0.0.1"]
}

plugin "docker" {
  config {
    volumes {
      enabled = true
    }
  }
}
