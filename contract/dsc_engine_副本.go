// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DSCEngineMetaData contains all meta data concerning the DSCEngine contract.
var DSCEngineMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"priceFeedAddress\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"dscAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MintDsc\",\"inputs\":[{\"name\":\"amountDscToMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnDsc\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositCollateral\",\"inputs\":[{\"name\":\"tokenCollateralAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositCollateralAndMintDsc\",\"inputs\":[{\"name\":\"tokenCollateralAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountDscToMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccountCollateralValue\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"totalCollateralValueInUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAccountInformation\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"totalDscMinted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralValueInUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollateralBalanceOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollateralTokenPriceFeed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHealthFactor\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSDscMinted\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenAmountFromUsd\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdAmountInWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUSDValue\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidate\",\"inputs\":[{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debtToCover\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemCollateral\",\"inputs\":[{\"name\":\"tokenCollateralAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemCollateralForDsc\",\"inputs\":[{\"name\":\"tokenCollateralAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CollateralDeposited\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollateralRedeemed\",\"inputs\":[{\"name\":\"redeemedFrom\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"redeemedTo\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DSCEngine__HealthFactorIsBroken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DSCEngine__HealthFactorIsNotBroken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DSCEngine__HealthFactorIsNotImproved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DSCEngine__NeedsMoreThanZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DSCEngine__OnlyAllowWETHAndWBTC\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DSCEngine__TokenAddressAndPriceFeedAddressMustBeSameLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DSCEngine__TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051611da5380380611da5833981810160405281019061003191906103af565b60015f819055508151835114610073576040517ff7fa022b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f90505b83518110156101b75782818151811061009457610093610437565b5b602002602001015160015f8684815181106100b2576100b1610437565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600484828151811061014557610144610437565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508080600101915050610078565b508073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050505050610464565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61024f82610209565b810181811067ffffffffffffffff8211171561026e5761026d610219565b5b80604052505050565b5f6102806101f4565b905061028c8282610246565b919050565b5f67ffffffffffffffff8211156102ab576102aa610219565b5b602082029050602081019050919050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102e9826102c0565b9050919050565b6102f9816102df565b8114610303575f5ffd5b50565b5f81519050610314816102f0565b92915050565b5f61032c61032784610291565b610277565b9050808382526020820190506020840283018581111561034f5761034e6102bc565b5b835b8181101561037857806103648882610306565b845260208401935050602081019050610351565b5050509392505050565b5f82601f83011261039657610395610205565b5b81516103a684826020860161031a565b91505092915050565b5f5f5f606084860312156103c6576103c56101fd565b5b5f84015167ffffffffffffffff8111156103e3576103e2610201565b5b6103ef86828701610382565b935050602084015167ffffffffffffffff8111156104105761040f610201565b5b61041c86828701610382565b925050604061042d86828701610306565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b60805161191b61048a5f395f81816106ab01528181611147015261121d015261191b5ff3fe608060405234801561000f575f5ffd5b50600436106100f3575f3560e01c80637d1a445011610095578063e90db8a311610064578063e90db8a3146102a4578063f6876608146102c0578063fa76dcf2146102dc578063fe6bcd7c1461030c576100f3565b80637d1a44501461020c5780639acd81b31461023c578063a5d5db0c14610258578063afea2e4814610274576100f3565b80632e10d41b116100d15780632e10d41b1461015f578063406e6fb91461018f57806349126ac2146101bf5780637be564fc146101db576100f3565b80631c08adda146100f757806326c01303146101275780632d3844e814610143575b5f5ffd5b610111600480360381019061010c91906113b2565b61033c565b60405161011e91906113ec565b60405180910390f35b610141600480360381019061013c9190611438565b6103a1565b005b61015d60048036038101906101589190611488565b6104e0565b005b610179600480360381019061017491906113b2565b610542565b60405161018691906114d5565b60405180910390f35b6101a960048036038101906101a491906114ee565b610588565b6040516101b691906114d5565b60405180910390f35b6101d960048036038101906101d4919061152c565b61060a565b005b6101f560048036038101906101f091906113b2565b61078a565b604051610203929190611557565b60405180910390f35b610226600480360381019061022191906113b2565b61079e565b60405161023391906114d5565b60405180910390f35b61025660048036038101906102519190611488565b610898565b005b610272600480360381019061026d9190611488565b6108fc565b005b61028e60048036038101906102899190611488565b610bb9565b60405161029b91906114d5565b60405180910390f35b6102be60048036038101906102b9919061157e565b610cc7565b005b6102da60048036038101906102d5919061152c565b610cdf565b005b6102f660048036038101906102f19190611488565b610d41565b60405161030391906114d5565b60405180910390f35b610326600480360381019061032191906113b2565b610e4f565b60405161033391906114d5565b60405180910390f35b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b805f81036103db576040517fcb1d3f7d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e3610e60565b5f6103ed84610ea4565b9050670de0b6b3a76400008110610430576040517f74433d2200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61043b8685610bb9565b90505f6064600a8361044d91906115fb565b6104579190611669565b90505f81836104669190611699565b905061047488828933610f2b565b61047f8688336110f1565b5f61048988610ea4565b90508481116104c4576040517fc422279900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104cd336112a7565b50505050506104da6112f9565b50505050565b805f810361051a576040517fcb1d3f7d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610522610e60565b61052b82610cdf565b6105358383610898565b61053d6112f9565b505050565b5f60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f60025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b805f8103610644576040517fcb1d3f7d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61064c610e60565b8160035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546106989190611699565b925050819055506106a8336112a7565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166340c10f1933856040518363ffffffff1660e01b81526004016107049291906116cc565b6020604051808303815f875af1158015610720573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107449190611728565b90508061077d576040517f4dc34d1700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506107866112f9565b5050565b5f5f61079583611302565b91509150915091565b5f5f5f90505b600480549050811015610892575f600482815481106107c6576107c5611753565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f60025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490506108768282610d41565b846108819190611699565b9350505080806001019150506107a4565b50919050565b805f81036108d2576040517fcb1d3f7d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108da610e60565b6108e683833333610f2b565b6108ef336112a7565b6108f76112f9565b505050565b805f8103610936576040517fcb1d3f7d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b825f73ffffffffffffffffffffffffffffffffffffffff1660015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036109f8576040517f439ac4bb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a00610e60565b8260025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610a879190611699565b925050819055508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167ff1c0dd7e9b98bbff859029005ef89b127af049cd18df1a8d79f0b7e019911e5685604051610aeb91906114d5565b60405180910390a35f8473ffffffffffffffffffffffffffffffffffffffff166323b872dd3330876040518463ffffffff1660e01b8152600401610b3193929190611780565b6020604051808303815f875af1158015610b4d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b719190611728565b905080610baa576040517f4dc34d1700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50610bb36112f9565b50505050565b5f5f60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f8173ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa158015610c64573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c88919061183b565b5050509150506402540be40081610c9f91906115fb565b670de0b6b3a764000085610cb391906115fb565b610cbd9190611669565b9250505092915050565b610cd183836108fc565b610cda8161060a565b505050565b805f8103610d19576040517fcb1d3f7d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d21610e60565b610d2c8233336110f1565b610d35336112a7565b610d3d6112f9565b5050565b5f5f60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f8173ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa158015610dec573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e10919061183b565b505050915050670de0b6b3a76400006402540be4008583610e3191906115fb565b610e3b91906115fb565b610e459190611669565b9250505092915050565b5f610e5982610ea4565b9050919050565b60025f5403610e9b576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f81905550565b5f5f5f610eb084611302565b915091505f8203610ee5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff92505050610f26565b5f6064603283610ef591906115fb565b610eff9190611669565b905082670de0b6b3a764000082610f1691906115fb565b610f209190611669565b93505050505b919050565b8260025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610fb291906118b2565b925050819055508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fb292db12b8dfa23b760ee5e3bd1d3be184cd8f0eeeacd27ce120d4de4e4721f68660405161102d91906114d5565b60405180910390a45f8473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83866040518363ffffffff1660e01b81526004016110719291906116cc565b6020604051808303815f875af115801561108d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110b19190611728565b9050806110ea576040517f4dc34d1700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b8260035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461113d91906118b2565b925050819055505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd8330876040518463ffffffff1660e01b81526004016111a293929190611780565b6020604051808303815f875af11580156111be573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111e29190611728565b90508061121b576040517f4dc34d1700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166342966c68856040518263ffffffff1660e01b815260040161127491906114d5565b5f604051808303815f87803b15801561128b575f5ffd5b505af115801561129d573d5f5f3e3d5ffd5b5050505050505050565b5f6112b182610ea4565b9050670de0b6b3a76400008110156112f5576040517f5c906f3900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b60015f81905550565b5f5f60035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054915061134d8361079e565b9050915091565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61138182611358565b9050919050565b61139181611377565b811461139b575f5ffd5b50565b5f813590506113ac81611388565b92915050565b5f602082840312156113c7576113c6611354565b5b5f6113d48482850161139e565b91505092915050565b6113e681611377565b82525050565b5f6020820190506113ff5f8301846113dd565b92915050565b5f819050919050565b61141781611405565b8114611421575f5ffd5b50565b5f813590506114328161140e565b92915050565b5f5f5f6060848603121561144f5761144e611354565b5b5f61145c8682870161139e565b935050602061146d8682870161139e565b925050604061147e86828701611424565b9150509250925092565b5f5f6040838503121561149e5761149d611354565b5b5f6114ab8582860161139e565b92505060206114bc85828601611424565b9150509250929050565b6114cf81611405565b82525050565b5f6020820190506114e85f8301846114c6565b92915050565b5f5f6040838503121561150457611503611354565b5b5f6115118582860161139e565b92505060206115228582860161139e565b9150509250929050565b5f6020828403121561154157611540611354565b5b5f61154e84828501611424565b91505092915050565b5f60408201905061156a5f8301856114c6565b61157760208301846114c6565b9392505050565b5f5f5f6060848603121561159557611594611354565b5b5f6115a28682870161139e565b93505060206115b386828701611424565b92505060406115c486828701611424565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61160582611405565b915061161083611405565b925082820261161e81611405565b91508282048414831517611635576116346115ce565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61167382611405565b915061167e83611405565b92508261168e5761168d61163c565b5b828204905092915050565b5f6116a382611405565b91506116ae83611405565b92508282019050808211156116c6576116c56115ce565b5b92915050565b5f6040820190506116df5f8301856113dd565b6116ec60208301846114c6565b9392505050565b5f8115159050919050565b611707816116f3565b8114611711575f5ffd5b50565b5f81519050611722816116fe565b92915050565b5f6020828403121561173d5761173c611354565b5b5f61174a84828501611714565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6060820190506117935f8301866113dd565b6117a060208301856113dd565b6117ad60408301846114c6565b949350505050565b5f69ffffffffffffffffffff82169050919050565b6117d3816117b5565b81146117dd575f5ffd5b50565b5f815190506117ee816117ca565b92915050565b5f819050919050565b611806816117f4565b8114611810575f5ffd5b50565b5f81519050611821816117fd565b92915050565b5f815190506118358161140e565b92915050565b5f5f5f5f5f60a0868803121561185457611853611354565b5b5f611861888289016117e0565b955050602061187288828901611813565b945050604061188388828901611827565b935050606061189488828901611827565b92505060806118a5888289016117e0565b9150509295509295909350565b5f6118bc82611405565b91506118c783611405565b92508282039050818111156118df576118de6115ce565b5b9291505056fea26469706673582212205faa5a8fd7a077d6f1dd88a11f3c07fd845b2bbfac58137454d54f9d8626feb064736f6c634300081c0033",
}

