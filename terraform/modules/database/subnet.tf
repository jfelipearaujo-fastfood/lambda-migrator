resource "aws_db_subnet_group" "db_subnet_group" {
  name       = "db-sng-${var.db_name}"
  subnet_ids = var.private_subnets
}
