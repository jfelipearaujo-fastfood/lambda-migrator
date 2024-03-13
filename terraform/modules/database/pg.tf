resource "aws_db_parameter_group" "db_parameter_group" {
  name   = "db-pg-${var.db_name}"
  family = "postgres16"

  parameter {
    name  = "rds.force_ssl"
    value = "0" # Disable SSL
  }
}
