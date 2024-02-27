# resource "aws_vpc_endpoint" "s3_endpoint" {
#   vpc_id       = module.vpc.vpc_id
#   service_name = "com.amazonaws.us-east-1.s3"

#   route_table_ids = module.vpc.private_route_table_ids

#   policy = jsonencode({
#     Version = "2012-10-17",
#     Statement = [{
#       Effect = "Allow",
#       Principal : "*",
#       Action = "s3:*",
#       Resource = [
#         data.aws_s3_bucket.bucket.arn,
#         "${data.aws_s3_bucket.bucket.arn}/*"
#       ]
#     }]
#   })
# }

# data "aws_prefix_list" "private_s3" {
#   prefix_list_id = aws_vpc_endpoint.s3_endpoint.prefix_list_id
# }

module "vpc_endpoints" {
  source  = "terraform-aws-modules/vpc/aws//modules/vpc-endpoints"
  version = "~> 5.0"

  vpc_id = module.vpc.vpc_id

  endpoints = {
    s3 = {
      service         = "s3"
      service_type    = "Gateway"
      route_table_ids = module.vpc.intra_route_table_ids
      policy          = data.aws_iam_policy_document.endpoint.json
    }
  }
}

data "aws_iam_policy_document" "endpoint" {
  statement {
    sid = "RestrictBucketAccessToIAMRole"

    principals {
      type        = "AWS"
      identifiers = ["*"]
    }

    actions = [
      "s3:GetObject",
      "s3:DeleteObject"
    ]

    resources = [
      data.aws_s3_bucket.bucket.arn,
      "${data.aws_s3_bucket.bucket.arn}/*",
    ]

    condition {
      test     = "ArnEquals"
      variable = "aws:PrincipalArn"
      values   = [aws_iam_role.lambda_role.arn]
    }
  }
}
