# Change these variables as necessary.
# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


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


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	go test -v -race ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## proto/gen: generate proto directories and files
.PHONY: proto/gen
proto/gen:
	@buf generate \
		./proto/steamdatabase/steam \
		--template '{"version":"v1","managed":{"enabled":true,"go_package_prefix":{"default":"steamdatabase/steam"}},"plugins":[{"plugin":"go","out":"proto/gen/go"}]}'
	@buf generate \
		./proto/steamdatabase/csgo \
		--template '{"version":"v1","managed":{"enabled":true,"go_package_prefix":{"default":"steamdatabase/csgo"}},"plugins":[{"plugin":"go","out":"proto/gen/go"}]}' \
		--exclude-path ./proto/steamdatabase/csgo/steammessages_base.proto \
		--exclude-path ./proto/steamdatabase/csgo/enums_clientserver.proto

## proto/clean: clean proto directories and files
.PHONY: proto/clean
proto/clean:
	@rm -rf ./proto/gen
