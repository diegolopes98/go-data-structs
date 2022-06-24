GO=go
GOTEST=$(GO) test
GOCOVER=$(GO) tool cover

test/cover:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out

clean:
	rm -rf coverage.out

.PHONY: test/cover
