package di

import (
	"github.com/soulcodex/karma-api/configs"
	httpserver "github.com/soulcodex/karma-api/pkg/http-server"
)

type RouterRegisters []RouterFunc
type RouterFunc func(services *CommonServices, cfg configs.Config)

type RouteRegisterer struct {
	router           *httpserver.Router
	routesToRegister RouterRegisters
}

func NewRouteRegisterer(router *httpserver.Router) *RouteRegisterer {
	return &RouteRegisterer{
		router: router,
	}
}

func (rr *RouteRegisterer) RegisterModuleRoutes(registerer RouterFunc) {
	rr.routesToRegister = append(rr.routesToRegister, registerer)
}

func (rr *RouteRegisterer) RegisterAllModulesRoutesOnRouter(c *CommonServices) {
	for _, routerFunc := range rr.routesToRegister {
		routerFunc(c, c.Config)
	}
}
