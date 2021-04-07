.PHONY: build run generate-pdf

build:
	docker build --build-arg USER_ID=1000 --build-arg GROUP_ID=1000 -t html-to-pdf-development .

run:
	docker run -it --rm -v $(PWD)/:/go/src/html-to-pdf html-to-pdf-development

generate-pdf:
	docker run -it --rm -v $(PWD)/:/go/src/html-to-pdf html-to-pdf-development go run cmd/main.go
