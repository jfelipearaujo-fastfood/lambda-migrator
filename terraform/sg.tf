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
