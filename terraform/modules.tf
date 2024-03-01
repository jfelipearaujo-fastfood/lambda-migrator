module "secret" {
  source = "./modules/secret"
}

module "database" {
  source = "./modules/database"

  region = var.region

  db_name     = "fastfood"
  db_port     = 5432
  db_username = module.secret.db_username
  db_password = module.secret.db_password

  vpc_id         = var.vpc_id
  vpc_cidr_block = var.vpc_cidr_block

  private_subnet_1a = var.private_subnet_1a
  private_subnet_1b = var.private_subnet_1b
  private_subnet_1c = var.private_subnet_1c
}

module "migrator" {
  source = "./modules/migrator"

  lambda_name = "migrator"

  db_host     = module.database.db_host
  db_name     = module.database.db_name
  db_username = module.secret.db_username
  db_password = module.secret.db_password

  private_subnet_1a = var.private_subnet_1a
  private_subnet_1b = var.private_subnet_1b
  private_subnet_1c = var.private_subnet_1c

  security_group_id = module.database.security_group_id
}
