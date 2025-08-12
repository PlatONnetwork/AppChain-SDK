// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package customchildchainmanager

import (
	"errors"
	"math/big"
	"strings"

	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = platon.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GenesisValidator is an auto generated low-level Go binding around an user-defined struct.
type GenesisValidator struct {
	Validator    common.Address
	Owner        common.Address
	InitialStake *big.Int
}

// Validator is an auto generated low-level Go binding around an user-defined struct.
type Validator struct {
	Owner    common.Address
	PubKey   []byte
	BlsKey   []byte
	IsActive bool
}

// CustomChildChainManagerMetaData contains all meta data concerning the CustomChildChainManager contract.
var CustomChildChainManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableStaking\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeGenesis\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"genesisSet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structGenesisValidator[]\",\"components\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidator\",\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structValidator\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newStakeManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newStateSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newChildStakeHandler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newExitHelper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onAddStake\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onDelegate\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onL2StateReceive\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onStake\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commissionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validators\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawSlashedStake\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"GenesisFinalized\",\"inputs\":[{\"name\":\"validatorsCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakingEnabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorDeactivated\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRegistered\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"pubKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"blsKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x608080604052346100bf575f549060ff8260081c1661006d575060ff80821603610033575b60405161208390816100c48239f35b60ff90811916175f557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160ff8152a15f610024565b62461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608490fd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c9081631904bb2e146115095750806344efbf841461146e578063601a943d146110af578063715018a61461105257806385758cc714610fa05780638da5cb5b14610f775780639e19e9c114610f00578063a18b105414610df0578063cc2a9a5b14610b2e578063d11aca6214610a29578063e6ce413e14610910578063f2fde38b14610882578063f43cda8b146101495763fa52c7d8146100b8575f80fd5b34610146576020366003190112610146576001600160a01b0390604090826100de6115e6565b168152609f60205220908154166100f760018301611738565b9161013a60ff600361010b60028501611738565b930154169161012c6040519586958652608060208701526080860190611626565b908482036040860152611626565b90151560608301520390f35b80fd5b5034610146576060366003190112610146576101636115fc565b67ffffffffffffffff919060443583811161028e57610186903690600401611664565b609b54919290916001600160a01b0316330361082b57609a546001600160a01b039081169116036107cc578060201161028e578135937f8ca9a95e41b5eece253c93f5b31eed1253aed6b145d8a6e14d913fdf8e7322938503610292575081929350906060918101031261027f57816001600160a01b0361020960208401611612565b606554911692906001600160a01b0316803b1561028e5760408051633651bb1d60e01b81526001600160a01b0386166004820152920135602483015282908290604490829084905af180156102835761026b575b505061026890611efc565b80f35b610274906116e6565b61027f57815f61025d565b5080fd5b6040513d84823e3d90fd5b8280fd5b91937f117f1d6f44fd34ccb7a58f1261fa59e5c4bf68e2712d65f246a8805167a9334481036106b3575083019060808483031261028e57602084013590811161028e5783019080601f8301121561028e578135906102ef826117da565b926102fd6040519485611716565b82845260208401906020829460051b82010192831161052957602001905b82821061069b575050508151938390610333866117da565b956103416040519788611716565b80875261034d816117da565b601f1901366020890137855b8181106105405750506103756004926060606493013590611b4a565b609954609b5460405163fc9c8d3960e01b81529493909204926001600160a01b03918216926020928692918391165afa9081156105355786916104e9575b6103bd935061184a565b60018060a01b03609854169160018060a01b03609a5416916040519160a08301907f117f1d6f44fd34ccb7a58f1261fa59e5c4bf68e2712d65f246a8805167a9334460208501526004356040850152608060608501525180915260c083019190865b8181106104ca57505050601f1982820301608083015260208087519283815201960190855b8181106104b45750505061046281859603601f198101835282611716565b823b156104af5761048c928492836040518096819582946316f1983160e01b845260048401611a19565b03925af180156102835761049f57505080f35b6104a8906116e6565b6101465780f35b505050fd5b8251885260209788019790920191600101610444565b82516001600160a01b031684526020938401939092019160010161041f565b919290506020813d60201161052d575b8161050660209383611716565b8101031261052957516001600160a01b0381168103610529576103bd92916103b3565b8580fd5b3d91506104f9565b6040513d88823e3d90fd5b6065546001600160a01b03908116949061055a8389611b22565b51604051630213119b60e51b815291166004820152602081602481895afa908115610654578991610663575b50610598606491604087013590611b4a565b04946001600160a01b036105ac848a611b22565b5116813b1561065f57604051638028a6db60e01b81526001600160a01b03919091166004820152602481018790529089908290604490829084905af18015610654578692918a91610637575b5060019392610625929091506106206001600160a01b03610619868d611b22565b5116611efc565b611b71565b94610630828b611b22565b5201610359565b610643919293506116e6565b610650578490885f6105f8565b8780fd5b6040513d8b823e3d90fd5b8980fd5b90506020813d602011610693575b8161067e60209383611716565b8101031261068f5751610598610586565b5f80fd5b3d9150610671565b602080916106a884611612565b81520191019061031b565b7f58e580ca1cdbe518f27d857873b615e807a3c395584a93d75e80c921c991e50f9193949250145f1461076d5760808184938101031261076a576106f960208201611612565b61070560408301611612565b6065546001600160a01b0316803b1561076657604051635a601d9160e01b81526001600160a01b03938416600482015291909216602482015260609290920135604483015282908290606490829084905af180156102835761049f57505080f35b8480fd5b50fd5b60405162461bcd60e51b815260206004820152603160248201527f437573746f6d4368696c64436861696e4d616e616765723a20494e56414c49446044820152705f4d4554484f445f5349474e415455524560781b6064820152608490fd5b60405162461bcd60e51b815260206004820152603160248201527f437573746f6d4368696c64436861696e4d616e616765723a204f4e4c595f434860448201527024a6222fa9aa20a5a2afa420a7222622a960791b6064820152608490fd5b60405162461bcd60e51b815260206004820152602960248201527f437573746f6d4368696c64436861696e4d616e616765723a204f4e4c595f455860448201526824aa2fa422a62822a960b91b6064820152608490fd5b50346101465760203660031901126101465761089c6115e6565b6108a46117f2565b6001600160a01b038116156108bc5761026890611a3e565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b50346101465760403660031901126101465761092a6115e6565b6001600160a01b03908116808352609f602052604083206003015490919060ff16156109eb57610958611b7e565b156109ae578291816098541691609a541690604051907f7f629647b0cf8231fa5380e25f7c9bf0685fecbdc41360b93da5b447cef9ee7360208301526040820152602435606082015260608152610462816116ca565b60405163973d02cb60e01b815260206004820152601060248201526f5761697420666f722067656e6573697360801b6044820152606490fd5b0390fd5b60405163973d02cb60e01b81526020600482015260156024820152742222a0a1aa24ab20aa22a2102b20a624a220aa27a960591b6044820152606490fd5b5034610146578060031936011261014657610a426117f2565b609d5460ff81166003811015610b1a578015610ad557600214610a905760ff1916600217609d557fda917aeab736a19e4ba54207413dbe4f8c7d558fde2c14d7e66fc8f7186ea8a38180a180f35b60405162461bcd60e51b815260206004820152601b60248201527f47656e657369734c69623a20616c726561647920656e61626c656400000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601960248201527f47656e657369734c69623a206e6f742066696e616c697a6564000000000000006044820152606490fd5b634e487b7160e01b83526021600452602483fd5b50346101465760c036600319011261014657610b486115e6565b610b506115fc565b6001600160a01b036044358181169391929084900361068f576064359083821680920361068f5760843584811680910361068f5760a4359285841680940361068f5787549660ff8860081c161597888099610de3575b8015610dcc575b15610d705760ff1981166001178a5588610d5f575b50868616151580610d54575b80610d4b575b80610d42575b80610d39575b80610d30575b15610cd15760ff895460081c1615610c7857610c3a966001600160601b0360a01b941684606554161760655583609854161760985582609954161760995581609a541617609a55609b541617609b55611a3e565b610c415780f35b61ff001981541681557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160018152a180f35b60405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608490fd5b60405162461bcd60e51b815260206004820152603160248201527f437573746f6d4368696c64436861696e4d616e616765723a20494e56414c494460448201527017d253925512505312569157d253941555607a1b6064820152608490fd5b50841515610be6565b50821515610be0565b50811515610bda565b50801515610bd4565b508684161515610bce565b61ffff19166101011789555f610bc2565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b50303b158015610bad5750600160ff821614610bad565b50600160ff821610610ba6565b503461014657606036600319011261014657610e0a6115e6565b610e126115fc565b6001600160a01b03918216808452609f602052604084206003015490919060ff16156109eb57610e40611b7e565b156109ae5782609854169280609a541691604051937fc7ddcf4441a1bb01353b38db832023115117943d28ad05b882de4ad99e94b8fc6020860152604085015216606083015260443560808301526080825260a082019180831067ffffffffffffffff841117610eec57849383604052803b15610766578385809482946316f1983160e01b8452610ed881609f199360a48201611a19565b0301925af180156102835761049f57505080f35b634e487b7160e01b5f52604160045260245ffd5b5034610146578060031936011261014657610f19611a86565b60405190602080830190808452825180925280604085019301945b828110610f415784840385f35b855180516001600160a01b0390811686528184015116858401526040908101519085015294810194606090930192600101610f34565b50346101465780600319360112610146576033546040516001600160a01b039091168152602090f35b5034610146578060031936011261014657610fb96117f2565b609d5460ff81166003811015610b1a5761100d5760ff1916600117609d557f87f41ee3facb6317b1c2811e539539ac1693525b4460699b4245e8aac9f590cb6020611002611a86565b51604051908152a180f35b60405162461bcd60e51b815260206004820152601d60248201527f47656e657369734c69623a20616c72656164792066696e616c697a65640000006044820152606490fd5b503461014657806003193601126101465761106b6117f2565b603380546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b503461068f5760c036600319011261068f576110c96115e6565b906110d26115fc565b6044356084359167ffffffffffffffff9283811161068f576110f8903690600401611664565b9360a43590811161068f57611111903690600401611664565b6065546001600160a01b0396919391908716330361141b5760ff609d54166003811015611407576112ee579161115061115892611160959436916119c3565b9236916119c3565b908388611b93565b60ff609d541660038110156112da576111799015611ea8565b6001600160a01b0385165f908152609e602052604090205461127457609c5460018101809111611260576001600160a01b0386165f908152609e602052604090205582604051956111c9876116fa565b16855282602086019116815260408501918252609c546801000000000000000081101561124c578060016112009201609c55611ec3565b939093611238578060029495965116906001600160601b0360a01b918287541617865560018601925116908254161790555191015580f35b634e487b7160e01b85526004859052602485fd5b634e487b7160e01b85526041600452602485fd5b634e487b7160e01b85526011600452602485fd5b509192611292915060018060a01b03165f52609e60205260405f2090565b5461129e811515611ea8565b5f1981019081116112c65760026112b76112c192611ec3565b5001918254611b71565b905580f35b634e487b7160e01b83526011600452602483fd5b634e487b7160e01b85526021600452602485fd5b6112ff989291939496959798611b7e565b156109ae576113946113a8926113b49661133061131d36868a6119c3565b6113283684866119c3565b908c89611b93565b88609854169a8980609a54169b6040519b8c997f1bcc0f4c3fad314e585165815f94ecca9b96690a26d6417d7876448a9a867a6960208c01521660408a0152166060880152608087015260643560a087015260e060c08701526101008601916119f9565b90601f1994858584030160e08601526119f9565b03908101835282611716565b823b1561068f576113de925f92836040518096819582946316f1983160e01b845260048401611a19565b03925af180156113fc576113f0575080f35b6113fa91506116e6565b005b6040513d5f823e3d90fd5b634e487b7160e01b5f52602160045260245ffd5b60405162461bcd60e51b815260206004820152602560248201527f4368696c64436861696e4d616e616765723a204f4e4c595f5354414b455f4d416044820152642720a3a2a960d91b6064820152608490fd5b3461068f57602036600319011261068f576114876115e6565b61148f6117f2565b6099546040516370a0823160e01b8152306004820152916001600160a01b0390911690602083602481855afa9182156113fc575f926114d3575b6113fa935061184a565b91506020833d602011611501575b816114ee60209383611716565b8101031261068f576113fa9251916114c9565b3d91506114e1565b3461068f5760208060031936011261068f575f60606115266115e6565b93611530816116ca565b8281528184820152816040820152015260018060a01b038092165f52609f815260405f20906115d960405192611565846116ca565b8481541684526115c661157a60018301611738565b84860190815260ff600361159060028601611738565b946040890195865201541694606087019515158652604051978897828952511690870152516080604087015260a0860190611626565b9051848203601f19016060860152611626565b9051151560808301520390f35b600435906001600160a01b038216820361068f57565b602435906001600160a01b038216820361068f57565b35906001600160a01b038216820361068f57565b91908251928382525f5b848110611650575050825f602080949584010152601f8019910116010190565b602081830181015184830182015201611630565b9181601f8401121561068f5782359167ffffffffffffffff831161068f576020838186019501011161068f57565b90600182811c921680156116c0575b60208310146116ac57565b634e487b7160e01b5f52602260045260245ffd5b91607f16916116a1565b6080810190811067ffffffffffffffff821117610eec57604052565b67ffffffffffffffff8111610eec57604052565b6060810190811067ffffffffffffffff821117610eec57604052565b90601f8019910116810190811067ffffffffffffffff821117610eec57604052565b9060405191825f825461174a81611692565b908184526020946001916001811690815f146117b8575060011461177a575b50505061177892500383611716565b565b5f90815285812095935091905b8183106117a057505061177893508201015f8080611769565b85548884018501529485019487945091830191611787565b9250505061177894925060ff191682840152151560051b8201015f8080611769565b67ffffffffffffffff8111610eec5760051b60200190565b6033546001600160a01b0316330361180657565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b6040805163a9059cbb60e01b60208083019182526001600160a01b0395909516602483015260448083019690965294815292939092909161188c606483611716565b60018060a01b031683519184830183811067ffffffffffffffff821117610eec5785528583527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564838701525161190f935f91829182855af13d1561199f573d916118f5836119a7565b9261190287519485611716565b83523d5f8885013e611fb4565b805183811591821561197f575b5050905015611929575050565b60849250519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152fd5b838092935001031261068f57820151801515810361068f5780835f61191c565b606091611fb4565b67ffffffffffffffff8111610eec57601f01601f191660200190565b9291926119cf826119a7565b916119dd6040519384611716565b82948184528183011161068f578281602093845f960137010152565b908060209392818452848401375f828201840152601f01601f1916010190565b6001600160a01b039091168152604060208201819052611a3b92910190611626565b90565b603380546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b609c5490611a93826117da565b916040611aa36040519485611716565b81845283602080910191609c5f527faf85b9071dfafeac1409d3f1d19bafc9bc7c37974cde8df0ee6168f0086e539c905f935b858510611ae557505050505050565b6003846001928451611af6816116fa565b848060a01b03808854168252858801541683820152600287015486820152815201930194019391611ad6565b8051821015611b365760209160051b010190565b634e487b7160e01b5f52603260045260245ffd5b81810292918115918404141715611b5d57565b634e487b7160e01b5f52601160045260245ffd5b91908201809211611b5d57565b60ff609d541660038110156114075760021490565b90919260018060a01b0380921693845f52602093609f855260405f2093600385019160ff835416611e62571693846001600160601b0360a01b82541617815560019081810184519167ffffffffffffffff91828411610eec57611bf68154611692565b938a601f95868111611e1d575b50508a858211600114611db957908060029493925f91611dae575b505f19600383901b1c191690871b1790555b01918651918211610eec57611c458354611692565b818111611d6b575b5088908211600114611ce15792611cd19592827f98ef40a5780e42524ba3086c03b8176598633b835c5e7acc9cfb2b383b85726b9a99989693611cc3965f91611cd6575b505f19600383901b1c191690831b1790555b60ff19825416179055606060405196879687528601526060850190611626565b908382036040850152611626565b0390a2565b90508701515f611c91565b601f19821690835f52895f20915f5b818110611d56575094838193611cd19996937f98ef40a5780e42524ba3086c03b8176598633b835c5e7acc9cfb2b383b85726b9d9c9b9996611cc39910611d3e575b5050811b019055611ca3565b8901515f1960f88460031b161c191690555f80611d32565b89830151845592860192918b01918b01611cf0565b835f52895f208280850160051c8201928c8610611da5575b0160051c019085905b828110611d9a575050611c4d565b5f8155018590611d8c565b92508192611d83565b90508901515f611c1e565b5f8381528c8120601f198416959493899390928f5b8d898310611e07575050508260029710611def575b5050811b019055611c30565b8b01515f1960f88460031b161c191690555f80611de3565b83015184558b959093019291820191018f611dce565b835f5286825f209181850160051c8301938510611e59575b0160051c019087905b828110611e4e57508c9150611c03565b5f8155018790611e3e565b92508192611e35565b60405163973d02cb60e01b815260048101889052601b60248201527f56414c494441544f5220414c52454144592041435449564154454400000000006044820152606490fd5b15611eaf57565b634e487b7160e01b5f52600160045260245ffd5b609c54811015611b3657600390609c5f52027faf85b9071dfafeac1409d3f1d19bafc9bc7c37974cde8df0ee6168f0086e539c01905f90565b606554604051630213119b60e51b81526001600160a01b039283166004820181905292909160209183916024918391165afa9081156113fc575f91611f82575b5015611f455750565b805f52609f602052600360405f200160ff1981541690557f23d934bfe7f1275bc6fd70432159c9cc1c0075d069f89da6a40f43bfe7a94ed35f80a2565b90506020813d602011611fac575b81611f9d60209383611716565b8101031261068f57515f611f3c565b3d9150611f90565b919290156120165750815115611fc8575090565b3b15611fd15790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b8251909150156120295750805190602001fd5b60405162461bcd60e51b8152602060048201529081906109e790602483019061162656fea26469706673582212200d5ec210726a8ec1d28fd104e83e842b5e1da2f6e6a6aba5b30bc6970c58ef2564736f6c63430008160033",
}

// CustomChildChainManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use CustomChildChainManagerMetaData.ABI instead.
var CustomChildChainManagerABI = CustomChildChainManagerMetaData.ABI

// CustomChildChainManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CustomChildChainManagerMetaData.Bin instead.
var CustomChildChainManagerBin = CustomChildChainManagerMetaData.Bin

// DeployCustomChildChainManager deploys a new platon contract, binding an instance of CustomChildChainManager to it.
func DeployCustomChildChainManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CustomChildChainManager, error) {
	parsed, err := CustomChildChainManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CustomChildChainManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CustomChildChainManager{CustomChildChainManagerCaller: CustomChildChainManagerCaller{contract: contract}, CustomChildChainManagerTransactor: CustomChildChainManagerTransactor{contract: contract}, CustomChildChainManagerFilterer: CustomChildChainManagerFilterer{contract: contract}}, nil
}

// CustomChildChainManager is an auto generated Go binding around an platon contract.
type CustomChildChainManager struct {
	CustomChildChainManagerCaller     // Read-only binding to the contract
	CustomChildChainManagerTransactor // Write-only binding to the contract
	CustomChildChainManagerFilterer   // Log filterer for contract events
}

// CustomChildChainManagerCaller is an auto generated read-only Go binding around an platon contract.
type CustomChildChainManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustomChildChainManagerTransactor is an auto generated write-only Go binding around an platon contract.
type CustomChildChainManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustomChildChainManagerFilterer is an auto generated log filtering Go binding around an platon contract events.
type CustomChildChainManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustomChildChainManagerSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type CustomChildChainManagerSession struct {
	Contract     *CustomChildChainManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CustomChildChainManagerCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type CustomChildChainManagerCallerSession struct {
	Contract *CustomChildChainManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// CustomChildChainManagerTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type CustomChildChainManagerTransactorSession struct {
	Contract     *CustomChildChainManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// CustomChildChainManagerRaw is an auto generated low-level Go binding around an platon contract.
type CustomChildChainManagerRaw struct {
	Contract *CustomChildChainManager // Generic contract binding to access the raw methods on
}

// CustomChildChainManagerCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type CustomChildChainManagerCallerRaw struct {
	Contract *CustomChildChainManagerCaller // Generic read-only contract binding to access the raw methods on
}

// CustomChildChainManagerTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type CustomChildChainManagerTransactorRaw struct {
	Contract *CustomChildChainManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCustomChildChainManager creates a new instance of CustomChildChainManager, bound to a specific deployed contract.
func NewCustomChildChainManager(address common.Address, backend bind.ContractBackend) (*CustomChildChainManager, error) {
	contract, err := bindCustomChildChainManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CustomChildChainManager{CustomChildChainManagerCaller: CustomChildChainManagerCaller{contract: contract}, CustomChildChainManagerTransactor: CustomChildChainManagerTransactor{contract: contract}, CustomChildChainManagerFilterer: CustomChildChainManagerFilterer{contract: contract}}, nil
}

// NewCustomChildChainManagerCaller creates a new read-only instance of CustomChildChainManager, bound to a specific deployed contract.
func NewCustomChildChainManagerCaller(address common.Address, caller bind.ContractCaller) (*CustomChildChainManagerCaller, error) {
	contract, err := bindCustomChildChainManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CustomChildChainManagerCaller{contract: contract}, nil
}

// NewCustomChildChainManagerTransactor creates a new write-only instance of CustomChildChainManager, bound to a specific deployed contract.
func NewCustomChildChainManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*CustomChildChainManagerTransactor, error) {
	contract, err := bindCustomChildChainManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CustomChildChainManagerTransactor{contract: contract}, nil
}

