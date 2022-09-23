## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_google"></a> [google](#provider\_google) | n/a |
| <a name="provider_google-beta"></a> [google-beta](#provider\_google-beta) | n/a |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_infrastructure_project"></a> [infrastructure\_project](#module\_infrastructure\_project) | ../terraform-google-project | n/a |

## Resources

| Name | Type |
|------|------|
| [google-beta_google_artifact_registry_repository.main](https://registry.terraform.io/providers/hashicorp/google-beta/latest/docs/resources/google_artifact_registry_repository) | resource |
| [google_cloudbuild_trigger.hello_cloud_run](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloudbuild_trigger) | resource |
| [google_sourcerepo_repository.main](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/sourcerepo_repository) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_billing_account_id"></a> [billing\_account\_id](#input\_billing\_account\_id) | n/a | `string` | n/a | yes |
| <a name="input_folder_id"></a> [folder\_id](#input\_folder\_id) | n/a | `string` | n/a | yes |
| <a name="input_main_region"></a> [main\_region](#input\_main\_region) | n/a | `any` | n/a | yes |
| <a name="input_main_repo_name"></a> [main\_repo\_name](#input\_main\_repo\_name) | n/a | `any` | n/a | yes |
| <a name="input_project_id"></a> [project\_id](#input\_project\_id) | n/a | `string` | n/a | yes |
| <a name="input_project_name"></a> [project\_name](#input\_project\_name) | n/a | `string` | n/a | yes |
| <a name="input_project_services"></a> [project\_services](#input\_project\_services) | n/a | `list(string)` | <pre>[<br>  "cloudbuild.googleapis.com",<br>  "sourcerepo.googleapis.com",<br>  "artifactregistry.googleapis.com"<br>]</pre> | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_infrastructure_project_id"></a> [infrastructure\_project\_id](#output\_infrastructure\_project\_id) | The infrastructure project id |
| <a name="output_infrastructure_project_number"></a> [infrastructure\_project\_number](#output\_infrastructure\_project\_number) | The infrastructure project nuber |
| <a name="output_main_repo_name"></a> [main\_repo\_name](#output\_main\_repo\_name) | The main repository name |
