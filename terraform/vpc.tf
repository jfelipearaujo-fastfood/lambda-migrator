module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 5.0"

  name = local.name
  cidr = local.vpc_cidr

  azs              = local.azs
  public_subnets   = [for k, v in local.azs : cidrsubnet(local.vpc_cidr, 8, k)]
  private_subnets  = [for k, v in local.azs : cidrsubnet(local.vpc_cidr, 8, k + 3)]
  database_subnets = [for k, v in local.azs : cidrsubnet(local.vpc_cidr, 8, k + 6)]
  intra_subnets    = [for k, v in local.azs : cidrsubnet(local.vpc_cidr, 8, k + 9)]

  create_database_subnet_group = true

  intra_dedicated_network_acl = true
  intra_inbound_acl_rules = concat(
    # NACL rule for local traffic
    [
      {
        rule_number = 100
        rule_action = "allow"
        from_port   = 0
        to_port     = 0
        protocol    = "-1"
        cidr_block  = "10.0.0.0/16"
      },
    ],
    # NACL rules for the response traffic from addresses in the AWS S3 prefix list
    [for k, v in zipmap(
      range(length(data.aws_ec2_managed_prefix_list.this.entries[*].cidr)),
      data.aws_ec2_managed_prefix_list.this.entries[*].cidr
      ) :
      {
        rule_number = 200 + k
        rule_action = "allow"
        from_port   = 1024
        to_port     = 65535
        protocol    = "tcp"
        cidr_block  = v
      }
    ]
  )
}