// NewCustomChildChainManagerFilterer creates a new log filterer instance of CustomChildChainManager, bound to a specific deployed contract.
func NewCustomChildChainManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*CustomChildChainManagerFilterer, error) {
	contract, err := bindCustomChildChainManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CustomChildChainManagerFilterer{contract: contract}, nil
}

// bindCustomChildChainManager binds a generic wrapper to an already deployed contract.
func bindCustomChildChainManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CustomChildChainManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustomChildChainManager *CustomChildChainManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustomChildChainManager.Contract.CustomChildChainManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustomChildChainManager *CustomChildChainManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.CustomChildChainManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustomChildChainManager *CustomChildChainManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.CustomChildChainManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustomChildChainManager *CustomChildChainManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustomChildChainManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustomChildChainManager *CustomChildChainManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustomChildChainManager *CustomChildChainManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.contract.Transact(opts, method, params...)
}

// GenesisSet is a free data retrieval call binding the contract method 0x9e19e9c1.
//
// Solidity: function genesisSet() view returns((address,address,uint256)[])
func (_CustomChildChainManager *CustomChildChainManagerCaller) GenesisSet(opts *bind.CallOpts) ([]GenesisValidator, error) {
	var out []interface{}
	err := _CustomChildChainManager.contract.Call(opts, &out, "genesisSet")

	if err != nil {
		return *new([]GenesisValidator), err
	}

	out0 := *abi.ConvertType(out[0], new([]GenesisValidator)).(*[]GenesisValidator)

	return out0, err

}

