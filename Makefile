.PHONY: all swagger modified update update-blueprint update-catalog update-deployment update-iaas clean

all:
	go build -o sdk-test

swagger:
	rm -rf pkg/client pkg/models
	swagger mixin -c=2 swagger/cas-iaas.json swagger/cas-blueprint.json swagger/cas-catalog.json swagger/cas-deployment.json > swagger/cas-combined.json
	swagger generate client -f swagger/cas-combined.json -t pkg
	./hack/fixup.sh

modified:
	git ls-files --modified | xargs git add

update: update-blueprint update-catalog update-deployment update-iaas

update-blueprint:
	curl 'https://api.mgmt.cloud.vmware.com/blueprint/api/swagger/swagger-api-docs?group=blueprint' | python3 -mjson.tool > swagger/cas-blueprint.json

update-catalog:
	curl 'https://api.mgmt.cloud.vmware.com/catalog/api/swagger/swagger/v2/api-docs?group=catalog' | python3 -mjson.tool > swagger/cas-catalog.json

update-deployment:
	curl 'https://api.mgmt.cloud.vmware.com/deployment/api/swagger/swagger/v2/api-docs?group=deployments' | python3 -mjson.tool > swagger/cas-deployment.json

update-iaas:
	curl 'https://api.mgmt.cloud.vmware.com/iaas/api/swagger/swagger/v2/api-docs?group=iaas' | python3 -mjson.tool > swagger/cas-iaas.json

clean:
	rm swagger/cas-combined.json
