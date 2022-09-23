## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 0.13 |
| <a name="requirement_google"></a> [google](#requirement\_google) | ~> 4.36 |
| <a name="requirement_random"></a> [random](#requirement\_random) | ~> 3.4.3 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_terraform"></a> [terraform](#provider\_terraform) | n/a |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_gcp_workspace_integration"></a> [gcp\_workspace\_integration](#module\_gcp\_workspace\_integration) | ../../../modules/workspace-gcp-integration | n/a |

## Resources

| Name | Type |
|------|------|
| [terraform_remote_state.google_infrastructure](https://registry.terraform.io/providers/hashicorp/terraform/latest/docs/data-sources/remote_state) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_hello_cloud_run_build_version"></a> [hello\_cloud\_run\_build\_version](#input\_hello\_cloud\_run\_build\_version) | n/a | `any` | n/a | yes |
| <a name="input_main_region"></a> [main\_region](#input\_main\_region) | n/a | `any` | n/a | yes |
| <a name="input_org_id"></a> [org\_id](#input\_org\_id) | n/a | `any` | n/a | yes |
| <a name="input_organisation_iam_audit_log_move_user_to_ou_sink"></a> [organisation\_iam\_audit\_log\_move\_user\_to\_ou\_sink](#input\_organisation\_iam\_audit\_log\_move\_user\_to\_ou\_sink) | n/a | `any` | n/a | yes |
| <a name="input_organisation_iam_audit_log_move_user_to_ou_sink_disabled"></a> [organisation\_iam\_audit\_log\_move\_user\_to\_ou\_sink\_disabled](#input\_organisation\_iam\_audit\_log\_move\_user\_to\_ou\_sink\_disabled) | n/a | `bool` | `true` | no |
| <a name="input_terraform_project_id"></a> [terraform\_project\_id](#input\_terraform\_project\_id) | n/a | `any` | n/a | yes |

## Outputs

No outputs.
