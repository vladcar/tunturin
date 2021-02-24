terraform {
  source = "../resources"

  extra_arguments "common_var" {
    commands  = get_terraform_commands_that_need_vars()
    arguments = ["-var-file=${get_terragrunt_dir()}/aws_config.tfvars"]
  }
}

include {
  path = find_in_parent_folders()
}

inputs = {

}
