variable "region" {
  type        = string
  description = "The default region to use for AWS"
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

variable "db_password" {
  type        = string
  description = "The password for the database"
}

variable "vpc_id" {
  type        = string
  description = "The ID of the VPC"
}

variable "vpc_cidr_block" {
  type        = string
  description = "The CIDR block of the VPC"
}

variable "private_subnet_1a" {
  type        = string
  description = "The ID of the private subnet in the first availability zone"
}

variable "private_subnet_1b" {
  type        = string
  description = "The ID of the private subnet in the second availability zone"
}

variable "private_subnet_1c" {
  type        = string
  description = "The ID of the private subnet in the third availability zone"
}
