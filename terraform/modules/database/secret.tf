resource "random_password" "password" {
  length           = 20
  special          = true
  override_special = "_%@"
}

resource "aws_secretsmanager_secret" "master_user_secret" {
  name = "db-${var.db_name}-secret"

  recovery_window_in_days = 0
}

resource "aws_secretsmanager_secret_version" "master_user_secret_version" {
  secret_id = aws_secretsmanager_secret.master_user_secret.id
  secret_string = jsonencode({
    host     = aws_db_instance.db.address,
    username = var.db_username,
    password = random_password.password.result
  })
}
