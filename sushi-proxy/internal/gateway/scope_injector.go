package gateway

import "github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"

type ScopeInjector struct {
}

func NewScopeInjector() *ScopeInjector {
	return &ScopeInjector{}
}

func (si *ScopeInjector) InjectServicePluginScopes(service model.Service) {

}
