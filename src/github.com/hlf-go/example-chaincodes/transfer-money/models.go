package main

import (
	"fmt"
)

/*
	Abstraction of money
*/

// MonetaryValue represents value in a particular currency
type MonetaryValue struct {
	Value           int64  `json:"value"`
	FractionalValue uint64 `json:"fraction"`
	Currency        string `json:"currency"`
}

// StringValue return money value combining non-fractional and fractional form
func (t *MonetaryValue) StringValue() string {
	return fmt.Sprintf("%v.%v", t.Value, t.FractionalValue)
}

/*
	Representation of a ledger in this form:

    [
			{
			"amount":{"value": 1111, "fraction": 111, "currency": "GBP"},
			"beneficiary": "someone"
			},
			{
				....
			}
		]

*/

// LedgerEntry represents an entry in the ledger.
// It presents 'amount' represented as MonetaryValue type.
// Beneficiary attribute is an unconstrained string type.
type LedgerEntry struct {
	Amount      MonetaryValue `json:"amount"`
	Beneficiary string        `json:"beneficiary"`
}

// Ledger represents a ledger showing records of beneficiaries to a sum of money
type Ledger struct {
	Entries []LedgerEntry `json:"entries"`
}
