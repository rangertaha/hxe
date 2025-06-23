// Service Configuration
//
group "example" {
  name = "example-service-group"
  description = "An example group"
  services = [service.example1, service.example2]
}

service "example1" {
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

service "example2" {
  name = "example-service-2"
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


task "example1" {
  name = "example-task"
  description = "An example task"
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

task "example2" {
  depends = [task.example1]
  name = "example-task-2"
  description = "An example task"
  command = "echo"
  args = ["Hello, World!"]
}


// Schedule execution of tasks, services, groups, and events
schedule "example1" {
  name = "example-schedule"
  description = "An example schedule"
  
  schedule = "0 0 * * *"
  groups = [group.example-task-group]
  tasks = [task.example-task1, task.example-task2]
  events = [event.example-event1]
  services = [service.example1, service.example2]
  enabled = true
}

event "example1" {
  name = "example-event"
  description = "An example event"
  publish = [service.example1, tasks.example2]
  data = {
    "key1" = "value1"
    "key2" = "value2"
  }
  payload = '{"key1": "value1", "key2": "value2"}'
}

