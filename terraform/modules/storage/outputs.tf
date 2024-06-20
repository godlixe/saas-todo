output "storage_host" {
  value = google_sql_database_instance.storage.public_ip_address
}

output "storage_port" {
  value = 5432
}

output "storage_name" {
  value = google_sql_database.storage.name
}

output "storage_user" {
  value = google_sql_user.new-user.name
}

output "storage_password" {
  value = google_sql_user.new-user.password
}
