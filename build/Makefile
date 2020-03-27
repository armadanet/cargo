IMAGE := nebula-cargo

build:
	cd .. && docker build -t $(IMAGE) -f build/Dockerfile .

tag: build
	docker tag $(IMAGE):latest codyperakslis/$(IMAGE):latest
	docker push codyperakslis/$(IMAGE):latest

run: build
	docker run -it $(IMAGE)
