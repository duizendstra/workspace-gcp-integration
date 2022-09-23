resource "google_pubsub_topic" "iam_audit_log_log_move_user_to_ou" {
  name    = "iam-audit-log-move-user-to-ou-topic"
  project = var.project_id

  message_retention_duration = "86600s"
}

resource "google_pubsub_topic_iam_member" "iam_audit_log_log_move_user_to_ou_writer" {
  project = var.project_id
  topic   = google_pubsub_topic.iam_audit_log_log_move_user_to_ou.name
  role    = "roles/pubsub.publisher"
  member  = google_logging_organization_sink.organisation_iam_audit_log_move_user_to_ou_sink.writer_identity
}
