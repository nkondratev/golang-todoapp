package core_http_server

import (
	"fmt"
	"net/http"
)

type APIVersion string

const (
	APIVersion1 APIVersion = "v1"
	APIVersion2 APIVersion = "v2"
	APIVersion3 APIVersion = "v3"
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion APIVersion
}

func NewAPIVersionRouter(apiVersion APIVersion) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:   http.NewServeMux(),
		apiVersion: apiVersion,
	}
}

func (r *APIVersionRouter) RegisterRouters(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		r.Handle(pattern, route.Handler)
	}
}
