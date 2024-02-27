# module "db" {
#   source  = "terraform-aws-modules/rds/aws"
#   version = "~> 5.0"

#   identifier = var.project_name

#   engine               = "postgres"
#   engine_version       = "14"
#   family               = "postgres14"
#   major_engine_version = "14"
#   instance_class       = "db.t3.micro"

#   allocated_storage     = 20
#   max_allocated_storage = 30

#   iam_database_authentication_enabled = false

#   db_name           = var.database_name
#   username          = local.db_username
#   password          = local.db_password
#   port              = var.database_port
#   apply_immediately = true

#   multi_az               = false
#   db_subnet_group_name   = module.vpc.database_subnet_group
#   vpc_security_group_ids = [module.security_group.security_group_id]

#   backup_retention_period = 0
#   deletion_protection     = false
# }
