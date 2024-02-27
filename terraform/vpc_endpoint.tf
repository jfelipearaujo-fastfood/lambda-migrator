resource "aws_vpc_endpoint" "s3_endpoint" {
  vpc_id       = module.vpc.vpc_id
  service_name = "com.amazonaws.us-east-1.s3"

  route_table_ids = [module.vpc.vpc_main_route_table_id]

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Effect = "Allow",
      Principal : "*",
      Action = "s3:*",
      Resource = [
        data.aws_s3_bucket.bucket.arn,
        "${data.aws_s3_bucket.bucket.arn}/*"
      ]
    }]
  })
}
