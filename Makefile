build:
	@echo "Building..."
	@env GOOS=linux GOARCH=arm64 go build -o terraform/bootstrap main.go

zip:
	@echo "Zipping..."
	@zip terraform/lambda.zip terraform/bootstrap

upload:
	@echo "Creating folder..."
	@aws s3api put-object --bucket jsfelipearaujo --key "migrations/" > /dev/null
	@echo "Uploading..."
	@aws s3 cp scripts/ s3://jsfelipearaujo/migrations --recursive

init:
	@echo "Initializing..."
	@cd terraform \
		&& terraform init -reconfigure

check:
	@echo "Checking..."
	make fmt && make validate && make plan

plan:
	@echo "Planning..."
	@cd terraform \
		&& terraform plan -out=plan \
		&& terraform show -json plan > plan.tfgraph

fmt:
	@echo "Formatting..."
	@cd terraform \
		&& terraform fmt -check

validate:
	@echo "Validating..."
	@cd terraform \
		&& terraform validate

apply:
	@echo "Applying..."
	@cd terraform \
		&& terraform apply plan

destroy:
	@echo "Destroying..."
	@cd terraform \
		&& terraform destroy -auto-approve

former:
	@echo "Initializing Terraform..."
	@cd terraformer \
		&& terraform init
	@echo "Running Terraformer..."
	@cd terraformer \
		&& terraformer import aws -o terraformer --regions=us-east-1 --resources=lambda