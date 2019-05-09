/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (

	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the book structure, with 5 properties.  Structure tags are used by encoding/json library
type Book struct {
	Bookname 	string `json:"bookname"`
	Author 		string `json:"author"`
	Publisher 	string `json: "publisher"`
	Location 	string `json: location`
	Library		string `json: library`
}

/*
 * The Init method is called when the Smart Contract "library" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "library"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryBook" {
		return s.queryBook(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createBook" {
		return s.createBook(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryBook(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	bookAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(bookAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	books := []Book{
		Book{Bookname: "연금술사", Author: "파울로 코엘료", Publisher: "문학동네", Location: "고양시", Library : "백석 도서관"},
	}

	i := 0
	for i < len(books) {
		fmt.Println("i is ", i)
		bookAsBytes, _ := json.Marshal(books[i])
		APIstub.PutState("BOOK"+strconv.Itoa(i), bookAsBytes)
		fmt.Println("Added", books[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createBook(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}
	fmt.Println(args)
	var book = Book{Bookname: args[1], Author: args[2], Publisher: args[3], Location: args[4], Library : args[5]}

	bookAsBytes, _ := json.Marshal(book)
	APIstub.PutState(args[0], bookAsBytes)

	return shim.Success(nil)
}
//
// func (s *SmartContract) queryAllCars(APIstub shim.ChaincodeStubInterface) sc.Response {
//
// 	startKey := "CAR0"
// 	endKey := "CAR999"
//
// 	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	defer resultsIterator.Close()
//
// 	// buffer is a JSON array containing QueryResults
// 	var buffer bytes.Buffer
// 	buffer.WriteString("[")
//
// 	bArrayMemberAlreadyWritten := false
// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()
// 		if err != nil {
// 			return shim.Error(err.Error())
// 		}
// 		// Add a comma before array members, suppress it for the first array member
// 		if bArrayMemberAlreadyWritten == true {
// 			buffer.WriteString(",")
// 		}
// 		buffer.WriteString("{\"Key\":")
// 		buffer.WriteString("\"")
// 		buffer.WriteString(queryResponse.Key)
// 		buffer.WriteString("\"")
//
// 		buffer.WriteString(", \"Record\":")
// 		// Record is a JSON object, so we write as-is
// 		buffer.WriteString(string(queryResponse.Value))
// 		buffer.WriteString("}")
// 		bArrayMemberAlreadyWritten = true
// 	}
// 	buffer.WriteString("]")
//
// 	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())
//
// 	return shim.Success(buffer.Bytes())
// }

// func (s *SmartContract) changeCarOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
//
// 	if len(args) != 2 {
// 		return shim.Error("Incorrect number of arguments. Expecting 2")
// 	}
//
// 	carAsBytes, _ := APIstub.GetState(args[0])
// 	car := Car{}
//
// 	json.Unmarshal(carAsBytes, &car)
// 	car.Owner = args[1]
//
// 	carAsBytes, _ = json.Marshal(car)
// 	APIstub.PutState(args[0], carAsBytes)
//
// 	return shim.Success(nil)
// }
//
// // The main function is only relevant in unit test mode. Only included here for completeness.
func main() {
	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err);
	}
}
