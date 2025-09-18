// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakemanager

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

// ValidatorStake is an auto generated low-level Go binding around an user-defined struct.
type ValidatorStake struct {
	Owner          common.Address
	StakeAmount    *big.Int
	CommissionRate *big.Int
	PubKey         []byte
	BlsKey         []byte
}

// StakeManagerMetaData contains all meta data concerning the StakeManager contract.
var StakeManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStake\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateFor\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegationOf\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_registryManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minDelegate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minDelegate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registryManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseDelegationOf\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseStakeOf\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashStakeOf\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeFor\",\"inputs\":[{\"name\":\"stake\",\"type\":\"tuple\",\"internalType\":\"structValidatorStake\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commissionRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeOf\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalDelegation\",\"inputs\":[],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalStake\",\"inputs\":[],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorDelegationOf\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawDelegation\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawStake\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawableDelegation\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawableStake\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DelegationAdded\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegationRemoved\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegationWithdrawn\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeAdded\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeRemoved\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeWithdrawn\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorStakeSlashed\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608080604052346100bf575f549060ff8260081c1661006d575060ff80821603610033575b6040516121b990816100c48239f35b60ff90811916175f557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160ff8152a15f610024565b62461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608490fd5b5f80fdfe60806040526004361015610011575f80fd5b5f803560e01c80630c1e8bf7146117fa5780630c63109e146117d25780633651bb1d1461167e578063375b3c0a146116615780633d9a6dcb146113e957806342623360146113b157806345255c0514611394578063512b91b614610e955780635a601d9114610b975780636374299e1461092657806374bbe1c8146106355780637a1ac61e146104ef5780638028a6db146102215780638b0e9f3f14610203578063d5364bbf146101ca578063e3c3ae58146101ac578063ef86e74c14610173578063f7b9a9a2146101305763f837123e146100eb575f80fd5b3461012d57602036600319011261012d576040906001600160a01b0361010f611815565b168152600560205220335f52602052602060405f2054604051908152f35b80fd5b503461012d57602036600319011261012d576040906001600160a01b03610155611815565b168152600860205220335f52602052602060405f2054604051908152f35b503461012d57602036600319011261012d576020906040906001600160a01b0361019b611815565b168152600483522054604051908152f35b503461012d578060031936011261012d576020600254604051908152f35b503461012d57602036600319011261012d576020906040906001600160a01b036101f2611815565b168152600683522054604051908152f35b503461012d578060031936011261012d576020600154604051908152f35b503461012d57604036600319011261012d5761023b611815565b603b54604051637c8211ff60e11b808252306004830152928492602092916001600160a01b0391602490813590841686868481845afa95861561042e5788966104bb575b508683916040519283809263199d7a1760e11b9a8b835260048301525afa801561042e5785908990610483575b6102b992501633146118f6565b80948484169889895260038852604089205480931161047a575b6102dd8386611dbe565b85603b54169160405191825230600483015288828681865afa91821561046f578a92610439575b50908489926040519485938492835260048301525afa801561042e578593899283926103ee575b5060405180978193630c825d9760e11b83528d6004840152165afa9384156103e3578589957fa018dcbc822f59fb0d0c3e7a86c8e4259b9676cdea9e5fc26279b9c4c5d86eef9589956103ad948c7f2fd10b18da60ce9090915fbb72edd4d6550168ad1915bd58038d803b9faba2d79d946103be575b50503392511690611b34565b604051908152a2604051908152a280f35b6103db9294503d8091833e6103d381836118c0565b810190611a9d565b915f8e6103a1565b6040513d89823e3d90fd5b92509350508681813d8311610427575b61040881836118c0565b8101031261042357849261041c89926118e2565b905f61032b565b8780fd5b503d6103fe565b6040513d8a823e3d90fd5b8980929b508193503d8311610468575b61045381836118c0565b8101031261046457518a9884610304565b5f80fd5b503d610449565b6040513d8c823e3d90fd5b955081956102d3565b50508681813d83116104b4575b61049a81836118c0565b8101031261042357846104af6102b9926118e2565b6102ac565b503d610490565b975094508587813d81116104e8575b6104d481836118c0565b81010312610464579551889690948661027f565b503d6104ca565b503461012d57606036600319011261012d57610509611815565b81549060ff8260081c161591828093610628575b8015610611575b156105b55760ff1981166001178455826105a4575b5060018060a01b03166bffffffffffffffffffffffff60a01b603b541617603b55602435603c55604435603d5561056d5780f35b61ff001981541681557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160018152a180f35b61ffff19166101011783555f610539565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b50303b1580156105245750600160ff821614610524565b50600160ff82161061051d565b503461012d576106443661182b565b9190603d5483106108d357603b54604051637c8211ff60e11b8152306004820152926020916001600160a01b039190602490831684878381845afa9687156108995789976108a4575b508482916040519283809263ea78803f60e01b82528b60048301525afa80156108995788908a9061085d575b6106c99250309086339116611ec8565b8282169586895260048552604089205488810180911161084a57878a526004865260408a20556005855260408920335f52855260405f205488810180911161084a57878a526005865260408a20335f52865260405f205560025488810180911161084a5789939291869160025585603b5416926040518094819363199d7a1760e11b835260048301525afa801561083f5784918491610807575b5016803b1561080357604051632862c41560e21b81526001600160a01b03928316600482015291861660248301526044820188905282908290606490829084905af180156107f8576107e0575b50507f52467f14b857734001c77e6f125dac41b45798837c9fc9adfe3a5b394c77a0e9926040519586521693a380f35b6107e990611860565b6107f457855f6107b0565b8580fd5b6040513d84823e3d90fd5b8280fd5b809250868092503d8311610838575b61082081836118c0565b810103126108035761083284916118e2565b5f610763565b503d610816565b6040513d85823e3d90fd5b634e487b7160e01b8a526011600452828afd5b50508481813d8311610892575b61087481836118c0565b8101031261088e57876108896106c9926118e2565b6106b9565b8880fd5b503d61086a565b6040513d8b823e3d90fd5b9096508481813d83116108cc575b6108bc81836118c0565b810103126104645751958461068d565b503d6108b2565b60405162461bcd60e51b815260206004820152602560248201527f5374616b654d616e616765723a20494e56414c49445f44454c45474154455f416044820152641353d5539560da1b6064820152608490fd5b503461012d57604036600319011261012d57610940611815565b60243560018060a01b038083169283855260209160038352604086205415610b525790818692603b5416604051637c8211ff60e11b81523060048201528581602481855afa918215610b135786918693610b1e575b50906024916040519283809263ea78803f60e01b82528660048301525afa8015610b135787908690610ad7575b6109d29250309085339116611ec8565b6109dc8684611f1b565b8482603b54169160246040518094819363199d7a1760e11b835260048301525afa908115610acc578491610a97575b5016803b1561080357604051637367209f60e11b81526001600160a01b039290921660048301526024820185905282908290604490829084905af180156107f857610a7f575b50507f7c717985ac273e663b7f3050f5b15a4388ff6ed952338954f650e2093e13937f91604051908152a280f35b610a8890611860565b610a9357835f610a51565b8380fd5b90508481813d8311610ac5575b610aae81836118c0565b81010312610a9357610abf906118e2565b5f610a0b565b503d610aa4565b6040513d86823e3d90fd5b50508581813d8311610b0c575b610aee81836118c0565b81010312610b085786610b036109d2926118e2565b6109c2565b8480fd5b503d610ae4565b6040513d87823e3d90fd5b9550915084813d8311610b4b575b610b3681836118c0565b81010312610464579251879385906024610995565b503d610b2c565b60405162461bcd60e51b815260048101849052601f60248201527f5374616b654d616e616765723a20494e56414c49445f56414c494441544f52006044820152606490fd5b503461012d57610ba63661182b565b603b54604051637c8211ff60e11b815230600482015292936001600160a01b0393602093928516908481602481855afa801561042e5785918991610e66575b5060246040518094819363199d7a1760e11b835260048301525afa80156103e35785918891610e1f575b5091610c2182610d87941633146118f6565b169384875260048452604087205490610c96604051610c3f816118a4565b603c8152857f72656d6f76652064656c65676174696f6e20616d6f756e7420657863656564739485898401527f2076616c696461746f722064656c65676174696f6e20616d6f756e74000000006040840152611fc0565b8689526004865260408920556005855260408820961695865f528452610d2e60405f20548460405191610cc883611888565b604983527f72656d6f76652064656c65676174696f6e20616d6f756e74206f662064656c65888401527f6761746f7220657863656564732064656c656761746f722064656c65676174696040840152681bdb88185b5bdd5b9d60ba1b6060840152611fc0565b8588526005855260408820875f52855260405f20558260025460405192610d54846118a4565b60388452868401527f20746f74616c2064656c65676174696f6e20616d6f756e7400000000000000006040840152611fc0565b600255828552600782526040852054818101809111610e0b578386526007835260408620556008825260408520845f52825260405f2054818101809111610e0b57907fbf340c6e47f6acc1fa5fcad9ef75c1e4bd8d91e7313667c3c9859f230fc7f88392918487526008835260408720865f52835260405f2055604051908152a380f35b634e487b7160e01b86526011600452602486fd5b809250858092503d8311610e5f575b610e3881836118c0565b81010312610e5b57610d8791610c2186610e5281946118e2565b92945050610c0f565b8680fd5b503d610e2e565b82819392503d8311610e8e575b610e7d81836118c0565b81010312610464578490515f610be5565b503d610e73565b50346104645760203660031901126104645760043567ffffffffffffffff81116104645760a0600319823603011261046457610ee66001600160a01b03610ede60048401611951565b163314611965565b603c546024820135106113445760646044820135116112f1576040610f1160648301836004016119bb565b9050036112a0576030610f2a60848301836004016119bb565b90500361125157603b54604051637c8211ff60e11b815230600482015291906001600160a01b0316602083602481845afa928315611164575f9361121c575b5060206024916040519283809263ea78803f60e01b82528760048301525afa8015611164575f906111e1575b610fb29150602483013590309033906001600160a01b0316611ec8565b610fc260648201826004016119bb565b610fcb816119ee565b91610fd960405193846118c0565b818352368282011161046457815f92602092838601378301015260408151036111a95780516020909101206001600160a01b03169161101c602483013584611f1b565b603b5460405163199d7a1760e11b81526004810192909252602090829060249082906001600160a01b03165afa908115611164575f9161116f575b5061106482600401611951565b9061107560648401846004016119bb565b92909161108860848601866004016119bb565b9092906001600160a01b0383163b15610464578761110588955f97936110f3899560449b6040519c8d9b8c9a8b9963601a943d60e01b8b5260048b015260018060a01b031660248a01526024810135828a01520135606488015260c0608488015260c4870191611a0a565b8481036003190160a486015291611a0a565b03926001600160a01b03165af1801561116457611150575b5060207f7c717985ac273e663b7f3050f5b15a4388ff6ed952338954f650e2093e13937f9160246040519101358152a280f35b61115b919350611860565b5f91602061111d565b6040513d5f823e3d90fd5b90506020813d6020116111a1575b8161118a602093836118c0565b810103126104645761119b906118e2565b5f611057565b3d915061117d565b60405162461bcd60e51b815260206004820152601060248201526f6e6f74206563647361207075624b657960801b6044820152606490fd5b506020813d602011611214575b816111fb602093836118c0565b810103126104645761120f610fb2916118e2565b610f95565b3d91506111ee565b9092506020813d602011611249575b81611238602093836118c0565b810103126104645751916020610f69565b3d915061122b565b60405162461bcd60e51b815260206004820152602160248201527f5374616b654d616e616765723a20494e56414c4944415f424c535f5055424b456044820152605960f81b6064820152608490fd5b60405162461bcd60e51b815260206004820152602360248201527f5374616b654d616e616765723a20494e56414c4944415f45434453415f5055426044820152624b455960e81b6064820152608490fd5b60405162461bcd60e51b815260206004820152602560248201527f5374616b654d616e616765723a20494e56414c49445f434f4d4d495353494f4e6044820152645f5241544560d81b6064820152608490fd5b60405162461bcd60e51b815260206004820152602260248201527f5374616b654d616e616765723a20494e56414c49445f5354414b455f414d4f55604482015261139560f21b6064820152608490fd5b34610464575f366003190112610464576020603d54604051908152f35b34610464576020366003190112610464576001600160a01b036113d2611815565b165f526003602052602060405f2054604051908152f35b34610464576113f73661182b565b919060018060a01b0380921691825f526020906007825261152760405f20546114a260405161142581611888565b605b8152887f72656d6f766520776974686472617761626c652064656c65676174696f6e20619384888401527f6d6f756e74206578636565647320746f74616c20616d6f756e74206f6620766160408401527f6c696461746f722077697468647261772064656c65676174696f6e00000000006060840152611fc0565b865f526007855260405f20556008845260405f20335f5284528660405f2054604051926114ce84611888565b605b8452868401527f6d6f756e74206578636565647320746f74616c20616d6f756e74206f6620646560408401527f6c656761746f722077697468647261772064656c65676174696f6e00000000006060840152611fc0565b845f526008835260405f20335f52835260405f205580603b541692604051637c8211ff60e11b81523060048201528381602481885afa80156111645784915f91611632575b5060246040518097819363ea78803f60e01b835260048301525afa80156111645786945f916115d0575b507faa0001b0e2e76b9a1b925257fefa58b756aebc52d2d7c8a85ea5beacc77a21009482846115c59316611f5b565b6040519586521693a3005b809550848092503d831161162b575b6115e981836118c0565b81010312610464576115c58682846116217faa0001b0e2e76b9a1b925257fefa58b756aebc52d2d7c8a85ea5beacc77a2100986118e2565b9350505094611596565b503d6115df565b82819392503d831161165a575b61164981836118c0565b81010312610464578390518861156c565b503d61163f565b34610464575f366003190112610464576020603c54604051908152f35b3461046457604036600319011261046457611697611815565b603b54604051637c8211ff60e11b815230600482015260248035936001600160a01b039384169360209390928490829081885afa80156111645784915f916117a3575b5060246040518097819363199d7a1760e11b835260048301525afa938415611164575f94611747575b50611732827fa018dcbc822f59fb0d0c3e7a86c8e4259b9676cdea9e5fc26279b9c4c5d86eef951633146118f6565b61173c8582611dbe565b6040519485521692a2005b93508284813d831161179c575b61175e81836118c0565b8101031261046457611732826117947fa018dcbc822f59fb0d0c3e7a86c8e4259b9676cdea9e5fc26279b9c4c5d86eef966118e2565b955050611703565b503d611754565b82819392503d83116117cb575b6117ba81836118c0565b8101031261046457839051876116da565b503d6117b0565b34610464575f36600319011261046457603b546040516001600160a01b039091168152602090f35b346104645761181361180b3661182b565b913390611b34565b005b600435906001600160a01b038216820361046457565b6060906003190112610464576001600160a01b0390600435828116810361046457916024359081168103610464579060443590565b67ffffffffffffffff811161187457604052565b634e487b7160e01b5f52604160045260245ffd5b6080810190811067ffffffffffffffff82111761187457604052565b6060810190811067ffffffffffffffff82111761187457604052565b90601f8019910116810190811067ffffffffffffffff82111761187457604052565b51906001600160a01b038216820361046457565b156118fd57565b60405162461bcd60e51b815260206004820152602660248201527f5374616b654d616e616765723a204f4e4c595f4348494c445f434841494e5f4d60448201526520a720a3a2a960d11b6064820152608490fd5b356001600160a01b03811681036104645790565b1561196c57565b60405162461bcd60e51b815260206004820152602160248201527f5374616b654d616e616765723a20494e56414c49445f5354414b455f4f574e456044820152602960f91b6064820152608490fd5b903590601e1981360301821215610464570180359067ffffffffffffffff82116104645760200191813603831361046457565b67ffffffffffffffff811161187457601f01601f191660200190565b908060209392818452848401375f828201840152601f01601f1916010190565b5f5b838110611a3b5750505f910152565b8181015183820152602001611a2c565b81601f82011215610464578051611a61816119ee565b92611a6f60405194856118c0565b8184526020828401011161046457611a8d9160208085019101611a2a565b90565b5190811515820361046457565b60208183031261046457805167ffffffffffffffff918282116104645701916080838203126104645760405192608084018481108482111761187457604052611ae5816118e2565b845260208101518381116104645782611aff918301611a4b565b6020850152604081015192831161046457611b21606092611b2c948301611a4b565b604085015201611a90565b606082015290565b603b5460408051637c8211ff60e11b81523060048201529695946001600160a01b0392831694602094919291858a6024818a5afa998a15611d46575f9a611d8f575b50858a60248551809a819363199d7a1760e11b835260048301525afa968715611d46575f97611d50575b505f856024819386519485938492630c825d9760e11b8452169b8c6004840152165afa908115611d465785611be4939281925f91611d2c575b505116911614611965565b845f5260068452611c4a815f205487835191611bff836118a4565b603583527f72656d6f766520776974686472617761626c65207374616b6520616d6f756e74888401527408195e18d959591cc81d1bdd185b08185b5bdd5b9d605a1b85840152611fc0565b855f5260068552815f20558383603b54169860248351809b819363ea78803f60e01b835260048301525afa8015611d22578697985f97969791611cc0575b507fb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3958385611cb79316611f5b565b519586521693a3565b809650858092503d8311611d1b575b611cd981836118c0565b8101031261046457611cb7878385611d117fb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3996118e2565b9350505095611c88565b503d611ccf565b50513d5f823e3d90fd5b611d4091503d805f833e6103d381836118c0565b5f611bd9565b83513d5f823e3d90fd5b9096508581813d8311611d88575b611d6881836118c0565b81010312610464575f856024611d7e82946118e2565b9993505050611ba0565b503d611d5e565b9099508581813d8311611db7575b611da781836118c0565b810103126104645751985f611b76565b503d611d9d565b60018060a01b0316805f5260209060038252604091611e2f835f205485855191611de7836118a4565b603283527f72656d6f7665207374616b6520616d6f756e7420657863656564732076616c69858401527119185d1bdc881cdd185ad948185b5bdd5b9d60721b87840152611fc0565b825f5260038252835f2055611e9160015485855191611e4d836118a4565b602e83527f72656d6f7665207374616b6520616d6f756e74206578636565647320746f7461858401526d1b081cdd185ad948185b5bdd5b9d60921b87840152611fc0565b600155815f5260068152825f2054938401809411611eb4576006915f52525f2055565b634e487b7160e01b5f52601160045260245ffd5b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064820192909252611f1991611f1482608481015b03601f1981018452836118c0565b611fee565b565b6001600160a01b03165f8181526003602052604090205482810191908210611eb4575f52600360205260405f2055600154908101809111611eb457600155565b60405163a9059cbb60e01b60208201526001600160a01b0390921660248301526044820192909252611f1991611f148260648101611f06565b60409160208252611fb48151809281602086015260208686019101611a2a565b601f01601f1916010190565b9091818311611fce57500390565b60405162461bcd60e51b8152908190611fea9060048301611f94565b0390fd5b60408051908101916001600160a01b031667ffffffffffffffff8311828410176118745761207b926040525f806020958685527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656487860152868151910182855af13d1561210a573d91612060836119ee565b9261206e60405194856118c0565b83523d5f8785013e61210e565b80518281159182156120eb575b50509050156120945750565b6084906040519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152fd5b838092935001031261046457816121029101611a90565b80825f612088565b6060915b919290156121705750815115612122575090565b3b1561212b5790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015611fce5750805190602001fdfea2646970667358221220dd859a585fecac0d3694e091befe65e6a15e4257a3f1eaa70dbe3d439854099764736f6c63430008160033",
}

// StakeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeManagerMetaData.ABI instead.
var StakeManagerABI = StakeManagerMetaData.ABI

// StakeManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakeManagerMetaData.Bin instead.
var StakeManagerBin = StakeManagerMetaData.Bin

// DeployStakeManager deploys a new platon contract, binding an instance of StakeManager to it.
func DeployStakeManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StakeManager, error) {
	parsed, err := StakeManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakeManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakeManager{StakeManagerCaller: StakeManagerCaller{contract: contract}, StakeManagerTransactor: StakeManagerTransactor{contract: contract}, StakeManagerFilterer: StakeManagerFilterer{contract: contract}}, nil
}

// StakeManager is an auto generated Go binding around an platon contract.
type StakeManager struct {
	StakeManagerCaller     // Read-only binding to the contract
	StakeManagerTransactor // Write-only binding to the contract
	StakeManagerFilterer   // Log filterer for contract events
}

// StakeManagerCaller is an auto generated read-only Go binding around an platon contract.
type StakeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerTransactor is an auto generated write-only Go binding around an platon contract.
type StakeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerFilterer is an auto generated log filtering Go binding around an platon contract events.
type StakeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeManagerSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type StakeManagerSession struct {
	Contract     *StakeManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeManagerCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type StakeManagerCallerSession struct {
	Contract *StakeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StakeManagerTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type StakeManagerTransactorSession struct {
	Contract     *StakeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakeManagerRaw is an auto generated low-level Go binding around an platon contract.
type StakeManagerRaw struct {
	Contract *StakeManager // Generic contract binding to access the raw methods on
}

// StakeManagerCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type StakeManagerCallerRaw struct {
	Contract *StakeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakeManagerTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type StakeManagerTransactorRaw struct {
	Contract *StakeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeManager creates a new instance of StakeManager, bound to a specific deployed contract.
func NewStakeManager(address common.Address, backend bind.ContractBackend) (*StakeManager, error) {
	contract, err := bindStakeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeManager{StakeManagerCaller: StakeManagerCaller{contract: contract}, StakeManagerTransactor: StakeManagerTransactor{contract: contract}, StakeManagerFilterer: StakeManagerFilterer{contract: contract}}, nil
}

// NewStakeManagerCaller creates a new read-only instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerCaller(address common.Address, caller bind.ContractCaller) (*StakeManagerCaller, error) {
	contract, err := bindStakeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeManagerCaller{contract: contract}, nil
}

// NewStakeManagerTransactor creates a new write-only instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeManagerTransactor, error) {
	contract, err := bindStakeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeManagerTransactor{contract: contract}, nil
}

// NewStakeManagerFilterer creates a new log filterer instance of StakeManager, bound to a specific deployed contract.
func NewStakeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeManagerFilterer, error) {
	contract, err := bindStakeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeManagerFilterer{contract: contract}, nil
}

// bindStakeManager binds a generic wrapper to an already deployed contract.
func bindStakeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeManager *StakeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeManager.Contract.StakeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeManager *StakeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeManager *StakeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeManager *StakeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeManager *StakeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeManager *StakeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeManager.Contract.contract.Transact(opts, method, params...)
}

