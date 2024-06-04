resource "google_sql_database_instance" "storage" {
  count = var.is_create_storage == true ? 1 : 0

  # Information
  name             = "cloud-sql-${var.storage_instance_name}"
  database_version = "POSTGRES_15"
  region           = "asia-southeast2"

  settings {
    activation_policy = "ALWAYS"
    availability_type = "ZONAL"
    backup_configuration {
      backup_retention_settings {
        retained_backups = 7
        retention_unit   = "COUNT"
      }
      enabled                        = false
      location                       = "asia"
      start_time                     = "01:00"
      transaction_log_retention_days = 7
    }
    disk_autoresize       = true
    disk_autoresize_limit = 0
    disk_size             = 10
    disk_type             = "PD_SSD"

    ip_configuration {
      authorized_networks {
        name  = "all"
        value = "0.0.0.0/0"
      }
      ipv4_enabled = true
    }
    location_preference {
      zone = "asia-southeast2-c"
    }
    pricing_plan = "PER_USE"
    tier         = "db-custom-1-3840"
  }

  deletion_protection = false
}

resource "google_sql_database" "storage" {
  count = var.is_create_storage == true ? 1 : 0

  name     = "cloud-sql-${var.storage_instance_name}"
  instance = google_sql_database_instance.storage[0].name
}

resource "google_sql_user" "new-user" {
  count = var.is_create_storage == true ? 1 : 0

  name     = "default-${var.storage_instance_name}"
  instance = google_sql_database_instance.storage[0].name
  password = var.storage_user_password
}
