package contracts

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/rlp"

	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
)

const (
	READY uint8 = 0
	DONE  uint8 = 1
)

var (
	ownerKey            = []byte("owner")
	planKeyPrefix       = "plan/"
	planHeightKeyPrefix = "plan/height/"

	ErrPlanNotFound = errors.New("upgrade plan not found")
)

type NameArarry []string

func (u *IUpgradePlan) String() string {
	buf, _ := json.Marshal(u)
	return string(buf)
}

func (c *Upgrade) setOwner(newOwner common.Address) {
	c.stateDb.SetState(c.contract.Address(), ownerKey, newOwner.Bytes())
}

func (c *Upgrade) getOwner() common.Address {
	var owner common.Address
	value := c.stateDb.GetState(c.contract.Address(), ownerKey)
	if len(value) != 0 {
		owner.SetBytes(value)
	}
	return owner
}

func (c *Upgrade) onlyOwner() {
	oldOwner := c.getOwner()
	caller := c.contract.CallerAddress
	addr := c.contract.Address()

	isOwner := oldOwner == caller || caller == addr
	contracts.Require(isOwner, "Ownable: caller is not a owner")
}

func (c *Upgrade) getUpgradePlan(height uint64) ([]IUpgradePlan, error) {
	return c.getUpgradePlanByHeight(height)
}

func (c *Upgrade) addUpgradePlan(plan IUpgradePlan) (err error) {
	defer func() {
		if err != nil {
			log.Warn("failed to add upgrade plan", "module", "upgrade", "plan", plan.String(), "err", err)
		} else {
			log.Info("Add upgrade plan success", "module", "upgrade", "plan", plan.String())
		}
	}()
	oldPlan, err := c.getUpgradePlanByName(plan.Name)
	if err != nil && err != ErrPlanNotFound {
		return err
	}
	contracts.Require(oldPlan.Name != plan.Name, "the adding plan already exists")
	if err = c.setUpgradePlan(plan); err != nil {
		return err
	}
	return c.addUpgradePlanToHeightArray(plan.Height, plan.Name)
}

func (c *Upgrade) setUpgradePlanDone(height uint64) error {
	names, _ := c.getUpgradePlanNameListByHeight(height)
	for _, name := range names {
		plan, err := c.getUpgradePlanByName(name)
		if err != nil {
			return err
		}
		plan.Status = DONE
		if err = c.setUpgradePlan(plan); err != nil {
			return err
		}
	}
	return nil
}

func (c *Upgrade) getUpgradePlanByHeight(height uint64) (plans []IUpgradePlan, err error) {
	log.Info("Get upgrade plan", "module", "upgrade", "height", height, "plans", len(plans), "err", err)
	names, _ := c.getUpgradePlanNameListByHeight(height)
	for _, name := range names {
		plan, err := c.getUpgradePlanByName(name)
		if err != nil {
			return plans, err
		}
		plans = append(plans, plan)
	}
	return
}

func (c *Upgrade) setUpgradePlan(plan IUpgradePlan) error {
	val, err := rlp.EncodeToBytes(&plan)
	if err != nil {
		return err
	}
	c.stateDb.SetState(c.contract.Address(), encodePlanKey(plan.Name), val)
	return nil
}

func (c *Upgrade) getUpgradePlanByName(name string) (plan IUpgradePlan, err error) {
	val := c.stateDb.GetState(c.contract.Address(), encodePlanKey(name))
	if len(val) != 0 {
		if err = rlp.DecodeBytes(val, &plan); err != nil {
			return
		}
	} else {
		err = ErrPlanNotFound
	}
	return
}

func (c *Upgrade) addUpgradePlanToHeightArray(height uint64, name string) error {
	val := c.stateDb.GetState(c.contract.Address(), encodePlanHeightKey(height))
	var names NameArarry
	if len(val) != 0 {
		if err := rlp.DecodeBytes(val, &names); err != nil {
			return err
		}
	}
	names = append(names, name)
	val, err := rlp.EncodeToBytes(&names)
	if err != nil {
		return err
	}
	c.stateDb.SetState(c.contract.Address(), encodePlanHeightKey(height), val)
	return nil
}

func (c *Upgrade) getUpgradePlanNameListByHeight(height uint64) (names NameArarry, err error) {
	val := c.stateDb.GetState(c.contract.Address(), encodePlanHeightKey(height))
	if len(val) != 0 {
		if err = rlp.DecodeBytes(val, &names); err != nil {
			return
		}
	}
	return
}

func encodePlanKey(name string) []byte {
	return []byte(fmt.Sprintf("%s%s", planKeyPrefix, name))
}

func encodePlanHeightKey(height uint64) []byte {
	return []byte(fmt.Sprintf("%s%d", planHeightKeyPrefix, height))
}
