
module "gateway" {
  source     = "./modules/gateway"
  api_name   = "tunturin-webhook-api-${var.env_config.env}"
  env_config = var.env_config
  tags       = var.tags
}

module "webhook_handler" {
  source           = "./modules/backend"
  api_gateway_id   = module.gateway.api_gateway_id
  env_config       = var.env_config
  telegram_bot_key = var.telegram_bot_key
  allowed_chats    = var.allowed_chats
}

resource "aws_apigatewayv2_route" "route" {
  api_id             = module.gateway.api_gateway_id
  route_key          = "POST /webhook"
  target             = "integrations/${aws_apigatewayv2_integration.integration.id}"
  authorization_type = "NONE"
}

resource "aws_apigatewayv2_integration" "integration" {
  api_id                 = module.gateway.api_gateway_id
  integration_type       = "AWS_PROXY"
  integration_method     = "POST"
  integration_uri        = module.webhook_handler.invoke_arn
  description            = "webhook handler endpoint"
  connection_type        = "INTERNET"
  payload_format_version = "2.0"

  lifecycle {
    ignore_changes = [passthrough_behavior]
  }
}