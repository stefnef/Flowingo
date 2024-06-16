/*
 * resource
 *
 * resource
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	sw "github.com/stefnef/Flowingo/m/internal/api/http"
	"github.com/stefnef/Flowingo/m/internal/api/http/handler"
	"github.com/stefnef/Flowingo/m/internal/core/service"
	"github.com/stefnef/Flowingo/m/internal/repository"
	"github.com/stefnef/Flowingo/m/pkg"
	"log"
)

func main() {
	log.Printf("Server started")

	var generator = pkg.NewGeneratorImpl()
	var resourceRepository = repository.NewInternalResourceRepository(generator)
	var infoHandler = handler.NewInfoHandler(service.NewInfoService())
	var resourceHandler = handler.NewResourceHandler(service.NewResourceService(resourceRepository, generator))
	var errorHandler = handler.NewErrorHandler()

	router := sw.NewRouter(infoHandler, resourceHandler, errorHandler)
	_ = router.SetTrustedProxies(nil)

	log.Fatal(router.Run(":8080"))
}
