pragma solidity ^0.4.18;


contract VotingBasic {
    uint256 public orgBalance; // Сумма, которая организация имеет
    address public orgAddress; // Адрес организации (назначает админ)
    mapping(address=>uint256) public investrosBal; // Баланс каждого инвестора
    address[] public investors; // Список всех, кто вложил деньги на контракт (инвестор)
    address public admin; // Админ, кто назначает организацию
    mapping (uint256=>bytes32) public proposalsWhy; // Конкретная причина (описание того, куда потратит организация),
    //  хранится по uint256 айди для каждой новой затраты  в виде хэша keccak256
    mapping (uint256=>uint256) public proposalsSum; // Конкретна сумма затраты, хранится по айди затраты
    mapping (uint256=>address) public proposalAddress; // Конкретный адрес каждой затраты 
    // на этот адрес перевод proposalsSum идет
    uint256 public proposalId; // Айди последней затраты
    mapping (uint256=>mapping(address=>bool)) proposalRes; // Результат голосования по затрате 
    // По айди затраты и адреса голосовальщика - получаем его результат голосования
    mapping (uint256=>int256) public investersWhoVotedRes;
}


contract VotingModificators is VotingBasic {

    modifier onlyOrg() {
        require(msg.sender == orgAddress);
        _;
    }

    modifier onlyInvestor() {
        bool res;
        for (uint256 i =0; i < investors.length; i++) {
            if (msg.sender == investors[i]) {
                res = true;
            }
        }
        require(res);
        _;
    }

    modifier onlyAdmin() {
        require(msg.sender == admin);
        _;
    }
}


contract VotingFunctions is VotingModificators {
    function setOrgAddress(address _orgAddress) public onlyAdmin returns(bool);
    function voteForProposal(uint256 _idProp, bool _vote) public onlyInvestor returns(bool);
    function countVotesOfProposal(uint256 _idProp) public returns(bool);
    function makeProposal(string _propWhy, uint256 _propSum, address _propAddress) public onlyOrg returns(bool);
    function updateTotalBalance(uint256 _value) internal;
    function setInvestor(address _invsetor) internal;
    function setInvestorPie(address _invsetor, uint256 _value) internal;
}


contract VoteMain is VotingFunctions {
    function VoteMain() public {
        admin = msg.sender;
    }

    function () public payable {
        updateTotalBalance(msg.value);
        setInvestor(msg.sender);
        setInvestorPie(msg.sender, msg.value);
    }
    
    function setOrgAddress(address _orgAddress) public onlyAdmin returns(bool) {
        orgAddress = _orgAddress;
    } 

    function makeProposal(string _propWhy, uint256 _propSum, address _propAddress) public onlyOrg returns(bool) {
        proposalId += 1;  
        proposalsWhy[proposalId] = keccak256(_propWhy);
        proposalsSum[proposalId] = _propSum;
        proposalAddress[proposalId] = _propAddress;
        return false;
    }

    function voteForProposal(uint256 _idProp, bool _vote) public onlyInvestor returns(bool) {
        proposalRes[_idProp][msg.sender] = _vote;
        if (_vote == true) {
            investersWhoVotedRes[_idProp] += 1; 
        } else {
            investersWhoVotedRes[_idProp] -= 1; 
        }
        return false;
    }
    
    function countVotesOfProposal(uint256 _idProp) public returns(bool) {
        if (investersWhoVotedRes[_idProp] > 0) {
            return true;
        } else {
            return false;
        }
    }
    // for test

    function showBalnAdr() public constant returns(uint256, address) {
        return (orgBalance, orgAddress);
    }

    function showInvestorslast() public constant returns(address) {
        return investors[investors.length];
    }

    function showProposal(uint256 _id) returns(bytes32, uint256, address) {
        return (proposalsWhy[_id], proposalsSum[_id], proposalAddress[_id]);
    }
    
    function showResofProposal(uint256 _id, address _voter) public constant returns(bool) {
        return proposalRes[_id][_voter];
    }
    // for test

    function updateTotalBalance(uint256 _value) internal {
        orgBalance += _value;
    }

    function setInvestor(address _invsetor) internal {
        investors.push(_invsetor);
    }

    function setInvestorPie(address _invsetor, uint256 _value) internal {
        investrosBal[_invsetor] = _value;
    }

}