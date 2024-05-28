package gov

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	abi2 "github.com/umbracle/ethgo/abi"
	"math/big"
)

type Proposal struct {
	ProposalId      *big.Int         `json:"proposalId"`
	Targets         []common.Address `json:"targets"`
	Values          []*big.Int       `json:"values"`
	Calldatas       []hexutil.Bytes  `json:"calldatas"`
	DescriptionHash common.Hash      `json:"descriptionHash"`
}
type RpcService struct {
	gov *Module
}

func NewRpcService(m *Module) *RpcService {
	return &RpcService{gov: m}
}

func (s *RpcService) ExecuteProposal(targets []common.Address, values []*big.Int, calldatas []hexutil.Bytes, descriptionHash common.Hash) error {
	data, _ := abi2.Encode([]interface{}{targets, values, calldatas, descriptionHash}, abi2.MustNewType("tuple(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash)"))
	proposalId := crypto.Keccak256Hash(data).Big()
	s.gov.addProposal(&Proposal{
		ProposalId:      proposalId,
		Targets:         targets,
		Values:          values,
		Calldatas:       calldatas,
		DescriptionHash: descriptionHash,
	})
	return nil
}

func (s *RpcService) PendingProposal() ([]*Proposal, error) {
	return s.gov.getProposals(), nil
}
