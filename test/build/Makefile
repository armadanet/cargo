IMAGE_NAME := armada-cargo-test
CONTAINER_NAME := armada-cargo-test


build:
	cd .. && docker build -t $(IMAGE_NAME) -f build/Dockerfile .


run: build
	docker run -it --name $(CONTAINER_NAME) $(IMAGE_NAME)
