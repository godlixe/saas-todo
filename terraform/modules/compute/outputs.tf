output "service_url" {
  value = google_cloud_run_service.compute[0].status[0].url
}
