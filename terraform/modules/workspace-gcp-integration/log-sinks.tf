resource "google_logging_organization_sink" "organisation_iam_audit_log_move_user_to_ou_sink" {
  name        = format("%s-%s", var.organisation_iam_audit_log_move_user_to_ou_sink, random_id.random_project_id_suffix.hex)
  description = "Captures the events of an account moved to another OU"
  org_id      = var.org_id

  destination = "pubsub.googleapis.com/${google_pubsub_topic.iam_audit_log_log_move_user_to_ou.id}"
  filter      = "resource.labels.method=\"google.admin.AdminService.moveUserToOrgUnit\""
  disabled    = var.organisation_iam_audit_log_move_user_to_ou_sink_disabled
}

resource "random_id" "random_project_id_suffix" {
  byte_length = 2
}