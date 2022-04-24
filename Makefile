TESTFLAGS	= -v -cover

test:
	@go test $(TESTFLAGS) ./...