# Sample Api for resource

### Api
The resource's server Api can be found [here](resource-api.yaml).

### Generate API Code
To generate the code run:

`docker run -u $(id -u):$(id -g) -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/resource-api.yaml -g go-gin-server -o /local/out/go`