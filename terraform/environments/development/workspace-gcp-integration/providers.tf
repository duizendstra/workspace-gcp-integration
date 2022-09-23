provider "google" {
  project = var.terraform_project_id
  region  = var.main_region
}

provider "google-beta" {
  project = var.terraform_project_id
  region  = var.main_region
}