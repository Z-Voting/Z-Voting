package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// device data model
type Device struct {
	SrcAddr  string `json:"srcAddr"`
	SrcPort  string `json:"srcPort"`
	InSnmp   string `json:"inSnmp"`
	DstAddr  string `json:"dstAddr"`
	DstPort  string `json:"dstPort"`
	OutSnmp  string `json:"outSnmp"`
	Protocol string `json:"protocol"`
}

type DeviceContract struct {
}

func (t *DeviceContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	_, args := stub.GetFunctionAndParameters()
	fmt.Println("INFO: init chaincode with args: %s", args)

	return shim.Success(nil)
}

func (t *DeviceContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Printf("INFO: invoke function: %s, args: %s\n", function, args)

	if function == "create" {
		// create device
		return t.create(stub, args)
	} else if function == "get" {
		// get devices
		return t.get(stub, args)
	} else if function == "delete" {
		// delete device
		return t.delete(stub, args)
	}

	return shim.Error("Invalid function name. Expecting [create, delete, get] funcations")
}

func (t *DeviceContract) Create(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: create device with args %s\n", args)

	if len(args) != 4 {
		return shim.Error("Invalid args to create device. Expecting 4 arguments")
	}

	// new device
	device := Device{}
	device.SrcAddr = args[0]
	device.SrcPort = args[1]
	device.InSnmp = ""
	device.DstAddr = args[2]
	device.DstPort = args[3]
	device.OutSnmp = ""
	device.Protocol = ""

	// save device on ledger
	marshMess, _ := json.Marshal(device)
	err := stub.PutState(device.SrcAddr, marshMess)
	if err != nil {
		return shim.Error("Error when create device")
	}

	return shim.Success(nil)
}

func (t *DeviceContract) Get(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: get device with args: %s\n", args)

	if len(args) != 1 {
		return shim.Error("Invalid args to get device. expecting device id")
	}

	// get device from ledger
	id := args[0]
	Avalbytes, err := stub.GetState(id)

	if err != nil {
		return shim.Error("Failed to get device: " + err.Error())
	}

	return shim.Success(Avalbytes)
}

func (t *DeviceContract) Search(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: get device with args: %s\n", args)

	if len(args) != 1 {
		return shim.Error("Invalid args to get device. expecting device id")
	}

	// get device from ledger
	id := args[0]
	Avalbytes, err := stub.GetState(id)

	if err != nil {
		return shim.Error("Failed to get device: " + err.Error())
	}

	return shim.Success(Avalbytes)
}

func (t *DeviceContract) Delete(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Printf("INFO: delete device with args: %s\n", args)

	if len(args) != 1 {
		return shim.Error("Invalid args for delete device. Expecting device id")
	}

	// delete device from ledger
	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error("Failed to delete device: " + err.Error())
	}

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(DeviceContract))
	if err != nil {
		fmt.Printf("ERROR: fail start device chaincode: %s\n", err.Error())
	}
}
