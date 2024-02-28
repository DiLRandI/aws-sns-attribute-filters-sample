.PHONY: build clean deploy
ARCH?=amd64

build:
	env GOARCH=$(ARCH) GOOS=linux CGO_ENABLED=0 go build -tags lambda.norpc -trimpath -ldflags="-s -w" -o bin/publisher/publisher publisher/main.go
	make write-bootstrap OUTPUT=publisher APP_NAME=publisher
	env GOARCH=$(ARCH) GOOS=linux CGO_ENABLED=0 go build -tags lambda.norpc -trimpath -ldflags="-s -w" -o bin/sqs-subscriber/sqs-subscriber sqs-subscriber/main.go
	make write-bootstrap OUTPUT=sqs-subscriber APP_NAME=sqs-subscriber
	env GOARCH=$(ARCH) GOOS=linux CGO_ENABLED=0 go build -tags lambda.norpc -trimpath -ldflags="-s -w" -o bin/sns-subscriber/sns-subscriber sns-subscriber/main.go
	make write-bootstrap OUTPUT=sns-subscriber APP_NAME=sns-subscriber

zip:
	zip -j -9 bin/publisher.zip bin/publisher/*
	zip -j -9 bin/sqs-subscriber.zip bin/sqs-subscriber/*
	zip -j -9 bin/sns-subscriber.zip bin/sns-subscriber/*

clean:
	rm -rf ./bin

deploy: clean build zip
	./node_modules/.bin/serverless -c serverless.yml deploy --verbose

write-bootstrap:
	echo "#!/bin/sh\n./$(APP_NAME)" > bin/$(OUTPUT)/bootstrap