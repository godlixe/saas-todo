# variable "service_account_file_path" {
#   description = "service account file path"
# }

variable "project_id" {
  description = "provider project id"
}

variable "region" {
  description = "provider region"
}
variable "tenant_name" {
  description = "tenant name"
}

variable "tenant_subdomain" {
  description = "tenant subdomain"
}

variable "tenant_password" {
  description = "tenant password"
}

variable "is_create_storage" {
  type        = bool
  description = "create a new storage instance"
  default     = true
}

variable "is_create_compute" {
  type        = bool
  description = "create a new storage instance"
  default     = true
}

variable "existing_host" {
  description = "existing storage host"
  default     = null
}

variable "existing_port" {
  description = "existing storage port"
  default     = null
}

variable "existing_name" {
  description = "existing storage name"
  default     = null
}

variable "existing_user" {
  description = "existing storage user"
  default     = null
}

variable "existing_password" {
  description = "existing storage password"
  default     = null
}
