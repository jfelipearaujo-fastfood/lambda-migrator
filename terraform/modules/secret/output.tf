output "db_username" {
  description = "The username for the database username"
  sensitive   = true
  value       = jsondecode(aws_secretsmanager_secret_version.superuser.secret_string)["username"]
}

output "db_password" {
  description = "The password for the database"
  sensitive   = true
  value       = jsondecode(aws_secretsmanager_secret_version.superuser.secret_string)["password"]
}
