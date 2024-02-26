resource "aws_secretsmanager_secret" "superuser" {
  name        = var.database_username
  description = "Database superuser, ${var.database_username}, database connection values"
  kms_key_id  = data.aws_kms_alias.secretsmanager.id
}

resource "aws_secretsmanager_secret_version" "superuser" {
  secret_id = aws_secretsmanager_secret.superuser.id
  secret_string = jsonencode({
    username = var.database_username
    password = local.db_password
  })
}
