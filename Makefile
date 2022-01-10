GO_Config := GO111MODULE=on



app:
		@$(MAKE) pb
		@$(MAKE) tidy
		CGO_ENABLED=0 $(GO_Config) go build -tags release -o bin/app cmd/main.go

run:
		@$(MAKE) pb
		@$(MAKE) tidy
		CGO_ENABLED=0 $(GO_Config) go run -tags release -o bin/app cmd/main.go
pb:
		bash ./proto/build_go.sh
		bash ./proto/build_doc.sh

tidy:
		go mod tidy

clean:
		rm -rf bin