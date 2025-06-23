// hxe Configuration
//

api {
  addr = "0.0.0.0"
  port = 8080
  username = "hxe"
  password = "hxe"

  client {
    url = "http://localhost:8080"
    token = "token"
    username = "username"
    password = "password"
  }
}

broker { 
  name = "hxe"
  addr = "0.0.0.0"
  port = 8080
}
   
database {
  type = "sqlite"
  migrate = true
}

