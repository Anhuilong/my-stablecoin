// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {Test, console} from "forge-std/Test.sol";
import {DSCEngine} from "../../src/DSCEngine.sol";
import {DecentralizedStableCoin} from "../../src/DecentralizedStable.sol";
import {ERC20Mock} from "../../test/mocks/ERC20Mock.sol";
import {MockV3Aggregator} from "../mocks/MockV3Aggregator.sol";

contract Handler is Test {
    DSCEngine dsce;
    DecentralizedStableCoin dsc;
    ERC20Mock weth;
    ERC20Mock wbtc;
    uint256 MAX = type(uint96).max;
    uint256 public count;
    address[] public users;
    MockV3Aggregator wethPriceFeed;

    constructor(DSCEngine _dsce, DecentralizedStableCoin _dsc) {
        dsce = _dsce;
        dsc = _dsc;
        address[] memory tokenAddresses = dsce.getCollateralTokens();
        weth = ERC20Mock(tokenAddresses[0]);
        wbtc = ERC20Mock(tokenAddresses[1]);
        wethPriceFeed = MockV3Aggregator(dsce.getCollateralTokenPriceFeed(address(weth)));
    }
    // 嘿！你只能存入有效的抵押品 而不是将任何地址传递给它  模糊测试会模拟随机的地址传入

    // function depositCollatera(address collateral, uint256 amountCollateral) public {
    //     dsce.depositCollateral(collateral, amountCollateral);
    // }

    //现在我们不使用随机地址 而是随机选择weth或wbtc 并且给dsce合约铸造token和授权使用token
    //然后不管我们多频繁的调用存入抵押品，不管我们存入了多少抵押品 我们都不会使InvariantsTest中的断言失败（抵押品的价值>dsc数量）

    function depositCollatera(uint256 collateralSeed, uint256 amountCollateral) public {
        ERC20Mock collateral = _getCollateralFromSeed(collateralSeed);
        amountCollateral = bound(amountCollateral, 1, MAX);
        vm.startPrank(msg.sender);
        collateral.mint(msg.sender, amountCollateral);
        collateral.approve(address(dsce), amountCollateral);
        dsce.depositCollateral(address(collateral), amountCollateral);
        vm.stopPrank();
        users.push(msg.sender);
    }
    //赎回抵押品

    function redeemCollateral(uint256 collateralSeed, uint256 amountCollateral) public {
        ERC20Mock collateral = _getCollateralFromSeed(collateralSeed);
        uint256 maxAmountCollateral = dsce.getCollateralBalanceOf(msg.sender, address(collateral));
        amountCollateral = bound(amountCollateral, 0, maxAmountCollateral);
        if (amountCollateral == 0) {
            return;
        }
        vm.startPrank(msg.sender);
        dsce.redeemCollateral(address(collateral), amountCollateral);
        vm.stopPrank();
    }

    //铸造dsc
    function mintDsc(uint256 amountDsc, uint256 addressSeed) public {
        console.log("mintDsc: ", amountDsc);
        if (users.length == 0) {
            return;
        }
        address msgsender = users[addressSeed % users.length];
        (uint256 totalDscMinted, uint256 collateralValueInUsd) = dsce.getAccountInformation(msgsender);

        int256 maxMintableDsc = (int256(collateralValueInUsd) / 2) - int256(totalDscMinted);
        if (maxMintableDsc < 0) {
            return;
        }
        amountDsc = bound(amountDsc, 1, uint256(maxMintableDsc));
        if (amountDsc == 0) {
            return;
        }
        vm.startPrank(msgsender);
        dsce.MintDsc(amountDsc);
        vm.stopPrank();
        count++;
    }
    //打破了不变测试
    // function updateCollateralPrice(uint96 newPrice) public {
    //     int256 newPriceInt = int256(uint256(newPrice));
    //     wethPriceFeed.updateAnswer(newPriceInt);
    // }

    //我们将从两个抵押品中随机选择
    function _getCollateralFromSeed(uint256 collateralSeed) private view returns (ERC20Mock) {
        if (collateralSeed % 2 == 0) {
            return weth;
        }
        return wbtc;
    }
}
