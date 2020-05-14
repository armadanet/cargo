IMAGE := armada-cargo
TEST_IMAGE := armada-cargo-test

build:
	cd .. && docker build -t $(IMAGE) -f build/Dockerfile .

tag: build
	docker tag $(IMAGE):latest codyperakslis/$(IMAGE):latest
	docker push codyperakslis/$(IMAGE):latest

run: build
	docker run --name $(IMAGE) -it $(IMAGE)

test-build:
	cd .. && docker build --build-arg location=test -t $(TEST_IMAGE) -f build/Dockerfile .

test-tag: test-build
	docker tag $(TEST_IMAGE):latest codyperakslis/$(TEST_IMAGE):latest
	docker push codyperakslis/$(TEST_IMAGE):latest
