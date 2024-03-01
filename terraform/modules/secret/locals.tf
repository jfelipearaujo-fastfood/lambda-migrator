data "aws_kms_alias" "secretsmanager" {
  name = "alias/aws/secretsmanager"
}

resource "random_pet" "users" {
  length    = 2
  separator = "_"
}

resource "random_password" "password" {
  length  = 16
  special = false
}

locals {
  db_username = random_pet.users.id
  db_password = random_password.password.result
}