// DSCEngineABI is the input ABI used to generate the binding from.
// Deprecated: Use DSCEngineMetaData.ABI instead.
var DSCEngineABI = DSCEngineMetaData.ABI

// DSCEngineBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DSCEngineMetaData.Bin instead.
var DSCEngineBin = DSCEngineMetaData.Bin

// DeployDSCEngine deploys a new Ethereum contract, binding an instance of DSCEngine to it.
func DeployDSCEngine(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddresses []common.Address, priceFeedAddress []common.Address, dscAddress common.Address) (common.Address, *types.Transaction, *DSCEngine, error) {
	parsed, err := DSCEngineMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DSCEngineBin), backend, tokenAddresses, priceFeedAddress, dscAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSCEngine{DSCEngineCaller: DSCEngineCaller{contract: contract}, DSCEngineTransactor: DSCEngineTransactor{contract: contract}, DSCEngineFilterer: DSCEngineFilterer{contract: contract}}, nil
}

// DSCEngine is an auto generated Go binding around an Ethereum contract.
type DSCEngine struct {
	DSCEngineCaller     // Read-only binding to the contract
	DSCEngineTransactor // Write-only binding to the contract
	DSCEngineFilterer   // Log filterer for contract events
}

// DSCEngineCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSCEngineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSCEngineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSCEngineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSCEngineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSCEngineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSCEngineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSCEngineSession struct {
	Contract     *DSCEngine        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSCEngineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSCEngineCallerSession struct {
	Contract *DSCEngineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DSCEngineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSCEngineTransactorSession struct {
	Contract     *DSCEngineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DSCEngineRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSCEngineRaw struct {
	Contract *DSCEngine // Generic contract binding to access the raw methods on
}

// DSCEngineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSCEngineCallerRaw struct {
	Contract *DSCEngineCaller // Generic read-only contract binding to access the raw methods on
}

// DSCEngineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSCEngineTransactorRaw struct {
	Contract *DSCEngineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSCEngine creates a new instance of DSCEngine, bound to a specific deployed contract.
func NewDSCEngine(address common.Address, backend bind.ContractBackend) (*DSCEngine, error) {
	contract, err := bindDSCEngine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSCEngine{DSCEngineCaller: DSCEngineCaller{contract: contract}, DSCEngineTransactor: DSCEngineTransactor{contract: contract}, DSCEngineFilterer: DSCEngineFilterer{contract: contract}}, nil
}

// NewDSCEngineCaller creates a new read-only instance of DSCEngine, bound to a specific deployed contract.
func NewDSCEngineCaller(address common.Address, caller bind.ContractCaller) (*DSCEngineCaller, error) {
	contract, err := bindDSCEngine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSCEngineCaller{contract: contract}, nil
}

// NewDSCEngineTransactor creates a new write-only instance of DSCEngine, bound to a specific deployed contract.
func NewDSCEngineTransactor(address common.Address, transactor bind.ContractTransactor) (*DSCEngineTransactor, error) {
	contract, err := bindDSCEngine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSCEngineTransactor{contract: contract}, nil
}

// NewDSCEngineFilterer creates a new log filterer instance of DSCEngine, bound to a specific deployed contract.
func NewDSCEngineFilterer(address common.Address, filterer bind.ContractFilterer) (*DSCEngineFilterer, error) {
	contract, err := bindDSCEngine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSCEngineFilterer{contract: contract}, nil
}

// bindDSCEngine binds a generic wrapper to an already deployed contract.
func bindDSCEngine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DSCEngineMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSCEngine *DSCEngineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DSCEngine.Contract.DSCEngineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSCEngine *DSCEngineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSCEngine.Contract.DSCEngineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSCEngine *DSCEngineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSCEngine.Contract.DSCEngineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSCEngine *DSCEngineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DSCEngine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSCEngine *DSCEngineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSCEngine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSCEngine *DSCEngineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSCEngine.Contract.contract.Transact(opts, method, params...)
}

