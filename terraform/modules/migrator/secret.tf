data "aws_secretsmanager_secret" "db_products_url_secret" {
  name = "db-products-url-secret"
}

data "aws_secretsmanager_secret_version" "db_products_url_secret_version" {
  secret_id = data.aws_secretsmanager_secret.db_products_url_secret.arn
}

data "aws_secretsmanager_secret" "db_orders_url_secret" {
  name = "db-orders-url-secret"
}

data "aws_secretsmanager_secret_version" "db_orders_url_secret_version" {
  secret_id = data.aws_secretsmanager_secret.db_orders_url_secret.arn
}

data "aws_secretsmanager_secret" "db_payments_url_secret" {
  name = "db-payments-url-secret"
}

data "aws_secretsmanager_secret_version" "db_payments_url_secret_version" {
  secret_id = data.aws_secretsmanager_secret.db_payments_url_secret.arn
}

data "aws_secretsmanager_secret" "db_productions_url_secret" {
  name = "db-productions-url-secret"
}

data "aws_secretsmanager_secret_version" "db_productions_url_secret_version" {
  secret_id = data.aws_secretsmanager_secret.db_productions_url_secret.arn
}

data "aws_secretsmanager_secret" "db_customers_url_secret" {
  name = "db-customers-url-secret"
}

data "aws_secretsmanager_secret_version" "db_customers_url_secret_version" {
  secret_id = data.aws_secretsmanager_secret.db_customers_url_secret.arn
}
