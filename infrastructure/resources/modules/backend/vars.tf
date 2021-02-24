variable "api_gateway_id" {
  type = string
}

variable "env_config" {
  type = object({
    account_id = string
    region = string
    env = string
  })
}

variable "tags" {
  type    = map(string)
  default = {}
}