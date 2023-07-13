.PHONY: all cli gui clean

GOARCH := amd64
DIST_DIR := dist
CLI_DIR := cli
GUI_DIR := gui
BINARY_NAME := certyclick_verify

all: cli gui

cli: cli-macos cli-macos-m1 cli-windows cli-linux

cli-macos:
	mkdir -p $(DIST_DIR)/macos
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/macos/$(BINARY_NAME)_cli ./$(CLI_DIR)
cli-macos-m1:
	mkdir -p $(DIST_DIR)/macos
	GOOS=darwin GOARCH=arm64 go build -o $(DIST_DIR)/macos/$(BINARY_NAME)_cli ./$(CLI_DIR)
cli-windows:
	mkdir -p $(DIST_DIR)/windows
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/windows/$(BINARY_NAME)_cli.exe ./$(CLI_DIR)
cli-linux:
	mkdir -p $(DIST_DIR)/ubuntu
	GOOS=linux GOARCH=amd64 go build -o $(DIST_DIR)/ubuntu/$(BINARY_NAME)_cli ./$(CLI_DIR)

gui: gui-macos gui-macos-m1 gui-windows gui-linux

gui-macos:
	mkdir -p $(DIST_DIR)/macos
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/macos/$(BINARY_NAME)_gui ./$(GUI_DIR)
gui-macos-m1:
	mkdir -p $(DIST_DIR)/macos
	GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/macos/$(BINARY_NAME)_gui ./$(GUI_DIR)
gui-windows:
	mkdir -p $(DIST_DIR)/windows
	GOOS=windows GOARCH=amd64 go build -o $(DIST_DIR)/windows/$(BINARY_NAME)_gui.exe ./$(GUI_DIR)
gui-linux:
	mkdir -p $(DIST_DIR)/ubuntu
	GOOS=linux GOARCH=amd64 go build -v -gcflags="all=-N -l" -o $(DIST_DIR)/ubuntu/$(BINARY_NAME)_gui ./$(GUI_DIR)

clean:
	rm -rf $(DIST_DIR)/*

test:
	go test ./...