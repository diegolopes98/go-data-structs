GO=go
GOTEST=$(GO) test
GOCOVER=$(GO) tool cover

test:
	$(GOTEST) -v -coverprofile=coverage.out ./...

cover:
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out

clean:
	rm -rf coverage.out

.PHONY: test cover clean
