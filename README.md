# Hyperledger Fabric Asset Management

This repository contains an internship assignment project implementing an **asset management system** on **Hyperledger Fabric**.

## 🚀 Features
- Fabric test network setup
- Smart contract in Go (Golang)
- REST API (Node.js + Fabric SDK)
- Dockerized REST API

## 🧩 Asset Schema
- DEALERID
- MSISDN
- MPIN
- BALANCE
- STATUS
- TRANSAMOUNT
- TRANSTYPE
- REMARKS

## 🔹 Level 1: Setup Test Network
```bash
cd fabric-samples/test-network
./network.sh up createChannel -ca
```

## 🔹 Level 2: Deploy Chaincode
```bash
./network.sh deployCC -ccn assetcc -ccp ../chaincode/asset-transfer-custom -ccl go
```

Invoke:
```bash
peer chaincode invoke -C mychannel -n assetcc -c '{"function":"CreateAsset","Args":["D001","999888777","1234","5000","ACTIVE","0","NA","Initial balance"]}'
```

## 🔹 Level 3: REST API
```bash
cd rest-api
npm install
node server.js
```

API Endpoints:
- POST /create
- GET /read/:id

## 🐳 Dockerize REST API
```bash
docker build -t fabric-rest .
docker run -p 3000:3000 fabric-rest
```
