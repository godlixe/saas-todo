output "compute_url" {
  value = module.compute.compute_url
}

output "storage_host" {
  value = module.storage.storage_host
}

output "storage_port" {
  value = module.storage.storage_port
}

output "storage_name" {
  value = module.storage.storage_name
}

output "storage_user" {
  value = module.storage.storage_user
}

output "storage_password" {
  value = module.storage.storage_password
}

output "metadata" {
  value = {
    compute = {
      api_url = module.compute.compute_url
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
