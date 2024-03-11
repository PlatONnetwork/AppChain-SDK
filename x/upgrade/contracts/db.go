package contracts

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/rlp"

	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
)

const (
	READY uint8 = 0
	DONE  uint8 = 1
)

var (
	ownerKey             = []byte("owner")
	planKeyPrefix        = "plan/"
	planHeightKeyPrefix  = "plan/height/"
	moduleValidNumberKey = []byte("module/valid/number")
	moduleVersionKey     = []byte("module/version")

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
	// TODO: limit plan.Info length
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
	if err := c.addUpgradePlanToHeightArray(plan.Height, plan.Name); err != nil {
		return err
	}

	vn, err := c.GetModuleValidNumberMap()
	if err != nil {
		return err
	}

	for _, mod := range plan.Modules {
		if _, ok := vn[mod.ModuleName]; !ok {
			vn[mod.ModuleName] = math.MaxUint64
			log.Info("Initialize module valid number", "module", mod.ModuleName, "validNumber", vn[mod.ModuleName])
		}
	}
	return c.SetModuleValidNumberMap(vn)
}

func (c *Upgrade) setUpgradePlanDone(height uint64) error {
	vn, err := c.GetModuleValidNumberMap()
	if err != nil {
		log.Warn("failed to get module valid number map", "module", "upgrade", "height", height)
		return err
	}
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

		for _, mod := range plan.Modules {
			if _, ok := vn[mod.ModuleName]; ok {
				if vn[mod.ModuleName] == math.MaxUint64 {
					vn[mod.ModuleName] = height
					log.Info("Set module valid number", "module", mod.ModuleName, "validNumber", vn[mod.ModuleName])
				}
			}
		}
	}
	return c.SetModuleValidNumberMap(vn)
}

func (c *Upgrade) getUpgradePlanByHeight(height uint64) (plans []IUpgradePlan, err error) {
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
	log.Info("Set upgrade plan", "plan", plan.String())
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

type ModuleValidNumberList []module.ModuleValidNumber

func (c *Upgrade) SetModuleValidNumberMap(vn module.ValidNumberMap) error {
	c.onlyOwner()
	contracts.Require(len(vn) > 0, "empty module valid number map")

	l := vn.AsSliceSorted()
	val, err := rlp.EncodeToBytes(&l)
	if err != nil {
		return err
	}
	c.stateDb.SetState(c.contract.Address(), moduleValidNumberKey, val)
	return nil
}

func (c *Upgrade) GetModuleValidNumberMap() (module.ValidNumberMap, error) {
	vn := make(module.ValidNumberMap, 0)
	val := c.stateDb.GetState(c.contract.Address(), moduleValidNumberKey)
	if len(val) != 0 {
		var l ModuleValidNumberList
		if err := rlp.DecodeBytes(val, &l); err != nil {
			return vn, err
		}
		for _, mvn := range l {
			vn[mvn.Name] = mvn.ValidNumber
		}
		return vn, nil
	}
	return vn, errors.New("empty module valid number map")
}

func (c *Upgrade) SetModuleVersionMap(vm module.VersionMap) error {
	c.onlyOwner()
	contracts.Require(len(vm) > 0, "empty module version map")

	l := vm.AsSliceSorted()
	val, err := rlp.EncodeToBytes(&l)
	if err != nil {
		return err
	}
	c.stateDb.SetState(c.contract.Address(), moduleVersionKey, val)
	return nil
}

func (c *Upgrade) GetModuleVersionMap() (module.VersionMap, error) {
	vm := make(module.VersionMap, 0)
	val := c.stateDb.GetState(c.contract.Address(), moduleVersionKey)
	if len(val) != 0 {
		var l module.ModuleVersionList
		if err := rlp.DecodeBytes(val, &l); err != nil {
			return vm, err
		}

		for _, mvm := range l {
			vm[mvm.Name] = mvm.Version
		}
		return vm, nil
	}
	return vm, errors.New("empty version map")
}

func encodePlanKey(name string) []byte {
	return []byte(fmt.Sprintf("%s%s", planKeyPrefix, name))
}

func encodePlanHeightKey(height uint64) []byte {
	return []byte(fmt.Sprintf("%s%d", planHeightKeyPrefix, height))
}
