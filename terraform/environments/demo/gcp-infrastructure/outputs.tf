output "infrastructure_project_id" {
  description = "The GCP infrastructure project id"
  value       = module.google_infrastructure.infrastructure_project_id
}

output "integration_project_id" {
  description = "The integration GCP project id"
  value       = module.workspace_gcp_integration_project.project_id
}

output "integration_project_number" {
  description = "The integration GCP project id"
  value       = module.workspace_gcp_integration_project.project_number
}

output "main_repo_name" {
  description = "The name of the main repository"
  value       = module.google_infrastructure.main_repo_name
}