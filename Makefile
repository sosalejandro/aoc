SRC = src/cmd/main.go
EXE = $(dir $(SRC))aoc

build:
	go build -o $(EXE) $(SRC)

run: build
	$(EXE) -day=$(day) -year=$(year)
