GO_Config := GO111MODULE=on



app:
		@$(MAKE) tidy
		CGO_ENABLED=0 $(GO_Config) go build -tags release -o bin/app cmd/main.go

run:
		@$(MAKE) tidy
		CGO_ENABLED=0 $(GO_Config) go run -tags release cmd/main.go

tidy:
		go mod tidy

clean:
		rm -rf bin