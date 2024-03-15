resource "random_password" "password" {
  length           = 20
  special          = true
  override_special = "_%@"
}

resource "aws_secretsmanager_secret" "master_user_secret" {
  name = "db-${var.db_name}-secret"
}

resource "aws_secretsmanager_secret_version" "master_user_secret_version" {
  secret_id = aws_secretsmanager_secret.master_user_secret.id
  secret_string = jsonencode({
    username = var.db_username,
    password = random_password.password.result
  })
}
