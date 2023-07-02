SRC = src/cmd/main.go
EXE = $(dir $(SRC))aoc

build:
	go build -o $(EXE) $(SRC)

run: build
	$(EXE) -day=$(day) -year=$(year)

create_day:
	# check if folder exist otherwise create folder
	[ -d "src/pkg/$(year)" ] || mkdir src/pkg/$(year)
	[ -d "src/pkg/$(year)/day_$(day)" ] || mkdir src/pkg/$(year)/day_$(day)

	# create files
	touch src/pkg/$(year)/day_$(day)/day_$(day).go
	touch src/pkg/$(year)/day_$(day)/input.txt

    # check length of src/pkg/$(year)/day_$(day)/day($(day).go, if it is empty,
    # add package line to the file
	if [ ! -s src/pkg/$(year)/day_$(day)/day_$(day).go ]; then \
        echo "package day_$$day" >> src/pkg/$$year/day_$$day/day_$$day.go; \
    fi

    # git add input and day file
	git add src/pkg/$(year)/day_$(day)/input.txt src/pkg/$(year)/day_$(day)/day_$(day).go
