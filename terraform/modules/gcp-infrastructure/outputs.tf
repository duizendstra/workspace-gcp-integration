output "infrastructure_project_id" {
  description = "The infrastructure project id"
  value       = module.infrastructure_project.project_id
}

output "infrastructure_project_number" {
  description = "The infrastructure project nuber"
  value       = module.infrastructure_project.project_number
}

output "main_repo_name" {
  description = "The main repository name"
  value       = google_sourcerepo_repository.main.name
}