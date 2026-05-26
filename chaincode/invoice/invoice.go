package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Invoice struct {
	ID          string  `json:"id"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	TaxWithheld float64 `json:"tax_withheld"`
	OriginOrg   string  `json:"origin_org"`
}

func (s *SmartContract) RegisterInvoice(ctx contractapi.TransactionContextInterface, id string, amount float64, currency string, tax float64) error {
	inv := Invoice{ID: id, Amount: amount, Currency: currency, TaxWithheld: tax}
	b, _ := json.Marshal(inv)
	return ctx.GetStub().PutState(id, b)
}

func (s *SmartContract) GetInvoice(ctx contractapi.TransactionContextInterface, id string) (*Invoice, error) {
	b, err := ctx.GetStub().GetState(id)
	if err != nil || b == nil {
		return nil, fmt.Errorf("invoice not found: %s", id)
	}
	var inv Invoice
	_ = json.Unmarshal(b, &inv)
	return &inv, nil
}

func main() {
	contractapi.Start(new(SmartContract))
}
