package contracts

import (
	"encoding/hex"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	vrftypes "github.com/PlatONnetwork/AppChain-SDK/x/vrf/types"
	vrfwrap "github.com/PlatONnetwork/AppChain-SDK/x/vrf/wrap"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/log"
)

// internal

func (c *VRFManager) SetStageModule(stage vrftypes.StageModuler) {
	c.stageModule = stage
}

func (c *VRFManager) SetStakeModule(stake vrftypes.StakeModuler) {
	c.stakeModule = stake
}

func (c *VRFManager) verifyNonceAndProof(validatorAddr basecommon.Address, blockNumber uint64, nonceAndProof []byte) error {
	previousNonce, err := vrfwrap.GetPreviousNonce(c.evm.StateDB, c.contract.Address(), blockNumber)
	if nil != err {
		log.Error("Failed to get previous vrf nonce", "blockNumber", blockNumber, "error", err)
		//return typesdk.NewRevertError("VRFManager: can not get previous vrf nonce")
		panic("VRFManager: can not get previous vrf nonce")
	}

	pubKey := c.stakeModule.GetValidatorECDSAPubKey(c.evm.StateDB, validatorAddr)

	if nil == pubKey {
		log.Error("Failed to get ecdsa pubkey of nonceAndProof provide validator", "validatorAddr", validatorAddr.Hex(), "blockNumber", blockNumber)
		return typesdk.NewRevertError("VRFManager: can not get ecdsa pubkey of nonceAndProof provide validator")
	}

	if err := vrfwrap.VerifyVrf(nonceAndProof, previousNonce, pubKey); nil != err {
		log.Error("Failed to verify vrf nonceAndProof", "nonceAndProof", hex.EncodeToString(nonceAndProof), "data", previousNonce.Hex(), "blockNumber", blockNumber, "error", err)
		return typesdk.NewRevertError("VRFManager: can not verify vrf nonceAndProof")
	}
	return nil
}

func (c *VRFManager) setNonceAndProof(blockNumber uint64, nonceAndProof []byte) {
	vrfwrap.StorageNonceAndProof(c.evm.StateDB, c.contract.Address(), blockNumber, nonceAndProof)
}
