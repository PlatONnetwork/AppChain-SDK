package module

import "encoding/json"

type ModuleGenesisConfig struct {
	Version     uint64 `json:"version"`
	CreateBlock uint64 `json:"createBlock"`
}

func GetVersionMapFromGenesis(modules map[string]json.RawMessage) (VersionMap, error) {
	vm := VersionMap{}
	for name, data := range modules {
		raw, err := data.MarshalJSON()
		if err != nil {
			return vm, err
		}

		var conf ModuleGenesisConfig
		if err := json.Unmarshal(raw, &conf); err != nil {
			return vm, err
		}
		vm[name]= conf.Version
	}
	return vm, nil
}
