package main

import (
	"fmt"
	"my-forked-chaincode-bootstrap/chaincode/component3/folder1"
	"my-forked-chaincode-bootstrap/chaincode/component3/folder2"
	"my-forked-chaincode-bootstrap/chaincode/component3/folder3"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//Component3Chaincode chaincode
type Component3Chaincode struct {
	folder1.F1Chaincode
	folder2.F2Chaincode
	folder3.F3Chaincode
}

// Init chaincode
func (t *Component3Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### Component3Chaincode Init ###########")
	return shim.Success(nil)
}

// Invoke chaincode functions
func (t *Component3Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### Component3Chaincode Invoke ###########")
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "Method1":
		fmt.Println("########## Calling Method1 ##############")
		return t.Method1(stub, args)
	case "F1Method1":
		fmt.Println("########## Calling F1Method1 ##############")
		return t.F1Method1(stub, args)
	case "F2Method1":
		fmt.Println("########## Calling F2Method1 ##############")
		return t.F2Method1(stub, args)
	case "F3Method1":
		fmt.Println("########## Calling F3Method1 ##############")
		return t.F3Method1(stub, args)
	}

	errorMsg := fmt.Sprintf("Unknown action '%s'. Please check the first argument, the function name.", function)
	return shim.Error(errorMsg)
}

//Method1 retrieves the citizen data, including any of its eligibilities and enrollments for all citizens
func (t *Component3Chaincode) Method1(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("Method1() called successfully")
	return shim.Success([]byte("Method1() called successfully"))
}
