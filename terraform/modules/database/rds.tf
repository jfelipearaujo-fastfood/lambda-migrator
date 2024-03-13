resource "aws_db_instance" "db" {
  identifier = var.db_name

  instance_class              = var.db_instance_class
  allocated_storage           = 10
  engine                      = var.db_engine
  engine_version              = var.db_engine_version
  db_name                     = var.db_name
  port                        = var.db_port
  username                    = var.db_username
  manage_master_user_password = true

  vpc_security_group_ids = [aws_security_group.db_security_group.id]
  db_subnet_group_name   = aws_db_subnet_group.db_subnet_group.name
  parameter_group_name   = aws_db_parameter_group.db_parameter_group.name

  skip_final_snapshot = true
  publicly_accessible = false

  deletion_protection = false
}

data "aws_secretsmanager_secret" "db" {
  arn = aws_db_instance.db.master_user_secret[0].secret_arn

  depends_on = [
    aws_db_instance.db
  ]
}

data "aws_secretsmanager_secret_version" "db" {
  secret_id = data.aws_secretsmanager_secret.db.id
}