// GenesisSet is a free data retrieval call binding the contract method 0x9e19e9c1.
//
// Solidity: function genesisSet() view returns((address,address,uint256)[])
func (_CustomChildChainManager *CustomChildChainManagerSession) GenesisSet() ([]GenesisValidator, error) {
	return _CustomChildChainManager.Contract.GenesisSet(&_CustomChildChainManager.CallOpts)
}

// GenesisSet is a free data retrieval call binding the contract method 0x9e19e9c1.
//
// Solidity: function genesisSet() view returns((address,address,uint256)[])
func (_CustomChildChainManager *CustomChildChainManagerCallerSession) GenesisSet() ([]GenesisValidator, error) {
	return _CustomChildChainManager.Contract.GenesisSet(&_CustomChildChainManager.CallOpts)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address _validator) view returns((address,bytes,bytes,bool))
func (_CustomChildChainManager *CustomChildChainManagerCaller) GetValidator(opts *bind.CallOpts, _validator common.Address) (Validator, error) {
	var out []interface{}
	err := _CustomChildChainManager.contract.Call(opts, &out, "getValidator", _validator)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address _validator) view returns((address,bytes,bytes,bool))
func (_CustomChildChainManager *CustomChildChainManagerSession) GetValidator(_validator common.Address) (Validator, error) {
	return _CustomChildChainManager.Contract.GetValidator(&_CustomChildChainManager.CallOpts, _validator)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address _validator) view returns((address,bytes,bytes,bool))
