resource "aws_vpc_endpoint" "s3_endpoint" {
  vpc_id       = module.vpc.vpc_id
  service_name = "com.amazonaws.us-east-1.s3"

  route_table_ids = module.vpc.private_route_table_ids

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

data "aws_prefix_list" "private_s3" {
  prefix_list_id = aws_vpc_endpoint.s3_endpoint.prefix_list_id
}
