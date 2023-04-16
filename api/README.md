# Sample Api for resource

### Generate Code

`docker run -u $(id -u):$(id -g) -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/resource-api.yaml -g go-gin-server -o /local/out/go`