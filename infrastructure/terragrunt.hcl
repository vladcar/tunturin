locals {
  aws_region = "eu-north-1"
}

remote_state {
  backend = "s3"
  generate = {
    path      = "backend.tf"
    if_exists = "overwrite"
  }
  config = {
    bucket         = "tunturin-remote-state"
    key            = "${path_relative_to_include()}/terraform.tfstate"
    region         = "eu-north-1"
    encrypt        = true
    dynamodb_table = "tunturin-terraform-lock-table"
  }
}

generate "provider" {
  path      = "provider.tf"
  if_exists = "skip"
  contents  = <<EOF
provider "aws" {
  region = "${local.aws_region}"
}
EOF
}

inputs = {
  terragrunt_parent_path = get_parent_terragrunt_dir()
}
