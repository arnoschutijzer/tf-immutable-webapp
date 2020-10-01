# tf-immutable-webapp

A Terraform module to provision immutable webapps.

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 0.13.3 |
| aws | >= 3.8 |

## Providers

| Name | Version |
|------|---------|
| archive | n/a |
| aws | >= 3.8 |
| null | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| bucket\_name | The name of the bucket. | `string` | n/a | yes |
| configuration\_files\_path | The relative path to the configuration files. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| s3\_uri | n/a |
| s3\_website\_endpoint | n/a |
