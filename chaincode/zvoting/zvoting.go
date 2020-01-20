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
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Election struct {
	Id        string
	Name      string
	StartTime string
	Duration  string
	Doctype   string
}

type Voter struct {
	Id string
	Name string
	Email string
	V1 string
	V2 string
	V3 string
	ElectionId string
	Doctype string
}

type Candidate struct {
	Id         string
	Name       string
	Sign       string
	ImgAddress string
	ElectionId string
	Doctype    string
}

func (election *Election) isRunning() bool {
	currentTime := time.Now().Unix()
	startTime, _ := strconv.ParseInt(election.StartTime, 10, 64)
	duration, _ := strconv.ParseInt(election.Duration, 10, 64)
	endTime := startTime + duration
	return currentTime >= startTime && currentTime <= endTime
}

func (election *Election) isFresh() bool {
	return election.StartTime == "0"
}

func (election *Election) isOver() bool {
	return !election.isFresh() && !election.isRunning()
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
	} else if function == "addCandidate" {
		return s.addCandidate(stub, args)
	} else if function == "delete" {
		return s.delete(stub, args)
	} else if function == "getElections" {
		return s.getElections(stub, args)
	} else if function == "getCandidates" {
		return s.getCandidates(stub, args)
	} else if function == "startElection" {
		return s.startElection(stub, args)
	} else if function == "registerVoter" {
		return s.registerVoter(stub, args)
	}

	return shim.Error("Invalid smart contract function")
}

func (s *ZVotingContract) registerVoter(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: Register Voter with args: %s\n", args)

	key := generateUID(20, stub, args)

	voter := Voter{
		Id:         key,
		Name:       args[0],
		Email:      args[1],
		V1:         args[2],
		V2:         args[3],
		V3:         args[4],
		ElectionId: args[5],
		Doctype:    "Voter",
	}

	voterJSON, _ := json.Marshal(voter)
	err := stub.PutState(key, voterJSON)
	if err != nil {
		fmt.Printf("ERROR: error PutState: %s\n", err.Error())
		shim.Error("error PutState: " + err.Error())
	}

	return shim.Success(nil)
}

func (s *ZVotingContract) getRandom(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	seed, _ := strconv.Atoi(args[0])
	rand.Seed(int64(seed))
	return shim.Success([]byte(strconv.Itoa(rand.Int())))
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

func generateUID(l int, stub shim.ChaincodeStubInterface, args []string) string {
	rand.Seed(time.Now().UnixNano())
	randStr := randomString(l)

	value, err := stub.GetState(randStr)
	if err != nil {
		return fmt.Sprintf("Failed to get asset: %s with error: %s", args[0], err)
	}
	for value != nil {
		randStr = randomString(l)
		value, err = stub.GetState(randStr)
	}
	return randStr
}

func (s *ZVotingContract) createId(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	randStr := generateUID(20, stub, args)

	return shim.Success([]byte(randStr))
}

func (s *ZVotingContract) createElection(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: create election with args: %s\n", args)

	key := generateUID(20, stub, args)

	election := Election{
		Id:        key,
		Name:      args[0],
		StartTime: "0",
		Duration:  args[1],
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

func (s *ZVotingContract) addCandidate(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: create election with args: %s\n", args)

	key := generateUID(20, stub, args)

	candidate := Candidate{
		Id:         key,
		Name:       args[0],
		Sign:       args[1],
		ImgAddress: args[2],
		ElectionId: args[3],
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

	queryString := newCouchQueryBuilder().addSelector("Doctype", "Election").getQueryString()

	iterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error("Error getting elections: " + err.Error())
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

func (s *ZVotingContract) getCandidates(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: Get Candidates")

	queryString := newCouchQueryBuilder().addSelector("Doctype", "Candidate").addSelector("ElectionId", args[0]).getQueryString()

	iterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error("Error getting elections: " + err.Error())
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

func (s *ZVotingContract) startElection(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: start election with args: %s\n", args)

	key := args[0]

	data, _ := stub.GetState(args[0])
	var election Election
	err := json.Unmarshal(data, &election)

	if !election.isFresh() {
		return shim.Error("This isn't a Fresh Election")
	}

	election.StartTime = strconv.Itoa(int(time.Now().Unix()))

	electionJSON, _ := json.Marshal(election)
	err = stub.PutState(key, electionJSON)
	if err != nil {
		fmt.Printf("ERROR: error PutState: %s\n", err.Error())
		shim.Error("error PutState: " + err.Error())
	}

	return shim.Success(nil)
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
