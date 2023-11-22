package main

import (
  "encoding/json"
  "fmt"
  "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Data struct {
  ID        string `json:"ID"`
  Value     string `json:"Value"`
}

type SmartContract struct {
  contractapi.Contract
}

// Store data in blockchain
func (s *SmartContract) StoreData(ctx contractapi.TransactionContext, id string, value string) error {
  data := Data{
    ID:        id,
    Value:     value,
  
  }

  dataBytes, err := json.Marshal(data)
  if err != nil {
    return err
  }

  err = ctx.GetStub().PutState(data.ID, dataBytes)
  if err != nil {
    return err
  }

  return nil
}

// Retrieve data from blockchain
func (s *SmartContract) RetrieveData(ctx contractapi.TransactionContext, id string) ([]byte, error) {
  dataBytes, err := ctx.GetStub().GetState(id)
  if err != nil {
    return nil, err
  }

  if dataBytes == nil {
    return nil, fmt.Errorf("Data with ID %s does not exist", id)
  }

  return dataBytes, nil
}

// Get history of data from blockchain
func (s *SmartContract) GetHistory(ctx contractapi.TransactionContext, id string) ([]byte, error) {
  historyIterator, err := ctx.GetStub().GetHistoryForKey(id)
  if err != nil {
    return nil, err
  }

  var historyData []Data
  for historyIterator.HasNext() {
    historyDataBytes, _ := historyIterator.Next()
    historyDataJSON := &Data{}
    err = json.Unmarshal(historyDataBytes, historyDataJSON)
    if err != nil {
      return nil, err
    }

    historyData = append(historyData, *historyDataJSON)
  }

  historyDataBytes, err := json.Marshal(historyData)
  if err != nil {
    return nil, err
  }

  return historyDataBytes, nil
}

// Get data based on a non-primary key using CouchDB rich query
func (s *SmartContract) GetByNonPrimaryKey(ctx contractapi.TransactionContext, query string) ([]byte, error) {
  queryResults, err := ctx.GetStub().GetQueryResult(query)
  if err != nil {
    return nil, err
  }

  var data []Data
  for queryResults.HasNext() {
    queryResult, err := queryResults.Next()
    if err != nil {
      return nil, err
    }

    dataBytes := queryResult.Value
    dataJSON := &Data{}
    err = json.Unmarshal(dataBytes, dataJSON)
    if err != nil {
      return nil, err
    }

    data = append(data, *dataJSON)
  }

  dataBytes, err := json.Marshal(data)
  if err != nil {
    return nil, err
  }

  return dataBytes, nil
}

func main() {
  chaincode, err := contractapi.NewChaincode(&SmartContract{})
  if err != nil {
    fmt.Printf("Error creating chaincode: %s\n", err)
    return
  }

  chaincode.Start()
}
