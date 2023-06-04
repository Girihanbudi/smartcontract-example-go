# smartm2m-technical-test

## Brief Description
This repository contain a sample API to manage file with smart contract. Visit [Postman documentation](https://documenter.getpostman.com/view/7070614/2s93sW9vXb) to seek a sample API request definition for test.

 
## Download
Download this repository using git

```bash
  git clone https://github.com/Girihanbudi/smartm2m-technical-test.git
```

## Installation

1. Install all required dependencies in project using command :

```bash
  go mod download
```

2. First, we need to make a connection to Ethereum Blockchain Environment. We can use ganache for local testing. If not exist yet, we can install it using provided setting in docker-compose. To install, run :

```bash
  make environments
```

3. After ganache running, inspect the log and it will show 10 account address and it private keys. Pick any private key, remove '0x' infront of it and paste the value printed in the log to the 'smartcontractholder' field in config.yaml file. After that, we need to deploy smart contract to Ethereum Blockchain Environment. Run :

```bash
  make contractdeployment
```

4. After smart contract successfully deployed, run the application :

```bash
  make run
```

or we can use docker to run the application as container

```bash
  make rundocker
```