// SPDX-License-Identifier: MIT

/**
 * 不变性文件将包含我们系统中的所有不变性
 * 我们的不变性是什么？
 * 1.dsc总量应该小于抵押品的总价值
 * 2.getter或view函数不应该回滚 -》这是一种永恒的不变性
 *
 */
pragma solidity ^0.8.18;

import {Test, console} from "forge-std/Test.sol";
import {StdInvariant} from "forge-std/StdInvariant.sol";
import {DeployDsc} from "../../script/DeployDsc.s.sol";
import {DSCEngine} from "../../src/DSCEngine.sol";
import {DecentralizedStableCoin} from "../../src/DecentralizedStable.sol";
import {HelperConfig} from "../../script/HelperConfig.s.sol";
import {IERC20} from "../../openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import {Handler} from "./Handler.t.sol";

contract InvariantsTest is StdInvariant, Test {
    DeployDsc deployer;
    DSCEngine dsce;
    DecentralizedStableCoin dsc;
    HelperConfig helperConfig;
    address weth;
    address wbtc;
    Handler public handler;

    function setUp() public {
        deployer = new DeployDsc();
        (dsce, dsc, helperConfig) = deployer.run();
        (,, weth, wbtc,) = helperConfig.activeNetworkConfig();
        //告诉foundry这是我们要测试的合约
        // targetContract(address(dsce));
        handler = new Handler(dsce, dsc);
        targetContract(address(handler));
    }
    //现在调用了handler的函数，里面只有一个存入抵押品的函数

    function invariant_protocoMustHaveMoreValyeThanTotalSupply() public view {
        //获取dsc总量
        uint256 totalSupply = dsc.totalSupply();
        //获取weth和wbtc数量
        uint256 totalwethDeposited = IERC20(weth).balanceOf(address(dsce));
        uint256 totalwbtcDeposited = IERC20(wbtc).balanceOf(address(dsce));
        //转换为价值
        uint256 wethValue = dsce.getUSDValue(weth, totalwethDeposited);
        uint256 wbtcValue = dsce.getUSDValue(wbtc, totalwbtcDeposited);

        console.log("wethValue: ", wethValue);
        console.log("wbtcValue: ", wbtcValue);
        console.log("totalSupply: ", totalSupply);
        console.log("totalwethDeposited: ", handler.count());
        assert(wethValue + wbtcValue >= totalSupply);
    }
}
