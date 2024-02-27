resource "aws_security_group" "security_group" {
  name_prefix = "db-sg-"
  description = "Default security group for ${var.project_name} PostgreSQL database allowing access from private network."
  vpc_id      = module.vpc.vpc_id

  ingress {
    from_port = var.database_port
    to_port   = var.database_port
    protocol  = "tcp"
    cidr_blocks = [
      module.vpc.vpc_cidr_block
    ]
  }

  dynamic "ingress" {
    for_each = [1]
    content {
      self      = true
      from_port = 0
      to_port   = 65535
      protocol  = "tcp"
    }
  }

  egress {
    from_port = 0
    to_port   = 0
    protocol  = "-1"
    cidr_blocks = [
      "0.0.0.0/0"
    ]
  }

  lifecycle {
    create_before_destroy = true
  }
}

# module "security_group" {
#   source  = "terraform-aws-modules/security-group/aws"
#   version = "~> 5.0"

#   name        = local.name
#   description = "PostgreSQL security group"
#   vpc_id      = module.vpc.vpc_id

#   # ingress
#   ingress_with_cidr_blocks = [
#     {
#       from_port   = 5432
#       to_port     = 5432
#       protocol    = "tcp"
#       description = "PostgreSQL access from within VPC"
#       cidr_blocks = module.vpc.vpc_cidr_block
#     }
#     # ,{
#     #   description = "Private subnet PostgreSQL access"
#     #   rule        = "postgresql-tcp"
#     #   cidr_blocks = join(",", module.vpc.private_subnets_cidr_blocks)
#     # }
#   ]
# }

# module "rds_proxy_security_group" {
#   source  = "terraform-aws-modules/security-group/aws"
#   version = "~> 5.0"

#   name        = "rds_proxy"
#   description = "PostgreSQL RDS Proxy security group"
#   vpc_id      = module.vpc.vpc_id

#   revoke_rules_on_delete = true

#   ingress_with_cidr_blocks = [
#     {
#       description = "Private subnet PostgreSQL access"
#       rule        = "postgresql-tcp"
#       cidr_blocks = join(",", module.vpc.private_subnets_cidr_blocks)
#     }
#   ]

#   egress_with_cidr_blocks = [
#     {
#       description = "Database subnet PostgreSQL access"
#       rule        = "postgresql-tcp"
#       cidr_blocks = join(",", module.vpc.database_subnets_cidr_blocks)
#     },
#   ]
# }