// GetAccountCollateralValue is a free data retrieval call binding the contract method 0x7d1a4450.
//
// Solidity: function getAccountCollateralValue(address user) view returns(uint256 totalCollateralValueInUsd)
func (_DSCEngine *DSCEngineCaller) GetAccountCollateralValue(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getAccountCollateralValue", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountCollateralValue is a free data retrieval call binding the contract method 0x7d1a4450.
//
// Solidity: function getAccountCollateralValue(address user) view returns(uint256 totalCollateralValueInUsd)
func (_DSCEngine *DSCEngineSession) GetAccountCollateralValue(user common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetAccountCollateralValue(&_DSCEngine.CallOpts, user)
}

// GetAccountCollateralValue is a free data retrieval call binding the contract method 0x7d1a4450.
//
// Solidity: function getAccountCollateralValue(address user) view returns(uint256 totalCollateralValueInUsd)
func (_DSCEngine *DSCEngineCallerSession) GetAccountCollateralValue(user common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetAccountCollateralValue(&_DSCEngine.CallOpts, user)
}

// GetAccountInformation is a free data retrieval call binding the contract method 0x7be564fc.
//
// Solidity: function getAccountInformation(address user) view returns(uint256 totalDscMinted, uint256 collateralValueInUsd)
func (_DSCEngine *DSCEngineCaller) GetAccountInformation(opts *bind.CallOpts, user common.Address) (struct {
	TotalDscMinted       *big.Int
	CollateralValueInUsd *big.Int
}, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getAccountInformation", user)

	outstruct := new(struct {
		TotalDscMinted       *big.Int
		CollateralValueInUsd *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalDscMinted = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CollateralValueInUsd = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAccountInformation is a free data retrieval call binding the contract method 0x7be564fc.
//
// Solidity: function getAccountInformation(address user) view returns(uint256 totalDscMinted, uint256 collateralValueInUsd)
func (_DSCEngine *DSCEngineSession) GetAccountInformation(user common.Address) (struct {
	TotalDscMinted       *big.Int
	CollateralValueInUsd *big.Int
}, error) {
	return _DSCEngine.Contract.GetAccountInformation(&_DSCEngine.CallOpts, user)
}

// GetAccountInformation is a free data retrieval call binding the contract method 0x7be564fc.
//
// Solidity: function getAccountInformation(address user) view returns(uint256 totalDscMinted, uint256 collateralValueInUsd)
func (_DSCEngine *DSCEngineCallerSession) GetAccountInformation(user common.Address) (struct {
	TotalDscMinted       *big.Int
	CollateralValueInUsd *big.Int
}, error) {
	return _DSCEngine.Contract.GetAccountInformation(&_DSCEngine.CallOpts, user)
}

// GetCollateralBalanceOf is a free data retrieval call binding the contract method 0x406e6fb9.
//
// Solidity: function getCollateralBalanceOf(address user, address token) view returns(uint256)
func (_DSCEngine *DSCEngineCaller) GetCollateralBalanceOf(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getCollateralBalanceOf", user, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralBalanceOf is a free data retrieval call binding the contract method 0x406e6fb9.
//
// Solidity: function getCollateralBalanceOf(address user, address token) view returns(uint256)
func (_DSCEngine *DSCEngineSession) GetCollateralBalanceOf(user common.Address, token common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetCollateralBalanceOf(&_DSCEngine.CallOpts, user, token)
}

// GetCollateralBalanceOf is a free data retrieval call binding the contract method 0x406e6fb9.
//
// Solidity: function getCollateralBalanceOf(address user, address token) view returns(uint256)
func (_DSCEngine *DSCEngineCallerSession) GetCollateralBalanceOf(user common.Address, token common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetCollateralBalanceOf(&_DSCEngine.CallOpts, user, token)
}

// GetCollateralTokenPriceFeed is a free data retrieval call binding the contract method 0x1c08adda.
//
// Solidity: function getCollateralTokenPriceFeed(address token) view returns(address)
func (_DSCEngine *DSCEngineCaller) GetCollateralTokenPriceFeed(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getCollateralTokenPriceFeed", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollateralTokenPriceFeed is a free data retrieval call binding the contract method 0x1c08adda.
//
// Solidity: function getCollateralTokenPriceFeed(address token) view returns(address)
func (_DSCEngine *DSCEngineSession) GetCollateralTokenPriceFeed(token common.Address) (common.Address, error) {
	return _DSCEngine.Contract.GetCollateralTokenPriceFeed(&_DSCEngine.CallOpts, token)
}

// GetCollateralTokenPriceFeed is a free data retrieval call binding the contract method 0x1c08adda.
//
// Solidity: function getCollateralTokenPriceFeed(address token) view returns(address)
func (_DSCEngine *DSCEngineCallerSession) GetCollateralTokenPriceFeed(token common.Address) (common.Address, error) {
	return _DSCEngine.Contract.GetCollateralTokenPriceFeed(&_DSCEngine.CallOpts, token)
}

// GetHealthFactor is a free data retrieval call binding the contract method 0xfe6bcd7c.
//
// Solidity: function getHealthFactor(address user) view returns(uint256)
func (_DSCEngine *DSCEngineCaller) GetHealthFactor(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getHealthFactor", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHealthFactor is a free data retrieval call binding the contract method 0xfe6bcd7c.
//
// Solidity: function getHealthFactor(address user) view returns(uint256)
func (_DSCEngine *DSCEngineSession) GetHealthFactor(user common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetHealthFactor(&_DSCEngine.CallOpts, user)
}

// GetHealthFactor is a free data retrieval call binding the contract method 0xfe6bcd7c.
//
// Solidity: function getHealthFactor(address user) view returns(uint256)
func (_DSCEngine *DSCEngineCallerSession) GetHealthFactor(user common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetHealthFactor(&_DSCEngine.CallOpts, user)
}

// GetSDscMinted is a free data retrieval call binding the contract method 0x2e10d41b.
//
// Solidity: function getSDscMinted(address user) view returns(uint256)
func (_DSCEngine *DSCEngineCaller) GetSDscMinted(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getSDscMinted", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSDscMinted is a free data retrieval call binding the contract method 0x2e10d41b.
//
// Solidity: function getSDscMinted(address user) view returns(uint256)
func (_DSCEngine *DSCEngineSession) GetSDscMinted(user common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetSDscMinted(&_DSCEngine.CallOpts, user)
}

// GetSDscMinted is a free data retrieval call binding the contract method 0x2e10d41b.
//
// Solidity: function getSDscMinted(address user) view returns(uint256)
func (_DSCEngine *DSCEngineCallerSession) GetSDscMinted(user common.Address) (*big.Int, error) {
	return _DSCEngine.Contract.GetSDscMinted(&_DSCEngine.CallOpts, user)
}

// GetTokenAmountFromUsd is a free data retrieval call binding the contract method 0xafea2e48.
//
// Solidity: function getTokenAmountFromUsd(address token, uint256 usdAmountInWei) view returns(uint256)
func (_DSCEngine *DSCEngineCaller) GetTokenAmountFromUsd(opts *bind.CallOpts, token common.Address, usdAmountInWei *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getTokenAmountFromUsd", token, usdAmountInWei)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenAmountFromUsd is a free data retrieval call binding the contract method 0xafea2e48.
//
// Solidity: function getTokenAmountFromUsd(address token, uint256 usdAmountInWei) view returns(uint256)
func (_DSCEngine *DSCEngineSession) GetTokenAmountFromUsd(token common.Address, usdAmountInWei *big.Int) (*big.Int, error) {
	return _DSCEngine.Contract.GetTokenAmountFromUsd(&_DSCEngine.CallOpts, token, usdAmountInWei)
}

// GetTokenAmountFromUsd is a free data retrieval call binding the contract method 0xafea2e48.
//
// Solidity: function getTokenAmountFromUsd(address token, uint256 usdAmountInWei) view returns(uint256)
func (_DSCEngine *DSCEngineCallerSession) GetTokenAmountFromUsd(token common.Address, usdAmountInWei *big.Int) (*big.Int, error) {
	return _DSCEngine.Contract.GetTokenAmountFromUsd(&_DSCEngine.CallOpts, token, usdAmountInWei)
}

// GetUSDValue is a free data retrieval call binding the contract method 0xfa76dcf2.
//
// Solidity: function getUSDValue(address token, uint256 amount) view returns(uint256)
func (_DSCEngine *DSCEngineCaller) GetUSDValue(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DSCEngine.contract.Call(opts, &out, "getUSDValue", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUSDValue is a free data retrieval call binding the contract method 0xfa76dcf2.
//
// Solidity: function getUSDValue(address token, uint256 amount) view returns(uint256)
func (_DSCEngine *DSCEngineSession) GetUSDValue(token common.Address, amount *big.Int) (*big.Int, error) {
	return _DSCEngine.Contract.GetUSDValue(&_DSCEngine.CallOpts, token, amount)
}

// GetUSDValue is a free data retrieval call binding the contract method 0xfa76dcf2.
//
// Solidity: function getUSDValue(address token, uint256 amount) view returns(uint256)
func (_DSCEngine *DSCEngineCallerSession) GetUSDValue(token common.Address, amount *big.Int) (*big.Int, error) {
	return _DSCEngine.Contract.GetUSDValue(&_DSCEngine.CallOpts, token, amount)
}

// MintDsc is a paid mutator transaction binding the contract method 0x49126ac2.
//
// Solidity: function MintDsc(uint256 amountDscToMint) returns()
func (_DSCEngine *DSCEngineTransactor) MintDsc(opts *bind.TransactOpts, amountDscToMint *big.Int) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "MintDsc", amountDscToMint)
}

// MintDsc is a paid mutator transaction binding the contract method 0x49126ac2.
//
// Solidity: function MintDsc(uint256 amountDscToMint) returns()
func (_DSCEngine *DSCEngineSession) MintDsc(amountDscToMint *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.MintDsc(&_DSCEngine.TransactOpts, amountDscToMint)
}

// MintDsc is a paid mutator transaction binding the contract method 0x49126ac2.
//
// Solidity: function MintDsc(uint256 amountDscToMint) returns()
func (_DSCEngine *DSCEngineTransactorSession) MintDsc(amountDscToMint *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.MintDsc(&_DSCEngine.TransactOpts, amountDscToMint)
}

// BurnDsc is a paid mutator transaction binding the contract method 0xf6876608.
//
// Solidity: function burnDsc(uint256 amount) returns()
func (_DSCEngine *DSCEngineTransactor) BurnDsc(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "burnDsc", amount)
}

// BurnDsc is a paid mutator transaction binding the contract method 0xf6876608.
//
// Solidity: function burnDsc(uint256 amount) returns()
func (_DSCEngine *DSCEngineSession) BurnDsc(amount *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.BurnDsc(&_DSCEngine.TransactOpts, amount)
}

// BurnDsc is a paid mutator transaction binding the contract method 0xf6876608.
//
// Solidity: function burnDsc(uint256 amount) returns()
func (_DSCEngine *DSCEngineTransactorSession) BurnDsc(amount *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.BurnDsc(&_DSCEngine.TransactOpts, amount)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_DSCEngine *DSCEngineTransactor) DepositCollateral(opts *bind.TransactOpts, tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "depositCollateral", tokenCollateralAddress, amountCollateral)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_DSCEngine *DSCEngineSession) DepositCollateral(tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.DepositCollateral(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_DSCEngine *DSCEngineTransactorSession) DepositCollateral(tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.DepositCollateral(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral)
}

// DepositCollateralAndMintDsc is a paid mutator transaction binding the contract method 0xe90db8a3.
//
// Solidity: function depositCollateralAndMintDsc(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountDscToMint) returns()
func (_DSCEngine *DSCEngineTransactor) DepositCollateralAndMintDsc(opts *bind.TransactOpts, tokenCollateralAddress common.Address, amountCollateral *big.Int, amountDscToMint *big.Int) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "depositCollateralAndMintDsc", tokenCollateralAddress, amountCollateral, amountDscToMint)
}

// DepositCollateralAndMintDsc is a paid mutator transaction binding the contract method 0xe90db8a3.
//
// Solidity: function depositCollateralAndMintDsc(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountDscToMint) returns()
func (_DSCEngine *DSCEngineSession) DepositCollateralAndMintDsc(tokenCollateralAddress common.Address, amountCollateral *big.Int, amountDscToMint *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.DepositCollateralAndMintDsc(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral, amountDscToMint)
}

// DepositCollateralAndMintDsc is a paid mutator transaction binding the contract method 0xe90db8a3.
//
// Solidity: function depositCollateralAndMintDsc(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountDscToMint) returns()
func (_DSCEngine *DSCEngineTransactorSession) DepositCollateralAndMintDsc(tokenCollateralAddress common.Address, amountCollateral *big.Int, amountDscToMint *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.DepositCollateralAndMintDsc(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral, amountDscToMint)
}

// Liquidate is a paid mutator transaction binding the contract method 0x26c01303.
//
// Solidity: function liquidate(address collateral, address user, uint256 debtToCover) returns()
func (_DSCEngine *DSCEngineTransactor) Liquidate(opts *bind.TransactOpts, collateral common.Address, user common.Address, debtToCover *big.Int) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "liquidate", collateral, user, debtToCover)
}

// Liquidate is a paid mutator transaction binding the contract method 0x26c01303.
//
// Solidity: function liquidate(address collateral, address user, uint256 debtToCover) returns()
func (_DSCEngine *DSCEngineSession) Liquidate(collateral common.Address, user common.Address, debtToCover *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.Liquidate(&_DSCEngine.TransactOpts, collateral, user, debtToCover)
}

// Liquidate is a paid mutator transaction binding the contract method 0x26c01303.
//
// Solidity: function liquidate(address collateral, address user, uint256 debtToCover) returns()
func (_DSCEngine *DSCEngineTransactorSession) Liquidate(collateral common.Address, user common.Address, debtToCover *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.Liquidate(&_DSCEngine.TransactOpts, collateral, user, debtToCover)
}

// RedeemCollateral is a paid mutator transaction binding the contract method 0x9acd81b3.
//
// Solidity: function redeemCollateral(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_DSCEngine *DSCEngineTransactor) RedeemCollateral(opts *bind.TransactOpts, tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "redeemCollateral", tokenCollateralAddress, amountCollateral)
}

// RedeemCollateral is a paid mutator transaction binding the contract method 0x9acd81b3.
//
// Solidity: function redeemCollateral(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_DSCEngine *DSCEngineSession) RedeemCollateral(tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.RedeemCollateral(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral)
}

// RedeemCollateral is a paid mutator transaction binding the contract method 0x9acd81b3.
//
// Solidity: function redeemCollateral(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_DSCEngine *DSCEngineTransactorSession) RedeemCollateral(tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.RedeemCollateral(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral)
}

// RedeemCollateralForDsc is a paid mutator transaction binding the contract method 0x2d3844e8.
//
// Solidity: function redeemCollateralForDsc(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_DSCEngine *DSCEngineTransactor) RedeemCollateralForDsc(opts *bind.TransactOpts, tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _DSCEngine.contract.Transact(opts, "redeemCollateralForDsc", tokenCollateralAddress, amountCollateral)
}