// DelegationOf is a free data retrieval call binding the contract method 0xf837123e.
//
// Solidity: function delegationOf(address validator) view returns(uint256 amount)
func (_StakeManager *StakeManagerCaller) DelegationOf(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "delegationOf", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationOf is a free data retrieval call binding the contract method 0xf837123e.
//
// Solidity: function delegationOf(address validator) view returns(uint256 amount)
func (_StakeManager *StakeManagerSession) DelegationOf(validator common.Address) (*big.Int, error) {
	return _StakeManager.Contract.DelegationOf(&_StakeManager.CallOpts, validator)
}

// DelegationOf is a free data retrieval call binding the contract method 0xf837123e.
//
// Solidity: function delegationOf(address validator) view returns(uint256 amount)
func (_StakeManager *StakeManagerCallerSession) DelegationOf(validator common.Address) (*big.Int, error) {
	return _StakeManager.Contract.DelegationOf(&_StakeManager.CallOpts, validator)
}

// MinDelegate is a free data retrieval call binding the contract method 0x45255c05.
//
// Solidity: function minDelegate() view returns(uint256)
func (_StakeManager *StakeManagerCaller) MinDelegate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "minDelegate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegate is a free data retrieval call binding the contract method 0x45255c05.
//
// Solidity: function minDelegate() view returns(uint256)
func (_StakeManager *StakeManagerSession) MinDelegate() (*big.Int, error) {
	return _StakeManager.Contract.MinDelegate(&_StakeManager.CallOpts)
}

// MinDelegate is a free data retrieval call binding the contract method 0x45255c05.
//
// Solidity: function minDelegate() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) MinDelegate() (*big.Int, error) {
	return _StakeManager.Contract.MinDelegate(&_StakeManager.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_StakeManager *StakeManagerCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_StakeManager *StakeManagerSession) MinStake() (*big.Int, error) {
	return _StakeManager.Contract.MinStake(&_StakeManager.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_StakeManager *StakeManagerCallerSession) MinStake() (*big.Int, error) {
	return _StakeManager.Contract.MinStake(&_StakeManager.CallOpts)
}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_StakeManager *StakeManagerCaller) RegistryManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "registryManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_StakeManager *StakeManagerSession) RegistryManager() (common.Address, error) {
	return _StakeManager.Contract.RegistryManager(&_StakeManager.CallOpts)
}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_StakeManager *StakeManagerCallerSession) RegistryManager() (common.Address, error) {
	return _StakeManager.Contract.RegistryManager(&_StakeManager.CallOpts)
}

