package common

import (
	"crypto/ecdsa"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
)

var (
	UserPrivateKeys = map[common.Address]*ecdsa.PrivateKey{
		common.HexToAddress("0x8a0f3F8389F79a05Dd03027dE0Bcbb06cADB3F9c"): crypto.HexMustToECDSA("5856fb978afcb48fd023a0e6ba76eb90ed35de98ca4a6873d28d1c5b9a47f39e"),
		common.HexToAddress("0x1FD13cA28ccc423e0565fD3C9b64187ad91Ae7D6"): crypto.HexMustToECDSA("fa4ca9b1cc136fbaab4ab1bf55a3718a5d9ff2cc0b6ffbb098682e5d257b416b"),
		common.HexToAddress("0x27E1a2e05186c9e4C4A3A957F8748fA35EA19076"): crypto.HexMustToECDSA("94c67767ede78befe7352bd14a8338a124147db0f01c1a011401340b94832af1"),
		common.HexToAddress("0x6151aE08d3acD4c8194bB029633Ac7059066C75F"): crypto.HexMustToECDSA("bc22d5de77e9b50f0635158edd282cfe6a5c944f3dee6b4d68a38dcfd5e68784"),
		common.HexToAddress("0xE9565da74A5149e43475e62646dC4f47c2FcEa6f"): crypto.HexMustToECDSA("f0014181639f603f938647aa7d38663c3e867b2c25c1da7784d38752deaca67e"),
		common.HexToAddress("0xd1391ba0Be0eECA4031e952402b92e26417D5F38"): crypto.HexMustToECDSA("6eb85bfd34859d81db7f3d0a0c3f72c662eb1ea92de435dffc41497bbb0fbbcc"),
		common.HexToAddress("0xEdA947963D73F7Ca96c94843E2bd091d31Ec9F90"): crypto.HexMustToECDSA("0d1d0a522c1765b2fadbe86ba2bf45b3ab9a9122f448c5f59bcee76dfe3790e1"),
		common.HexToAddress("0xD03ae6Da0708D073f3f920Ea1AAEbfdc847971b3"): crypto.HexMustToECDSA("6fc120b574690e6019a0c91fef94ad6fb0c196264fb986497ce4b46e662c2251"),
		common.HexToAddress("0xD9bE8f83736b508514183363a1fec4c0Bd1E80Db"): crypto.HexMustToECDSA("b3b71d3e1c48e20b9b6d548a20a5ba4db5a7a3666d2e35ec911538be4dfd006d"),
		common.HexToAddress("0x0504d04808022aC69bBb37206FD97Edaef9268e9"): crypto.HexMustToECDSA("9eb4fda04c360a5b3ea046f2317006bb2f05802215ac75ae316b9dd7d03e2c15"),
	}
	UserAddrs = []common.Address{
		common.HexToAddress("0x8a0f3F8389F79a05Dd03027dE0Bcbb06cADB3F9c"),
		common.HexToAddress("0x1FD13cA28ccc423e0565fD3C9b64187ad91Ae7D6"),
		common.HexToAddress("0x27E1a2e05186c9e4C4A3A957F8748fA35EA19076"),
		common.HexToAddress("0x6151aE08d3acD4c8194bB029633Ac7059066C75F"),
		common.HexToAddress("0xE9565da74A5149e43475e62646dC4f47c2FcEa6f"),
		common.HexToAddress("0xd1391ba0Be0eECA4031e952402b92e26417D5F38"),
		common.HexToAddress("0xEdA947963D73F7Ca96c94843E2bd091d31Ec9F90"),
		common.HexToAddress("0xD03ae6Da0708D073f3f920Ea1AAEbfdc847971b3"),
		common.HexToAddress("0xD9bE8f83736b508514183363a1fec4c0Bd1E80Db"),
		common.HexToAddress("0x0504d04808022aC69bBb37206FD97Edaef9268e9"),
	}
)
