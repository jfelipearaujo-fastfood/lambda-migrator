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
      DB_HOST     = module.db.db_instance_endpoint
      DB_PORT     = var.database_port
      DB_NAME     = var.database_name
      DB_USERNAME = local.db_username
      DB_PASSWORD = jsondecode(aws_secretsmanager_secret_version.superuser.secret_string)["password"]
    }
  }

  source_code_hash = filebase64sha256("./lambda.zip")

  # vpc_config {
  #   ipv6_allowed_for_dual_stack = false
  #   subnet_ids                  = [module.vpc.private_subnets.0, module.vpc.private_subnets.1, module.vpc.private_subnets.2]
  #   security_group_ids          = [module.vpc.default_security_group_id]
  # }
}

resource "aws_s3_bucket_notification" "lambda_trigger" {
  bucket = data.aws_s3_bucket.bucket.id

  lambda_function {
    lambda_function_arn = aws_lambda_function.lambda_function.arn
    events              = ["s3:ObjectCreated:*"]
    filter_prefix       = "migrations/"
    filter_suffix       = ".sql"
  }
}

resource "aws_lambda_permission" "lambda_permission" {
  statement_id  = "AllowS3Invoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_function.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = data.aws_s3_bucket.bucket.arn
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

resource "aws_iam_policy" "lambda_policy_s3" {
  name = "policy-s3-${var.lambda_name}"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action = [
        "s3:ListBucket",
        "s3:GetObject",
        "s3:DeleteObject"
      ],
      Effect = "Allow",
      Resource = [
        data.aws_s3_bucket.bucket.arn,
        "${data.aws_s3_bucket.bucket.arn}/*"
      ]
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

resource "aws_iam_policy_attachment" "lambda_policy_attachment_s3" {
  name       = "policy-attachment-s3-${var.lambda_name}"
  roles      = [aws_iam_role.lambda_role.name]
  policy_arn = aws_iam_policy.lambda_policy_s3.arn
}
