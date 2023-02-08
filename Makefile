MONGO_IMAGE := mongo:4.4
MONGO_INITDB_ROOT_USERNAME := admin
MONGO_INITDB_ROOT_PASSWORD := password
MONGO_CONTAINER_NAME := mongo
MONGO_PORT := 27017

dev:
	@echo "check docker"
	@if ! command -v docker &> /dev/null; then \
		echo "install docker..."; \
		brew install docker; \
	fi
	@echo "check colima"
	@if ! command -v colima &> /dev/null; then \
		echo "install colima"; \
		brew instasll colima; \
	fi
	@echo "start docker runtime"
	@colima start
	@echo "start mongo service on port ${MONGO_PORT}"
	@docker run -d \
		-p 27017:27017 \
			--name ${MONGO_CONTAINER_NAME} \
			-e MONGO_INITDB_ROOT_USERNAME=${MONGO_INITDB_ROOT_USERNAME} \
			-e MONGO_INITDB_ROOT_PASSWORD=${MONGO_INITDB_ROOT_PASSWORD} \
		${MONGO_IMAGE}