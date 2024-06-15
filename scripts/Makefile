.PHONY:

VERSION ?= 1.0.0
USER_NAME = "Piorosen"
NAMESPACE = "piorosen"

login:
	docker login ghcr.io -u $(USER_NAME);

deploy_collection:
# make deploy_collection VERSION=1.0.0

	cd ../collection_from_telegram && \
	docker build -t ghcr.io/$(NAMESPACE)/colllection-telegram:$(VERSION) .
	docker push ghcr.io/$(NAMESPACE)/colllection-telegram:$(VERSION)

deploy_converter:
# make deploy_collection VERSION=1.0.0

	cd ../collection_from_telegram && \
	docker build -t ghcr.io/$(NAMESPACE)/colllection-converter:$(VERSION) .
	docker push ghcr.io/$(NAMESPACE)/colllection-converter:$(VERSION)
