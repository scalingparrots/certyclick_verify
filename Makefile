.PHONY: all cli gui clean

GOOS ?= $(shell go env GOOS)
GOARCH := amd64
DIST_DIR := dist
CLI_DIR := cli
GUI_DIR := gui
BINARY_NAME := certyclick_verify

all: cli gui

cli:
	mkdir -p $(DIST_DIR)/macos
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/macos/$(BINARY_NAME)_cli ./$(CLI_DIR)
	mkdir -p $(DIST_DIR)/windows
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/windows/$(BINARY_NAME)_cli.exe ./$(CLI_DIR)
	mkdir -p $(DIST_DIR)/linux
	GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/linux/$(BINARY_NAME)_cli ./$(CLI_DIR)

gui:
	mkdir -p $(DIST_DIR)/macos
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/macos/$(BINARY_NAME)_gui ./$(GUI_DIR)
	mkdir -p $(DIST_DIR)/windows
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/windows/$(BINARY_NAME)_gui.exe ./$(GUI_DIR)
	mkdir -p $(DIST_DIR)/linux
	GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/linux/$(BINARY_NAME)_gui ./$(GUI_DIR)

clean:
	rm -rf $(DIST_DIR)/*
