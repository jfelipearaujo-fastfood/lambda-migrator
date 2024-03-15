output "db_host" {
  description = "The hostname of the database"
  value       = aws_db_instance.db.address
}

output "db_port" {
  description = "The port of the database"
  value       = aws_db_instance.db.port
}

output "db_name" {
  description = "The name of the database"
  value       = aws_db_instance.db.db_name
}

output "db_username" {
  description = "The username for the database"
  value       = aws_db_instance.db.username
}

output "security_group_id" {
  description = "The ID of the security group"
  value       = aws_security_group.db_security_group.id
}
