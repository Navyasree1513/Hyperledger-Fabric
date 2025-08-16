# Hyperledger Fabric Asset Management

## Project Description
This project implements a blockchain-based system using Hyperledger Fabric to manage and track assets for a financial institution. The system supports creating assets, updating asset values, querying the world state, and retrieving asset transaction history. The assets represent accounts with specific attributes such as DEALERID, MSISDN, MPIN, BALANCE, STATUS, TRANSAMOUNT, TRANSTYPE, and REMARKS.

## Prerequisites
Before you begin, ensure you have met the following requirements:
- **Docker**: Ensure Docker is installed and running on your machine.
- **Docker Compose**: Required for managing multi-container Docker applications.
- **Go**: Install Go if you are using Golang for the smart contract.
- **Node.js**: Install Node.js if you are using JavaScript for the REST API.
- **Git**: Required for cloning the repository.

##  Installation

##  Step 1: Clone the Repository
Clone this repository to your local machine:
```bash
git clone https://github.com/yourusername/hyperledger-fabric-asset-management.git
cd hyperledger-fabric-asset-management
```

##  Step 2: Install Hyperledger Fabric Binaries
- Download the Hyperledger Fabric binaries from the Hyperledger Fabric Releases page.
- Extract the binaries and move them to a directory in your PATH:
```bash
tar -xvf hyperledger-fabric-linux-amd64.tar.gz
sudo mv bin/* /usr/local/bin/
```

##  Step 3: Smart Contract Development
## Create the Smart Contract
 1. Navigate to the Chaincode Directory:
- Go to the asset-transfer-custom/chaincode-go directory in your cloned repository:
```bash
cd fabric-samples/asset-transfer-basic
```
 2. Create a New File for the Smart Contract:
 - Create a new file named chaincode.go:
 ```bash
 touch chaincode.go
 ```

##  Step 4: Set Up the Test Network
1. Navigate to the test network directory:
```bash
cd fabric-samples/test-network
```
2. Start the test network:
```bash
./network.sh up
```
3. Create a channel:
```bash
./network.sh createChannel
```
4. Deploy the chaincode:
```bash
./network.sh deployCC
```

##  Step 5: Test the Smart Contract
- Open a new terminal window to interact with the deployed chaincode.
- Set the environment variables for the peer CLI:
```bash
export PATH=$PATH:$(pwd)/../bin
export FABRIC_CFG_PATH=$(pwd)/../config/
```
## Create an Asset:
Use the following command to create an asset:
```bash
peer chaincode invoke -o localhost:7050 --channelID mychannel --name asset-transfer-basic -c '{"function":"CreateAsset","Args":["dealer1","1234567890","1234","1000","active","500","credit","Initial deposit"]}'
```
## Query the Asset:
Use the following command to query the asset:
```bash
peer chaincode query -o localhost:7050 --channelID mychannel --name asset-transfer-basic -c '{"function":"QueryAsset","Args":["dealer1"]}'
```

##  Step 6: Set Up the REST API Project
1. Create a New Directory for the REST API:
- Navigate to your project root directory and create a new directory for the REST API:
```bash
mkdir rest-api
cd rest-api
```
2. Initialize a New Node.js Project:
```bash
npm init -y
```
3. Install Required Packages:
```bash
npm install express body-parser fabric-network
```

## Step 7: Create the REST API Code
## Create the Main Application File:
1. Create a new file named server.js:
```bash
touch server.js
```
2. Implement the REST API Code:
- Open server.js in your preferred text editor and add the code

## Step 8: Running the REST API
1. Start the REST API Server:
- In the rest-api directory, run the following command to start the server:
```bash
node server.js
```
2. Test the API Endpoints:
- Use Postman or any API testing tool to test the following endpoints:
- Create Asset: Send a POST request to http://localhost:3000/createAsset with the asset details in the request body.
- Query Asset: Send a GET request to http://localhost:3000/queryAsset/{dealerId} to retrieve asset details.

## 🐳 Dockerize REST API
```bash
docker build -t fabric-rest .
docker run -p 3000:3000 fabric-rest
```
