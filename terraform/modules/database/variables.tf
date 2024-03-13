variable "region" {
  type        = string
  description = "The default region to use for AWS"
}

variable "db_instance_class" {
  type        = string
  description = "The instance class to use for the database"
}

variable "db_engine" {
  type        = string
  description = "The database engine to use"
}

variable "db_engine_version" {
  type        = string
  description = "The version of the database engine to use"
}

variable "db_name" {
  type        = string
  description = "The name of the database"
}

variable "db_port" {
  type        = number
  description = "The port the database will listen on"
}

variable "db_username" {
  type        = string
  description = "The username for the database"
}

variable "vpc_id" {
  type        = string
  description = "The ID of the VPC"
}

variable "vpc_cidr_block" {
  type        = string
  description = "The CIDR block of the VPC"
}

variable "private_subnets" {
  type        = list(string)
  description = "The IDs of the private subnets"
}
