package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type RahasakContract struct {
}

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (s *RahasakContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	fmt.Printf("INFO: init chaincode args: %s\n", args)

	return shim.Success(nil)
}

func (s *RahasakContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	fmt.Printf("INFO: invoke function: %s, args: %s\n", function, args)

	if function == "create" {
		return s.Create(stub, args)
	} else if function == "get" {
		return s.Get(stub, args)
	} else if function == "search" {
		return s.Search(stub, args)
	}

	return shim.Error("Invalid smart contract function")
}

func (s *RahasakContract) Create(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: create with args: %s\n", args)

	// user
	usr := User{
		Id:    args[0],
		Name:  args[1],
		Email: args[2],
	}
	usrJsn, _ := json.Marshal(usr)
	err := stub.PutState(args[0], usrJsn)
	if err != nil {
		fmt.Printf("ERROR: error PutState: %s\n", err.Error())
		shim.Error("error PutState: " + err.Error())
	}

	return shim.Success(nil)
}

func (s *RahasakContract) Get(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: get with args: %s\n", args)

	usr, _ := stub.GetState(args[0])
	if usr == nil {
		return shim.Error("Could not get user with id: " + args[0])
	}

	return shim.Success(usr)
}

func (s *RahasakContract) Search(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: search with args: %s\n", args)

	// from, to range comes with args
	frm := args[0]
	to := args[1]

	// search by range
	iterator, err := stub.GetStateByRange(frm, to)
	if err != nil {
		return shim.Error("Error search by range: " + err.Error())
	}
	defer iterator.Close()

	// build json respone
	buffer, err := buildResponse(iterator)
	if err != nil {
		return shim.Error("Error constract response: " + err.Error())
	}
	fmt.Printf("INFO: search response:%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func buildResponse(iterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing query results
	var buffer bytes.Buffer
	buffer.WriteString("[")

	written := false
	for iterator.HasNext() {
		resp, err := iterator.Next()
		if err != nil {
			return nil, err
		}

		// add a comma before array members, suppress it for the first array member
		if written == true {
			buffer.WriteString(",")
		}

		// record is a JSON object, so we write as it is
		buffer.WriteString(string(resp.Value))
		written = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

func main() {
	err := shim.Start(new(RahasakContract))
	if err != nil {
		fmt.Printf("ERROR: error creating rahasak contact: %s\n", err.Error())
	}
}
