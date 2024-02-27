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
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -buildvcs -vet=off ./...


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests
.PHONY: test
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

## run: run the  application
.PHONY: run
run: build 
	docker compose run -d -p 8080:8080 --name ${BINARY_NAME} ${DEV_BACKEND}

## run/live: run the applicatitn with reloading on file changes
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
## build: build the dev docker image
.PHONY: build
build:
	docker compose build ${DEV_BACKEND}

## build/dev/image: build the dev docker image
.PHONY: build/prod
build/prod:
	docker compose build ${PROD_BACKEND} 

# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #

## push: push changes to the remote Git repository
.PHONY: push
push: tidy audit no-dirty
	git push

## production/deploy: deploy the application to production
.PHONY: production/deploy
production/deploy: confirm tidy audit no-dirty build/prod 
	# Include additional deployment steps here...
	
## build/docker: build the application in docker
.PHONY: build/docker
build/docker:
    # Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	CGO_ENABLED=0 GOOS=linux go build -o /${BINARY_NAME} 

