pragma solidity ^0.4.18;


contract GetMoney {
    uint totalBalance;
    bool allowance;
    address[] investors;
    mapping(address => uint) investorsbal;
    mapping(address => int) votres;

    modifier allow {
        for (uint i = 0; i < investors.length; i++) {
            if (investors[i] == msg.sender) {
                allowance = true;
                break;
            }
        }
        _;
    }

    function () public payable {
        updateTotalbalance();
        investorsbal[msg.sender] += msg.value;
    }
    
    function investPart() public allow payable returns(uint) {
        return investorsbal[msg.sender];
    }
    
    function vote(int res) public allow {
        if (res == 1) {
            votres[msg.sender] = res;
        } else if (res == -1) {
            votres[msg.sender] = res;
        }
    }
    
    function showVotres() public  payable allow  returns(int) {
        return votres[msg.sender];
    }
    
    function updateTotalbalance() internal {
        totalBalance += msg.value;
    }
}