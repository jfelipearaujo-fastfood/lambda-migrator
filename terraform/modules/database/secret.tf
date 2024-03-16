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

  depends_on = [
    aws_db_instance.db
  ]
}

resource "aws_iam_policy" "db_secret_policy" {
  name = "db-secret-policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "secretsmanager:GetSecretValue",
          "secretsmanager:DescribeSecret"
        ]
        Resource = aws_secretsmanager_secret.master_user_secret.arn
      },
    ]
  })
}

resource "aws_iam_role_policy_attachment" "db_secret_policy_attachment" {
  role       = "fastfood-service-account-role"
  policy_arn = aws_iam_policy.db_secret_policy.arn
}
