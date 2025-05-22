// SPDX-License-Identifier: MIT

pragma solidity ^0.8.18;

import {Script} from "../forge-std/Script.sol";
import {DSCEngine} from "../src/DSCEngine.sol";
import {DecentralizedStableCoin} from "../src/DecentralizedStable.sol";
import {HelperConfig} from "./HelperConfig.s.sol";

contract DeployDsc is Script {
    address[] public tokenAddresses;
    address[] public priceFeedAddress;

    function run() public returns (DSCEngine, DecentralizedStableCoin, HelperConfig) {
        HelperConfig helperConfig = new HelperConfig();
        (address wethUsdPriceFeed, address wbtcUsdPriceFeed, address weth, address wbtc, uint256 deployerKey) =
            helperConfig.activeNetworkConfig();
        tokenAddresses = [weth, wbtc];
        priceFeedAddress = [wethUsdPriceFeed, wbtcUsdPriceFeed];
        vm.startBroadcast(deployerKey);
        DecentralizedStableCoin dsc = new DecentralizedStableCoin();
        DSCEngine engine = new DSCEngine(tokenAddresses, priceFeedAddress, address(dsc));
        dsc.transferOwnership(address(engine));
        vm.stopBroadcast();
        return (engine, dsc, helperConfig);
    }
}
