NAME = solution
BIN_DIR = bin
INPUT_FILE_PATH = ../input.txt
BIN_PATH = $(BIN_DIR)/$(NAME)

build:
	go build -o $(BIN_PATH) cmd/$(NAME)/main.go

run: build
	./$(BIN_PATH) $(INPUT_FILE_PATH)

clean:
	rm -rf $(BIN_DIR)/*