data "aws_caller_identity" "current" {}

data "aws_availability_zones" "available" {}

data "aws_s3_bucket" "bucket" {
  bucket = var.bucket_name
}

data "aws_kms_alias" "secretsmanager" {
  name = "alias/aws/secretsmanager"
}

resource "random_password" "password" {
  length  = 16
  special = false
}

locals {
  name = var.project_name

  db_password = random_password.password.result

  vpc_cidr = "10.0.0.0/16"
  azs      = slice(data.aws_availability_zones.available.names, 0, 3)
}
