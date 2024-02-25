data "aws_availability_zones" "available" {}

module "db" {
  source = "terraform-aws-modules/rds/aws"

  identifier = "db-${local.name}"

  engine               = "postgres"
  engine_version       = "15"
  family               = "postgres15"
  major_engine_version = "15"
  instance_class       = "db.t3.micro"

  allocated_storage     = 20
  max_allocated_storage = 30

  db_name  = var.project_name
  username = var.database_username
  port     = var.database_port

  db_subnet_group_name   = module.vpc.database_subnet_group
  vpc_security_group_ids = [module.security_group.security_group_id]
}
