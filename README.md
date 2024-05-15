# Fast-Food Database

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
| <a name="input_region"></a> [region](#input\_region) | The default region to use for AWS | `string` | `"us-east-1"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | The default tags to use for AWS resources | `map(string)` | <pre>{<br>  "App": "lambda"<br>}</pre> | no |
| <a name="input_vpc_name"></a> [vpc\_name](#input\_vpc\_name) | The name of the VPC | `string` | `"vpc-fastfood"` | no |
## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_migrator"></a> [migrator](#module\_migrator) | ./modules/migrator | n/a |
## Resources

No resources.
## Outputs

No outputs.
<!-- END_TF_DOCS -->

## Modules

- [database](./docs/database.md)
- [migrator](./docs/migrator.md)