package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestMarshalMonetaryValue(t *testing.T) {

	money := &MonetaryValue{
		1234,
		456,
		"GBP",
	}

	expectedResult := "1234.456"
	actualResult := money.StringValue()
	if expectedResult != actualResult {
		t.Fatalf("Expected: %v Got: %v", expectedResult, actualResult)
	}

}

func TestUnMarshalMontaryValue(t *testing.T) {

	data := []byte(`{"value":10,"fraction":1,"currency":"GBP"}`)
	var actualResult MonetaryValue
	err := json.Unmarshal(data, &actualResult)
	if err != nil {
		t.Fatalf("Unable to unmarshal")
	}

	fmt.Println(actualResult)

	expectedResult := MonetaryValue{
		10,
		1,
		"GBP",
	}

	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Fatalf("Expected %v Got %v", expectedResult, actualResult)
	}

}

func TestLedgerEntry(t *testing.T) {

	money := &MonetaryValue{
		1234,
		456,
		"GBP",
	}

	ledgerEntry := &LedgerEntry{
		*money,
		"Unspent",
	}

	encoding, err := json.Marshal(ledgerEntry)
	if err != nil {
		t.Fatalf("Failed to marshal ledgerEntry")
	}

	expectedResult := `{"amount":{"value":1234,"fraction":456,"currency":"GBP"},"beneficiary":"Unspent"}`
	actualResult := string(encoding)

	if actualResult != expectedResult {
		t.Fatalf("Expected: %v Got: %v", expectedResult, actualResult)
	}

}

func TestMarshalLedgerWithEmptyValue(t *testing.T) {
	ledger := new(Ledger)

	encoding, err := json.Marshal(ledger)
	if err != nil {
		t.Fatalf("Failure to marhsal ledger")
	}

	expectedResult := `{"entries":null}`
	actualResult := string(encoding)
	if actualResult != expectedResult {
		t.Fatalf("Expected: %v Got: %v", expectedResult, actualResult)
	}
}

func TestMarhalLedgerWithInitialValue(t *testing.T) {

	// Ledger with initial values
	money := &MonetaryValue{
		1234,
		456,
		"GBP",
	}

	ledgerEntry := &LedgerEntry{
		*money,
		"Unspent",
	}

	ledger := &Ledger{
		[]LedgerEntry{
			*ledgerEntry,
		},
	}

	encoding, err := json.Marshal(ledger)
	if err != nil {
		t.Fatalf("Json marshal failed")
	}

	expectedResult := `{"entries":[{"amount":{"value":1234,"fraction":456,"currency":"GBP"},"beneficiary":"Unspent"}]}`
	actualResult := string(encoding)
	if actualResult != expectedResult {
		t.Fatalf("Expected: %v Got: %v", expectedResult, actualResult)
	}
}
