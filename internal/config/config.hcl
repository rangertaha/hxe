// hxe Configuration
//

// REST API server
api {
  addr = "0.0.0.0"
  port = 8080
  username = "admin"
  password = "password"
}
 
// Message Queue: Embedded NATS server to use for messaging
// default to IPC internal messaging
mq "nats" {
    ipc  = true
    name = "hxe"
    addr = "0.0.0.0"
    port = 7070
}

// Relational Database: (Required) RDS client connection
// Relational database is used to store the state of the system
// By default, SQLite is used. Other options include PostgreSQL, MySQL
rdb "sqlite" {
  # path    = "./hxe.db"
  # dsn     = ""       # For postgres: "host=localhost user=postgres dbname=hxe sslmode=disable"
  migrate = true
}

// Timeseries Database: (Optional) Timeseries database client connection
// A timeseries database used to store historical metrics
// By default, InfluxDB is used. In the future, other options may be supported
tdb "influxdb" {
  host = "localhost"
  port = 8086
  username = "influxdb"
  password = "influxdb"
  token = "influxdb"
}


