output "db_host" {
  description = "The hostname of the database"
  value       = module.db.db_instance_endpoint
}

output "db_instance_password" {
  description = "The password for the database"
  sensitive   = true
  value       = module.db.db_instance_password
}

output "security_group_id" {
  description = "The ID of the security group"
  value       = module.db.security_group_id
}
