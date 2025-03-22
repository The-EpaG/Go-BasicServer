# it is mandatory to use make targets within a POSIX-compliant shell (e.g. bash, zsh, etc.)

ifeq ($(OS),Windows_NT)
# for Windows
	EXE := .exe
else
# for Unix / Unix-like systems
	EXE :=
endif

BUILD_DIR			= .
BINARY_NAME		= main
BUILD_FILE 		= $(BUILD_DIR)/$(BINARY_NAME)$(EXE)
ENTRY_FILE		= main.go
RUN_CMD				= serve
BIN						= bin

HOOK_PRE_COMMIT 							= .git/hooks/pre-commit
HOOK_PREPARE_COMMIT_MSG				= .git/hooks/prepare-commit-msg
TOOL_WATCH 										= $(BIN)/CompileDaemon
TOOL_TEST											= $(BIN)/gotestsum
TOOL_LINT											=	$(BIN)/golangci-lint

.PHONY: build run tools hooks watch lint hooks clean deepclean test test-e2e

all: watch

$(BUILD_FILE): $(ENTRY_FILE)
	go build -o $@ $<

build:
	go build -o $(BUILD_FILE) $(ENTRY_FILE)

run: build
	./$(BUILD_FILE) $(RUN_CMD)

watch: $(BUILD_FILE) $(TOOL_WATCH)
	./$(TOOL_WATCH) -build="make build" -command="./$(BUILD_FILE) $(RUN_CMD)"

test: test-e2e

test-e2e: $(TOOL_TEST)
	./$(TOOL_TEST) --format=testname -- ./test/e2e

lint: $(TOOL_LINT)
	go fmt ./...
	./$(TOOL_LINT) run ./...

tools: $(TOOL_WATCH) $(TOOL_TEST) $(TOOL_LINT) hooks

$(TOOL_WATCH):
	@mkdir -p $(BIN)
	GOBIN=$(PWD)/$(BIN) go install github.com/githubnemo/CompileDaemon@latest

$(TOOL_TEST):
	@mkdir -p $(BIN)
	GOBIN=$(PWD)/$(BIN) go install gotest.tools/gotestsum@latest

$(TOOL_LINT):
	@mkdir -p $(BIN)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s v1.64.5

hooks: $(HOOK_PRE_COMMIT) $(HOOK_PREPARE_COMMIT_MSG)

$(HOOK_PRE_COMMIT):
	cp pre-commit .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit

$(HOOK_PREPARE_COMMIT_MSG):
	cp prepare-commit-msg .git/hooks/prepare-commit-msg
	chmod +x .git/hooks/prepare-commit-msg

clean:
	go clean

deepclean:
	go clean -testcache -cache -modcache
	rm -rf $(BIN)