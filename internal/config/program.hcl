// Program Configuration
//

program "example" {
  name = "example-service"
  description = "An example service"
  command = "echo"
  args = ["Hello, World!"]
  working_dir = "/tmp"
  user = "nobody"
  group = "nobody"
  autostart = false
  enabled = true
  retries = 3
  max_retries = 5
}