// StakeOf is a free data retrieval call binding the contract method 0x42623360.
//
// Solidity: function stakeOf(address validator) view returns(uint256 amount)
func (_StakeManager *StakeManagerCaller) StakeOf(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "stakeOf", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeOf is a free data retrieval call binding the contract method 0x42623360.
//
// Solidity: function stakeOf(address validator) view returns(uint256 amount)
func (_StakeManager *StakeManagerSession) StakeOf(validator common.Address) (*big.Int, error) {
	return _StakeManager.Contract.StakeOf(&_StakeManager.CallOpts, validator)
}

// StakeOf is a free data retrieval call binding the contract method 0x42623360.
//
// Solidity: function stakeOf(address validator) view returns(uint256 amount)
func (_StakeManager *StakeManagerCallerSession) StakeOf(validator common.Address) (*big.Int, error) {
	return _StakeManager.Contract.StakeOf(&_StakeManager.CallOpts, validator)
}

// TotalDelegation is a free data retrieval call binding the contract method 0xe3c3ae58.
//
// Solidity: function totalDelegation() view returns(uint256 amount)
func (_StakeManager *StakeManagerCaller) TotalDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "totalDelegation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDelegation is a free data retrieval call binding the contract method 0xe3c3ae58.
//
// Solidity: function totalDelegation() view returns(uint256 amount)
func (_StakeManager *StakeManagerSession) TotalDelegation() (*big.Int, error) {
	return _StakeManager.Contract.TotalDelegation(&_StakeManager.CallOpts)
}

// TotalDelegation is a free data retrieval call binding the contract method 0xe3c3ae58.
//
// Solidity: function totalDelegation() view returns(uint256 amount)
func (_StakeManager *StakeManagerCallerSession) TotalDelegation() (*big.Int, error) {
	return _StakeManager.Contract.TotalDelegation(&_StakeManager.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256 amount)
func (_StakeManager *StakeManagerCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "totalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256 amount)
func (_StakeManager *StakeManagerSession) TotalStake() (*big.Int, error) {
	return _StakeManager.Contract.TotalStake(&_StakeManager.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256 amount)
func (_StakeManager *StakeManagerCallerSession) TotalStake() (*big.Int, error) {
	return _StakeManager.Contract.TotalStake(&_StakeManager.CallOpts)
}

// ValidatorDelegationOf is a free data retrieval call binding the contract method 0xef86e74c.
//
// Solidity: function validatorDelegationOf(address validator) view returns(uint256 amount)
func (_StakeManager *StakeManagerCaller) ValidatorDelegationOf(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "validatorDelegationOf", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorDelegationOf is a free data retrieval call binding the contract method 0xef86e74c.
//
// Solidity: function validatorDelegationOf(address validator) view returns(uint256 amount)
func (_StakeManager *StakeManagerSession) ValidatorDelegationOf(validator common.Address) (*big.Int, error) {
	return _StakeManager.Contract.ValidatorDelegationOf(&_StakeManager.CallOpts, validator)
}

// ValidatorDelegationOf is a free data retrieval call binding the contract method 0xef86e74c.
//
// Solidity: function validatorDelegationOf(address validator) view returns(uint256 amount)
func (_StakeManager *StakeManagerCallerSession) ValidatorDelegationOf(validator common.Address) (*big.Int, error) {
	return _StakeManager.Contract.ValidatorDelegationOf(&_StakeManager.CallOpts, validator)
}

// WithdrawableDelegation is a free data retrieval call binding the contract method 0xf7b9a9a2.
//
// Solidity: function withdrawableDelegation(address validator) view returns(uint256 amount)
func (_StakeManager *StakeManagerCaller) WithdrawableDelegation(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "withdrawableDelegation", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableDelegation is a free data retrieval call binding the contract method 0xf7b9a9a2.
//
// Solidity: function withdrawableDelegation(address validator) view returns(uint256 amount)
func (_StakeManager *StakeManagerSession) WithdrawableDelegation(validator common.Address) (*big.Int, error) {
	return _StakeManager.Contract.WithdrawableDelegation(&_StakeManager.CallOpts, validator)
}

// WithdrawableDelegation is a free data retrieval call binding the contract method 0xf7b9a9a2.
//
// Solidity: function withdrawableDelegation(address validator) view returns(uint256 amount)
func (_StakeManager *StakeManagerCallerSession) WithdrawableDelegation(validator common.Address) (*big.Int, error) {
	return _StakeManager.Contract.WithdrawableDelegation(&_StakeManager.CallOpts, validator)
}

// WithdrawableStake is a free data retrieval call binding the contract method 0xd5364bbf.
//
// Solidity: function withdrawableStake(address validator) view returns(uint256 amount)
func (_StakeManager *StakeManagerCaller) WithdrawableStake(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeManager.contract.Call(opts, &out, "withdrawableStake", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableStake is a free data retrieval call binding the contract method 0xd5364bbf.
//
// Solidity: function withdrawableStake(address validator) view returns(uint256 amount)
func (_StakeManager *StakeManagerSession) WithdrawableStake(validator common.Address) (*big.Int, error) {
	return _StakeManager.Contract.WithdrawableStake(&_StakeManager.CallOpts, validator)
}

// WithdrawableStake is a free data retrieval call binding the contract method 0xd5364bbf.
//
// Solidity: function withdrawableStake(address validator) view returns(uint256 amount)
func (_StakeManager *StakeManagerCallerSession) WithdrawableStake(validator common.Address) (*big.Int, error) {
	return _StakeManager.Contract.WithdrawableStake(&_StakeManager.CallOpts, validator)
}

// AddStake is a paid mutator transaction binding the contract method 0x6374299e.
//
// Solidity: function addStake(address validator, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactor) AddStake(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "addStake", validator, amount)
}

// AddStake is a paid mutator transaction binding the contract method 0x6374299e.
//
// Solidity: function addStake(address validator, uint256 amount) returns()
func (_StakeManager *StakeManagerSession) AddStake(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.AddStake(&_StakeManager.TransactOpts, validator, amount)
}

// AddStake is a paid mutator transaction binding the contract method 0x6374299e.
//
// Solidity: function addStake(address validator, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactorSession) AddStake(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.AddStake(&_StakeManager.TransactOpts, validator, amount)
}

// DelegateFor is a paid mutator transaction binding the contract method 0x74bbe1c8.
//
// Solidity: function delegateFor(address validator, address delegator, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactor) DelegateFor(opts *bind.TransactOpts, validator common.Address, delegator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "delegateFor", validator, delegator, amount)
}

// DelegateFor is a paid mutator transaction binding the contract method 0x74bbe1c8.
//
// Solidity: function delegateFor(address validator, address delegator, uint256 amount) returns()
func (_StakeManager *StakeManagerSession) DelegateFor(validator common.Address, delegator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.DelegateFor(&_StakeManager.TransactOpts, validator, delegator, amount)
}

// DelegateFor is a paid mutator transaction binding the contract method 0x74bbe1c8.
//
// Solidity: function delegateFor(address validator, address delegator, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactorSession) DelegateFor(validator common.Address, delegator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.DelegateFor(&_StakeManager.TransactOpts, validator, delegator, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _registryManager, uint256 _minStake, uint256 _minDelegate) returns()
func (_StakeManager *StakeManagerTransactor) Initialize(opts *bind.TransactOpts, _registryManager common.Address, _minStake *big.Int, _minDelegate *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "initialize", _registryManager, _minStake, _minDelegate)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _registryManager, uint256 _minStake, uint256 _minDelegate) returns()
func (_StakeManager *StakeManagerSession) Initialize(_registryManager common.Address, _minStake *big.Int, _minDelegate *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Initialize(&_StakeManager.TransactOpts, _registryManager, _minStake, _minDelegate)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _registryManager, uint256 _minStake, uint256 _minDelegate) returns()
func (_StakeManager *StakeManagerTransactorSession) Initialize(_registryManager common.Address, _minStake *big.Int, _minDelegate *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.Initialize(&_StakeManager.TransactOpts, _registryManager, _minStake, _minDelegate)
}

// ReleaseDelegationOf is a paid mutator transaction binding the contract method 0x5a601d91.
//
// Solidity: function releaseDelegationOf(address validator, address delegator, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactor) ReleaseDelegationOf(opts *bind.TransactOpts, validator common.Address, delegator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "releaseDelegationOf", validator, delegator, amount)
}

// ReleaseDelegationOf is a paid mutator transaction binding the contract method 0x5a601d91.
//
// Solidity: function releaseDelegationOf(address validator, address delegator, uint256 amount) returns()
func (_StakeManager *StakeManagerSession) ReleaseDelegationOf(validator common.Address, delegator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ReleaseDelegationOf(&_StakeManager.TransactOpts, validator, delegator, amount)
}

// ReleaseDelegationOf is a paid mutator transaction binding the contract method 0x5a601d91.
//
// Solidity: function releaseDelegationOf(address validator, address delegator, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactorSession) ReleaseDelegationOf(validator common.Address, delegator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ReleaseDelegationOf(&_StakeManager.TransactOpts, validator, delegator, amount)
}

// ReleaseStakeOf is a paid mutator transaction binding the contract method 0x3651bb1d.
//
// Solidity: function releaseStakeOf(address validator, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactor) ReleaseStakeOf(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "releaseStakeOf", validator, amount)
}

// ReleaseStakeOf is a paid mutator transaction binding the contract method 0x3651bb1d.
//
// Solidity: function releaseStakeOf(address validator, uint256 amount) returns()
func (_StakeManager *StakeManagerSession) ReleaseStakeOf(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ReleaseStakeOf(&_StakeManager.TransactOpts, validator, amount)
}

// ReleaseStakeOf is a paid mutator transaction binding the contract method 0x3651bb1d.
//
// Solidity: function releaseStakeOf(address validator, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactorSession) ReleaseStakeOf(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.ReleaseStakeOf(&_StakeManager.TransactOpts, validator, amount)
}

