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
	go fmt ./cmd/...
	go fmt ./infrastructure/aws/sdk/...
	go fmt ./infrastructure/persistence/...
	go fmt ./infrastructure/transport/...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./core/...
	go vet ./cmd/...
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

## run: DEPRICATED run the  application
.PHONY: run
run: build prune
	docker compose run -p 8080:8080 --name ${BINARY_NAME} ${DEV_BACKEND} 
	
## run: DEPRICATED run the  application
.PHONY: run/prod
run/prod: build/prod prune
	docker compose run -p 8080:8080 --name ${BINARY_NAME} ${PROD_BACKEND} 

## run/live: DEPRICATED run the applicatitn with reloading on file changes
.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "/tmp/bin/${BINARY_NAME}" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

# ==================================================================================== #
# DOCKER
# ==================================================================================== #
#
## build: DEPRICATED build the dev docker image
.PHONY: build
build:
	docker compose build ${DEV_BACKEND}

## build/prod: DEPRICATED build the prod docker image
.PHONY: build/prod
build/prod:
	docker compose build ${PROD_BACKEND}


## prune: DEPRICATED remove all stopped containers and unused images
.PHONY: prune
prune:
	docker system prune -f
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
production/deploy: confirm tidy no-dirty deploy/prod
	# Include additional deployment steps here...
	
## build/docker: DEPRICATED build the application in docker
.PHONY: build/docker
build/docker:
    # Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	CGO_ENABLED=0 GOOS=linux go build -o /${BINARY_NAME} ${MAIN_PACKAGE_PATH}

# ==================================================================================== #
# AWS CDK
# ==================================================================================== #
#
## deploy/prod: deploy the application to production
.PHONY: deploy/prod
deploy/prod:
	cd ./infrastructure/aws/cdk; npm run build; cdk deploy --all

## destroy/prod: destroy the application in production
.PHONY: destroy/prod
destroy/prod:
	cd ./infrastructure/aws/cdk; cdk destroy --all
