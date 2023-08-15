.PHONY: build checkenv lint run push

# Before running make push, export DOCKERHUB_USER_NAME=aar1 (or whatever your user name is)
checkenv:
ifndef DOCKERHUB_USER_NAME
	$(error DOCKERHUB_USER_NAME is undefined)
endif

build: lint checkenv
	# Note: passing a secret via an arg is bad practice!
	#   1. It's written into a file in GitHub where everyone can see it
	#   2. It will likely appear in the build output, CI/CD logs so it could leak that way
	#   3. There's a fair chance it could end up in the final Docker image unless you're super careful
	# Instead use a mechanism like https://docs.docker.com/engine/reference/builder/#run---mounttypesecret
	docker build --build-arg SECRET1=hBFKd/3krLQCHK85WpOgMsKLplnCK7s/NvBP --build-arg BUILD_ARG=qwerty -t ${DOCKERHUB_USER_NAME}/hash-service .

lint:
	hadolint Dockerfile
	# To do: add golang linting here

run: checkenv
	docker run -p 8080:8080 ${DOCKERHUB_USER_NAME}/hash-service

push: checkenv
	# Note: using the latest tag is not best practice. Use a version or commit ID.
	docker push ${DOCKERHUB_USER_NAME}/hash-service:latest