// SlashStakeOf is a paid mutator transaction binding the contract method 0x8028a6db.
//
// Solidity: function slashStakeOf(address validator, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactor) SlashStakeOf(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "slashStakeOf", validator, amount)
}

// SlashStakeOf is a paid mutator transaction binding the contract method 0x8028a6db.
//
// Solidity: function slashStakeOf(address validator, uint256 amount) returns()
func (_StakeManager *StakeManagerSession) SlashStakeOf(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SlashStakeOf(&_StakeManager.TransactOpts, validator, amount)
}

// SlashStakeOf is a paid mutator transaction binding the contract method 0x8028a6db.
//
// Solidity: function slashStakeOf(address validator, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactorSession) SlashStakeOf(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.SlashStakeOf(&_StakeManager.TransactOpts, validator, amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x512b91b6.
//
// Solidity: function stakeFor((address,uint256,uint256,bytes,bytes) stake) returns()
func (_StakeManager *StakeManagerTransactor) StakeFor(opts *bind.TransactOpts, stake ValidatorStake) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "stakeFor", stake)
}

// StakeFor is a paid mutator transaction binding the contract method 0x512b91b6.
//
// Solidity: function stakeFor((address,uint256,uint256,bytes,bytes) stake) returns()
func (_StakeManager *StakeManagerSession) StakeFor(stake ValidatorStake) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeFor(&_StakeManager.TransactOpts, stake)
}

