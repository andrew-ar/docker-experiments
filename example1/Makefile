.PHONY: build checkenv lint run push

# Before running make push, export DOCKERHUB_USER_NAME=aar1 (or whatever your user name is)
checkenv:
ifndef DOCKERHUB_USER_NAME
	$(error DOCKERHUB_USER_NAME is undefined)
endif

build: lint checkenv
	docker build -t ${DOCKERHUB_USER_NAME}/hello-service .

lint:
	hadolint Dockerfile
	# To do: add golang linting here

run: checkenv
	docker run -p 8080:8080 ${DOCKERHUB_USER_NAME}/hello-service

push: checkenv
	docker push ${DOCKERHUB_USER_NAME}/hello-service:latest
