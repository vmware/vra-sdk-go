.PHONY: all swagger modified update update-blueprint update-catalog-deployment update-iaas update-content update-pipeline update-project update-cmx clean
SWAGGER_VERSION=0.29.0
SWAGGER_ENDPOINT ?= api.mgmt.cloud.vmware.com

all:
	go build -o sdk-test

swagger: check-swagger
	rm -rf pkg/client pkg/models
	./hack/fix_iaas_swagger
	./hack/fix_blueprint_swagger
	./hack/fix_catalog_deployment_swagger
	./hack/fix_pipeline_swagger
	./hack/fix_project_swagger
	./hack/fix_cmx_swagger
	swagger mixin -c=1 swagger/vra-iaas-fixed.json swagger/vra-blueprint-fixed.json swagger/vra-catalog-deployment-fixed.json swagger/vra-content.json swagger/vra-pipeline-fixed.json swagger/vra-project-fixed.json swagger/vra-cmx-fixed.json | python3 -mjson.tool > swagger/vra-combined.json
	./hack/fix_vra_swagger --omit-security
	swagger generate client -f swagger/vra-combined.json -t pkg
	./hack/fixup.sh

check-swagger:
	@swagger version | grep ${SWAGGER_VERSION} > /dev/null || { echo "Wrong version of swagger command. Install go-swagger ${SWAGGER_VERSION}"; exit 1; }

modified:
	git ls-files --modified | xargs git add

update: update-blueprint update-catalog-deployment update-iaas update-content update-pipeline update-project update-cmx

update-blueprint:
	curl --insecure 'https://${SWAGGER_ENDPOINT}/blueprint/api/swagger/swagger-api-docs?group=2019-09-12' | python3 -mjson.tool > swagger/vra-blueprint.json

update-catalog-deployment:
	curl --insecure 'https://${SWAGGER_ENDPOINT}/deployment/api/swagger/swagger/v2/api-docs?group=2020-08-25' | python3 -mjson.tool > swagger/vra-catalog-deployment.json

update-iaas:
	curl --insecure 'https://${SWAGGER_ENDPOINT}/iaas/api/swagger?apiVersion=2021-07-15' | python3 -mjson.tool > swagger/vra-iaas.json

update-content:
	curl --insecure 'https://${SWAGGER_ENDPOINT}/content/api/swagger/v2/api-docs?group=2019-01-15' | python3 -mjson.tool > swagger/vra-content.json

update-pipeline:
	curl --insecure 'https://${SWAGGER_ENDPOINT}/pipeline/api/swagger/v2/api-docs?group=2019-10-17' | python3 -mjson.tool > swagger/vra-pipeline.json

update-project:
	curl --insecure 'https://${SWAGGER_ENDPOINT}/project/api/swagger/v2/api-docs?group=2019-01-15' | python3 -mjson.tool > swagger/vra-project.json

update-cmx:
	curl --insecure 'https://${SWAGGER_ENDPOINT}/cmx/v2/api-docs' | python3 -mjson.tool > swagger/vra-cmx.json

test:
	go build -o sdk-test

clean:
	rm swagger/vra-combined.json