// StakeFor is a paid mutator transaction binding the contract method 0x512b91b6.
//
// Solidity: function stakeFor((address,uint256,uint256,bytes,bytes) stake) returns()
func (_StakeManager *StakeManagerTransactorSession) StakeFor(stake ValidatorStake) (*types.Transaction, error) {
	return _StakeManager.Contract.StakeFor(&_StakeManager.TransactOpts, stake)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x3d9a6dcb.
//
// Solidity: function withdrawDelegation(address validator, address to, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactor) WithdrawDelegation(opts *bind.TransactOpts, validator common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdrawDelegation", validator, to, amount)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x3d9a6dcb.
//
// Solidity: function withdrawDelegation(address validator, address to, uint256 amount) returns()
func (_StakeManager *StakeManagerSession) WithdrawDelegation(validator common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawDelegation(&_StakeManager.TransactOpts, validator, to, amount)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0x3d9a6dcb.
//
// Solidity: function withdrawDelegation(address validator, address to, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactorSession) WithdrawDelegation(validator common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawDelegation(&_StakeManager.TransactOpts, validator, to, amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x0c1e8bf7.
//
// Solidity: function withdrawStake(address validator, address to, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactor) WithdrawStake(opts *bind.TransactOpts, validator common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.contract.Transact(opts, "withdrawStake", validator, to, amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x0c1e8bf7.
//
// Solidity: function withdrawStake(address validator, address to, uint256 amount) returns()
func (_StakeManager *StakeManagerSession) WithdrawStake(validator common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawStake(&_StakeManager.TransactOpts, validator, to, amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x0c1e8bf7.
//
// Solidity: function withdrawStake(address validator, address to, uint256 amount) returns()
func (_StakeManager *StakeManagerTransactorSession) WithdrawStake(validator common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakeManager.Contract.WithdrawStake(&_StakeManager.TransactOpts, validator, to, amount)
}

// StakeManagerDelegationAddedIterator is returned from FilterDelegationAdded and is used to iterate over the raw logs and unpacked data for DelegationAdded events raised by the StakeManager contract.
type StakeManagerDelegationAddedIterator struct {
	Event *StakeManagerDelegationAdded // Event containing the contract specifics and raw log

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
func (it *StakeManagerDelegationAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerDelegationAdded)
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
		it.Event = new(StakeManagerDelegationAdded)
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
func (it *StakeManagerDelegationAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerDelegationAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerDelegationAdded represents a DelegationAdded event raised by the StakeManager contract.
type StakeManagerDelegationAdded struct {
	Validator common.Address
	Delegator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegationAdded is a free log retrieval operation binding the contract event 0x52467f14b857734001c77e6f125dac41b45798837c9fc9adfe3a5b394c77a0e9.
//
// Solidity: event DelegationAdded(address indexed validator, address indexed delegator, uint256 amount)
func (_StakeManager *StakeManagerFilterer) FilterDelegationAdded(opts *bind.FilterOpts, validator []common.Address, delegator []common.Address) (*StakeManagerDelegationAddedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "DelegationAdded", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerDelegationAddedIterator{contract: _StakeManager.contract, event: "DelegationAdded", logs: logs, sub: sub}, nil
}

// WatchDelegationAdded is a free log subscription operation binding the contract event 0x52467f14b857734001c77e6f125dac41b45798837c9fc9adfe3a5b394c77a0e9.
//
// Solidity: event DelegationAdded(address indexed validator, address indexed delegator, uint256 amount)
func (_StakeManager *StakeManagerFilterer) WatchDelegationAdded(opts *bind.WatchOpts, sink chan<- *StakeManagerDelegationAdded, validator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "DelegationAdded", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerDelegationAdded)
				if err := _StakeManager.contract.UnpackLog(event, "DelegationAdded", log); err != nil {
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

// ParseDelegationAdded is a log parse operation binding the contract event 0x52467f14b857734001c77e6f125dac41b45798837c9fc9adfe3a5b394c77a0e9.
//
// Solidity: event DelegationAdded(address indexed validator, address indexed delegator, uint256 amount)
func (_StakeManager *StakeManagerFilterer) ParseDelegationAdded(log types.Log) (*StakeManagerDelegationAdded, error) {
	event := new(StakeManagerDelegationAdded)
	if err := _StakeManager.contract.UnpackLog(event, "DelegationAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerDelegationRemovedIterator is returned from FilterDelegationRemoved and is used to iterate over the raw logs and unpacked data for DelegationRemoved events raised by the StakeManager contract.
type StakeManagerDelegationRemovedIterator struct {
	Event *StakeManagerDelegationRemoved // Event containing the contract specifics and raw log

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
func (it *StakeManagerDelegationRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerDelegationRemoved)
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
		it.Event = new(StakeManagerDelegationRemoved)
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
func (it *StakeManagerDelegationRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerDelegationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerDelegationRemoved represents a DelegationRemoved event raised by the StakeManager contract.
type StakeManagerDelegationRemoved struct {
	Validator common.Address
	Delegator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegationRemoved is a free log retrieval operation binding the contract event 0xbf340c6e47f6acc1fa5fcad9ef75c1e4bd8d91e7313667c3c9859f230fc7f883.
//
// Solidity: event DelegationRemoved(address indexed validator, address indexed delegator, uint256 amount)
func (_StakeManager *StakeManagerFilterer) FilterDelegationRemoved(opts *bind.FilterOpts, validator []common.Address, delegator []common.Address) (*StakeManagerDelegationRemovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "DelegationRemoved", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerDelegationRemovedIterator{contract: _StakeManager.contract, event: "DelegationRemoved", logs: logs, sub: sub}, nil
}

// WatchDelegationRemoved is a free log subscription operation binding the contract event 0xbf340c6e47f6acc1fa5fcad9ef75c1e4bd8d91e7313667c3c9859f230fc7f883.
//
// Solidity: event DelegationRemoved(address indexed validator, address indexed delegator, uint256 amount)
func (_StakeManager *StakeManagerFilterer) WatchDelegationRemoved(opts *bind.WatchOpts, sink chan<- *StakeManagerDelegationRemoved, validator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "DelegationRemoved", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerDelegationRemoved)
				if err := _StakeManager.contract.UnpackLog(event, "DelegationRemoved", log); err != nil {
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

// ParseDelegationRemoved is a log parse operation binding the contract event 0xbf340c6e47f6acc1fa5fcad9ef75c1e4bd8d91e7313667c3c9859f230fc7f883.
//
// Solidity: event DelegationRemoved(address indexed validator, address indexed delegator, uint256 amount)
func (_StakeManager *StakeManagerFilterer) ParseDelegationRemoved(log types.Log) (*StakeManagerDelegationRemoved, error) {
	event := new(StakeManagerDelegationRemoved)
	if err := _StakeManager.contract.UnpackLog(event, "DelegationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerDelegationWithdrawnIterator is returned from FilterDelegationWithdrawn and is used to iterate over the raw logs and unpacked data for DelegationWithdrawn events raised by the StakeManager contract.
type StakeManagerDelegationWithdrawnIterator struct {
	Event *StakeManagerDelegationWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakeManagerDelegationWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerDelegationWithdrawn)
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
		it.Event = new(StakeManagerDelegationWithdrawn)
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
func (it *StakeManagerDelegationWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerDelegationWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerDelegationWithdrawn represents a DelegationWithdrawn event raised by the StakeManager contract.
type StakeManagerDelegationWithdrawn struct {
	Validator common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegationWithdrawn is a free log retrieval operation binding the contract event 0xaa0001b0e2e76b9a1b925257fefa58b756aebc52d2d7c8a85ea5beacc77a2100.
//
// Solidity: event DelegationWithdrawn(address indexed validator, address indexed recipient, uint256 amount)
func (_StakeManager *StakeManagerFilterer) FilterDelegationWithdrawn(opts *bind.FilterOpts, validator []common.Address, recipient []common.Address) (*StakeManagerDelegationWithdrawnIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "DelegationWithdrawn", validatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerDelegationWithdrawnIterator{contract: _StakeManager.contract, event: "DelegationWithdrawn", logs: logs, sub: sub}, nil
}

// WatchDelegationWithdrawn is a free log subscription operation binding the contract event 0xaa0001b0e2e76b9a1b925257fefa58b756aebc52d2d7c8a85ea5beacc77a2100.
//
// Solidity: event DelegationWithdrawn(address indexed validator, address indexed recipient, uint256 amount)
func (_StakeManager *StakeManagerFilterer) WatchDelegationWithdrawn(opts *bind.WatchOpts, sink chan<- *StakeManagerDelegationWithdrawn, validator []common.Address, recipient []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "DelegationWithdrawn", validatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerDelegationWithdrawn)
				if err := _StakeManager.contract.UnpackLog(event, "DelegationWithdrawn", log); err != nil {
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

// ParseDelegationWithdrawn is a log parse operation binding the contract event 0xaa0001b0e2e76b9a1b925257fefa58b756aebc52d2d7c8a85ea5beacc77a2100.
//
// Solidity: event DelegationWithdrawn(address indexed validator, address indexed recipient, uint256 amount)
func (_StakeManager *StakeManagerFilterer) ParseDelegationWithdrawn(log types.Log) (*StakeManagerDelegationWithdrawn, error) {
	event := new(StakeManagerDelegationWithdrawn)
	if err := _StakeManager.contract.UnpackLog(event, "DelegationWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakeManager contract.
type StakeManagerInitializedIterator struct {
	Event *StakeManagerInitialized // Event containing the contract specifics and raw log

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
func (it *StakeManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerInitialized)
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
		it.Event = new(StakeManagerInitialized)
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
func (it *StakeManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerInitialized represents a Initialized event raised by the StakeManager contract.
type StakeManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakeManager *StakeManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakeManagerInitializedIterator, error) {

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakeManagerInitializedIterator{contract: _StakeManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StakeManager *StakeManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakeManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerInitialized)
				if err := _StakeManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StakeManager *StakeManagerFilterer) ParseInitialized(log types.Log) (*StakeManagerInitialized, error) {
	event := new(StakeManagerInitialized)
	if err := _StakeManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerStakeAddedIterator is returned from FilterStakeAdded and is used to iterate over the raw logs and unpacked data for StakeAdded events raised by the StakeManager contract.
type StakeManagerStakeAddedIterator struct {
	Event *StakeManagerStakeAdded // Event containing the contract specifics and raw log

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
func (it *StakeManagerStakeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerStakeAdded)
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
		it.Event = new(StakeManagerStakeAdded)
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
func (it *StakeManagerStakeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerStakeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerStakeAdded represents a StakeAdded event raised by the StakeManager contract.
type StakeManagerStakeAdded struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeAdded is a free log retrieval operation binding the contract event 0x7c717985ac273e663b7f3050f5b15a4388ff6ed952338954f650e2093e13937f.
//
// Solidity: event StakeAdded(address indexed validator, uint256 amount)
func (_StakeManager *StakeManagerFilterer) FilterStakeAdded(opts *bind.FilterOpts, validator []common.Address) (*StakeManagerStakeAddedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "StakeAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerStakeAddedIterator{contract: _StakeManager.contract, event: "StakeAdded", logs: logs, sub: sub}, nil
}

// WatchStakeAdded is a free log subscription operation binding the contract event 0x7c717985ac273e663b7f3050f5b15a4388ff6ed952338954f650e2093e13937f.
//
// Solidity: event StakeAdded(address indexed validator, uint256 amount)
func (_StakeManager *StakeManagerFilterer) WatchStakeAdded(opts *bind.WatchOpts, sink chan<- *StakeManagerStakeAdded, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "StakeAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerStakeAdded)
				if err := _StakeManager.contract.UnpackLog(event, "StakeAdded", log); err != nil {
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

// ParseStakeAdded is a log parse operation binding the contract event 0x7c717985ac273e663b7f3050f5b15a4388ff6ed952338954f650e2093e13937f.
//
// Solidity: event StakeAdded(address indexed validator, uint256 amount)
func (_StakeManager *StakeManagerFilterer) ParseStakeAdded(log types.Log) (*StakeManagerStakeAdded, error) {
	event := new(StakeManagerStakeAdded)
	if err := _StakeManager.contract.UnpackLog(event, "StakeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerStakeRemovedIterator is returned from FilterStakeRemoved and is used to iterate over the raw logs and unpacked data for StakeRemoved events raised by the StakeManager contract.
type StakeManagerStakeRemovedIterator struct {
	Event *StakeManagerStakeRemoved // Event containing the contract specifics and raw log

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
func (it *StakeManagerStakeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerStakeRemoved)
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
		it.Event = new(StakeManagerStakeRemoved)
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
func (it *StakeManagerStakeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerStakeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerStakeRemoved represents a StakeRemoved event raised by the StakeManager contract.
type StakeManagerStakeRemoved struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeRemoved is a free log retrieval operation binding the contract event 0xa018dcbc822f59fb0d0c3e7a86c8e4259b9676cdea9e5fc26279b9c4c5d86eef.
//
// Solidity: event StakeRemoved(address indexed validator, uint256 amount)
func (_StakeManager *StakeManagerFilterer) FilterStakeRemoved(opts *bind.FilterOpts, validator []common.Address) (*StakeManagerStakeRemovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "StakeRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerStakeRemovedIterator{contract: _StakeManager.contract, event: "StakeRemoved", logs: logs, sub: sub}, nil
}

// WatchStakeRemoved is a free log subscription operation binding the contract event 0xa018dcbc822f59fb0d0c3e7a86c8e4259b9676cdea9e5fc26279b9c4c5d86eef.
//
// Solidity: event StakeRemoved(address indexed validator, uint256 amount)
func (_StakeManager *StakeManagerFilterer) WatchStakeRemoved(opts *bind.WatchOpts, sink chan<- *StakeManagerStakeRemoved, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "StakeRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerStakeRemoved)
				if err := _StakeManager.contract.UnpackLog(event, "StakeRemoved", log); err != nil {
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

// ParseStakeRemoved is a log parse operation binding the contract event 0xa018dcbc822f59fb0d0c3e7a86c8e4259b9676cdea9e5fc26279b9c4c5d86eef.
//
// Solidity: event StakeRemoved(address indexed validator, uint256 amount)
func (_StakeManager *StakeManagerFilterer) ParseStakeRemoved(log types.Log) (*StakeManagerStakeRemoved, error) {
	event := new(StakeManagerStakeRemoved)
	if err := _StakeManager.contract.UnpackLog(event, "StakeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the StakeManager contract.
type StakeManagerStakeWithdrawnIterator struct {
	Event *StakeManagerStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakeManagerStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerStakeWithdrawn)
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
		it.Event = new(StakeManagerStakeWithdrawn)
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
func (it *StakeManagerStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerStakeWithdrawn represents a StakeWithdrawn event raised by the StakeManager contract.
type StakeManagerStakeWithdrawn struct {
	Validator common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed validator, address indexed recipient, uint256 amount)
func (_StakeManager *StakeManagerFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, validator []common.Address, recipient []common.Address) (*StakeManagerStakeWithdrawnIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "StakeWithdrawn", validatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerStakeWithdrawnIterator{contract: _StakeManager.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed validator, address indexed recipient, uint256 amount)
func (_StakeManager *StakeManagerFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *StakeManagerStakeWithdrawn, validator []common.Address, recipient []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "StakeWithdrawn", validatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerStakeWithdrawn)
				if err := _StakeManager.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed validator, address indexed recipient, uint256 amount)
func (_StakeManager *StakeManagerFilterer) ParseStakeWithdrawn(log types.Log) (*StakeManagerStakeWithdrawn, error) {
	event := new(StakeManagerStakeWithdrawn)
	if err := _StakeManager.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeManagerValidatorStakeSlashedIterator is returned from FilterValidatorStakeSlashed and is used to iterate over the raw logs and unpacked data for ValidatorStakeSlashed events raised by the StakeManager contract.
type StakeManagerValidatorStakeSlashedIterator struct {
	Event *StakeManagerValidatorStakeSlashed // Event containing the contract specifics and raw log

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
func (it *StakeManagerValidatorStakeSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeManagerValidatorStakeSlashed)
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
		it.Event = new(StakeManagerValidatorStakeSlashed)
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
func (it *StakeManagerValidatorStakeSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeManagerValidatorStakeSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeManagerValidatorStakeSlashed represents a ValidatorStakeSlashed event raised by the StakeManager contract.
type StakeManagerValidatorStakeSlashed struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorStakeSlashed is a free log retrieval operation binding the contract event 0x2fd10b18da60ce9090915fbb72edd4d6550168ad1915bd58038d803b9faba2d7.
//
// Solidity: event ValidatorStakeSlashed(address indexed validator, uint256 amount)
func (_StakeManager *StakeManagerFilterer) FilterValidatorStakeSlashed(opts *bind.FilterOpts, validator []common.Address) (*StakeManagerValidatorStakeSlashedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakeManager.contract.FilterLogs(opts, "ValidatorStakeSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakeManagerValidatorStakeSlashedIterator{contract: _StakeManager.contract, event: "ValidatorStakeSlashed", logs: logs, sub: sub}, nil
}

// WatchValidatorStakeSlashed is a free log subscription operation binding the contract event 0x2fd10b18da60ce9090915fbb72edd4d6550168ad1915bd58038d803b9faba2d7.
//
// Solidity: event ValidatorStakeSlashed(address indexed validator, uint256 amount)
func (_StakeManager *StakeManagerFilterer) WatchValidatorStakeSlashed(opts *bind.WatchOpts, sink chan<- *StakeManagerValidatorStakeSlashed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakeManager.contract.WatchLogs(opts, "ValidatorStakeSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeManagerValidatorStakeSlashed)
				if err := _StakeManager.contract.UnpackLog(event, "ValidatorStakeSlashed", log); err != nil {
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

// ParseValidatorStakeSlashed is a log parse operation binding the contract event 0x2fd10b18da60ce9090915fbb72edd4d6550168ad1915bd58038d803b9faba2d7.
//
// Solidity: event ValidatorStakeSlashed(address indexed validator, uint256 amount)
func (_StakeManager *StakeManagerFilterer) ParseValidatorStakeSlashed(log types.Log) (*StakeManagerValidatorStakeSlashed, error) {
	event := new(StakeManagerValidatorStakeSlashed)
	if err := _StakeManager.contract.UnpackLog(event, "ValidatorStakeSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
