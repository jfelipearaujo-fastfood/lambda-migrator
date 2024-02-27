resource "aws_lambda_function" "lambda_function" {
  function_name = var.lambda_name

  filename      = "./lambda.zip"
  role          = aws_iam_role.lambda_role.arn
  handler       = "bootstrap"
  runtime       = "provided.al2023"
  architectures = ["arm64"]
  memory_size   = 128
  timeout       = 30

  environment {
    variables = {
      # DB_HOST     = module.rds_proxy.proxy_endpoint
      DB_HOST = module.db.db_instance_endpoint
      DB_NAME = var.database_name
      DB_USER = local.db_username
      DB_PASS = module.db.db_instance_password
    }
  }

  source_code_hash = filebase64sha256("./lambda.zip")

  vpc_config {
    ipv6_allowed_for_dual_stack = false
    subnet_ids                  = module.vpc.private_subnets
    security_group_ids          = [aws_security_group.security_group.id]
  }
}

resource "aws_iam_role" "lambda_role" {
  name = "role-${var.lambda_name}"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action = "sts:AssumeRole",
      Effect = "Allow",
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
}

resource "aws_iam_policy_attachment" "lambda_policy_attachment" {
  name       = "policy-attachment-${var.lambda_name}"
  roles      = [aws_iam_role.lambda_role.name]
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_policy_attachment" "iam_role_policy_attachment_vpc" {
  name       = "policy-attachment-vpc-${var.lambda_name}"
  roles      = [aws_iam_role.lambda_role.name]
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
}
