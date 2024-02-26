module "rds_proxy" {
  source = "terraform-aws-modules/rds-proxy/aws"

  name          = "rds-proxy"
  iam_role_name = "rds-proxy-role"

  vpc_subnet_ids         = module.vpc.private_subnets
  vpc_security_group_ids = [module.rds_proxy_sg.security_group_id]

  endpoints = {
    read_write = {
      name                   = "read-write-endpoint"
      vpc_subnet_ids         = module.vpc.private_subnets
      vpc_security_group_ids = [module.rds_proxy_sg.security_group_id]
    },
    read_only = {
      name                   = "read-only-endpoint"
      vpc_subnet_ids         = module.vpc.private_subnets
      vpc_security_group_ids = [module.rds_proxy_sg.security_group_id]
      target_role            = "READ_ONLY"
    }
  }

  auth = {
    (local.db_username) = {
      description = aws_secretsmanager_secret.superuser.description
      secret_arn  = aws_secretsmanager_secret.superuser.arn
    }
  }

  engine_family = "POSTGRESQL"
  debug_logging = false

  # Target RDS instance
  target_db_instance     = true
  db_instance_identifier = "db-${var.database_name}"
}
