resource "google_artifact_registry_repository" "main" {
  provider = google-beta

  project       = module.infrastructure_project.project_id
  location      = var.main_region
  repository_id = var.main_repo_name
  description   = "Docker repository" //TODO make variable
  format        = "DOCKER"
  
  depends_on = [
    module.infrastructure_project.google_project_service
  ]
}