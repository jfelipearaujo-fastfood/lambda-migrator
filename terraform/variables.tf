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

variable "bucket_name" {
}

variable "project_name" {
  description = "The name of the project"
  default     = "fastfood"
}

variable "database_name" {
  description = "The name of the database"
  default     = "fastfood"
}

variable "database_username" {
  description = "The username of the database"
  default     = "fastfood"
}

variable "database_port" {
  description = "The port the database will listen on"
  default     = 5432
}

variable "lambda_name" {
  description = "The name of the lambda function"
  default     = "lambda_migrator"
}
