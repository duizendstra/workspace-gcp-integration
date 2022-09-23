variable "project_name" {
  type = string
}

variable "project_id" {
  type = string
}

variable "folder_id" {
  type = string
}
variable "billing_account_id" {
  type = string
}

variable "project_services" {
  default = ["cloudbuild.googleapis.com", "sourcerepo.googleapis.com", "artifactregistry.googleapis.com"]
  type    = list(string)
}

variable "main_region" {}

variable "main_repo_name" {}
