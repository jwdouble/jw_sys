GO_Config := GO111MODULE=on


app:
		CGO_ENABLED=0 $(GO_Config) go build -tags release -o bin/app  cmd/main.go

run:
		CGO_ENABLED=1 $(GO_Config) go run -race -tags release  cmd/main.go


tidy:
		go mod tidy

clean:
		rm -rf bin