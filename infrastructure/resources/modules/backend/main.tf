
locals {
  service_prefix = "tunturin"
}

module "prospect_creator_lambda" {
  source         = "vladcar/serverless-common-basic-lambda/aws"
  function_name  = "tunturin-tg-webhook-handler-${var.env_config.env}"
  source_path    = "${path.module}/function.zip"
  handler        = "telegram-send-data-handler"
  memory_size    = 256
  runtime        = "go1.x"
  layers         = []
  create_role    = false
  execution_role = module.service_lambda_execution_role.role_arn
  tags           = var.tags

  env_vars = {
    BOT_KEY = var.telegram_bot_key
  }
}

module "service_lambda_execution_role" {
  source      = "github.com/vladcar/terraform-aws-service-prefixed-lambda-role.git?ref=v1.0.0"
  name_prefix = "tunturin-exec-${var.env_config.env}"
  additional_policies = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
  ]

  allowed_service_prefixes = [
    local.service_prefix
  ]
}

resource "aws_lambda_permission" "api_gateway_permission" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"
  source_arn    = "arn:aws:execute-api:${var.env_config.region}:${var.env_config.account_id}:${var.api_gateway_id}/*"
  function_name = module.prospect_creator_lambda.lambda_function_name
}