// RedeemCollateralForDsc is a paid mutator transaction binding the contract method 0x2d3844e8.
//
// Solidity: function redeemCollateralForDsc(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_DSCEngine *DSCEngineSession) RedeemCollateralForDsc(tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.RedeemCollateralForDsc(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral)
}

// RedeemCollateralForDsc is a paid mutator transaction binding the contract method 0x2d3844e8.
//
// Solidity: function redeemCollateralForDsc(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_DSCEngine *DSCEngineTransactorSession) RedeemCollateralForDsc(tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _DSCEngine.Contract.RedeemCollateralForDsc(&_DSCEngine.TransactOpts, tokenCollateralAddress, amountCollateral)
}

// DSCEngineCollateralDepositedIterator is returned from FilterCollateralDeposited and is used to iterate over the raw logs and unpacked data for CollateralDeposited events raised by the DSCEngine contract.
type DSCEngineCollateralDepositedIterator struct {
	Event *DSCEngineCollateralDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DSCEngineCollateralDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSCEngineCollateralDeposited)
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
		it.Event = new(DSCEngineCollateralDeposited)
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
func (it *DSCEngineCollateralDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSCEngineCollateralDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSCEngineCollateralDeposited represents a CollateralDeposited event raised by the DSCEngine contract.
type DSCEngineCollateralDeposited struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCollateralDeposited is a free log retrieval operation binding the contract event 0xf1c0dd7e9b98bbff859029005ef89b127af049cd18df1a8d79f0b7e019911e56.
//
// Solidity: event CollateralDeposited(address indexed user, address indexed token, uint256 amount)
func (_DSCEngine *DSCEngineFilterer) FilterCollateralDeposited(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*DSCEngineCollateralDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DSCEngine.contract.FilterLogs(opts, "CollateralDeposited", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &DSCEngineCollateralDepositedIterator{contract: _DSCEngine.contract, event: "CollateralDeposited", logs: logs, sub: sub}, nil
}

// WatchCollateralDeposited is a free log subscription operation binding the contract event 0xf1c0dd7e9b98bbff859029005ef89b127af049cd18df1a8d79f0b7e019911e56.
//
// Solidity: event CollateralDeposited(address indexed user, address indexed token, uint256 amount)
func (_DSCEngine *DSCEngineFilterer) WatchCollateralDeposited(opts *bind.WatchOpts, sink chan<- *DSCEngineCollateralDeposited, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DSCEngine.contract.WatchLogs(opts, "CollateralDeposited", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSCEngineCollateralDeposited)
				if err := _DSCEngine.contract.UnpackLog(event, "CollateralDeposited", log); err != nil {
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

// ParseCollateralDeposited is a log parse operation binding the contract event 0xf1c0dd7e9b98bbff859029005ef89b127af049cd18df1a8d79f0b7e019911e56.
//
// Solidity: event CollateralDeposited(address indexed user, address indexed token, uint256 amount)
func (_DSCEngine *DSCEngineFilterer) ParseCollateralDeposited(log types.Log) (*DSCEngineCollateralDeposited, error) {
	event := new(DSCEngineCollateralDeposited)
	if err := _DSCEngine.contract.UnpackLog(event, "CollateralDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DSCEngineCollateralRedeemedIterator is returned from FilterCollateralRedeemed and is used to iterate over the raw logs and unpacked data for CollateralRedeemed events raised by the DSCEngine contract.
type DSCEngineCollateralRedeemedIterator struct {
	Event *DSCEngineCollateralRedeemed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DSCEngineCollateralRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DSCEngineCollateralRedeemed)
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
		it.Event = new(DSCEngineCollateralRedeemed)
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
func (it *DSCEngineCollateralRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DSCEngineCollateralRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DSCEngineCollateralRedeemed represents a CollateralRedeemed event raised by the DSCEngine contract.
type DSCEngineCollateralRedeemed struct {
	RedeemedFrom common.Address
	RedeemedTo   common.Address
	Token        common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCollateralRedeemed is a free log retrieval operation binding the contract event 0xb292db12b8dfa23b760ee5e3bd1d3be184cd8f0eeeacd27ce120d4de4e4721f6.
//
// Solidity: event CollateralRedeemed(address indexed redeemedFrom, address indexed redeemedTo, address indexed token, uint256 amount)
func (_DSCEngine *DSCEngineFilterer) FilterCollateralRedeemed(opts *bind.FilterOpts, redeemedFrom []common.Address, redeemedTo []common.Address, token []common.Address) (*DSCEngineCollateralRedeemedIterator, error) {

	var redeemedFromRule []interface{}
	for _, redeemedFromItem := range redeemedFrom {
		redeemedFromRule = append(redeemedFromRule, redeemedFromItem)
	}
	var redeemedToRule []interface{}
	for _, redeemedToItem := range redeemedTo {
		redeemedToRule = append(redeemedToRule, redeemedToItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DSCEngine.contract.FilterLogs(opts, "CollateralRedeemed", redeemedFromRule, redeemedToRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &DSCEngineCollateralRedeemedIterator{contract: _DSCEngine.contract, event: "CollateralRedeemed", logs: logs, sub: sub}, nil
}

// WatchCollateralRedeemed is a free log subscription operation binding the contract event 0xb292db12b8dfa23b760ee5e3bd1d3be184cd8f0eeeacd27ce120d4de4e4721f6.
//
// Solidity: event CollateralRedeemed(address indexed redeemedFrom, address indexed redeemedTo, address indexed token, uint256 amount)
func (_DSCEngine *DSCEngineFilterer) WatchCollateralRedeemed(opts *bind.WatchOpts, sink chan<- *DSCEngineCollateralRedeemed, redeemedFrom []common.Address, redeemedTo []common.Address, token []common.Address) (event.Subscription, error) {

	var redeemedFromRule []interface{}
	for _, redeemedFromItem := range redeemedFrom {
		redeemedFromRule = append(redeemedFromRule, redeemedFromItem)
	}
	var redeemedToRule []interface{}
	for _, redeemedToItem := range redeemedTo {
		redeemedToRule = append(redeemedToRule, redeemedToItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _DSCEngine.contract.WatchLogs(opts, "CollateralRedeemed", redeemedFromRule, redeemedToRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DSCEngineCollateralRedeemed)
				if err := _DSCEngine.contract.UnpackLog(event, "CollateralRedeemed", log); err != nil {
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

// ParseCollateralRedeemed is a log parse operation binding the contract event 0xb292db12b8dfa23b760ee5e3bd1d3be184cd8f0eeeacd27ce120d4de4e4721f6.
//
// Solidity: event CollateralRedeemed(address indexed redeemedFrom, address indexed redeemedTo, address indexed token, uint256 amount)
func (_DSCEngine *DSCEngineFilterer) ParseCollateralRedeemed(log types.Log) (*DSCEngineCollateralRedeemed, error) {
	event := new(DSCEngineCollateralRedeemed)
	if err := _DSCEngine.contract.UnpackLog(event, "CollateralRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
