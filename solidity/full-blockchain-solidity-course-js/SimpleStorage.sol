// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

// contracts can be thought of as a class in OOP
contract SimpleStorage {
    uint256 favoriteNumber; //internal by default
    //People public person = People({favoriteNumber: 2, name: "Mike"});
    People[] public people;

    mapping(string => uint256) public nameToFavoriteNumber;

    struct People {
        uint256 favoriteNumber;
        string name;
    }

    function store(uint256 _favoriteNumber) public virtual {
        favoriteNumber = _favoriteNumber;
        //retrieve(); // calling a view function is free unless we call it internally from another func
    }

    function retrieve() public view returns (uint256) {
        return favoriteNumber;
    }

    // calldata, memory, storage
    // calldata = temporary vars that can't be modified
    // memory   = tempororary vars that can be modified
    // storage  = permanent vars that can be modified
    function addPerson(string memory _name, uint256 _favoriteNumber) public {
        //People memory newPerson = People({favoriteNumber: _favoriteNumber, name: _name});
        //people.push(newPerson);

        people.push(People(_favoriteNumber, _name)); // short-hand notation
        nameToFavoriteNumber[_name] = _favoriteNumber;
    }
}
