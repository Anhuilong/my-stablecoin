// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

import {Test, console} from "forge-std/Test.sol";
import {DeployDsc} from "../../script/DeployDsc.s.sol";
import {DSCEngine} from "../../src/DSCEngine.sol";
import {DecentralizedStableCoin} from "../../src/DecentralizedStable.sol";
import {HelperConfig} from "../../script/HelperConfig.s.sol";
import {ERC20Mock} from "../mocks/ERC20Mock.sol";
import {MockV3Aggregator} from "../mocks/MockV3Aggregator.sol";
import {IERC20} from "../../openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";

contract DSCEngineTest is Test {
    DeployDsc deployer;
    DSCEngine engine;
    DecentralizedStableCoin dsc;
    HelperConfig helperConfig;
    address ethUsdPriceFeed;
    address btcUsdPriceFeed;
    address weth;
    address wbtc;
    address[] public tokenAddresses;
    address[] public priceFeedAddresses;
    address USER = makeAddr("USER");
    address USER2 = makeAddr("USER2");
    uint8 public constant DECIMALS = 8;
    int256 public constant ETH_USD_PRICE = 2000e8;
    int256 public constant BTC_USD_PRICE = 50000e8;
    uint256 public constant AMOUNT_COLLATERAL = 10 ether;
    uint256 public constant AMOUNT_DSC = 1e18;

    event CollateralDeposited(address indexed user, address indexed token, uint256 amount);

    function setUp() external {
        deployer = new DeployDsc();
        (engine, dsc, helperConfig) = deployer.run();
        (ethUsdPriceFeed, btcUsdPriceFeed, weth, wbtc,) = helperConfig.activeNetworkConfig();

        //为USER用户铸造一些WETH
        ERC20Mock(weth).mint(USER, AMOUNT_COLLATERAL);
    }

    ///////////////////////////////
    //// Constructor Tests //////
    ///////////////////////////////

    /* 测试当tokenAddresses和priceFeedAddresses长度不匹配时会回滚 */
    function test_RevertIfTokenAndPriceFeedLengthMismatch() public {
        tokenAddresses = [weth];
        priceFeedAddresses = [ethUsdPriceFeed, btcUsdPriceFeed];
        vm.expectRevert(DSCEngine.DSCEngine__TokenAddressAndPriceFeedAddressMustBeSameLength.selector);
        new DSCEngine(tokenAddresses, priceFeedAddresses, address(dsc));
    }
    ///////////////////////
    //// Price Tests //////
    ///////////////////////

    /* 测试获取美元价值是否正确 */
    function testGetUsdValue() public view {
        uint256 ethAmount = 15e18;
        uint256 expectedUsd = 30000e18;
        uint256 actualUsd = engine.getUSDValue(weth, ethAmount);
        assertEq(actualUsd, expectedUsd);
    }
    /* 测试通过给定的代币价值，来获取代币的数量 */

    function testGetTokenAmountFromUsd() public view {
        //价值100eth的代币
        uint256 amount = 100 ether;
        uint256 expectedWeth = 0.05 ether;
        uint256 actualWeth = engine.getTokenAmountFromUsd(weth, amount);
        assertEq(actualWeth, expectedWeth);
    }

    ///////////////////////
    ////存入抵押品测试 //////
    ///////////////////////

    /* 测试当抵押品数量为零时是否会回滚（revert） */
    function testRevertIfCollateralZero() public {
        vm.startPrank(USER);
        ERC20Mock(weth).approve(address(engine), AMOUNT_COLLATERAL);
        vm.expectRevert(DSCEngine.DSCEngine__NeedsMoreThanZero.selector);
        engine.depositCollateral(weth, 0);
        vm.stopPrank();
    }

    /* 测试只能 存入以太坊和比特币 */
    function testRevertIfCollateralIsNotAllowed() public {
        //创建一个ERC20模拟的随机代币  将它存入到抵押品
        ERC20Mock allowedToken = new ERC20Mock("Allowed", "ALW", USER, AMOUNT_COLLATERAL);
        vm.startPrank(USER);
        vm.expectRevert(DSCEngine.DSCEngine__OnlyAllowWETHAndWBTC.selector);
        engine.depositCollateral(address(allowedToken), AMOUNT_COLLATERAL);
        vm.stopPrank();
    }

    modifier depositCollateral(address tokenAddress, uint256 amount) {
        vm.startPrank(USER);
        ERC20Mock(tokenAddress).approve(address(engine), amount);
        engine.depositCollateral(tokenAddress, amount);
        vm.stopPrank();
        _;
    }
    /* 测试抵押品余额是否更新 */

    function test_UpdatesCollateralBalance() public depositCollateral(weth, AMOUNT_COLLATERAL) {
        vm.prank(USER);
        uint256 actualBalance = engine.getCollateralBalanceOf(USER, weth);
        assertEq(actualBalance, AMOUNT_COLLATERAL);
    }
    /* 测试可以存入抵押品并获取账户信息 */

    function testCanDepositCollateralAndGetAccountInfo() public depositCollateral(weth, AMOUNT_COLLATERAL) {
        //获取账户信息 dsc总量 以及抵押品总价值(美元) 检查数量是否正确
        (uint256 totalDscMinted, uint256 collataerValueInUsd) = engine.getAccountInformation(USER);
        assertEq(totalDscMinted, 0);
        //根据代币的美元价值来获取代币数量
        uint256 expectedCollateralValue = engine.getTokenAmountFromUsd(weth, collataerValueInUsd);
        assertEq(AMOUNT_COLLATERAL, expectedCollateralValue);
    }
    ///////////////////////
    ////铸造dsc测试 ////////
    ///////////////////////
    /* 测试账户余额 */

    function testDscMintedZero() public {
        vm.startPrank(USER);
        assertEq(engine.getSDscMinted(USER), 0);
        vm.stopPrank();
    }

    /* 测试铸造数量是否正确*/
    function testDscMinted() public depositCollateral(weth, AMOUNT_COLLATERAL) {
        vm.startPrank(USER);
        engine.MintDsc(10000e18);
        //10000000000000000000
        // assertEq(engine.getSDscMinted(USER), AMOUNT_DSC);ß
        vm.stopPrank();
    }

    /* 测试健康因子不足1 回滚*/
    function testRevertIfHealthFactorIsBroken() public {
        vm.startPrank(USER);
        vm.expectRevert(DSCEngine.DSCEngine__HealthFactorIsBroken.selector);
        engine.MintDsc(AMOUNT_DSC);
        vm.stopPrank();
    }

    ///////////////////////
    ////赎回抵押品测试 //////
    ///////////////////////
    /* 取出抵押品测试 */
    /**
     *  这里测试出的问题，如果我们只是存入了抵押品之后，想要赎回抵押品，中间并没有铸造代币，那么我们在计算健康因子时，totalDscMinted（DSC数量）会是0，因为除数是0，所以会报错。
     *  解决方案：在计算健康因子时，应该判断totalDscMinted是否为0，如果为0，则直接返回1e18（健康因子正常）
     *  这里的代码已经在DSCEngine.sol中修改了
     */
    function testRevertIfWithdrawCollateralMoreThanBalance() public depositCollateral(weth, AMOUNT_COLLATERAL) {
        vm.startPrank(USER);
        engine.redeemCollateral(weth, 1 ether);
        vm.stopPrank();
    }

    ///////////////////////
    ////销毁DSC测试 ////////
    ///////////////////////

    /* 销毁DSC */
    function testBurnDsc() public depositCollateral(weth, AMOUNT_COLLATERAL) {
        vm.startPrank(USER);
        //这里需要把dsc还给engine合约，但是engine不能直接拿，你需要给他授权
        dsc.approve(address(engine), AMOUNT_DSC);
        engine.MintDsc(AMOUNT_DSC);
        engine.burnDsc(AMOUNT_DSC);
        assertEq(engine.getSDscMinted(USER), 0);
        vm.stopPrank();
    }

    ///////////////////////
    ////清算测试 ///////////
    ///////////////////////

    /* 测试健康因子大于1 回滚*/
    function testRevertIfHealthFactorGood() public depositCollateral(weth, AMOUNT_COLLATERAL) {
        vm.startPrank(USER);
        engine.MintDsc(AMOUNT_DSC);
        vm.expectRevert(DSCEngine.DSCEngine__HealthFactorIsNotBroken.selector);
        engine.liquidate(weth, USER, AMOUNT_COLLATERAL);
        vm.stopPrank();
    }

    /* 测试给予清算者债务10%的奖励*/

    function testLiquidate() public depositCollateral(weth, AMOUNT_COLLATERAL) {
        /*.给被清算人存入抵押品，铸造，使健康因子 >1*/
        vm.startPrank(USER);
        engine.MintDsc(9000e18);
        vm.stopPrank();

        /*2. mock 降低 ETH 价格，让 health factor < 1*/
        MockV3Aggregator ethUsdPriceFeed = MockV3Aggregator(engine.getCollateralTokenPriceFeed(weth));
        ethUsdPriceFeed.updateAnswer(1000e8);

        //给清算者一些eth
        deal(address(weth), USER, 10000 ether);

        //模拟清算者
        vm.startPrank(USER2);

        /* 给清算者存入抵押品 并铸造dsc */
        ERC20Mock(weth).mint(USER2, 100000 ether);
        ERC20Mock(weth).approve(address(engine), 10000 ether);
        engine.depositCollateral(weth, 10000 ether);
        engine.MintDsc(9000e18);

        /* 因为清算函数里面要销毁dsc，所以把dsc还给engine ，需要给engine授权，它才可以拿你的dsc */
        dsc.approve(address(engine), type(uint256).max);

        /* 获取清算者在清算之前的weth余额*/
        IERC20 i_weth = IERC20(weth);
        uint256 collateralBefore = i_weth.balanceOf(USER2);
        engine.liquidate(weth, USER, 9000e18);
        /* 获取清算者在清算之后的weth余额 如果清算后的余额大于清算前的余额证明得到了奖励*/
        uint256 collateralAfter = i_weth.balanceOf(USER2);
        assertGt(collateralAfter, collateralBefore);

        /* 有没有改变被清算人的健康因子*/
        uint256 userHealthAfter = engine.getHealthFactor(USER);
        assertGt(userHealthAfter, 1e18);

        /* 清算人自己的健康因子有没有被破坏*/
        uint256 liquidatorHealth = engine.getHealthFactor(USER2);
        assertGt(liquidatorHealth, 1e18);

        vm.stopPrank();
    }
}
