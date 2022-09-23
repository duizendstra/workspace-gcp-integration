module "infrastructure_project" {
  source             = "../terraform-google-project"
  main_region        = var.main_region
  project_name       = var.project_name
  project_id         = var.project_id
  folder_id          = var.folder_id
  billing_account_id = var.billing_account_id
  project_services   = var.project_services
}
