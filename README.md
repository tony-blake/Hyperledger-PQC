# Hyperleger-Fabric-PQC-minimal-example
Using hybrid x509 certs  with Leighton-Mikhali signatures on top of ECDSA in HLF


Quick and Dirty Set Up
-----------------------

0. Clone this directory into your $HOME folder
1. Start Docker (assuming Docker is installed) and remove existing docker containers by running ```docker rm -f $(docker ps -aq)```
2. Navigate to ``` src/github.com/hlf-go/writing-chaincode/fabric ```
3. run ``` curl -sSL https://raw.githubusercontent.com/hyperledger/fabric/release-1.1/scripts/bootstrap-1.1.0-preview.sh -o setup_script.sh ```
4. run ``` bash setup_script.sh ``` (Note - This pulls more docker images than is necessary for this example. However this is suffcient for the example)
5. Change GOPATH varible by typing ```GOPATH= ``` and hit return. Then type ```GOPATH=$HOME/go``` and hit return.
6. Then export the path to the go bin folder by returning ```export PATH=$GOPATH/bin:$PATH```
7. Copy the up to date contents of the bin folder to the new location. ```cp bin/* $HOME/go/bin/```
8. Start minimal fabric network (1 peer, 1 orderer, 1 CLI) by running the command 
   ```bash
   ./fabricOPS.sh start
   ```


Patching OpenSSl amd LibEST for post quantum cryptography
---------------------------------------------------------

1. Visit this website (Isara corp PQC test server). 
2. Clone repo and Follow instructions for patching
3. export path to new libraries by returning ```export PATH=/usr/local/pqpki-openssl1.0.2o/bin:$PATH``` 
4. To install on CLI container run  ```./fabricOPS.sh cli```, then ```cd```, 
then run ```cp $GOPATH/github.com/hyperledger/fabric/examples/chaincode/go/minimalcc2/sslPatches.sh ~/```
then run ```bash sslPatches.sh```




Deploy Chaincode
----------------

1. Interact with peer by using CLI container. Run ```./fabricOPS.sh cli``` to enter CLI conatiner as root.
2. Instantiate chaincode by running ```./scripts/instantiate.sh -c minimalcc2 -v 1.0 -p minimalcc2```
3. Invoke chaincode by running ``` ./scripts/invoke.sh -c minimalcc2 ```





Current Issues with Example
===========================
1. Cannot seem to run bash script from chaincode located at ```src/github.com/hlf-go/example-chaincodes/minimalcc2/chaincode.go```

Chaincode is instantiated and invoked but there is no indication that the bash script executed.
I tested running the bash script from a simple go file and it works perfectly in that case.
