variable "tags" {
  type        = map(string)
  description = "The default tags to use for AWS resources"
  default = {
    App = "database"
  }
}

variable "region" {
  type        = string
  description = "The default region to use for AWS"
  default     = "us-east-1"
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
