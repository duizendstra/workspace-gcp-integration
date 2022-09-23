output "project_number" {
  description = "The GCP project number"
  value       = google_project.main.number
}

output "project_id" {
  description = "The GCP project id"
  value       = google_project.main.project_id
}