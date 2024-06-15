variable "compute_name" {
    description = "compute instance name"
    default = null
}

variable "storage_host" {
    description = "database storage host"
}

variable "storage_port" {
    description = "database storage port"
}

variable "storage_name" {
    description = "database storage name"
}

variable "storage_user" {
    description = "database storage user"
}

variable "storage_password" {
    description = "database storage password"
}

variable "compute_url" {
    description = "compute instance url"
    default = null
}