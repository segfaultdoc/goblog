package service

import (
	"net/http"
	"github.com/callistaenterprise/goblog/common/tracing"
	"github.com/gorilla/mux"

)

// NewRouter creates a mux.Router pointer.
func NewRouter() *mux.Router {

	initQL(&LiveGraphQLResolvers{})

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(loadTracing(route.HandlerFunc, route.Name))
	}
	return router
}



func loadTracing(next http.Handler, handlerName string) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		span := tracing.StartHTTPTrace(req, handlerName)
		defer span.Finish()

		ctx := tracing.UpdateContext(req.Context(), span)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
