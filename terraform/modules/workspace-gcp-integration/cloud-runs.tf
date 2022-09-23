resource "google_cloud_run_service" "hello_cloud_run" {
  name     = "hello-cloud-run"
  project  = var.project_id
  location = var.main_region

  template {
    spec {
      containers {
        image = "europe-west3-docker.pkg.dev/${var.infrastructure_project_id}/workspace-gcp-integration/hello-cloud-run:${var.hello_cloud_run_build_version}" //TODO
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
}

resource "google_artifact_registry_repository_iam_member" "workspace-run-admin-account" {
  provider = google-beta
  project  = var.infrastructure_project_id

  location   = var.main_region
  repository = "workspace-gcp-integration" //TODO Makevariable
  role       = "roles/artifactregistry.reader"
  member     = "serviceAccount:service-${var.project_number}@serverless-robot-prod.iam.gserviceaccount.com"
}

resource "google_cloud_run_service_iam_binding" "binding" {
  location = google_cloud_run_service.hello_cloud_run.location
  project  = google_cloud_run_service.hello_cloud_run.project
  service  = google_cloud_run_service.hello_cloud_run.name
  role     = "roles/run.invoker"
  members = [
    "serviceAccount:service-${var.project_number}@serverless-robot-prod.iam.gserviceaccount.com",
  ]
}
