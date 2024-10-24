package model

type LoadBalancingAlgorithm string

const (
	RoundRobin LoadBalancingAlgorithm = "round_robin"
	Weighted   LoadBalancingAlgorithm = "weighted"
	IPHash     LoadBalancingAlgorithm = "ip_hash"
)

func (alg LoadBalancingAlgorithm) IsValid() bool {
	switch alg {
	case RoundRobin, Weighted, IPHash:
		return true
	default:
		return false
	}
}
