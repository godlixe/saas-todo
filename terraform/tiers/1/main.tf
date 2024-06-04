provider "google" {
  project = var.project_id
  region  = var.region
}

terraform {
  backend "gcs" {}
}

module "storage" {
  source = "../../modules/storage"

  storage_instance_name = var.tenant_name
  storage_user_password = var.tenant_password

  # Pool configuration
  is_create_storage = var.is_create_storage
  existing_host     = var.storage_host
  existing_port     = var.storage_port
  existing_name     = var.storage_name
  existing_user     = var.storage_user
  existing_password = var.storage_password
}

module "compute" {
  source = "../../modules/compute"

  # Pool configuration
  is_create_compute = var.is_create_compute
  compute_name      = var.tenant_name
  storage_host      = module.storage.host
  storage_port      = module.storage.port
  storage_name      = module.storage.name
  storage_user      = module.storage.user
  storage_password  = module.storage.password
}
