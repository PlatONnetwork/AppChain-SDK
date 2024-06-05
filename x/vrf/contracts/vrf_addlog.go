package contracts

import (
	"math/big"
)

func (c *VRFManager) addLogVRFNonceAddedEvent(block *big.Int, nonce []byte) error {
	c.EmitVRFNonceAddedEvent(block, nonce)
	return nil
}
