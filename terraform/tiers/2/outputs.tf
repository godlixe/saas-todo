output "metadata" {
  sensitive = true
  value = {
    compute = {
      compute_url = module.compute.compute_url
    }
    storage = {
      host     = module.storage.storage_host
      port     = module.storage.storage_port
      name     = module.storage.storage_name
      user     = module.storage.storage_user
      password = module.storage.storage_password
    }
  }
}
