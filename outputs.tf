output "s3_uri" {
  value = local.s3_uri
}

output "s3_website_endpoint" {
  value = aws_s3_bucket.storage.website_endpoint
}