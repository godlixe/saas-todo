output "host" {
  value = var.is_create_storage == true ? google_sql_database_instance.storage[0].public_ip_address : var.existing_host
}

output "port" {
  value = 5432
}

output "name" {
  value = var.is_create_storage == true ? google_sql_database.storage[0].name : var.existing_name
}

output "user" {
  value = var.is_create_storage == true ? google_sql_user.new-user[0].name : var.existing_user
}

output "password" {
  value = var.is_create_storage == true ? var.storage_instance_name : var.existing_password
}
