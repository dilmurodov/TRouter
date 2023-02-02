package trouter

import (
	"net/http"
	"sync"
)

type Context struct {
	Request *http.Request
	Writer  ResponseWriter

	Params   Params
	handlers HandlersChain
	index    int8
	fullPath string

	engine *TRouter
	params *Params

	// This mutex protects Keys map.
	mu sync.RWMutex

	// Keys is a key/value pair exclusively for the context of each request.
	Keys map[string]any

	// Errors is a list of errors attached to all the handlers/middlewares who used this context.
	Errors error

	// Accepted defines a list of manually accepted formats for content negotiation.
	Accepted []string
}
