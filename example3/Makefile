.PHONY: build checkenv lint run push

# Before running make push, export DOCKERHUB_USER_NAME=aar1 (or whatever your user name is)
checkenv:
ifndef DOCKERHUB_USER_NAME
	$(error DOCKERHUB_USER_NAME is undefined)
endif

build: lint checkenv
	docker build --build-arg BUILD_ARG=qwerty -t ${DOCKERHUB_USER_NAME}/nicer-hash-service .

lint:
	hadolint Dockerfile
	# To do: add golang linting here

run: checkenv
ifndef SECRET1
	$(error Environment variable SECRET1 is undefined)
endif
ifndef SECRET2
	$(error Environment variable SECRET2 is undefined)
endif
	docker run -p 8080:8080 -e SECRET1 -e SECRET2 ${DOCKERHUB_USER_NAME}/nicer-hash-service

push: checkenv
	# Note: using the latest tag is not best practice. Use a version or commit ID.
	docker push ${DOCKERHUB_USER_NAME}/nicer-hash-service:latest
