resource "google_cloudbuild_trigger" "hello_cloud_run" {
  name    = "hello-cloud-run"
  project = module.infrastructure_project.project_id

  trigger_template {
    branch_name = "main"
    repo_name   = var.main_repo_name
  }

  substitutions = {
    _DEPLOY_REGION : "europe-west3"
    _GCR_HOSTNAME : "europe-west3-docker.pkg.dev"
    _GCR_PROJECT_ID : module.infrastructure_project.project_id
    _ARTIFACT_REPO_NAME : var.main_repo_name
    _IMAGE_NAME : "hello-cloud-run"
  }

  included_files = ["gcp/cloud-runs/hello-cloud-run/version"]

  filename = "gcp/cloud-runs/hello-cloud-run/cloudbuild.yaml"
  depends_on = [
    module.infrastructure_project.service
  ]
}
