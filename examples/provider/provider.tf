terraform {
  required_providers {
    task-list = {
      source = "svanellewee/task-list"
    }
  }
}

provider "task-list" {
  # example configuration here
}

resource "task" "example" {
  provider = task-list
  description = "Yes this works!"
}
