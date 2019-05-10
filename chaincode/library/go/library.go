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
	"bytes"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the book structure, with 5 properties.  Structure tags are used by encoding/json library
type Book struct {
	Bookno		int	   `json: "bookno"`
	Bookname 	string `json: "bookname"`
	Author 		string `json: "author"`
	Publisher 	string `json: "publisher"`
	Location 	string `json: "location"`
	Library		string `json: "library"`
	Rent        bool   `json: "rent"`
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

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	bookname := args[0]
	location := args[1]
	fmt.Println("bookname, location : ", bookname, location)

	results, err := APIstub.GetStateByPartialCompositeKey(bookname, []string{location})

	if err != nil{
		return shim.Error(err.Error())
	}
	defer results.Close()
	var i int
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	for i = 0; results.HasNext(); i++{
		responseRange, err := results.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(responseRange.Value))
		//objectType, compositeKeyParts, err := APIstub.SplitCompositeKey(responseRange.Key)
		//bookAsBytes, _ := APIstub.GetState(responseRange.Key)
		fmt.Printf("- found a marble from Key:%s Value:%s \n", responseRange.Key, responseRange.Value)
		//return shim.Success(responseRange.Value)
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	books := []Book{
		Book{Bookno: 0,Bookname: "연금술사", Author: "파울로 코엘료", Publisher: "문학동네", Location: "고양시", Library : "백석 도서관0", Rent: true},
		Book{Bookno: 1,Bookname: "연금술사", Author: "파울로 코엘료", Publisher: "문학동네", Location: "고양시", Library : "백석 도서관1", Rent: true},
		Book{Bookno: 2,Bookname: "연금술사", Author: "파울로 코엘료", Publisher: "문학동네", Location: "고양시", Library : "백석 도서관2", Rent: true},
		Book{Bookno: 3,Bookname: "연금술사", Author: "파울로 코엘료", Publisher: "문학동네", Location: "고양시", Library : "백석 도서관3", Rent: true},
		Book{Bookno: 4,Bookname: "연금술사", Author: "파울로 코엘료", Publisher: "문학동네", Location: "고양시", Library : "백석 도서관4", Rent: true},
		Book{Bookno: 5,Bookname: "연금술사", Author: "파울로 코엘료", Publisher: "문학동네", Location: "고양시", Library : "백석 도서관5", Rent: true},
		Book{Bookno: 6,Bookname: "연금술사", Author: "파울로 코엘료", Publisher: "문학동네", Location: "고양시", Library : "백석 도서관6", Rent: true},
		Book{Bookno: 7,Bookname: "연금술사", Author: "파울로 코엘료", Publisher: "문학동네", Location: "고양시", Library : "백석 도서관7", Rent: true},
		Book{Bookno: 8,Bookname: "연금술사", Author: "파울로 코엘료", Publisher: "문학동네", Location: "고양시", Library : "백석 도서관8", Rent: true},
	}

	i := 0
	for i < len(books) {
		fmt.Println("i is ", i)
		Key, _ := APIstub.CreateCompositeKey(books[i].Bookname, []string{books[i].Location, books[i].Library})
		bookAsBytes, _ := json.Marshal(books[i])
		APIstub.PutState(Key, bookAsBytes)
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
	var book = Book{Bookno: args[0], Bookname: args[1], Author: args[2], Publisher: args[3], Location: args[4], Library : args[5], Rent: true}
	Key, _ := APIstub.CreateCompositeKey(args[0], []string{args[4], args[5]})
	bookAsBytes, _ := json.Marshal(book)
	APIstub.PutState(Key, bookAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {
	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err);
	}
}
