## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 0.13 |
| <a name="requirement_google"></a> [google](#requirement\_google) | ~> 4.36 |
| <a name="requirement_random"></a> [random](#requirement\_random) | ~> 3.4.3 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_google_infrastructure"></a> [google\_infrastructure](#module\_google\_infrastructure) | ../../../modules/gcp-infrastructure | n/a |
| <a name="module_workspace_gcp_integration_project"></a> [workspace\_gcp\_integration\_project](#module\_workspace\_gcp\_integration\_project) | ../../../modules/terraform-google-project | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_billing_account_id"></a> [billing\_account\_id](#input\_billing\_account\_id) | n/a | `any` | n/a | yes |
| <a name="input_folder_id"></a> [folder\_id](#input\_folder\_id) | n/a | `any` | n/a | yes |
| <a name="input_main_region"></a> [main\_region](#input\_main\_region) | n/a | `any` | n/a | yes |
| <a name="input_org_id"></a> [org\_id](#input\_org\_id) | n/a | `any` | n/a | yes |
| <a name="input_terraform_project_id"></a> [terraform\_project\_id](#input\_terraform\_project\_id) | n/a | `any` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_infrastructure_project_id"></a> [infrastructure\_project\_id](#output\_infrastructure\_project\_id) | The GCP infrastructure project id |
| <a name="output_integration_project_id"></a> [integration\_project\_id](#output\_integration\_project\_id) | The integration GCP project id |
| <a name="output_integration_project_number"></a> [integration\_project\_number](#output\_integration\_project\_number) | The integration GCP project id |
| <a name="output_main_repo_name"></a> [main\_repo\_name](#output\_main\_repo\_name) | The name of the main repository |
