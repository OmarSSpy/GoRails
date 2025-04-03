package router

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, map[string]string)

type Route struct {
	pattern *regexp.Regexp
	params  []string
	handler HandlerFunc
}
type Router struct {
	routes map[string][]Route
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string][]Route),
	}
}

func (r *Router) Handle(method, path string, handler HandlerFunc) {
	parts := strings.Split(path, "/")
	var regexParts []string
	var params []string

	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			paramName := part[1:]
			params = append(params, paramName)
			regexParts = append(regexParts, `(?P<`+paramName+`>[^/]+)`)
		} else {
			regexParts = append(regexParts, part)
		}
	}

	pattern := "^" + strings.Join(regexParts, "/") + "$"
	compiledRegex := regexp.MustCompile(pattern)

	r.routes[method] = append(r.routes[method], Route{
		pattern: compiledRegex,
		params:  params,
		handler: handler,
	})
}

func (r *Router) GET(path string, handler HandlerFunc) {
	r.Handle("GET", path, handler)
}

func (r *Router) POST(path string, handler HandlerFunc) {
	r.Handle("POST", path, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path

	for _, route := range r.routes[method] {
		if matches := route.pattern.FindStringSubmatch(path); matches != nil {
			params := make(map[string]string)
			for i, name := range route.pattern.SubexpNames() {
				if i > 0 && name != "" {
					params[name] = matches[i]
				}
			}
			route.handler(w, req, params)
			return
		}
	}
	http.NotFound(w, req)
	fmt.Fprintf(w, "GORAILS: 404 Not Found -> %s %s", method, path)
}
