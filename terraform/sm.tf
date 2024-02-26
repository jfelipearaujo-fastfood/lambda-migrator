data "aws_kms_alias" "secretsmanager" {
  name = "alias/aws/secretsmanager"
}

resource "aws_secretsmanager_secret" "superuser" {
  name        = local.db_username
  description = "Database superuser, ${local.db_username}, database connection values"
  kms_key_id  = data.aws_kms_alias.secretsmanager.id
}

resource "aws_secretsmanager_secret_version" "superuser" {
  secret_id = aws_secretsmanager_secret.superuser.id

  secret_string = jsonencode({
    username = local.db_username
    password = local.db_password
  })
}