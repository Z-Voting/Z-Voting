This is the chaincode repository for Z-Voting, a privacy-preserving e-voting platform that utilizes Zero-Knowledge Proofs and Multi-Party Computation.

# Chaincode dev mode

Development Network for Hyperledger Fabric.

Just run the following commands from the project root directory to get started. You will have to open 3 terminals:

# Terminal 1 (Start up network)
```bash
docker-compose up
```

Wait for 20 seconds before you move to terminal 2.
# Terminal 2 (Build and run chaincode)
Here, we have to set the Chaincode Name first (the CC_N variable).

```bash
export CC_N=zvoting && docker exec -it chaincode bash -c "cd $CC_N && go build -o $CC_N && CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=$CC_N:1 ./$CC_N"
```

# Terminal 3 (Install and instantiate chaincode)
The following code works with the provided chaincode **zvoting**. You may need to modify it according to your own needs. But if you are a beginner, you can just modify the zvoting code and keep using the following commands.

```bash
# connect to cli container
docker exec -it cli bash

# install chaincode
# we are given unique name and version of the chaincode
peer chaincode install -p chaincodedev/chaincode/zvoting -n zvoting -v 1

# instantiate chaincode
# we are giving channel name and chaincode name
peer chaincode instantiate -n zvoting -v 1 -c '{"Args":[]}' -C myc
```

# Test chaincode
To test the chaincode we can execute invoke, query transactions. Following is the way to do that.

```bash
#If you are not in terminal 3, connect to cli container with the next command. But if you are already there, you should skip it.
docker exec -it cli bash

# invoke transactions with 'create'  
peer chaincode invoke -n zvoting -c '{"Args":["create", "001", "lambda", "l@zvoting.com"]}' -C myc
peer chaincode invoke -n zvoting -c '{"Args":["create", "002", "ops", "o@zvoting.com"]}' -C myc

# query transactions with 'get'
# output - {"id":"001","name":"lambda","email":"l@zvoting.com"}
peer chaincode query -n zvoting -c '{"Args":["get", "001"]}' -C myc

# query transactions with 'search'
# output
[{"id":"001","name":"lambda","email":"l@zvoting.com"},{"id":"002","name":"ops","email":"o@zvoting.com"}]
peer chaincode query -n zvoting -c '{"Args":["search", "001", "005"]}' -C myc
```


Read more infomarions from [here](https://medium.com/@itseranga/test-hyperledger-fabric-chaincode-in-dev-environment-8794096b5df2)
