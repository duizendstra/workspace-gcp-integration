resource "google_project" "main" {
  name            = var.project_name
  project_id      = format("%s-%s", var.project_id, random_id.random_project_id_suffix.hex)
  folder_id       = var.folder_id
  billing_account = var.billing_account_id
}

resource "google_project_service" "main" {
  for_each = toset(var.project_services)
  project  = google_project.main.id
  service  = each.key

  disable_dependent_services = true
  disable_on_destroy = false
}

resource "random_id" "random_project_id_suffix" {
  byte_length = 2
}
