variable "storage_instance_name" {
  type    = string
  default = "tenant"
}

variable "storage_user_password" {
  type    = string
  default = "123root987"
}

variable "is_create_storage" {
  type    = bool
  default = true
}

variable "existing_host" {
  description = "existing storage host"
}

variable "existing_port" {
  description = "existing storage port"
}

variable "existing_name" {
  description = "existing storage name"
}

variable "existing_user" {
  description = "existing storage user"
}

variable "existing_password" {
  description = "existing storage password"
  default = "123root987"
}
