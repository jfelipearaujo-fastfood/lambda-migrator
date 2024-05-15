module "migrator" {
  source = "./modules/migrator"

  lambda_name = "migrator"

  vpc_name = var.vpc_name
}
