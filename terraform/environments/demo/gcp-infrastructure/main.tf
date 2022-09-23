module "google_infrastructure" {
  source             = "../../../modules/gcp-infrastructure"
  main_region        = var.main_region
  folder_id          = var.folder_id
  billing_account_id = var.billing_account_id

  project_name   = "Infrastructure Demo"
  project_id     = "infrastructure"
  main_repo_name = "workspace-gcp-integration"
}

module "workspace_gcp_integration_project" {
  source             = "../../../modules/terraform-google-project"
  main_region        = var.main_region
  folder_id          = var.folder_id
  billing_account_id = var.billing_account_id

  project_name = "Workspace GCP integration Demo"
  project_id   = "integration"
  project_services = ["workflows.googleapis.com","run.googleapis.com","eventarc.googleapis.com"]
}