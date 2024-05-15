resource "aws_lambda_function" "lambda_function" {
  function_name = "lambda_${var.lambda_name}"

  filename      = "./lambda.zip"
  role          = aws_iam_role.lambda_role.arn
  handler       = "bootstrap"
  runtime       = "provided.al2023"
  architectures = ["arm64"]
  memory_size   = 128
  timeout       = 30

  environment {
    variables = {
      DB_PRODUCTS_NAME    = "products"
      DB_PRODUCTS_URL     = data.aws_secretsmanager_secret_version.db_products_url_secret_version.secret_string
      DB_ORDERS_NAME      = "orders"
      DB_ORDERS_URL       = data.aws_secretsmanager_secret_version.db_orders_url_secret_version.secret_string
      DB_PAYMENTS_NAME    = "payments"
      DB_PAYMENTS_URL     = data.aws_secretsmanager_secret_version.db_payments_url_secret_version.secret_string
      DB_PRODUCTIONS_NAME = "productions"
      DB_PRODUCTIONS_URL  = data.aws_secretsmanager_secret_version.db_productions_url_secret_version.secret_string
    }
  }

  source_code_hash = filebase64sha256("./lambda.zip")

  vpc_config {
    ipv6_allowed_for_dual_stack = false
    subnet_ids                  = data.aws_subnets.private_subnets.ids
    security_group_ids          = data.aws_security_groups.dbs_security_groups.ids
  }
}
