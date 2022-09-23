# https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/eventarc_trigger

resource "google_eventarc_trigger" "workflow" {
  name     = "workflow-trigger"
  project  = var.project_id
  location = var.main_region

  matching_criteria {
    attribute = "type"
    value     = "google.cloud.pubsub.topic.v1.messagePublished"
  }

  destination {
    workflow = google_workflows_workflow.hello_workflows.id
  }

  transport {
    pubsub {
      topic = google_pubsub_topic.iam_audit_log_log_move_user_to_ou.id
    }
  }
  service_account = "${var.project_number}-compute@developer.gserviceaccount.com"
}

resource "google_eventarc_trigger" "cloud_run" {
  name     = "cloud-run-trigger"
  project  = var.project_id
  location = var.main_region

  matching_criteria {
    attribute = "type"
    value     = "google.cloud.pubsub.topic.v1.messagePublished"
  }

  destination {
    cloud_run_service {
      service = google_cloud_run_service.hello_cloud_run.name
      region  = google_cloud_run_service.hello_cloud_run.location
    }
  }

  transport {
    pubsub {
      topic = google_pubsub_topic.iam_audit_log_log_move_user_to_ou.id
    }
  }
  service_account = "${var.project_number}-compute@developer.gserviceaccount.com"
}