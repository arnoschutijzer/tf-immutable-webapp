terraform {
  required_version = ">= 0.13.3"
  required_providers {
    aws = {
      version = ">= 3.8"
    }
    archive = {
      version = ">= 1.3"
    }
    null = {
      version = ">= 2.1"
    }
  }
}
