
resource "aws_apigatewayv2_api" "gateway" {
  name                         = var.api_name
  protocol_type                = "HTTP"
  disable_execute_api_endpoint = false
  description                  = "Tunturin telegram bot api"
  tags                         = var.tags
}

resource "aws_apigatewayv2_stage" "stage" {
  api_id      = aws_apigatewayv2_api.gateway.id
  name        = var.env_config.env
  auto_deploy = true

  lifecycle {
    ignore_changes = [deployment_id]
  }
}

