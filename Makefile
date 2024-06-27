# Include the .env file and export the variables
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

test:
	@go test -v -count=1 -ldflags '-extldflags -w' ./...