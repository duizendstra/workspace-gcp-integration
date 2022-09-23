module "gcp_workspace_integration" {
  source = "../../../modules/workspace-gcp-integration"

  project_id                = data.terraform_remote_state.google_infrastructure.outputs.integration_project_id
  project_number            = data.terraform_remote_state.google_infrastructure.outputs.integration_project_number
  infrastructure_project_id = data.terraform_remote_state.google_infrastructure.outputs.infrastructure_project_id
  main_repo_name            = data.terraform_remote_state.google_infrastructure.outputs.main_repo_name

  org_id                                                   = var.org_id
  main_region                                              = var.main_region
  organisation_iam_audit_log_move_user_to_ou_sink          = var.organisation_iam_audit_log_move_user_to_ou_sink
  organisation_iam_audit_log_move_user_to_ou_sink_disabled = var.organisation_iam_audit_log_move_user_to_ou_sink_disabled
  hello_cloud_run_build_version                            = var.hello_cloud_run_build_version
}
