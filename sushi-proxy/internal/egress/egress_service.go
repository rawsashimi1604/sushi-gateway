package egress

import (
	"net/http"
)

type EgressService struct {
}

func NewEgressService() *EgressService {
	return &EgressService{}
}

func (s *EgressService) ForwardRequest(req *http.Request) ([]byte, int, *EgressError) {
	return nil, 0, nil
}
