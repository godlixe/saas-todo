output "storage_host" {
  value = google_sql_database_instance.storage[0].public_ip_address
}

output "storage_port" {
  value = 5432
}

output "storage_name" {
  value = google_sql_database.storage[0].name
}

output "storage_user" {
  value = google_sql_user.new-user[0].name

output "storage_password" {
  value = var.storage_instance_name
}
