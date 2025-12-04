BIN_DIR = bin
SOLUTION_BIN_PATH = $(BIN_DIR)/solution

build:
	go build -o $(SOLUTION_BIN_PATH) cmd/solution/main.go

run: build
	./$(SOLUTION_BIN_PATH)

clean:
	rm -rf $(BIN_DIR)/*