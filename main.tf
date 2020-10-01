locals {
  s3_uri = "s3://${var.bucket_name}"
}

resource "aws_s3_bucket" "storage" {
  bucket = var.bucket_name
  acl    = "public-read"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "PublicReadGetObject",
            "Effect": "Allow",
            "Principal": "*",
            "Action": [
                "s3:GetObject"
            ],
            "Resource": [
                "arn:aws:s3:::${var.bucket_name}/*"
            ]
        }
    ]
}
EOF

  website {
    index_document = "index.html"
    error_document = "error.html"
  }
}

# calculate the hash of the configuration dir
# so we can use it as a trigger to upload the config files
data "archive_file" "dotfiles" {
  type        = "zip"
  output_path = "${path.module}/configuration.zip"

  source_dir = var.configuration_files_path
}

resource "null_resource" "upload_configuration_files" {
  triggers = {
    configuration_files_hash = data.archive_file.dotfiles.output_sha
  }

  provisioner "local-exec" {
    command = "aws s3 cp ${var.configuration_files_path} ${local.s3_uri} --recursive"
  }
}
