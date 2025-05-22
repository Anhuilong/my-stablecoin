// SPDX-License-Identifier: MIT

pragma solidity ^0.8.18;

import {DecentralizedStableCoin} from "./DecentralizedStable.sol";
import {ReentrancyGuard} from "../openzeppelin-contracts/contracts/utils/ReentrancyGuard.sol";
import {IERC20} from "../openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import {AggregatorV3Interface} from
    "../chainlink-brownie-contracts/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";
import {OracleLib} from "./libraies/OracleLib.sol";
/**
 * @title DSC引擎
 * @author Anhuilong
 * 这个系统被设计的尽可能简洁，并且让代币保持1美元的锚定
 *
 * 我们的系统应该始终是过度抵押的 任何时候所有抵押品的价值都应该大于所有DSC的价值
 *
 * 这个稳定币是外部抵押品的，算法稳定币，类似于DAI 没有治理，没有费用，并且只支持WETH 和WBTC
 * @notice 这个合约是一个DSC系统的核心，它处理了所有有关铸币和赎回的逻辑，以及有关存入和提取抵押品的逻辑
 * @notice 这个合约受到了MakerDAO的启发，但是它是一个简化的版本
 */

contract DSCEngine is ReentrancyGuard {
    /////////////////////////
    //////////Errors/////////
    /////////////////////////
    error DSCEngine__NeedsMoreThanZero();
    error DSCEngine__TokenAddressAndPriceFeedAddressMustBeSameLength();
    error DSCEngine__OnlyAllowWETHAndWBTC();
    error DSCEngine__TransferFailed();
    error DSCEngine__HealthFactorIsBroken();
    error DSCEngine__HealthFactorIsNotBroken();
    error DSCEngine__HealthFactorIsNotImproved();

    using OracleLib for AggregatorV3Interface;
    ////////////////////////////
    //////////State Variables/////////
    ////////////////////////////

    uint256 private constant ADDITIONAL_FEED_PRECISION = 1e10;
    uint256 private constant PRECISION = 1e18;
    uint256 private constant LIQUIDATION_THRESHOLD = 50; //50%
    mapping(address token => address priceFeed) private s_PriceFeeds; //tokenToPriceFeed
    DecentralizedStableCoin private immutable i_dsc;
    mapping(address user => mapping(address token => uint256 amount)) private s_CollateralDeposited; //抵押品余额
    mapping(address user => uint256 amountDscMinted) private s_DscMinted;
    address[] private s_collateralTokens;
    uint256 private constant MIN_HEALTH_FACTOR = 1e18;
    uint256 private constant LIQUIDATION_BONUS = 10; //10%
    /////////////////////////
    //////////Events/////////
    /////////////////////////

    event CollateralDeposited(address indexed user, address indexed token, uint256 amount);
    event CollateralRedeemed(
        address indexed redeemedFrom, address indexed redeemedTo, address indexed token, uint256 amount
    );
    ////////////////////////////
    //////////Modifiers/////////
    ////////////////////////////
    /**
     * 存入的数量不能为0
     */

    modifier moreThanZero(uint256 _amount) {
        if (_amount == 0) {
            revert DSCEngine__NeedsMoreThanZero();
        }
        _;
    }

    /**
     * 只能存入以太坊和比特币
     */
    modifier isAllowedToken(address _tokenAddress) {
        if (s_PriceFeeds[_tokenAddress] == address(0)) {
            revert DSCEngine__OnlyAllowWETHAndWBTC();
        }
        _;
    }

    ////////////////////////////
    //////////Functions/////////
    ////////////////////////////
    constructor(address[] memory tokenAddresses, address[] memory priceFeedAddress, address dscAddress) {
        if (tokenAddresses.length != priceFeedAddress.length) {
            revert DSCEngine__TokenAddressAndPriceFeedAddressMustBeSameLength();
        }
        // ETH / USD BTC /USD

        for (uint256 i = 0; i < tokenAddresses.length; i++) {
            //将对应的价格源地址映射到对应的token地址上
            s_PriceFeeds[tokenAddresses[i]] = priceFeedAddress[i];
            //将对应的token地址添加到抵押品列表中
            s_collateralTokens.push(tokenAddresses[i]);
        }

        i_dsc = DecentralizedStableCoin(dscAddress);
    }

    //////////////////////////////
    ///External functions/////////
    //////////////////////////////
    /**
     *
     * @param tokenCollateralAddress 抵押品地址
     * @param amountCollateral 抵押品数量
     * @param amountDscToMint 铸造的DSC数量
     * @notice 存入抵押品的同时铸造DSC
     */
    function depositCollateralAndMintDsc(
        address tokenCollateralAddress,
        uint256 amountCollateral,
        uint256 amountDscToMint
    ) external {
        depositCollateral(tokenCollateralAddress, amountCollateral);
        MintDsc(amountDscToMint);
    }

    //存入抵押品
    function depositCollateral(address tokenCollateralAddress, /* 抵押品地址*/ uint256 amountCollateral /* 抵押品数量*/ )
        public
        moreThanZero(amountCollateral)
        isAllowedToken(tokenCollateralAddress)
        nonReentrant /* 不可重入修饰符 需使用openzeppelin*/
    {
        s_CollateralDeposited[msg.sender][tokenCollateralAddress] += amountCollateral;
        emit CollateralDeposited(msg.sender, tokenCollateralAddress, amountCollateral);
        bool success = IERC20(tokenCollateralAddress).transferFrom(msg.sender, address(this), amountCollateral);
        if (!success) {
            revert DSCEngine__TransferFailed();
        }
    }

    //铸造DSC
    function MintDsc(uint256 amountDscToMint) public moreThanZero(amountDscToMint) nonReentrant {
        //追踪每个人铸造了多少DSC
        s_DscMinted[msg.sender] += amountDscToMint;
        //判断一下铸造的DSC是否超过了抵押品的价值
        _revertIfHealthFactorIsBroken(msg.sender);
        //铸造DSC
        bool success = i_dsc.mint(msg.sender, amountDscToMint);
        if (!success) {
            revert DSCEngine__TransferFailed();
        }
    }

    /**
     *
     * @notice 赎回抵押品
     * 赎回后检查健康因子是否被破坏
     */
    function redeemCollateral(address tokenCollateralAddress, uint256 amountCollateral)
        public
        moreThanZero(amountCollateral)
        nonReentrant
    {
        _redeemCollateral(tokenCollateralAddress, amountCollateral, msg.sender, msg.sender);
        //检查健康因子是否被破坏
        _revertIfHealthFactorIsBroken(msg.sender);
    }

    //人们可以销毁自己的DSC 原因是如果知道自己的抵押品已经不足了，那么可以销毁自己所持有的DSC，从而减轻债务，但是不退还抵押品
    function burnDsc(uint256 amount) public moreThanZero(amount) nonReentrant {
        _burnDsc(amount, msg.sender, msg.sender);
        _revertIfHealthFactorIsBroken(msg.sender);
    }

    //赎回抵押品的同时销毁dsc
    function redeemCollateralForDsc(address tokenCollateralAddress, uint256 amountCollateral)
        external
        moreThanZero(amountCollateral)
        nonReentrant
    {
        //销毁
        burnDsc(amountCollateral);
        ///赎回
        redeemCollateral(tokenCollateralAddress, amountCollateral);
    }

    /*其他人销毁你的DSC。我们应该有一个阈值，如果你的抵押品价值低于这个阈值，那么其他人就可以通过偿还你的DSC来获得你的抵押品 
      比如说这个阈值是50% 你抵押了100美元的以太坊，那么你可以铸造50美元的DSC，如果你的抵押品价值低于75美元，比如74美元，那么其他人就可以通过偿还你的DSC来获得你的抵押品。
      这样他就赚到了24美金
    */
    /**
     *
     * @param collateral 将要清算的抵押品地址
     * @param user 将要清算的用户地址
     * @param debtToCover 需要偿还的债务 要销毁的dsc数量
     * @notice 清算
     * @notice 只要你改善了他们的健康因子，你将获得清算奖金，也就是拿走用户的资金
     * @notice 协议的理想状态是 抵押率200%左右的超额抵押下才能正常工作 即铸造了50美元的dsc 那么抵押品价值要100美元，如果抵押率低于150% 就会启动清算
     * @notice 如果协议的抵押率达到100%或更低，那么我们将无法激励清算员 比如铸造了50美元的dsc 而抵押品的价值就只剩20美元，没有人会支付50美元来获得20美元。
     *
     */
    function liquidate(address collateral, address user, uint256 debtToCover)
        external
        moreThanZero(debtToCover)
        nonReentrant
    {
        //检查用户的健康因子 如果健康因子大于1 则不能清算
        uint256 startingUserHealthFactor = _healthFactor(user);
        if (startingUserHealthFactor >= MIN_HEALTH_FACTOR) {
            revert DSCEngine__HealthFactorIsNotBroken();
        }
        //偿还债务 并获取抵押品
        //假设这个用户拥有140美元的抵押品 ，拥有100美元的dsc，偿还债务为100美元
        //需要用ETH偿还债务 那么就需要知道100美元的dsc是多少ETH
        uint256 tokenAmountFromDebtCovered = getTokenAmountFromUsd(collateral, debtToCover);
        //给予清算者10%的奖励
        uint256 bonusCollateral = (tokenAmountFromDebtCovered * LIQUIDATION_BONUS) / 100;
        uint256 totalCollateralToRedeem = tokenAmountFromDebtCovered + bonusCollateral;
        _redeemCollateral(collateral, totalCollateralToRedeem, user, msg.sender);
        //销毁dsc
        _burnDsc(debtToCover, user, msg.sender);
        // 如果没有改善健康因子 则回滚
        uint256 endingUserHealthFactor = _healthFactor(user);
        if (endingUserHealthFactor <= startingUserHealthFactor) {
            revert DSCEngine__HealthFactorIsNotImproved();
        }
        // 如果因为这个过程，清算者的健康因子被破坏 我们也应该回滚
        _revertIfHealthFactorIsBroken(msg.sender);
    }
    ///////////////////////////////////////////////////
    //////////Private & Internal View functions/////////
    ///////////////////////////////////////////////////
    /**
     *
     * @param amountDscToBurn 需要销毁的DSC数量
     * @param onBehalfOf 代币的接收者
     * @param dscFrom 代币的来源
     * @notice 像这样的低级内部函数 不要调用
     */

    function _burnDsc(uint256 amountDscToBurn, address onBehalfOf, address dscFrom) private {
        s_DscMinted[onBehalfOf] -= amountDscToBurn;
        //先将代币转回铸币地址 然后在销毁
        bool sucess = i_dsc.transferFrom(dscFrom, address(this), amountDscToBurn);
        if (!sucess) {
            revert DSCEngine__TransferFailed();
        }
        i_dsc.burn(amountDscToBurn);
    }
    /**
     *
     * @param tokenCollateralAddress 抵押品地址
     * @param amountCollateral 抵押品数量
     * @param from 抵押品的来源
     * @param to 抵押品的接收者
     */

    function _redeemCollateral(address tokenCollateralAddress, uint256 amountCollateral, address from, address to)
        private
    {
        //在新版的solidity中 如果余额只有100 而你要取出1000 会直接回滚。
        s_CollateralDeposited[from][tokenCollateralAddress] -= amountCollateral;
        //当状态发生改变时，触发事件
        emit CollateralRedeemed(from, to, tokenCollateralAddress, amountCollateral);
        bool success = IERC20(tokenCollateralAddress).transfer(to, amountCollateral);
        if (!success) {
            revert DSCEngine__TransferFailed();
        }
    }
    /**
     *
     * @param user 用户地址
     * @notice 获取用户铸造的DSC总量和抵押品的总价值
     */

    function _getAccountInformation(address user)
        private
        view
        returns (uint256 totalDscMinted, uint256 collataerValueInUsd)
    {
        totalDscMinted = s_DscMinted[user];
        collataerValueInUsd = getAccountCollateralValue(user);
        return (totalDscMinted, collataerValueInUsd);
    }
    /**
     *
     * @notice 返回健康因子
     */

    function _healthFactor(address user) private view returns (uint256 healthFactor) {
        //创建一个函数   获取铸造的DSC总量  获取总抵押品的价值
        (uint256 totalDscMinted, uint256 collataerValueInUsd) = _getAccountInformation(user);
        if (totalDscMinted == 0) {
            return type(uint256).max;
        }

        uint256 collateralAdjustedForThreshold = (collataerValueInUsd * LIQUIDATION_THRESHOLD) / 100;
        return (collateralAdjustedForThreshold * PRECISION) / totalDscMinted;
    }
    /**
     *
     * @param user 用户地址
     * @notice 检查健康因子，如果健康因子小于1，就会被清算
     */

    function _revertIfHealthFactorIsBroken(address user) internal view {
        uint256 healthFactor = _healthFactor(user);
        if (healthFactor < MIN_HEALTH_FACTOR) {
            revert DSCEngine__HealthFactorIsBroken();
        }
    }

    /////////////////////////////////////////
    //////////public  View functions/////////
    /////////////////////////////////////////
    /**
     *
     * @param token 代币地址
     * @param usdAmountInWei 美元数量
     * @notice 根据给定代币的美元价值，来获取代币数量   美元价值/代币价格
     */
    function getTokenAmountFromUsd(address token, uint256 usdAmountInWei) public view returns (uint256) {
        AggregatorV3Interface priceFeed = AggregatorV3Interface(s_PriceFeeds[token]);
        (, int256 price,,,) = priceFeed.stalePriceCheck();
        return (usdAmountInWei * PRECISION) / (uint256(price) * ADDITIONAL_FEED_PRECISION);
    }
    /**
     *
     * @param user 用户地址
     * @notice 获取用户抵押品的价值（多个抵押代币的总价值）
     */

    function getAccountCollateralValue(address user) public view returns (uint256 totalCollateralValueInUsd) {
        //我们只需要循环遍历每个抵押代币，获取用户存入的数量，然后乘以每个代币的价格，以获取价值。
        for (uint256 i = 0; i < s_collateralTokens.length; i++) {
            address token = s_collateralTokens[i];
            uint256 amount = s_CollateralDeposited[user][token];
            totalCollateralValueInUsd += getUSDValue(token, amount);
        }
        return totalCollateralValueInUsd;
    }
    /**
     *
     * @param token 代币地址
     * @param amount 代币数量
     * @notice 获取代币的实际价值 价格*数量
     */

    function getUSDValue(address token, uint256 amount) public view returns (uint256) {
        AggregatorV3Interface priceFeed = AggregatorV3Interface(s_PriceFeeds[token]);
        (, int256 price,,,) = priceFeed.latestRoundData();
        return ((uint256(price) * amount) * ADDITIONAL_FEED_PRECISION) / PRECISION;
    }

    function getCollateralBalanceOf(address user, address token) public view returns (uint256) {
        return s_CollateralDeposited[user][token];
    }

    function getAccountInformation(address user)
        public
        view
        returns (uint256 totalDscMinted, uint256 collateralValueInUsd)
    {
        return _getAccountInformation(user);
    }

    function getSDscMinted(address user) public view returns (uint256) {
        return s_DscMinted[user];
    }

    function getCollateralTokenPriceFeed(address token) external view returns (address) {
        return s_PriceFeeds[token];
    }
    //10000000000000000000000

    function getHealthFactor(address user) public view returns (uint256) {
        return _healthFactor(user);
    }

    function getCollateralTokens() public view returns (address[] memory) {
        return s_collateralTokens;
    }
}
