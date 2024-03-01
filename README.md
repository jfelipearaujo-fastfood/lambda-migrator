# Fast-Food Database Provisioning

This project its responsible for provisioning a database in AWS using Terraform. The database will be a PostgreSQL instance and it will be created in a VPC. The database will be created in a private subnet and it will be accessed by a lambda function that will be responsible for running the database migrations.

## Diagram

![diagram](./docs/diagrams/cloud_aws_database_migrations.png)

## Terraform

Bellow it is described all the resources used to create the architecture proposed for this project.

<!-- BEGIN_TF_DOCS -->

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | 1.7.4 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | 5.38.0 |
## Providers

No providers.
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_private_subnet_1a"></a> [private\_subnet\_1a](#input\_private\_subnet\_1a) | The ID of the private subnet in the first availability zone | `string` | n/a | yes |
| <a name="input_private_subnet_1b"></a> [private\_subnet\_1b](#input\_private\_subnet\_1b) | The ID of the private subnet in the second availability zone | `string` | n/a | yes |
| <a name="input_private_subnet_1c"></a> [private\_subnet\_1c](#input\_private\_subnet\_1c) | The ID of the private subnet in the third availability zone | `string` | n/a | yes |
| <a name="input_region"></a> [region](#input\_region) | The default region to use for AWS | `string` | `"us-east-1"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | The default tags to use for AWS resources | `map(string)` | <pre>{<br>  "App": "database"<br>}</pre> | no |
| <a name="input_vpc_cidr_block"></a> [vpc\_cidr\_block](#input\_vpc\_cidr\_block) | The CIDR block of the VPC | `string` | n/a | yes |
| <a name="input_vpc_id"></a> [vpc\_id](#input\_vpc\_id) | The ID of the VPC | `string` | n/a | yes |
## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_database"></a> [database](#module\_database) | ./modules/database | n/a |
| <a name="module_migrator"></a> [migrator](#module\_migrator) | ./modules/migrator | n/a |
| <a name="module_secret"></a> [secret](#module\_secret) | ./modules/secret | n/a |
## Resources

No resources.
## Outputs

No outputs.
<!-- END_TF_DOCS -->

## Modules

- [database](./docs/database.md)
- [migrator](./docs/migrator.md)
- [secret](./docs/secret.md)