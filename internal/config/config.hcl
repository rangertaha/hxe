// hxe Configuration
//


// Logging
log {
  level = "info"
  format = "json"
}

// REST API server
api {
  addr = "0.0.0.0"
  port = 8080
  username = "admin"
  password = "password"
  # cors     = true
  # jwt_secret = "change-this-secret"
}
 
// Messaging broker
broker { 
  name = "hxe"
  addr = "0.0.0.0"
  port = 7070
}

// (Required) RDS client connection
// Relational database is used to store the state of the system
// By default, SQLite is used. Other options include PostgreSQL, MySQL
db "sqlite" {
  # path    = "./hxe.db"
  # dsn     = ""       # For postgres: "host=localhost user=postgres dbname=hxe sslmode=disable"
  migrate = true
}

// (Optional) Timeseries database client connection
// A timeseries database used to store historical metrics
// By default, InfluxDB is used. In the future, other options may be supported
tsdb "influxdb" {
  host = "localhost"
  port = 8086
  username = "influxdb"
  password = "influxdb"
  token = "influxdb"
}


