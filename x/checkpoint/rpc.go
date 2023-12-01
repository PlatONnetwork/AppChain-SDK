package checkpoint

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/types"
)

type RpcService struct {
	checkpoint *Module
}

func NewRpcService(m *Module) *RpcService {
	return &RpcService{checkpoint: m}
}

func (s *RpcService) GenerateExitProof(exitID uint64) (types.Proof, error) {
	return s.checkpoint.generateExitProof(exitID)
}
