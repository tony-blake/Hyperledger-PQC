package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

/*
	Chaincode
*/

// SimpleChaincode representing a class of chaincode
type SimpleChaincode struct{}

// Init to initiate the SimpleChaincode class
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	data, err := stub.GetState("money-transfer")
	if err != nil {
		return shim.Error("Unable to interact with state")
	}

	result := fmt.Sprintf("data %v", data)
	fmt.Println(result)

	return shim.Success([]byte("Init called"))
}

// Invoke a method specified in the SimpleChaincode class
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Hello Invoke")
	function, _ := stub.GetFunctionAndParameters()
	if function == "methodName" {
		fmt.Println(function)
	}
	return shim.Success([]byte("Invoke"))
}

/*
	Chaincode process entry point
*/
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}