func (_CustomChildChainManager *CustomChildChainManagerCallerSession) GetValidator(_validator common.Address) (Validator, error) {
	return _CustomChildChainManager.Contract.GetValidator(&_CustomChildChainManager.CallOpts, _validator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CustomChildChainManager *CustomChildChainManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustomChildChainManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CustomChildChainManager *CustomChildChainManagerSession) Owner() (common.Address, error) {
	return _CustomChildChainManager.Contract.Owner(&_CustomChildChainManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CustomChildChainManager *CustomChildChainManagerCallerSession) Owner() (common.Address, error) {
	return _CustomChildChainManager.Contract.Owner(&_CustomChildChainManager.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(address owner, bytes pubKey, bytes blsKey, bool isActive)
func (_CustomChildChainManager *CustomChildChainManagerCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Owner    common.Address
	PubKey   []byte
	BlsKey   []byte
	IsActive bool
}, error) {
	var out []interface{}
	err := _CustomChildChainManager.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Owner    common.Address
		PubKey   []byte
		BlsKey   []byte
		IsActive bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PubKey = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.BlsKey = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(address owner, bytes pubKey, bytes blsKey, bool isActive)
func (_CustomChildChainManager *CustomChildChainManagerSession) Validators(arg0 common.Address) (struct {
	Owner    common.Address
	PubKey   []byte
	BlsKey   []byte
	IsActive bool
}, error) {
	return _CustomChildChainManager.Contract.Validators(&_CustomChildChainManager.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(address owner, bytes pubKey, bytes blsKey, bool isActive)
func (_CustomChildChainManager *CustomChildChainManagerCallerSession) Validators(arg0 common.Address) (struct {
	Owner    common.Address
	PubKey   []byte
	BlsKey   []byte
	IsActive bool
}, error) {
	return _CustomChildChainManager.Contract.Validators(&_CustomChildChainManager.CallOpts, arg0)
}

// EnableStaking is a paid mutator transaction binding the contract method 0xd11aca62.
//
// Solidity: function enableStaking() returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactor) EnableStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomChildChainManager.contract.Transact(opts, "enableStaking")
}

// EnableStaking is a paid mutator transaction binding the contract method 0xd11aca62.
//
// Solidity: function enableStaking() returns()
func (_CustomChildChainManager *CustomChildChainManagerSession) EnableStaking() (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.EnableStaking(&_CustomChildChainManager.TransactOpts)
}

// EnableStaking is a paid mutator transaction binding the contract method 0xd11aca62.
//
// Solidity: function enableStaking() returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactorSession) EnableStaking() (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.EnableStaking(&_CustomChildChainManager.TransactOpts)
}

// FinalizeGenesis is a paid mutator transaction binding the contract method 0x85758cc7.
//
// Solidity: function finalizeGenesis() returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactor) FinalizeGenesis(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomChildChainManager.contract.Transact(opts, "finalizeGenesis")
}

// FinalizeGenesis is a paid mutator transaction binding the contract method 0x85758cc7.
//
// Solidity: function finalizeGenesis() returns()
func (_CustomChildChainManager *CustomChildChainManagerSession) FinalizeGenesis() (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.FinalizeGenesis(&_CustomChildChainManager.TransactOpts)
}

// FinalizeGenesis is a paid mutator transaction binding the contract method 0x85758cc7.
//
// Solidity: function finalizeGenesis() returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactorSession) FinalizeGenesis() (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.FinalizeGenesis(&_CustomChildChainManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address newOwner, address newStakeManager, address newStateSender, address newToken, address newChildStakeHandler, address newExitHelper) returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactor) Initialize(opts *bind.TransactOpts, newOwner common.Address, newStakeManager common.Address, newStateSender common.Address, newToken common.Address, newChildStakeHandler common.Address, newExitHelper common.Address) (*types.Transaction, error) {
	return _CustomChildChainManager.contract.Transact(opts, "initialize", newOwner, newStakeManager, newStateSender, newToken, newChildStakeHandler, newExitHelper)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address newOwner, address newStakeManager, address newStateSender, address newToken, address newChildStakeHandler, address newExitHelper) returns()
func (_CustomChildChainManager *CustomChildChainManagerSession) Initialize(newOwner common.Address, newStakeManager common.Address, newStateSender common.Address, newToken common.Address, newChildStakeHandler common.Address, newExitHelper common.Address) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.Initialize(&_CustomChildChainManager.TransactOpts, newOwner, newStakeManager, newStateSender, newToken, newChildStakeHandler, newExitHelper)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address newOwner, address newStakeManager, address newStateSender, address newToken, address newChildStakeHandler, address newExitHelper) returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactorSession) Initialize(newOwner common.Address, newStakeManager common.Address, newStateSender common.Address, newToken common.Address, newChildStakeHandler common.Address, newExitHelper common.Address) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.Initialize(&_CustomChildChainManager.TransactOpts, newOwner, newStakeManager, newStateSender, newToken, newChildStakeHandler, newExitHelper)
}

