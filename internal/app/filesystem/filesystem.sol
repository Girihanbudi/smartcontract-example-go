// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract FileSystem { 
  mapping (address  => File) files;

  struct File {
    string name;
    string file;
  }

  function addFile(string memory _name, string memory _file) public { 
    files[msg.sender] = File(_name, _file);
  }

  function deleteFile() public {
    delete files[msg.sender];
  }

  function getFile() public view returns (File memory) {
    return files[msg.sender];
  }
}