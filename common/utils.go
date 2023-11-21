package common

import (
	"os"

	"github.com/PlatONnetwork/PlatON-Go/accounts/keystore"
)

func DecryptKey(ksFile, pwdFile string) (*keystore.Key, error) {
	json, err := os.ReadFile(ksFile)
	if err != nil {
		return nil, err
	}
	passphrase, err := os.ReadFile(pwdFile)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(json, string(passphrase))
	if err != nil {
		return nil, err
	}
	return key, nil
}
