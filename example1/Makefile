# Override this to your own Docker Hub user name:
DOCKER_USER_NAME=aar1

build:
	docker build -t ${DOCKER_USER_NAME}/hello-service .

run:
	docker run -p 8080:8080 ${DOCKER_USER_NAME}/hello-service

push:
	docker push ${DOCKER_USER_NAME}/hello-service:latest