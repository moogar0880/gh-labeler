IMAGE_NAME=moogar0880/ghlabeler
VERSION=$(shell git describe --tags 2> /dev/null || echo 'latest')

.PHONY: image vendor

image: vendor
	docker build -t $(IMAGE_NAME):$(VERSION) .

vendor:
	glide install
