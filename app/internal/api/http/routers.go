package http

import (
	"github.com/stefnef/Flowingo/m/internal/api/http/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. examples GET, POST etc.
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter(infoHandler handler.InfoHandler, resourceHandler handler.ResourceHandler, errorHandler handler.ErrorHandler) *gin.Engine {
	router := gin.Default()
	routes := getRoutes(infoHandler, resourceHandler)

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, errorHandler.HandleErrors, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, errorHandler.HandleErrors, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

func getRoutes(infoHandler handler.InfoHandler, resourceHandler handler.ResourceHandler) Routes {
	return Routes{
		{
			"InfoGet",
			http.MethodGet,
			"/info",
			infoHandler.GetInfo,
		},

		{
			"ResourceGet",
			http.MethodGet,
			"/resource",
			resourceHandler.GetResources,
		},

		{
			"ResourceIdGet",
			http.MethodGet,
			"/resource/:id",
			resourceHandler.GetResource,
		},

		{
			"ResourcePost",
			http.MethodPost,
			"/resource",
			resourceHandler.PostResource,
		},
	}
}
