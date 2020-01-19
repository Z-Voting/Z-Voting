# Chaincode dev mode

Development Network for Hyperledger Fabric.

Just run the following commands from the project root directory:

**Start up network:** `docker-compose up -d`

**Build and run code:** Here, we have to set the Chaincode Name first (the CC_N variable).

`
export CC_N=rahasak && docker exec -e -it chaincode bash -c "cd $CC_N && go build -o $CC_N && CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=$CC_N:1 ./$CC_N"
`


Read more infomarions from [here](https://medium.com/@itseranga/test-hyperledger-fabric-chaincode-in-dev-environment-8794096b5df2)
