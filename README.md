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

## Getting started

First create the infrastructure:
```bash
# plan the stack...
$ terraform plan

# check the output and apply it
$ terraform apply
```

This will create the S3 bucket and upload the configuration files.  
The output will contain the `s3_uri` that you will need to upload the application artefacts to the bucket and the `s3_website_endpoint`.  
If you navigate to that endpoint in a browser you will see an error since the application itself is not deployed yet.  

## Deploy the example app
```bash
# navigate to the example_app directory
$ cd example_app

# deploy the application artefact
$ ./deploy_app.sh s3://BUCKET_NAME
```

You can now navigate to the `s3_website_endpoint` you have received from provisioning the stack.  