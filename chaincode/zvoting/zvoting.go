package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"math/rand"
	"strconv"
	"time"
)

type ZVotingContract struct {
}

type User struct {
	Id string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Election struct {
	Id string
	Name string
	StartTime string
	EndTime string
	Doctype string
}

type Candidate struct {
	Id string
	Name string
	Sign string
	ImgAddress string
	ElectionId string
	Doctype string
}

func (s *ZVotingContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	fmt.Printf("INFO: init chaincode args: %s\n", args)

	return shim.Success(nil)
}

func (s *ZVotingContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	fmt.Printf("INFO: invoke function: %s, args: %s\n", function, args)

	if function == "create" {
		return s.Create(stub, args)
	} else if function == "get" {
		return s.Get(stub, args)
	} else if function == "search" {
		return s.Search(stub, args)
	} else if function == "getRandom" {
		return s.getRandom(stub, args)
	} else if function == "createId" {
		return s.createId(stub, args)
	} else if function == "createElection" {
		return s.createElection(stub, args)
	} else if function == "createCandidate" {
		return s.createCandidate(stub, args)
	} else if function == "delete" {
		return s.delete(stub, args);
	}

	return shim.Error("Invalid smart contract function")
}

func (s *ZVotingContract) getRandom(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	seed, _ := strconv.Atoi(args[0])
	rand.Seed( int64(seed) )
	return shim.Success( []byte(strconv.Itoa(rand.Int()))  )
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randomString(l int) string {
	b := make([]rune, l)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (s *ZVotingContract) createId(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	rand.Seed(time.Now().UnixNano())
	var l int = 20
	randStr := randomString(l)

	value, err := stub.GetState(randStr)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to get asset: %s with error: %s", args[0], err))
	}
	for value != nil {
		randStr = randomString(l)
		value, err = stub.GetState(randStr)
	}

	return shim.Success( []byte(randStr) )
}

func (s *ZVotingContract) createElection(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: create election with args: %s\n", args)

	key := args[0]

	election := Election{
		Id:        key,
		Name:      args[1],
		StartTime: args[2],
		EndTime:   args[3],
		Doctype:   "Election",
	}

	electionJSON, _ := json.Marshal(election)
	err := stub.PutState(key, electionJSON)
	if err != nil {
		fmt.Printf("ERROR: error PutState: %s\n", err.Error())
		shim.Error("error PutState: " + err.Error())
	}

	return shim.Success(nil)
}

func (s *ZVotingContract) createCandidate(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: create election with args: %s\n", args)

	key := args[0]

	candidate := Candidate{
		Id:         key,
		Name:       args[1],
		Sign:       args[2],
		ImgAddress: args[3],
		ElectionId: args[4],
		Doctype:    "Candidate",
	}
	candidateJSON, _ := json.Marshal(candidate)
	err := stub.PutState(key, candidateJSON)
	if err != nil {
		fmt.Printf("ERROR: error PutState: %s\n", err.Error())
		shim.Error("error PutState: " + err.Error())
	}

	return shim.Success(nil)
}

func (s *ZVotingContract) delete(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: delete with key: %s\n", args)

	err := stub.PutState(args[0], nil)
	if err != nil {
		fmt.Printf("ERROR: error PutState: %s\n", err.Error())
		shim.Error("error PutState: " + err.Error())
	}

	return shim.Success(nil)
}

func (s *ZVotingContract) getElections(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: Get Elections")


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

func (s *ZVotingContract) Create(stub shim.ChaincodeStubInterface, args []string) peer.Response {
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

func (s *ZVotingContract) Get(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: get with args: %s\n", args)

	data, _ := stub.GetState(args[0])
	if data == nil {
		return shim.Error("Could not get record with id: " + args[0])
	}

	return shim.Success(data)
}

func (s *ZVotingContract) Search(stub shim.ChaincodeStubInterface, args []string) peer.Response {
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
	err := shim.Start(new(ZVotingContract))
	if err != nil {
		fmt.Printf("ERROR: error creating rahasak contact: %s\n", err.Error())
	}
}