// OnAddStake is a paid mutator transaction binding the contract method 0xe6ce413e.
//
// Solidity: function onAddStake(address validator, uint256 amount) returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactor) OnAddStake(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomChildChainManager.contract.Transact(opts, "onAddStake", validator, amount)
}

// OnAddStake is a paid mutator transaction binding the contract method 0xe6ce413e.
//
// Solidity: function onAddStake(address validator, uint256 amount) returns()
func (_CustomChildChainManager *CustomChildChainManagerSession) OnAddStake(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.OnAddStake(&_CustomChildChainManager.TransactOpts, validator, amount)
}

// OnAddStake is a paid mutator transaction binding the contract method 0xe6ce413e.
//
// Solidity: function onAddStake(address validator, uint256 amount) returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactorSession) OnAddStake(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.OnAddStake(&_CustomChildChainManager.TransactOpts, validator, amount)
}

// OnDelegate is a paid mutator transaction binding the contract method 0xa18b1054.
//
// Solidity: function onDelegate(address validator, address delegator, uint256 amount) returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactor) OnDelegate(opts *bind.TransactOpts, validator common.Address, delegator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomChildChainManager.contract.Transact(opts, "onDelegate", validator, delegator, amount)
}

// OnDelegate is a paid mutator transaction binding the contract method 0xa18b1054.
//
// Solidity: function onDelegate(address validator, address delegator, uint256 amount) returns()
func (_CustomChildChainManager *CustomChildChainManagerSession) OnDelegate(validator common.Address, delegator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.OnDelegate(&_CustomChildChainManager.TransactOpts, validator, delegator, amount)
}

// OnDelegate is a paid mutator transaction binding the contract method 0xa18b1054.
//
// Solidity: function onDelegate(address validator, address delegator, uint256 amount) returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactorSession) OnDelegate(validator common.Address, delegator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.OnDelegate(&_CustomChildChainManager.TransactOpts, validator, delegator, amount)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 id, address sender, bytes data) returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactor) OnL2StateReceive(opts *bind.TransactOpts, id *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _CustomChildChainManager.contract.Transact(opts, "onL2StateReceive", id, sender, data)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 id, address sender, bytes data) returns()
func (_CustomChildChainManager *CustomChildChainManagerSession) OnL2StateReceive(id *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.OnL2StateReceive(&_CustomChildChainManager.TransactOpts, id, sender, data)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 id, address sender, bytes data) returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactorSession) OnL2StateReceive(id *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.OnL2StateReceive(&_CustomChildChainManager.TransactOpts, id, sender, data)
}

// OnStake is a paid mutator transaction binding the contract method 0x601a943d.
//
// Solidity: function onStake(address validator, address owner, uint256 amount, uint256 commissionRate, bytes pubKey, bytes blsKey) returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactor) OnStake(opts *bind.TransactOpts, validator common.Address, owner common.Address, amount *big.Int, commissionRate *big.Int, pubKey []byte, blsKey []byte) (*types.Transaction, error) {
	return _CustomChildChainManager.contract.Transact(opts, "onStake", validator, owner, amount, commissionRate, pubKey, blsKey)
}

// OnStake is a paid mutator transaction binding the contract method 0x601a943d.
//
// Solidity: function onStake(address validator, address owner, uint256 amount, uint256 commissionRate, bytes pubKey, bytes blsKey) returns()
func (_CustomChildChainManager *CustomChildChainManagerSession) OnStake(validator common.Address, owner common.Address, amount *big.Int, commissionRate *big.Int, pubKey []byte, blsKey []byte) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.OnStake(&_CustomChildChainManager.TransactOpts, validator, owner, amount, commissionRate, pubKey, blsKey)
}

// OnStake is a paid mutator transaction binding the contract method 0x601a943d.
//
// Solidity: function onStake(address validator, address owner, uint256 amount, uint256 commissionRate, bytes pubKey, bytes blsKey) returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactorSession) OnStake(validator common.Address, owner common.Address, amount *big.Int, commissionRate *big.Int, pubKey []byte, blsKey []byte) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.OnStake(&_CustomChildChainManager.TransactOpts, validator, owner, amount, commissionRate, pubKey, blsKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomChildChainManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CustomChildChainManager *CustomChildChainManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.RenounceOwnership(&_CustomChildChainManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.RenounceOwnership(&_CustomChildChainManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CustomChildChainManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CustomChildChainManager *CustomChildChainManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.TransferOwnership(&_CustomChildChainManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.TransferOwnership(&_CustomChildChainManager.TransactOpts, newOwner)
}

// WithdrawSlashedStake is a paid mutator transaction binding the contract method 0x44efbf84.
//
// Solidity: function withdrawSlashedStake(address to) returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactor) WithdrawSlashedStake(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CustomChildChainManager.contract.Transact(opts, "withdrawSlashedStake", to)
}

// WithdrawSlashedStake is a paid mutator transaction binding the contract method 0x44efbf84.
//
// Solidity: function withdrawSlashedStake(address to) returns()
func (_CustomChildChainManager *CustomChildChainManagerSession) WithdrawSlashedStake(to common.Address) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.WithdrawSlashedStake(&_CustomChildChainManager.TransactOpts, to)
}

// WithdrawSlashedStake is a paid mutator transaction binding the contract method 0x44efbf84.
//
// Solidity: function withdrawSlashedStake(address to) returns()
func (_CustomChildChainManager *CustomChildChainManagerTransactorSession) WithdrawSlashedStake(to common.Address) (*types.Transaction, error) {
	return _CustomChildChainManager.Contract.WithdrawSlashedStake(&_CustomChildChainManager.TransactOpts, to)
}

