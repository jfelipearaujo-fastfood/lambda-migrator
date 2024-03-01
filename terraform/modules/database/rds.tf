module "db" {
  source  = "terraform-aws-modules/rds/aws"
  version = "5.9.0"

  identifier = var.db_name

  engine               = "postgres"
  engine_version       = "14"
  family               = "postgres14"
  major_engine_version = "14"
  instance_class       = "db.t3.micro"

  allocated_storage     = 20
  max_allocated_storage = 30

  iam_database_authentication_enabled = false

  db_name  = var.db_name
  username = var.db_username
  password = var.db_password
  # username = local.db_username
  # password          = jsondecode(aws_secretsmanager_secret_version.superuser.secret_string)["password"]
  port              = var.db_port
  apply_immediately = true

  multi_az = false

  create_db_subnet_group = false
  subnet_ids = [
    var.private_subnet_1a,
    var.private_subnet_1b,
    var.private_subnet_1c
  ]

  vpc_security_group_ids = [
    aws_security_group.security_group.id
  ]

  backup_retention_period = 0
  deletion_protection     = false
}
