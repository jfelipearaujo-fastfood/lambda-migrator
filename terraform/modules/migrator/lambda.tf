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
    # subnet_ids = [
    #   "subnet-06d859d415c3a68b2",
    #   "subnet-0a4bc525bd9701233",
    #   "subnet-0a2e43567e6918221"
    # ]
    #   security_group_ids = [
    #     "sg-07884e7e9e1b33c59",
    #     "sg-00debc7e7628d0e8d",
    #     "sg-063a1fc97fbdcc94b",
    #     "sg-0aae885c8c3d9857e"
    #   ]
  }
}