// CustomChildChainManagerGenesisFinalizedIterator is returned from FilterGenesisFinalized and is used to iterate over the raw logs and unpacked data for GenesisFinalized events raised by the CustomChildChainManager contract.
type CustomChildChainManagerGenesisFinalizedIterator struct {
	Event *CustomChildChainManagerGenesisFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  platon.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CustomChildChainManagerGenesisFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomChildChainManagerGenesisFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CustomChildChainManagerGenesisFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CustomChildChainManagerGenesisFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustomChildChainManagerGenesisFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustomChildChainManagerGenesisFinalized represents a GenesisFinalized event raised by the CustomChildChainManager contract.
type CustomChildChainManagerGenesisFinalized struct {
	ValidatorsCount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterGenesisFinalized is a free log retrieval operation binding the contract event 0x87f41ee3facb6317b1c2811e539539ac1693525b4460699b4245e8aac9f590cb.
//
// Solidity: event GenesisFinalized(uint256 validatorsCount)
func (_CustomChildChainManager *CustomChildChainManagerFilterer) FilterGenesisFinalized(opts *bind.FilterOpts) (*CustomChildChainManagerGenesisFinalizedIterator, error) {

	logs, sub, err := _CustomChildChainManager.contract.FilterLogs(opts, "GenesisFinalized")
	if err != nil {
		return nil, err
	}
	return &CustomChildChainManagerGenesisFinalizedIterator{contract: _CustomChildChainManager.contract, event: "GenesisFinalized", logs: logs, sub: sub}, nil
}

// WatchGenesisFinalized is a free log subscription operation binding the contract event 0x87f41ee3facb6317b1c2811e539539ac1693525b4460699b4245e8aac9f590cb.
//
// Solidity: event GenesisFinalized(uint256 validatorsCount)
func (_CustomChildChainManager *CustomChildChainManagerFilterer) WatchGenesisFinalized(opts *bind.WatchOpts, sink chan<- *CustomChildChainManagerGenesisFinalized) (event.Subscription, error) {

	logs, sub, err := _CustomChildChainManager.contract.WatchLogs(opts, "GenesisFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustomChildChainManagerGenesisFinalized)
				if err := _CustomChildChainManager.contract.UnpackLog(event, "GenesisFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGenesisFinalized is a log parse operation binding the contract event 0x87f41ee3facb6317b1c2811e539539ac1693525b4460699b4245e8aac9f590cb.
//
// Solidity: event GenesisFinalized(uint256 validatorsCount)
func (_CustomChildChainManager *CustomChildChainManagerFilterer) ParseGenesisFinalized(log types.Log) (*CustomChildChainManagerGenesisFinalized, error) {
	event := new(CustomChildChainManagerGenesisFinalized)
	if err := _CustomChildChainManager.contract.UnpackLog(event, "GenesisFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustomChildChainManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CustomChildChainManager contract.
type CustomChildChainManagerInitializedIterator struct {
	Event *CustomChildChainManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  platon.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CustomChildChainManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomChildChainManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CustomChildChainManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CustomChildChainManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustomChildChainManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustomChildChainManagerInitialized represents a Initialized event raised by the CustomChildChainManager contract.
type CustomChildChainManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CustomChildChainManager *CustomChildChainManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*CustomChildChainManagerInitializedIterator, error) {

	logs, sub, err := _CustomChildChainManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CustomChildChainManagerInitializedIterator{contract: _CustomChildChainManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CustomChildChainManager *CustomChildChainManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CustomChildChainManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _CustomChildChainManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustomChildChainManagerInitialized)
				if err := _CustomChildChainManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CustomChildChainManager *CustomChildChainManagerFilterer) ParseInitialized(log types.Log) (*CustomChildChainManagerInitialized, error) {
	event := new(CustomChildChainManagerInitialized)
	if err := _CustomChildChainManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustomChildChainManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CustomChildChainManager contract.
type CustomChildChainManagerOwnershipTransferredIterator struct {
	Event *CustomChildChainManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  platon.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CustomChildChainManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomChildChainManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CustomChildChainManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CustomChildChainManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustomChildChainManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustomChildChainManagerOwnershipTransferred represents a OwnershipTransferred event raised by the CustomChildChainManager contract.
type CustomChildChainManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CustomChildChainManager *CustomChildChainManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CustomChildChainManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CustomChildChainManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CustomChildChainManagerOwnershipTransferredIterator{contract: _CustomChildChainManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CustomChildChainManager *CustomChildChainManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CustomChildChainManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CustomChildChainManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustomChildChainManagerOwnershipTransferred)
				if err := _CustomChildChainManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CustomChildChainManager *CustomChildChainManagerFilterer) ParseOwnershipTransferred(log types.Log) (*CustomChildChainManagerOwnershipTransferred, error) {
	event := new(CustomChildChainManagerOwnershipTransferred)
	if err := _CustomChildChainManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustomChildChainManagerStakingEnabledIterator is returned from FilterStakingEnabled and is used to iterate over the raw logs and unpacked data for StakingEnabled events raised by the CustomChildChainManager contract.
type CustomChildChainManagerStakingEnabledIterator struct {
	Event *CustomChildChainManagerStakingEnabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  platon.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CustomChildChainManagerStakingEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomChildChainManagerStakingEnabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CustomChildChainManagerStakingEnabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CustomChildChainManagerStakingEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustomChildChainManagerStakingEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustomChildChainManagerStakingEnabled represents a StakingEnabled event raised by the CustomChildChainManager contract.
type CustomChildChainManagerStakingEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingEnabled is a free log retrieval operation binding the contract event 0xda917aeab736a19e4ba54207413dbe4f8c7d558fde2c14d7e66fc8f7186ea8a3.
//
// Solidity: event StakingEnabled()
func (_CustomChildChainManager *CustomChildChainManagerFilterer) FilterStakingEnabled(opts *bind.FilterOpts) (*CustomChildChainManagerStakingEnabledIterator, error) {

	logs, sub, err := _CustomChildChainManager.contract.FilterLogs(opts, "StakingEnabled")
	if err != nil {
		return nil, err
	}
	return &CustomChildChainManagerStakingEnabledIterator{contract: _CustomChildChainManager.contract, event: "StakingEnabled", logs: logs, sub: sub}, nil
}

// WatchStakingEnabled is a free log subscription operation binding the contract event 0xda917aeab736a19e4ba54207413dbe4f8c7d558fde2c14d7e66fc8f7186ea8a3.
//
// Solidity: event StakingEnabled()
func (_CustomChildChainManager *CustomChildChainManagerFilterer) WatchStakingEnabled(opts *bind.WatchOpts, sink chan<- *CustomChildChainManagerStakingEnabled) (event.Subscription, error) {

	logs, sub, err := _CustomChildChainManager.contract.WatchLogs(opts, "StakingEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustomChildChainManagerStakingEnabled)
				if err := _CustomChildChainManager.contract.UnpackLog(event, "StakingEnabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakingEnabled is a log parse operation binding the contract event 0xda917aeab736a19e4ba54207413dbe4f8c7d558fde2c14d7e66fc8f7186ea8a3.
//
// Solidity: event StakingEnabled()
func (_CustomChildChainManager *CustomChildChainManagerFilterer) ParseStakingEnabled(log types.Log) (*CustomChildChainManagerStakingEnabled, error) {
	event := new(CustomChildChainManagerStakingEnabled)
	if err := _CustomChildChainManager.contract.UnpackLog(event, "StakingEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustomChildChainManagerValidatorDeactivatedIterator is returned from FilterValidatorDeactivated and is used to iterate over the raw logs and unpacked data for ValidatorDeactivated events raised by the CustomChildChainManager contract.
type CustomChildChainManagerValidatorDeactivatedIterator struct {
	Event *CustomChildChainManagerValidatorDeactivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  platon.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CustomChildChainManagerValidatorDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomChildChainManagerValidatorDeactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CustomChildChainManagerValidatorDeactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CustomChildChainManagerValidatorDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustomChildChainManagerValidatorDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustomChildChainManagerValidatorDeactivated represents a ValidatorDeactivated event raised by the CustomChildChainManager contract.
type CustomChildChainManagerValidatorDeactivated struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeactivated is a free log retrieval operation binding the contract event 0x23d934bfe7f1275bc6fd70432159c9cc1c0075d069f89da6a40f43bfe7a94ed3.
//
// Solidity: event ValidatorDeactivated(address indexed validator)
func (_CustomChildChainManager *CustomChildChainManagerFilterer) FilterValidatorDeactivated(opts *bind.FilterOpts, validator []common.Address) (*CustomChildChainManagerValidatorDeactivatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CustomChildChainManager.contract.FilterLogs(opts, "ValidatorDeactivated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &CustomChildChainManagerValidatorDeactivatedIterator{contract: _CustomChildChainManager.contract, event: "ValidatorDeactivated", logs: logs, sub: sub}, nil
}

// WatchValidatorDeactivated is a free log subscription operation binding the contract event 0x23d934bfe7f1275bc6fd70432159c9cc1c0075d069f89da6a40f43bfe7a94ed3.
//
// Solidity: event ValidatorDeactivated(address indexed validator)
func (_CustomChildChainManager *CustomChildChainManagerFilterer) WatchValidatorDeactivated(opts *bind.WatchOpts, sink chan<- *CustomChildChainManagerValidatorDeactivated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CustomChildChainManager.contract.WatchLogs(opts, "ValidatorDeactivated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustomChildChainManagerValidatorDeactivated)
				if err := _CustomChildChainManager.contract.UnpackLog(event, "ValidatorDeactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorDeactivated is a log parse operation binding the contract event 0x23d934bfe7f1275bc6fd70432159c9cc1c0075d069f89da6a40f43bfe7a94ed3.
//
// Solidity: event ValidatorDeactivated(address indexed validator)
func (_CustomChildChainManager *CustomChildChainManagerFilterer) ParseValidatorDeactivated(log types.Log) (*CustomChildChainManagerValidatorDeactivated, error) {
	event := new(CustomChildChainManagerValidatorDeactivated)
	if err := _CustomChildChainManager.contract.UnpackLog(event, "ValidatorDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustomChildChainManagerValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the CustomChildChainManager contract.
type CustomChildChainManagerValidatorRegisteredIterator struct {
	Event *CustomChildChainManagerValidatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  platon.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CustomChildChainManagerValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomChildChainManagerValidatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CustomChildChainManagerValidatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CustomChildChainManagerValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustomChildChainManagerValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustomChildChainManagerValidatorRegistered represents a ValidatorRegistered event raised by the CustomChildChainManager contract.
type CustomChildChainManagerValidatorRegistered struct {
	Validator common.Address
	Owner     common.Address
	PubKey    []byte
	BlsKey    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0x98ef40a5780e42524ba3086c03b8176598633b835c5e7acc9cfb2b383b85726b.
//
// Solidity: event ValidatorRegistered(address indexed validator, address owner, bytes pubKey, bytes blsKey)
func (_CustomChildChainManager *CustomChildChainManagerFilterer) FilterValidatorRegistered(opts *bind.FilterOpts, validator []common.Address) (*CustomChildChainManagerValidatorRegisteredIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CustomChildChainManager.contract.FilterLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return &CustomChildChainManagerValidatorRegisteredIterator{contract: _CustomChildChainManager.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0x98ef40a5780e42524ba3086c03b8176598633b835c5e7acc9cfb2b383b85726b.
//
// Solidity: event ValidatorRegistered(address indexed validator, address owner, bytes pubKey, bytes blsKey)
func (_CustomChildChainManager *CustomChildChainManagerFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *CustomChildChainManagerValidatorRegistered, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _CustomChildChainManager.contract.WatchLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustomChildChainManagerValidatorRegistered)
				if err := _CustomChildChainManager.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorRegistered is a log parse operation binding the contract event 0x98ef40a5780e42524ba3086c03b8176598633b835c5e7acc9cfb2b383b85726b.
//
// Solidity: event ValidatorRegistered(address indexed validator, address owner, bytes pubKey, bytes blsKey)
func (_CustomChildChainManager *CustomChildChainManagerFilterer) ParseValidatorRegistered(log types.Log) (*CustomChildChainManagerValidatorRegistered, error) {
	event := new(CustomChildChainManagerValidatorRegistered)
	if err := _CustomChildChainManager.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
