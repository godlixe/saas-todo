output "compute_url" {
  value = var.is_create_compute == true ? google_cloud_run_service.compute[0].status[0].url : var.compute_url
}
