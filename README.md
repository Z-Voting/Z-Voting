# Chaincode dev mode

Development Network for Hyperledger Fabric.

Just run the following commands from the project root directory:

**Start up network:** `docker-compose up -d`

**Build and run code:**

`
docker exec -it chaincode bash -c "cd rahasak && go build -o rahasak && CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=rahasak:1 ./rahasak"
`


Read more infomarions from [here](https://medium.com/@itseranga/test-hyperledger-fabric-chaincode-in-dev-environment-8794096b5df2)
