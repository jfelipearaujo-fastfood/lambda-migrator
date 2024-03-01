<!-- BEGIN_TF_DOCS -->

## Requirements

No requirements.
## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | n/a |
| <a name="provider_random"></a> [random](#provider\_random) | n/a |
## Inputs

No inputs.
## Modules

No modules.
## Resources

| Name | Type |
|------|------|
| [aws_secretsmanager_secret.superuser](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/secretsmanager_secret) | resource |
| [aws_secretsmanager_secret_version.superuser](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/secretsmanager_secret_version) | resource |
| [random_password.password](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/password) | resource |
| [random_pet.users](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/pet) | resource |
| [aws_kms_alias.secretsmanager](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/kms_alias) | data source |
## Outputs

| Name | Description |
|------|-------------|
| <a name="output_db_password"></a> [db\_password](#output\_db\_password) | The password for the database |
| <a name="output_db_username"></a> [db\_username](#output\_db\_username) | The username for the database username |
<!-- END_TF_DOCS -->