/*
esya Chaincode
--------------------

Right now, consider it as simple key-value store like redis :)

The chaincode must implement the three methods required by the Fabric chaincode
API: (1) Init, (2) Invoke, and (3) Query. In this example, the functionality of
each of these methods is described below.

(1) Init:
	The method that is triggered when the chaincode is deployed.

(2) Invoke:
	The method that is triggered when a chaincode receives an invocation
	transaction.

(3) Query:
	The method that is triggered when a chaincode receives a query transaction.

The chaincode must also contain the main() method, which starts the chaincode
when it is first deployed. This is basically a smart contract.
*/

package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// CrowdFundChaincode implementation
type CrowdFundChaincode struct {
}

//
// Init creates the state variable with name "VoteChainKey" and stores the value
// from the incoming request into this variable. We now have a key/value pair
// for VoteChainKey --> Json object as string.
//
func (t *CrowdFundChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// State variable "voteChainKey"
	var voteChainKey string
	// The value stored inside the state variable "voteChainKey"
	var jsonString string
	// Any error to be reported back to the client
	var err error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}

	// Initialize the state variable name
	voteChainKey = args[0]
	// Initialize the state variable value
	jsonString = args[1]

	fmt.Printf("jsonString = %q\n", jsonString)

	// Write the state to the ledger
	err = stub.PutState(voteChainKey, []byte(jsonString))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

//
// Invoke retrieves the state variable "VoteChainKey" and replaces by data
// specified in the incoming request. Then it stores the new value back, thus
// updating the ledger.
//
func (t *CrowdFundChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// State variable "VoteChainKey"
	var voteChainKey string
	// The json stored inside the state variable "VoteChainKey"
	var jsonString string

	// Any error to be reported back to the client
	var err error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}

	// Read in the name of the state variable to be updated
	voteChainKey = args[0]
	jsonString = args[1]

	fmt.Printf("jsonString = %q\n", jsonString)

	// Write the state back to the ledger
	err = stub.PutState(voteChainKey, []byte(jsonString))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

//
// Query retrieves the state variable "voteChainKey" and returns its current value
// in the response.
//
func (t *CrowdFundChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting \"query\".")
	}

	// State variable "voteChainKey"
	var voteChainKey string
	// Any error to be reported back to the client
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the state variable to query.")
	}

	// Read in the name of the state variable to be returned
	voteChainKey = args[0]

	// Get the current value of the state variable
	jsonValueBytes, err := stub.GetState(voteChainKey)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + voteChainKey + "\"}"
		return nil, errors.New(jsonResp)
	}
	if jsonValueBytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + voteChainKey + "\"}"
		return nil, errors.New(jsonResp)
	}

	jsonResp := "{\"Key\":\"" + voteChainKey + "\",\"jsonKey\":\"" + string(jsonValueBytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return jsonValueBytes, nil
}

func main() {
	err := shim.Start(new(CrowdFundChaincode))

	if err != nil {
		fmt.Printf("Error starting CrowdFundChaincode: %s", err)
	}
}
