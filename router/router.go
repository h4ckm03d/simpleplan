package router

import (
	"net/http"
	"path"
)

// Middleware type defines the function signature for middleware implementation
type Middleware func(http.Handler) http.Handler

// Router implements the needed methods for the Dispatcher
// to be able to match and execute requests.
type Router interface {
	// Add takes a route path and a handler to store for further matching
	Add(method string, path string, handler http.Handler)

	// Wrap takes a Middleware to wrap all handlers in order (from inside out) at router level.
	Wrap(Middleware)

	// Match checks if a request matches this router.
	// If so, adds the route parameters to the request context and returns the corresponding handler.
	// If route doesn't matches, the response is nil
	Match(*http.Request) http.Handler
}

// New creates a new Router with the provided prefix
func New(prefix string) Router {
	// Create router
	return &router{
		prefix:     prefix,
		tree:       rootNode("GET", "/", nil),
		middleware: make([]Middleware, 0),
	}
}

// router implements Router interface
type router struct {
	// Routes prefix for this router
	prefix string

	// Routes tree
	tree *node

	// Middlewares collection
	middleware []Middleware
}

func (r *router) Add(method, route string, h http.Handler) {
	r.tree.add(method, path.Join(r.prefix, route), h)
}

func (r *router) Wrap(m Middleware) {
	r.middleware = append(r.middleware, m)
}

func (r *router) Match(req *http.Request) http.Handler {
	h := r.tree.match(req)
	if h == nil {
		return nil
	}
	for _, m := range r.middleware {
		h = m(h)
	}

	return h
}

type routeParamsKey struct{}

// Params returns a map[string]string containing all route parameters
func Params(req *http.Request) map[string]string {
	params := req.Context().Value(routeParamsKey{})
	if _, ok := params.(map[string]string); ok {
		return params.(map[string]string)
	}

	return nil
}

// Param is a convenience function to retrieve a route param from the current request.
func Param(req *http.Request, key string) string {
	params := Params(req)
	if params != nil {
		if v, ok := params[key]; ok {
			return v
		}
	}

	return ""
}
