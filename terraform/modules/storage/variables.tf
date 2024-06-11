variable "storage_instance_name" {
  type    = string
  default = "tenant"
}

variable "is_create_storage" {
  type    = bool
  default = true
}

variable "storage_host" {
  description = "existing storage host"
}

variable "storage_port" {
  description = "existing storage port"
}

variable "storage_name" {
  description = "existing storage name"
}

variable "storage_user" {
  description = "existing storage user"
}

variable "storage_password" {
  description = "existing storage password"
  default     = "123root987"
}
