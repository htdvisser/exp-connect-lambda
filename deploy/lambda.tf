data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "lambda_execution_role" {
  name               = "${local.project_name}-execution"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
  tags               = local.common_tags
}

resource "aws_iam_role_policy_attachment" "lambda_execution_role_attachment" {
  role       = aws_iam_role.lambda_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "lambda_function" {
  function_name    = "${local.project_name}-lambda"
  runtime          = "go1.x"
  architectures    = ["x86_64"] # Runtime go1.x does not support "arm64" yet.
  handler          = "lambda"
  filename         = "../dist/lambda.zip"
  source_code_hash = filebase64sha256("../dist/lambda.zip")
  timeout          = 15
  role             = aws_iam_role.lambda_execution_role.arn
  tags             = local.common_tags
}

resource "aws_lambda_function_url" "lambda_function_url" {
  function_name      = aws_lambda_function.lambda_function.function_name
  authorization_type = "NONE"
}

output "lambda_function_url" {
  description = "Lambda Function URL"
  value       = aws_lambda_function_url.lambda_function_url.function_url
}
