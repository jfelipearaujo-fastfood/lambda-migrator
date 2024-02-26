resource "random_pet" "users" {
  length    = 2
  separator = "_"
}

resource "random_password" "password" {
  length  = 16
  special = false
}

locals {
  name = var.project_name

  # using random here due to secrets taking at least 7 days before fully deleting from account
  db_username = random_pet.users.id
  db_password = random_password.password.result

  vpc_cidr = "10.0.0.0/16"
  azs      = slice(data.aws_availability_zones.available.names, 0, 3)
}
