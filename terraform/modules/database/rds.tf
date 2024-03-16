resource "aws_db_instance" "db" {
  identifier = var.db_name

  instance_class    = var.db_instance_class
  allocated_storage = 10
  engine            = var.db_engine
  engine_version    = var.db_engine_version
  db_name           = var.db_name
  port              = var.db_port

  username = var.db_username
  password = random_password.password.result

  vpc_security_group_ids = [aws_security_group.db_security_group.id]
  db_subnet_group_name   = aws_db_subnet_group.db_subnet_group.name
  parameter_group_name   = aws_db_parameter_group.db_parameter_group.name

  skip_final_snapshot = true
  publicly_accessible = false

  deletion_protection = false
}
