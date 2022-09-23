resource "google_sourcerepo_repository" "main" {
  name    = var.main_repo_name
  project = module.infrastructure_project.project_id
  depends_on = [
    module.infrastructure_project.service
  ]
}
