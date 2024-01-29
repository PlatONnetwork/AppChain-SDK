# ERC-20

基于AppChain-SDK的L2为开发者提供了两种代币发行方案：L1发行和L2发行。目前仅支持标准ERC-20代币。

## L1发行

为了满足在L1上发行的资产转移到L2，AppChain-SDK为开发者提供了L1发行机制。结合状态同步机制，开发者可以自由将资产在L1和L2之间转移。

### DepositManager

部署在L1的`DepositManager`合约负责锁定和解锁用户的代币。

`DepositManager`合约对外提供`deposit`接口，任何人都可以调用该接口将指定数量的代币存入L2中。调用该接口后，用户指定数量的代币将被锁定，同时该合约调用`StateSender`合约的`syncState`接口发出状态同步事件，以将代币转移到L2中。

`DepositManager`合约同时也处理来自L2的状态同步事件（即取款），解锁指定用户的指定数量的代币。

### DepositHandler

与之对应的是，在L2上`DepositHandler`合约负责接收用户转移的代币和提取用户的代币到L1。

`DepositHandler`合约处理来自`DepositManager`合约的存款事件，为指定的用户增加指定数量的代币。当用户想将代币提取到L1时，可调用合约的`withdraw`接口，该接口将减少用户指定数量的代币，同时调用`L2StateSender`合约的`syncState`接口，发出状态同步事件（也称为退出事件），将用户指定数量的代币转移到L1中。

当包含有该用户的退出事件的checkpoint提交到L1上时，用户需要主动调用ExitHelper合约的`exit`接口完成取款动作，以解锁锁定在`DepositManager`合约的指定数量的代币。

## L2发行

与L1发行对应，AppChain-SDK也为开发者提供了L2发行机制。结合状态同步机制，L2发行的代币也可以自由在L2和L1之间转移。

### WithdrawManager

与`DepositManager`合约类似，部署在L2的`WithdrawManager`合约负责锁定和解锁用户的代币。

当用户需要将代币转移到L1时，可以调用`deposit`接口，传人代币接收地址和代币数量，这些数量的代币将锁定在合约中。同时，合约会调用`L2StateSender`合约的`syncState`接口发出状态同步事件（也称为退出事件）。包含有该事件的checkpoint提交到L1时，用户需要主动调用ExitHelper合约的`exit`接口完成代币转移动作。

`DepositManager`合约同时处理来自L1的状态同步事件，将事件对应的用户的指定数量的代币解锁。

### WithdrawHandler

`WithdrawHandler`部署在L1上，负责存取用户从L2转移的代币。

用户将代币从L2转移到L1，产生的退出事件将在`WithdrawHandler`合约中处理，该合约调用代币合约的接口`mint`去铸造指定数量的代币给用户。

而当用户需要将代币转移回L2时，可调用合约提供的`withdraw`接口，该接口销毁用户指定数量的代币，并调用`StateSender`合约的`syncState`接口发出状态同步事件。该事件将由部署在L2的`WithdrawManager`合约处理，以解锁事件中指定数量的代币。
