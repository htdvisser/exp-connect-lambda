locals {
  project_name = "exp-connect-lambda"
  common_tags = {
    TerraformModule = local.project_name
  }
}

terraform {
  backend "s3" {
    region         = "eu-west-1"
    bucket         = "htdvisser-ops"
    key            = "exp-connect-lambda.tfstate"
    dynamodb_table = "htdvisser-ops"
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

provider "aws" {
  region = "eu-west-1"
}
