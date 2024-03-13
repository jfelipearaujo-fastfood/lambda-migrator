module "database" {
  source = "./modules/database"

  region = var.region

  db_name           = "fastfood"
  db_engine         = "postgres"
  db_engine_version = "16"
  db_instance_class = "db.t3.micro"
  db_port           = 5432
  db_username       = var.db_username

  vpc_id         = var.vpc_id
  vpc_cidr_block = var.vpc_cidr_block

  private_subnets = var.private_subnets
}

module "migrator" {
  source = "./modules/migrator"

  lambda_name = "migrator"

  db_host     = module.database.db_host
  db_name     = module.database.db_name
  db_username = module.database.db_username
  db_password = module.database.db_pass

  private_subnets   = var.private_subnets
  security_group_id = module.database.security_group_id
}
