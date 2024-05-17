package eip

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
)

const prefix = "eip712"

type EIP712 struct {
}

// bytes32 hashedName = keccak256(bytes(name));
// bytes32 hashedVersion = keccak256(bytes(version));
// bytes32 typeHash = keccak256(
//
//	"EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
//
// );
// _HASHED_NAME = hashedName;
// _HASHED_VERSION = hashedVersion;
// _CACHED_CHAIN_ID = block.chainid;
// _CACHED_DOMAIN_SEPARATOR = _buildDomainSeparator(typeHash, hashedName, hashedVersion);
// _CACHED_THIS = address(this);
// _TYPE_HASH = typeHash;
func (e *EIP712) Init(name, version string) {
	//hashedName := crypto.Keccak256Hash([]byte(name))
	//hashedVersion := crypto.Keccak256Hash([]byte(version))
	//typeHash := crypto.Keccak256Hash([]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"))
}

func (e *EIP712) DomainSeparatorV4() common.Hash {
	return common.ZeroHash

}

func (e *EIP712) BuildDomainSeparator(typeHash, nameHash, versionHash common.Hash) common.Hash {
	return common.ZeroHash

}

func (e *EIP712) hashTypedDataV4(structHash common.Hash) common.Hash {
	return common.ZeroHash
}
