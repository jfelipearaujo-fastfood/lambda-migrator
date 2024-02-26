data "aws_s3_bucket" "bucket" {
  bucket = var.bucket_name
}

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
      DB_HOST     = module.db.db_instance_address
      DB_PORT     = var.database_port
      DB_NAME     = module.db.db_instance_identifier
      DB_USERNAME = var.database_username
      DB_PASSWORD = jsondecode(aws_secretsmanager_secret_version.sm_database_credentials_version.secret_string)["password"]
    }
  }

  source_code_hash = filebase64sha256("./lambda.zip")
}

resource "aws_s3_bucket_notification" "lambda_trigger" {
  bucket = data.aws_s3_bucket.bucket.id

  lambda_function {
    lambda_function_arn = aws_lambda_function.lambda_function.arn
    events              = ["s3:ObjectCreated:*"]
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
  name = "policy-${var.lambda_name}"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action   = "s3:GetObject",
      Effect   = "Allow",
      Resource = "${data.aws_s3_bucket.bucket.arn}/migrations/*"
      }, {
      Action   = "s3:DeleteObject",
      Effect   = "Allow",
      Resource = "${data.aws_s3_bucket.bucket.arn}/migrations/*"
    }]
  })
}

resource "aws_iam_policy_attachment" "lambda_policy_attachment" {
  name       = "policy-attachment-${var.lambda_name}"
  roles      = [aws_iam_role.lambda_role.name]
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_policy_attachment" "lambda_policy_attachment_s3" {
  name       = "policy-attachment-s3-${var.lambda_name}"
  roles      = [aws_iam_role.lambda_role.name]
  policy_arn = aws_iam_policy.lambda_policy_s3.arn
}
