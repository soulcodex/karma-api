package httpserver

import (
	"context"

	"github.com/gorilla/mux"

	"net/http"
	"time"
)

type Router struct {
	muxRouter *mux.Router
	server    *http.Server
}

type RouterOptions struct {
	WriteTimeout int
	ReadTimeout  int
	Middlewares  []Middleware
}

type Middleware func(next http.Handler) http.Handler

func NewRouter(options RouterOptions) Router {
	readTimeout := 30 * time.Second
	writeTimeout := 30 * time.Second

	if options.ReadTimeout != 0 {
		readTimeout = time.Duration(options.ReadTimeout) * time.Second
	}

	if options.WriteTimeout != 0 {
		writeTimeout = time.Duration(options.WriteTimeout) * time.Second
	}

	m := muxRouter()
	m.Use(enableCorsMiddleware)

	for _, middleware := range options.Middlewares {
		m.Use(mux.MiddlewareFunc(middleware))
	}

	h := &http.Server{
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}

	return Router{m, h}
}

func muxRouter() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}

func DefaultRouter(serverWriteTimeOut int, serverReadTimeout int, middlewares ...Middleware) *Router {
	options := RouterOptions{
		WriteTimeout: serverReadTimeout,
		ReadTimeout:  serverWriteTimeOut,
		Middlewares:  middlewares,
	}

	rt := NewRouter(options)

	return &rt
}

func (router *Router) ListenAndServe(addr string) error {
	router.server.Addr = addr
	router.server.Handler = router.muxRouter

	// Use default options
	return router.server.ListenAndServe()
}

func (router *Router) Shutdown(ctx context.Context) error {
	return router.server.Shutdown(ctx)
}

func (router *Router) GetMuxRouter() *mux.Router {
	return router.muxRouter
}

func (router *Router) AddMiddleware(middlewares ...Middleware) {
	for _, middleware := range middlewares {
		router.muxRouter.Use(mux.MiddlewareFunc(middleware))
	}
}

func (router *Router) Handle(method string, path string, handler http.Handler, middlewares ...Middleware) {
	router.handleMultipleMethods([]string{method}, path, handler, middlewares...)
}

func (router *Router) handleMultipleMethods(
	methods []string,
	path string,
	handler http.Handler,
	middlewares ...Middleware,
) {
	subRouter := router.muxRouter.NewRoute().Subrouter()
	for _, middleware := range middlewares {
		subRouter.Use(mux.MiddlewareFunc(middleware))
	}

	subRouter.HandleFunc(path, handler.ServeHTTP).Methods(methods...)
}

func (router *Router) Get(path string, handler http.Handler, middlewares ...Middleware) {
	router.Handle(http.MethodGet, path, handler, middlewares...)
}

func (router *Router) Post(path string, handler http.Handler, middlewares ...Middleware) {
	router.Handle(http.MethodPost, path, handler, middlewares...)
}

func (router *Router) Put(path string, handler http.Handler, middlewares ...Middleware) {
	router.Handle(http.MethodPut, path, handler, middlewares...)
}

func (router *Router) Patch(path string, handler http.Handler, middlewares ...Middleware) {
	router.Handle(http.MethodPatch, path, handler, middlewares...)
}

func (router *Router) Delete(path string, handler http.Handler, middlewares ...Middleware) {
	router.Handle(http.MethodDelete, path, handler, middlewares...)
}

func (router *Router) Head(path string, handler http.Handler, middlewares ...Middleware) {
	router.Handle(http.MethodHead, path, handler, middlewares...)
}

func (router *Router) Options(path string, handler http.Handler, middlewares ...Middleware) {
	router.Handle(http.MethodOptions, path, handler, middlewares...)
}

func (router *Router) Route(methods []string, path string, handler http.Handler, middlewares ...Middleware) {
	router.handleMultipleMethods(methods, path, handler, middlewares...)
}
