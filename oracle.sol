// SPDX-License-Identifier: CC-BY-1.0
pragma solidity ^0.8.17;

contract GHOracle {
    uint256 queryId = 0;

    mapping(uint256 => bool) responses;

    event Query(uint256 indexed id, string prUrl, string ghUsername);
    event Response(uint256 indexed id, bool response);

    // Asks the backend the question
    // "was the PR at prURL merged and did it include commits from a user with ghUsername?"
    function query(string memory prUrl, string memory ghUsername) public {
        uint256 id = queryId;
        queryId += 1;

        emit Query(id, prUrl, ghUsername);
    }

    // Function the backend uses to respond to a query
    function respond(uint256 id, bool response) public {
        responses[id] = response;

        emit Response(id, response);
    }

}