all: game

game:
	go build .

clean:
	go clean -v

.PHONY:
	all