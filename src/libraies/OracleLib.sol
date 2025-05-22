// SPDX-License-Identifier: MIT

pragma solidity ^0.8.18;

import {AggregatorV3Interface} from
    "../../chainlink-brownie-contracts/contracts/src/v0.8/shared/interfaces/AggregatorV2V3Interface.sol";
/**
 * @title OracleLib
 * @author Ahl
 * @notice 这个库用于检查 Chainlink 预言机中是否存在过时的数据
 * 如果价格过期 函数将回滚并使dsce引擎无法使用 如果价格过期 我们应该立即停止，不让任何事情发生
 * 如果chainlink网络瘫痪，你在协议中锁定的大量资金 那就太糟糕了
 *
 */

library OracleLib {
    error OracleLib__StalePriceCheck();
    //硬编码的价格源心跳  即过期时间

    uint256 constant STALE_PRICE_DELAY = 3 hours;
    //检查过期价格的函数

    function stalePriceCheck(AggregatorV3Interface priceFeed)
        public
        view
        returns (uint80, int256, uint256, uint256, uint80)
    {
        (uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound) =
            priceFeed.latestRoundData();
        //获得价格源被更后的秒数
        uint256 secondsSince = block.timestamp - updatedAt;
        if (secondsSince > STALE_PRICE_DELAY) {
            revert OracleLib__StalePriceCheck();
        }
        return (roundId, answer, startedAt, updatedAt, answeredInRound);
    }
}
