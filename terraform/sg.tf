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

module "security_group_lambda" {
  name = var.lambda_name

  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  description = "Security Group for Lambda Egress"

  vpc_id = module.vpc.vpc_id

  egress_cidr_blocks      = []
  egress_ipv6_cidr_blocks = []

  # Prefix list ids to use in all egress rules in this module
  egress_prefix_list_ids = [module.vpc_endpoints.endpoints["s3"]["prefix_list_id"]]

  egress_rules = ["https-443-tcp"]
}
