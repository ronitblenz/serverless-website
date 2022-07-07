.PHONY: build

build:
	sam build

deploy-infra:
	sam build && aws-vault exec ronitblenz --no-session -- sam deploy

deploy-site:
	aws-vault exec ronitblenz --no-session -- aws s3 sync ./resume-site s3://ronit-demo-site