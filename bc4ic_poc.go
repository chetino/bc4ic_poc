/*
 * BC4IC - Blockchain Care for Instutionalized Children (POC - Proof of Concept)
 * Develop by:
 * 				FABIAN CHIERA (fabian.chiera@gmail.com)
 *				@chetino (Github) 
 *				Blockfactory (blockfactory.io)
 */

 package main

 import (
	 "bytes"
	 "encoding/json"
	 "fmt"
 
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 sc "github.com/hyperledger/fabric/protos/peer"
 )
 
 // Define the Smart Contract structure
 type SmartContract struct {
 }
 
 // Define the child structure. Structure tags are used by encoding/json library
 type Child struct {
	 DU     string `json:"du"`
	 Name   string `json:"name"`
	 Agency string `json:"agency"`
 }
 
 func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 // Retrieve the requested Smart Contract function and arguments
	 function, args := APIstub.GetFunctionAndParameters()
	 // Route to the appropriate handler function to interact with the ledger appropriately
	 if function == "queryChild" {
		 return s.queryChild(APIstub, args)
	 } else if function == "initLedger" {
		 return s.initLedger(APIstub)
	 } else if function == "registerChild" {
		 return s.registerChild(APIstub, args)
	 } else if function == "queryChildren" {
		 return s.queryChildren(APIstub)
	 } else if function == "transferChild" {
		 return s.transferChild(APIstub, args)
	 }
 
	 return shim.Error("Invalid Smart Contract function name.")
 }
 
 func (s *SmartContract) queryChild(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }
 
	 childAsBytes, _ := APIstub.GetState(args[0])
	 return shim.Success(childAsBytes)
 }
 
 func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) registerChild(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 4 {
		 return shim.Error("Incorrect number of arguments. Expecting 4")
	 }
 
	 var child = Child{DU: args[1], Name: args[2], Agency: args[3]}
 
	 childAsBytes, _ := json.Marshal(child)
	 APIstub.PutState(args[0], childAsBytes)
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) queryChildren(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 startKey := ""
	 endKey := ""
 
	 resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	 if err != nil {
		 return shim.Error(err.Error())
	 }
	 defer resultsIterator.Close()
 
	 // buffer is a JSON array containing QueryResults
	 var buffer bytes.Buffer
	 buffer.WriteString("[")
 
	 bArrayMemberAlreadyWritten := false
	 for resultsIterator.HasNext() {
		 queryResponse, err := resultsIterator.Next()
		 if err != nil {
			 return shim.Error(err.Error())
		 }
		 // Add a comma before array members, suppress it for the first array member
		 if bArrayMemberAlreadyWritten == true {
			 buffer.WriteString(",")
		 }
		 buffer.WriteString("{\"Key\":")
		 buffer.WriteString("\"")
		 buffer.WriteString(queryResponse.Key)
		 buffer.WriteString("\"")
 
		 buffer.WriteString(", \"Record\":")
		 // Record is a JSON object, so we write as-is
		 buffer.WriteString(string(queryResponse.Value))
		 buffer.WriteString("}")
		 bArrayMemberAlreadyWritten = true
	 }
	 buffer.WriteString("]")
 
	 fmt.Printf("- queryChildren:\n%s\n", buffer.String())
 
	 return shim.Success(buffer.Bytes())
 }
 
 func (s *SmartContract) transferChild(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }
 
	 childAsBytes, _ := APIstub.GetState(args[0])
	 child := Child{}
 
	 json.Unmarshal(childAsBytes, &child)
	 child.Agency = args[1]
 
	 childAsBytes, _ = json.Marshal(child)
	 APIstub.PutState(args[0], childAsBytes)
 
	 return shim.Success(nil)
 }
 
 // The main function is only relevant in unit test mode. Only included here for completeness.
 func main() {
 
	 // Create a new Smart Contract
	 err := shim.Start(new(SmartContract))
	 if err != nil {
		 fmt.Printf("Error creating new Smart Contract: %s", err)
	 }
 }
 
 