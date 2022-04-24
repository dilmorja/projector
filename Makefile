TESTFLAGS	= -v -cover

test:
	@go test $(TESTFLAGS) ./...
install: test
	@go install