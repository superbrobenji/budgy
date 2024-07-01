MAIN_PACKAGE_PATH := ./cmd
BINARY_NAME := budgy
DEV_BACKEND := backend_dev
PROD_BACKEND := backend_prod
# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	git diff --exit-code


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./core/...
	go fmt ./infrastructure/aws/sdk/...
	go fmt ./infrastructure/persistence/...
	go fmt ./infrastructure/transport/...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./core/...
	go vet ./infrastructure/aws/sdk/...
	go vet ./infrastructure/persistence/...
	go vet ./infrastructure/transport/...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -buildvcs -vet=off ./...


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests
.PHONY:test
test:
	go test -v -buildvcs ./...

## test: run all tests
.PHONY: test/race
test/race:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #
#

## push: push changes to the remote Git repository
.PHONY: push
push: tidy audit no-dirty
	git push

## production/deploy: deploy the application to production
.PHONY: production/deploy
production/deploy: confirm tidy no-dirty deploy/all
	# Include additional deployment steps here...
	
j ==================================================================================== #
# AWS CDK
# ==================================================================================== #
#
## npm/install: install dependencies for cdk
.PHONY: npm/install
npm/install:
	cd ./infrastructure/aws/cdk; npm install

## deploy/all: deploy the application to production
.PHONY: deploy/all
deploy/all:
	cd ./infrastructure/aws/cdk; npm run build; cdk deploy --all

## deploy/prod: deploy the application to production
.PHONY: deploy/db
deploy/db:
	cd ./infrastructure/aws/cdk; npm run build; cdk deploy BudgyDatabaseStack

## deploy/prod: deploy the application to production
.PHONY: deploy/api
deploy/api:
	cd ./infrastructure/aws/cdk; npm run build; cdk deploy BudgyApiStack

## deploy/prod: deploy the application to production
.PHONY: deploy/auth
deploy/auth:
	cd ./infrastructure/aws/cdk; npm run build; cdk deploy BudgyAuthStack
## destroy/prod: destroy the application in production
.PHONY: destroy/all
destroy/all:
	cd ./infrastructure/aws/cdk; cdk destroy --all
