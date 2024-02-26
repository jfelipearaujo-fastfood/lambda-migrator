resource "aws_lambda_permission" "tfer--AllowS3Invoke" {
  action        = "lambda:InvokeFunction"
  function_name = "arn:aws:lambda:us-east-1:167192228103:function:lambda_migrator"
  principal     = "s3.amazonaws.com"
  source_arn    = "arn:aws:s3:::jsfelipearaujo"
  statement_id  = "AllowS3Invoke"
}
