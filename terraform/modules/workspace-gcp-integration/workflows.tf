# https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/workflows_workflow
resource "google_workflows_workflow" "hello_workflows" {
  project     = var.project_id
  name        = "hello-workflows"
  region      = var.main_region
  description = "Example workflow"
  source_contents = file("../../../../gcp/workflows/hello-workflows.yaml")
}