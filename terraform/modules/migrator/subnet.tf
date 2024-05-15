data "aws_subnets" "private_subnets" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.vpc.id]
  }
}

data "aws_db_subnet_group" "private_subnet_group_products_db" {
  name = "db-sng-products"
}

data "aws_db_subnet_group" "private_subnet_group_orders_db" {
  name = "db-sng-orders"
}

data "aws_db_subnet_group" "private_subnet_group_payments_db" {
  name = "db-sng-payments"
}

data "aws_db_subnet_group" "private_subnet_group_productions_db" {
  name = "db-sng-productions"
}
