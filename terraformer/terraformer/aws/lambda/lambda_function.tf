resource "aws_lambda_function" "tfer--lambda_migrator" {
  architectures = ["arm64"]

  environment {
    variables = {
      DB_HOST     = "rds-proxy.proxy-c4ozchbz74dh.us-east-1.rds.amazonaws.com"
      DB_NAME     = "db-fastfood"
      DB_PASSWORD = "aaiicl59ZU2QbVYj"
      DB_PORT     = "5432"
      DB_USERNAME = "definite_warthog"
    }
  }

  ephemeral_storage {
    size = "512"
  }

  function_name = "lambda_migrator"
  handler       = "bootstrap"

  logging_config {
    log_format = "Text"
    log_group  = "/aws/lambda/lambda_migrator"
  }

  memory_size                    = "128"
  package_type                   = "Zip"
  reserved_concurrent_executions = "-1"
  role                           = "arn:aws:iam::167192228103:role/role-lambda_migrator"
  runtime                        = "provided.al2023"
  skip_destroy                   = "false"
  source_code_hash               = "UaRSaW7hA24X6xu3YS1IWjRNGB6JKZ79w0E0Q311mjs="

  tags = {
    App = "database"
  }

  tags_all = {
    App = "database"
  }

  timeout = "30"

  tracing_config {
    mode = "PassThrough"
  }

  vpc_config {
    ipv6_allowed_for_dual_stack = "false"
    security_group_ids          = ["sg-0bc1779db14112ddd"]
    subnet_ids                  = ["subnet-00507713021d2d926", "subnet-00a743fd7a2650956", "subnet-0a030cad947ddd1c1"]
  }
}
