GO=go
GOCOVER=$(GO) tool cover

test/cover:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out

.PHONY: test/cover
