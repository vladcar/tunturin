output "api_gateway_id" {
  value = aws_apigatewayv2_api.gateway.id
}

output "api_uri" {
  value = aws_apigatewayv2_stage.stage.invoke_url
}
