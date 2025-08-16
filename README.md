# Hyperledger Fabric Asset Management

# Hyperledger Fabric Asset Management System
## Project Description
This project implements a blockchain-based system using Hyperledger Fabric to manage and track assets for a financial institution. The system supports creating assets, updating asset values, querying the world state, and retrieving asset transaction history. The assets represent accounts with specific attributes such as DEALERID, MSISDN, MPIN, BALANCE, STATUS, TRANSAMOUNT, TRANSTYPE, and REMARKS.

## Prerequisites
Before you begin, ensure you have met the following requirements:
- Docker: Ensure Docker is installed and running on your machine.
- Docker Compose: Required for managing multi-container Docker applications.
- Go: Install Go if you are using Golang for the smart contract.
- Node.js: Install Node.js if you are using JavaScript for the REST API.
- Git: Required for cloning the repository.

## Installation

### Step 1: Clone the Repository
Clone this repository to your local machine:
```bash
git clone https://github.com/yourusername/hyperledger-fabric-asset-management.git
cd hyperledger-fabric-asset-management
