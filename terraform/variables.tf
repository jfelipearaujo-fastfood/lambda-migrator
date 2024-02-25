variable "region" {
  description = "The default region to use for AWS"
  default     = "us-east-1"
}

variable "tags" {
  description = "The default tags to use for AWS resources"
  type        = map(string)
  default = {
    App = "database"
  }
}

variable "project_name" {
  description = "The name of the project"
  default     = "fast-food"
}

variable "database_port" {
  description = "The port the database will listen on"
  default     = 5432
}

variable "database_username" {
  description = "The username to use for the database"
  default     = "fastfood"
}
