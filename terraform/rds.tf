data "aws_availability_zones" "available" {}

module "rds" {
  source  = "terraform-aws-modules/rds/aws"
  version = "~> 5.0"

  username = local.db_username
  password = local.db_password

  # When using RDS Proxy w/ IAM auth - Database must be username/password auth, not IAM
  iam_database_authentication_enabled = false

  identifier            = "db-${var.database_name}"
  engine                = "postgres"
  engine_version        = "15"
  family                = "postgres15"
  major_engine_version  = "15"
  instance_class        = "db.t3.micro"
  port                  = var.database_port
  apply_immediately     = true
  allocated_storage     = 20
  max_allocated_storage = 30

  db_subnet_group_name   = module.vpc.database_subnet_group
  vpc_security_group_ids = [module.rds_sg.security_group_id]
  multi_az               = true

  backup_retention_period = 0
  deletion_protection     = false
}
