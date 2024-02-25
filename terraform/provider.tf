terraform {
  required_version = ">= 1.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.33"
    }
  }

  backend "s3" {
    region = "us-east-1"
    key    = "terraform/database/terraform.tfstate"
  }
}

provider "aws" {
  region = var.region

  default_tags {
    tags = var.tags
  }
}
