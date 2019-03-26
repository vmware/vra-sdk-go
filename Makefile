.PHONY: swagger

all:
	go build -o sdk-test

swagger:
	swagger mixin -c=2 swagger/cas-iaas.json swagger/cas-blueprint.json swagger/cas-catalog.json swagger/cas-deployment.json > swagger/cas-combined.json
	swagger generate client -f swagger/cas-combined.json -t pkg
	./hack/fixup.sh

modified:
	git ls-files --modified | xargs git add
