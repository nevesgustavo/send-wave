
# GENERAL VARS
SEND-WAVE-LAMBDA_FILE=send-wave-v0.0.1.zip

CURRENT_DIR = $(shell pwd)

hooks:
ifeq ($(shell which $(CURRENT_DIR)/.git/hooks/pre-push), )
	@echo "Installing pre-push git hooks"
	$(shell chmod +x $(CURRENT_DIR)/githooks/pre-push)
	$(shell ln -s $(CURRENT_DIR)/githooks/pre-push $(CURRENT_DIR)/.git/hooks/pre-push)
else
	@echo "Git pre-push hooks already installed"
endif

# LINT / VET / TEST
lint:
	@echo "Linting ..."
	@golangci-lint run --timeout=10m

vet:
	@echo "Vetting ..."
	@go vet ./...

test:
	@echo "Running unit tests..."
	@go clean -testcache go test ./... -cover -v -coverprofile=cover.out

coverage: test
	@go tool cover -html=cover.out -o cover.html

#########################################
# LAMBDA ACTIONS #
#########################################

build-binaries-send-wave:
	@echo "Building binaries..."
	@GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o cmd/send-wave cmd/main.go

zip-binaries-send-wave:
	@echo "Zipping binaries..."
	@zip -j cmd/${SEND-WAVE-LAMBDA_FILE} cmd/send-wave

build-and-zip: build-binaries-send-wave zip-binaries-send-wave