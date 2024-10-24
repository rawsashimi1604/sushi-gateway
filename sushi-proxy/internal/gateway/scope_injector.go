package gateway

type ScopeInjector struct {
}

func NewScopeInjector() *ScopeInjector {
	return &ScopeInjector{}
}

func (si *ScopeInjector) InjectServicePluginScopes(service Service) {
	
}
