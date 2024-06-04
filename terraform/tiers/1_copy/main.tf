# provider "google" {
#   # credentials = file(var.service_account_file_path)
#   project = var.project_id
#   region  = var.region
# }

# # terraform {
# #   backend "gcs" {
# #     # credentials = service_account_file_path
# #   }
# # }

# module "storage" {
#   source = "../../modules/storage"

#   storage_instance_name = var.tenant_name
#   storage_user_password = var.tenant_password

#   # Pool configuration
#   is_create_storage = var.is_create_storage
#   existing_host     = var.existing_host
#   existing_port     = var.existing_port
#   existing_name     = var.existing_name
#   existing_user     = var.existing_user
#   existing_password = var.existing_password
# }

# module "compute" {
#   source = "../../modules/compute"

#   # Pool configuration
#   is_create_compute = var.is_create_compute
#   compute_name      = var.tenant_name
#   storage_host      = module.storage.host
#   storage_port      = module.storage.port
#   storage_name      = module.storage.name
#   storage_user      = module.storage.user
#   storage_password  = module.storage.password
# }
