# 自定义RPC接口

`RpcModule`接口为开发者开发模块提供了自定义RPC接口的能力。

## RpcModule

!!! info "APIs"

    `APIs()[]rpc.API`

    返回值： 返回模块实现的RPC列表。

```go title="x/checkpoint/module.go"
func (m *Module) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: types.ModuleName,
			Version:   "1.0",
			Service:   NewRpcService(m),
			Public:    true,
		},
	}
}
```

```go title="x/checkpoint/rpc.go"
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
```
