PROJECT_NAME := editor
GO := go
GOFMT := gofmt
GOLINT := golint
STATICCHECK := staticcheck

LDFLAGS := -w -s
BUILD_FLAGS := -ldflags="$(LDFLAGS)"

CMD_DIR := cmd
PKG_DIR := pkg
TEST_DIR := pkg/tests

BINARY_NAME := editor
BINARY_PATH := bin/$(BINARY_NAME)

RED := \033[0;31m
GREEN := \033[0;32m
YELLOW := \033[0;33m
BLUE := \033[0;34m
NC := \033[0m

.PHONY: all build test lint fmt clean help

all: fmt test build

build:
	@echo "$(BLUE)Building $(PROJECT_NAME)...$(NC)"
	@mkdir -p bin
	$(GO) build $(BUILD_FLAGS) -o $(BINARY_PATH) ./$(CMD_DIR)
	@echo "$(GREEN)Build completed: $(BINARY_PATH)$(NC)"

run: build
	@echo "$(BLUE)Running $(PROJECT_NAME)...$(NC)"
	@./$(BINARY_PATH)

test:
	@echo "$(BLUE)Running tests...$(NC)"
	$(GO) test -v -race ./$(TEST_DIR)/...
	@echo "$(GREEN)Tests completed$(NC)"

lint:
	@echo "$(BLUE)Running golint...$(NC)"
	@which $(GOLINT) > /dev/null 2>&1 || (echo "$(YELLOW)golint not installed, installing...$(NC)" && $(GO) install golang.org/x/lint/golint@latest)
	$(GOLINT) ./$(PKG_DIR)/... ./$(CMD_DIR)/...
	@echo "$(GREEN)Lint completed$(NC)"

fmt:
	@echo "$(BLUE)Formatting code...$(NC)"
	$(GOFMT) -w -s ./$(PKG_DIR) ./$(CMD_DIR)
	@echo "$(GREEN)Formatting completed$(NC)"

check-fmt:
	@echo "$(BLUE)Checking code format...$(NC)"
	@if [ -n "$$($(GOFMT) -l ./$(INTERNAL_DIR) ./$(CMD_DIR))" ]; then \
		echo "$(RED)Code is not properly formatted!$(NC)"; \
		$(GOFMT) -d ./$(PKG_DIR) ./$(CMD_DIR); \
		exit 1; \
	else \
		echo "$(GREEN)Code is properly formatted$(NC)"; \
	fi

clean:
	@echo "$(BLUE)Cleaning up...$(NC)"
	rm -rf bin/ coverage.out coverage.html
	$(GO) clean
	@echo "$(GREEN)Clean completed$(NC)"

install-tools:
	@echo "$(BLUE)Installing development tools...$(NC)"
	$(GO) install golang.org/x/lint/golint@latest
	$(GO) install honnef.co/go/tools/cmd/staticcheck@latest
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "$(GREEN)Tools installed$(NC)"

golangci-lint:
	@echo "$(BLUE)Running golangci-lint...$(NC)"
	@which golangci-lint > /dev/null 2>&1 || (echo "$(YELLOW)golangci-lint not installed, installing...$(NC)" && $(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./$(PKG_DIR)/... ./$(CMD_DIR)/...
	@echo "$(GREEN)GolangCI-Lint completed$(NC)"

help:
	@echo "$(BLUE)$(PROJECT_NAME) Makefile$(NC)"
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  $(GREEN)all$(NC)          - Run fmt, lint, vet, test and build"
	@echo "  $(GREEN)build$(NC)        - Build the project"
	@echo "  $(GREEN)run$(NC)          - Build and run the project"
	@echo "  $(GREEN)test$(NC)         - Run tests"
	@echo "  $(GREEN)lint$(NC)         - Run golint"
	@echo "  $(GREEN)fmt$(NC)          - Format Go code"
	@echo "  $(GREEN)check-fmt$(NC)    - Check if code is properly formatted"
	@echo "  $(GREEN)clean$(NC)        - Remove build artifacts"
	@echo "  $(GREEN)install-tools$(NC) - Install required development tools"
	@echo "  $(GREEN)golangci-lint$(NC) - Run golangci-lint (comprehensive)"
	@echo "  $(GREEN)help$(NC)         - Show this help message"