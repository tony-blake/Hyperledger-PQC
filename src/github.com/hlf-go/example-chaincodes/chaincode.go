package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("minimalcc2")

// SimpleChaincode representing a class of chaincode
type SimpleChaincode struct{}

// Init to initiate the SimpleChaincode class
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("Hello Init")
	return shim.Success([]byte("Init called"))
}

// Invoke a method specified in the SimpleChaincode class
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("Hello Invoke")
	function := stub.GetFunctionAndParameters()
	fmt.Println(" ")
	fmt.Println("starting invoke, for - " + function)

	if function == "init" { //initialize the chaincode state, used as reset
		return t.Init(stub)
	} else if function == "call_script" {//call bash script to start PQC
		return call_script(stub)

	}





	return shim.Success([]byte("Invoke"))
}



func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Debugf("Error: %s", err)
	}
}
