.PHONY: all cli gui clean

GOARCH := amd64
DIST_DIR := dist
CLI_DIR := cli
GUI_DIR := gui
BINARY_NAME := certyclick_verify

all: cli gui

cli:
	mkdir -p $(DIST_DIR)/macos_intel
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/macos_intel/$(BINARY_NAME)_cli ./$(CLI_DIR)
	mkdir -p $(DIST_DIR)/macos_m1
	GOOS=darwin GOARCH=arm64 go build -o $(DIST_DIR)/macos_m1/$(BINARY_NAME)_cli ./$(CLI_DIR)
	mkdir -p $(DIST_DIR)/windows
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/windows/$(BINARY_NAME)_cli.exe ./$(CLI_DIR)
	mkdir -p $(DIST_DIR)/linux
	GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/linux/$(BINARY_NAME)_cli ./$(CLI_DIR)

gui:
	mkdir -p $(DIST_DIR)
	go build -o $(DIST_DIR)/$(BINARY_NAME)_gui ./$(GUI_DIR)

clean:
	rm -rf $(DIST_DIR)/*
