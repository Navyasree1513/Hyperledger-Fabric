package main

import (
    "encoding/json"
    "fmt"
    "strconv"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Asset struct {
    DealerID   string  `json:"dealerId"`
    MSISDN     string  `json:"msisdn"`
    MPIN       string  `json:"mpin"`
    Balance    float64 `json:"balance"`
    Status     string  `json:"status"`
    TransAmt   float64 `json:"transAmount"`
    TransType  string  `json:"transType"`
    Remarks    string  `json:"remarks"`
}

type SmartContract struct {
    contractapi.Contract
}

func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, dealerID, msisdn, mpin, balance, status, transAmt, transType, remarks string) error {
    _, exists := s.AssetExists(ctx, dealerID)
    if exists {
        return fmt.Errorf("asset with DealerID %s already exists", dealerID)
    }

    bal, _ := strconv.ParseFloat(balance, 64)
    amt, _ := strconv.ParseFloat(transAmt, 64)

    asset := Asset{
        DealerID:  dealerID,
        MSISDN:    msisdn,
        MPIN:      mpin,
        Balance:   bal,
        Status:    status,
        TransAmt:  amt,
        TransType: transType,
        Remarks:   remarks,
    }

    assetJSON, _ := json.Marshal(asset)
    return ctx.GetStub().PutState(dealerID, assetJSON)
}

func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, dealerID string) (*Asset, error) {
    assetJSON, err := ctx.GetStub().GetState(dealerID)
    if err != nil || assetJSON == nil {
        return nil, fmt.Errorf("asset %s not found", dealerID)
    }
    var asset Asset
    _ = json.Unmarshal(assetJSON, &asset)
    return &asset, nil
}

func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, dealerID string, balance, status string) error {
    asset, err := s.ReadAsset(ctx, dealerID)
    if err != nil {
        return err
    }
    bal, _ := strconv.ParseFloat(balance, 64)
    asset.Balance = bal
    asset.Status = status

    assetJSON, _ := json.Marshal(asset)
    return ctx.GetStub().PutState(dealerID, assetJSON)
}

func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, dealerID string) (bool, error) {
    assetJSON, err := ctx.GetStub().GetState(dealerID)
    if err != nil {
        return false, err
    }
    return assetJSON != nil, nil
}

func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
    resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
    if err != nil {
        return nil, err
    }
    defer resultsIterator.Close()

    var assets []*Asset
    for resultsIterator.HasNext() {
        queryResponse, _ := resultsIterator.Next()
        var asset Asset
        _ = json.Unmarshal(queryResponse.Value, &asset)
        assets = append(assets, &asset)
    }
    return assets, nil
}

func (s *SmartContract) GetHistory(ctx contractapi.TransactionContextInterface, dealerID string) ([]map[string]interface{}, error) {
    resultsIterator, err := ctx.GetStub().GetHistoryForKey(dealerID)
    if err != nil {
        return nil, err
    }
    defer resultsIterator.Close()

    var history []map[string]interface{}
    for resultsIterator.HasNext() {
        mod, _ := resultsIterator.Next()
        var asset Asset
        _ = json.Unmarshal(mod.Value, &asset)
        history = append(history, map[string]interface{}{
            "txId":   mod.TxId,
            "value":  asset,
            "isDel":  mod.IsDelete,
            "ts":     mod.Timestamp,
        })
    }
    return history, nil
}

func main() {
    chaincode, err := contractapi.NewChaincode(new(SmartContract))
    if err != nil {
        panic("Error creating chaincode: " + err.Error())
    }
    if err := chaincode.Start(); err != nil {
        panic("Error starting chaincode: " + err.Error())
    }
}
