build:
	@echo "Building..."
	@env GOOS=linux GOARCH=arm64 go build -o terraform/bootstrap main.go

zip:
	@echo "Zipping..."
	@zip terraform/lambda.zip terraform/bootstrap

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