output "storage_host" {
  value = var.is_create_storage == true ? google_sql_database_instance.storage[0].public_ip_address : var.storage_host
}

output "storage_port" {
  value = 5432
}

output "storage_name" {
  value = var.is_create_storage == true ? google_sql_database.storage[0].name : var.storage_name
}

output "storage_user" {
  value = var.is_create_storage == true ? google_sql_user.new-user[0].name : var.storage_user
}

output "storage_password" {
  value = var.is_create_storage == true ? var.storage_instance_name : var.storage_password
}
