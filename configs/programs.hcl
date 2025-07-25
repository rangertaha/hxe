# Example program configurations for HXE

program "web-server" {
  description = "Nginx web server"
  exec        = "nginx -g 'daemon off;'"
  directory   = "/var/www"
  user        = "www-data"
  group       = "www-data"
  autostart   = true
  enabled     = true
  retries     = 3
}

program "api-server" {
  description = "Go API server"
  exec        = "go run main.go"
  directory   = "/opt/api"
  user        = "api"
  group       = "api"
  autostart   = true
  enabled     = true
  retries     = 5
}

program "database" {
  description = "PostgreSQL database"
  exec        = "postgres -D /var/lib/postgresql/data"
  directory   = "/var/lib/postgresql"
  user        = "postgres"
  group       = "postgres"
  autostart   = true
  enabled     = true
  retries     = 3
}

program "monitoring" {
  description = "Prometheus monitoring"
  exec        = "prometheus --config.file=/etc/prometheus/prometheus.yml"
  directory   = "/opt/prometheus"
  user        = "prometheus"
  group       = "prometheus"
  autostart   = false
  enabled     = true
  retries     = 2
} 