BINARY_NAME=main.out

build:
	go build -o ${BINARY_NAME} ./cmd/app/main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}