package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SupplyChainContract struct {
	contractapi.Contract
}

type Product struct {
	ProductID   string `json:"productID"`
	ProductName string `json:"productName"`
	Owner       string `json:"owner"`
	Status      string `json:"status"`
}

func (s *SupplyChainContract) CreateProduct(ctx contractapi.TransactionContextInterface, productID string, productName string, owner string) error {
	product := Product{
		ProductID:   productID,
		ProductName: productName,
		Owner:       owner,
		Status:      "created",
	}

	productAsBytes, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("failed to marshal product: %s", err.Error())
	}

	return ctx.GetStub().PutState(productID, productAsBytes)
}

func (s *SupplyChainContract) TransferProduct(ctx contractapi.TransactionContextInterface, productID string, newOwner string) error {
	productAsBytes, err := ctx.GetStub().GetState(productID)
	if err != nil {
		return fmt.Errorf("failed to get product: %s", err.Error())
	} else if productAsBytes == nil {
		return fmt.Errorf("product %s does not exist", productID)
	}

	product := new(Product)
	err = json.Unmarshal(productAsBytes, product)
	if err != nil {
		return fmt.Errorf("failed to unmarshal product: %s", err.Error())
	}

	product.Owner = newOwner

	productAsBytes, err = json.Marshal(product)
	if err != nil {
		return fmt.Errorf("failed to marshal product: %s", err.Error())
	}

	return ctx.GetStub().PutState(productID, productAsBytes)
}

func (s *SupplyChainContract) GetProduct(ctx contractapi.TransactionContextInterface, productID string) (*Product, error) {
	productAsBytes, err := ctx.GetStub().GetState(productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %s", err.Error())
	} else if productAsBytes == nil {
		return nil, fmt.Errorf("product %s does not exist", productID)
	}

	product := new(Product)
	err = json.Unmarshal(productAsBytes, product)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal product: %s", err.Error())
	}

	return product, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SupplyChainContract))
	if err != nil {
		fmt.Printf("Error creating supply chain chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting supply chain chaincode: %s", err.Error())
	}
}