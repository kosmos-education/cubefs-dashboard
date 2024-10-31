# cfs-gui Makefile

phony := build 
phony += clean
phony += backend
phony += frontend

DOCKER_IMAGE="cubefs-dashboard"

build: frontend backend

docker-build:
	docker build . -t "${DOCKER_IMAGE}"

docker-push:
	docker push "${DOCKER_IMAGE}"

frontend:
	sh build.sh --frontend

backend:
	sh build.sh --backend
clean: bin
	rm -rf bin

.PHONY: $(phony)
