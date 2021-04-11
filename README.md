## bc4ic_poc

/*
 * BC4IC - Blockchain Care for Instutionalized Children (POC - Proof of Concept)
 * Develop by:
 * 				FABIAN CHIERA (fabian.chiera@gmail.com)
 *				@chetino (Github) 
 *				Blockfactory (blockfactory.io)
 */


## How to use it

#chaincode install

peer chaincode install -n bc4ic_poc -v 1.0 -p github.com/bc4ic_poc -l golang

#chaincode instantiate
peer chaincode instantiate -n bc4ic_poc -v 1.0 -C childrenchannel -c '{"Args":[]}' -o orderer.bc4ic.com:7050

#register new child on the blockchain
peer chaincode invoke -n bc4ic_poc -C peer chaincode instantiate -n bc4ic_poc -v 1.0 -C childrenchannel -c '{"Args":[]}' -o orderer.bc4ic.com:7050
 -c '{"Args":["registerChild","1","daniel fernandez","",""]}' -o orderer.bc4ic.com:7050


peer chaincode invoke -n bc4ic_poc -C peer chaincode instantiate -n bc4ic_poc -v 1.0 -C childrenchannel -c '{"Args":[]}' -o orderer.bc4ic.com:7050
 -c '{"Args":["registerChild","2","lorena sinz","",""]}' -o orderer.bc4ic.com:7050


peer chaincode invoke -n bc4ic_poc -C peer chaincode instantiate -n bc4ic_poc -v 1.0 -C childrenchannel -c '{"Args":[]}' -o orderer.bc4ic.com:7050
 -c '{"Args":["registerChild","3","clarisa peterson","",""]}' -o orderer.bc4ic.com:7050



peer chaincode invoke -n bc4ic_poc -C peer chaincode instantiate -n bc4ic_poc -v 1.0 -C childrenchannel -c '{"Args":[]}' -o orderer.bc4ic.com:7050
 -c '{"Args":["registerChild","4","diego parli","",""]}' -o orderer.bc4ic.com:7050


#get the children on the blockchain network
peer chaincode invoke -n bc4ic_poc -C childrenchannel -c '{"Args":["queryChildren"]}' -o orderer.bc4ic.com:7050

#get an specific child
peer chaincode invoke -n bc4ic_poc -C childrenchannel -c '{"Args":["queryChild","2"]}' -o orderer.bc4ic.com:7050

#child secure transfer to an specific agency
peer chaincode invoke -n bc4ic_poc -C childrenchannel -c '{"Args":["transferChild","2","PPNA"]}' -o orderer.bc4ic.com:7050

