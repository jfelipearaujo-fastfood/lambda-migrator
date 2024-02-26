resource "random_password" "database_password" {
  length           = 16
  special          = true
  upper            = false
  override_special = "_#$"
}

resource "aws_secretsmanager_secret" "sm_database_credentials" {
  name = "sm-${var.project_name}"
}

resource "aws_secretsmanager_secret_version" "sm_database_credentials_version" {
  secret_id = aws_secretsmanager_secret.sm_database_credentials.id

  secret_string = jsonencode({
    username = var.database_username,
    password = random_password.database_password.result,
  })
}
