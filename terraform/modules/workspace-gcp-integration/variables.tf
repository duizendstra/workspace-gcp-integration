variable "project_id" {
  type = string
}
variable "org_id" {}
variable "project_number" {}

variable "main_region" {}
variable "organisation_iam_audit_log_move_user_to_ou_sink" {}
variable "infrastructure_project_id" {}
variable "main_repo_name" {}
variable "hello_cloud_run_build_version" {}
variable "organisation_iam_audit_log_move_user_to_ou_sink_disabled" {
  type    = bool
  default = true
}
