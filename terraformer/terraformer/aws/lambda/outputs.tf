output "aws_lambda_function_tfer--lambda_migrator_id" {
  value = "${aws_lambda_function.tfer--lambda_migrator.id}"
}

output "aws_lambda_permission_tfer--AllowS3Invoke_id" {
  value = "${aws_lambda_permission.tfer--AllowS3Invoke.id}"
}
