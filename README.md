## About The Project

### Google Workspace automation with Eventarc

The solution in this project was presented at [DevFest Craiova 2022](https://gdg.community.dev/events/details/google-gdg-craiova-presents-devfest-craiova-2022/)

## Get Ready to integrate Workspace with the Google Cloud Platform

### Workspace configuration
Google Workspace needs to [Share data with Google Cloud services](https://support.google.com/a/answer/9320190) for this solution. Once enabled the events will flow to the [organization](https://cloud.google.com/resource-manager/docs/creating-managing-organization) logs. You can view the logs by opening the log viewer on organisation level [ organisation level](`https://console.cloud.google.com/logs/query?organizationId=`).

On the GCP organisation you need the following roles:
	- Logging Admin
	- Organisation Administrator

### GCP billinhg
A [billing account](https://cloud.google.com/billing/docs/concepts) is required. You need Billing Account User access on the billing account.

### Terraform
The solution uses [Terraform](https://registry.terraform.io/providers/hashicorp/google/latest/docs) to deploy the infrastructure. The [state](https://www.terraform.io/language/settings/backends/gcs) is maintained in a [cloud bucket](https://cloud.google.com/storage/docs/creating-buckets). For this solution you need such a bucket and write rights on that bucket. Note down both the name of this bucket and the project id.

Terraform will create two Google Cloud Projects in an existing [GCP Folder](https://cloud.google.com/resource-manager/docs/creating-managing-folders). Create the folder in the GCP organisation, or use an existing one. Note down down the id of the folder. On this folder you need the following roles:
	-   Folder Viewer
	-   Project Owner
	-   Project Creator

To deploy the solution you first need to clone this repository. When you are in the root directory of the repository go to the `gcp-infrastructure` directory.

```
cd terraform/environments/development/gcp-infrastructure/
```

- Create the file `backend.tf` with the contents
```
terraform {
  backend "gcs" {
	 bucket = "[the name of the bucket for the terraform state]"
	 prefix = "development/infrastructure"
}
```
- Create the file `terraform.tfvars` with the contents:
```
terraform_project_id = "[the project id]"
org_id = "[the id of the GCP organisation]"
folder_id = "[the id of the folder]"
billing_account_id = "[the id of the billing account]"
main_region = "[the region to deploy the solution]"
```
- run `terraform init`
- run `terraform plan`
- review the plan
- run `terraform apply`

**You will get an API error here, wait 5 minutes and run Run `terraform apply` again**

- run `git remote add development https://source.developers.google.com/p/[infrastructure_project_id]/r/workspace-gcp-integration`
- run `git push development`

The git push will trigger the build of a container. Check the Cloud Build in the infrastructure project. Once the build is done, a container is available in the Artifact Registry

- goto the `workspace-gcp-integration` directory
- Create the file `backend.tf` with the contents
```
terraform {
  backend "gcs" {
    bucket = "[the name of the bucket for the terraform state]"
    prefix = "development/workspace-gcp-sync"
  }
}
data "terraform_remote_state" "google_infrastructure" {
  backend = "gcs"
  config = {
    bucket = "[the name of the bucket for the terraform state]"
    prefix = "development/infrastructure"
  }
}
```
- Create the file `terraform.tfvars` with the contents:
```
terraform_project_id = "[the project id]"
main_region = "[the region to deploy the solution]"
organisation_iam_audit_log_move_user_to_ou_sink = "iam-audit-log-move-user-to-ou-sink-development"
organisation_iam_audit_log_move_user_to_ou_sink_disabled = false
hello_cloud_run_build_version = "0.0.1"
org_id = "[the GCP organisation id]"
```

- run `terraform init`
- run `terraform plan`
- review the plan
- run `terraform apply`

## Contact

Jasper Duizendstra - [@duizendstra](https://twitter.com/@duizendstra) - jasper@duizendstra.com

Project Link: [https://github.com/duizendstra/workspace-gcp-integration](https://github.com/duizendstra/workspace-gcp-integration)