// SPDX-License-Identifier: MIT

pragma solidity ^0.8.18;

import {ERC20, ERC20Burnable} from "../openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import {Ownable} from "../openzeppelin-contracts/contracts/access/Ownable.sol";

/*
 *  @title 去中心化稳定币
 *  @dev 该合约是一个去中心化稳定币的合约
 *  @auther: Anhuilong
 *  外部抵押品： ETH BTC
 *  类型：算法稳定币
 *  去中心化的相对稳定：将它挂钩到美元
 *
 *  这个合约是打算由DSC引擎管理的，这个合约我打算使用ERC20实现的稳定币系统
 */
contract DecentralizedStableCoin is ERC20Burnable, Ownable {
    error DecentralizedStableCoin__MustBeMoreThanZero();
    error DecentralizedStableCoin__BurnAmountExceedsBalance();
    error DecentralizedStableCoin__NoZeroAddress();

    constructor() ERC20("Decentralized Stable Coin", "DSC") Ownable(msg.sender) {}

    //销毁
    function burn(uint256 _amount) public override onlyOwner {
        uint256 balance = balanceOf(msg.sender);
        //销毁的数量必须大于0
        if (_amount <= 0) {
            revert DecentralizedStableCoin__MustBeMoreThanZero();
        }
        //如果用户的余额小于销毁的数量。
        if (balance < _amount) {
            revert DecentralizedStableCoin__BurnAmountExceedsBalance();
        }
        super.burn(_amount);
    }

    //铸造
    function mint(address _to, uint256 _amount) external onlyOwner returns (bool) {
        //不让人们铸币到0地址，因为这种情况经常发生
        if (_to == address(0)) {
            revert DecentralizedStableCoin__NoZeroAddress();
        }

        if (_amount <= 0) {
            revert DecentralizedStableCoin__MustBeMoreThanZero();
        }
        _mint(_to, _amount);
        return true;
    }

    function burnFrom(address account, uint256 value) public override {
        _spendAllowance(account, _msgSender(), value);
        _burn(account, value);
    }
}
