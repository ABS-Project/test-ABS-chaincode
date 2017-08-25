//
//  Copyright Tongji University. All Rights Reserved.
//  用于操作记录的添加和查询
//  SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("TxRecorder")

type SimpleChaincode struct {
}

// ============================================================================================================================
// TxInfo struct
// ============================================================================================================================
type TxInfoStruct struct {
	TxID          string    `json:"TxID"`       //交易ID
	TxProposer    string    `json:"TxProposer"` //交易发起人
	TxProductID   string    `json:"TxProductID"`
	TxTime        time.Time `json:"TxTime"`        //交易时间
	TxChaincode   string    `json:"TxChaincode"`   //链码名称
	TxFunction    string    `json:"TxFunction"`    //所调函数
	TxArguments   string    `json:"TxArguments"`   //所传参数
	TxDescription string    `json:"TxDescription"` //交易描述
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### TxRecorder Init ###########")
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### TxRecorder Invoke ###########")

	function, args := stub.GetFunctionAndParameters()
	if function == "add" {
		// Deletes an entity from its state
		return t.add(stub, args)
	}

	if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	}

	if function == "query" {
		// queries an entity state
		return t.query(stub, args)
	}

	if function == "queryAllTxRecord" {
		// queries an entity state
		return t.queryAllTxRecord(stub, args)
	}

	logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'queryAllTxRecord'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'queryAllTxRecord'. But got: %v", args[0]))
}

// #======================================================================
// # function：添加一条操作记录
// # input： TxInfoStruct
// #======================================================================
func (t *SimpleChaincode) add(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8. ")
	}
  //交易ID
	TxID := args[0]
	TxTest, _ := stub.GetState(TxID)
	if TxTest != nil {
		return shim.Error("the Transaction is sexisted")
	}
	//交易人，需要到BusinessPartnerInfo合约中检验操作人是否存在
	TxProposer := args[1]
	functionName := "addOperateLog"
	invokeArgs := util.ToChaincodeArgs(functionName, TxProposer, TxID)
	response := stub.InvokeChaincode("BusinessPartnerInfo", invokeArgs, "mychannel")
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	var TxInfoObj TxInfoStruct
	TxInfoObj.TxID = args[0]
	TxInfoObj.TxProposer = args[1]
	TxInfoObj.TxProductID = args[2]
	TxInfoObj.TxTime, _ = time.Parse("2006-01-02T15:04:05.000Z", args[3])
	TxInfoObj.TxChaincode = args[4]
	TxInfoObj.TxFunction = args[5]
	TxInfoObj.TxArguments = args[6]
	TxInfoObj.TxDescription = args[7]

	jsonAsBytes, _ := json.Marshal(TxInfoObj)
	err = stub.PutState(TxID, []byte(jsonAsBytes))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

// 空操作，不允许删除操作记录
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	return shim.Success(nil)
}

// #======================================================================
// # function：查询某个用户的所有操作
// # input： 用户属性下的操作记录ID列表
// #======================================================================
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var TxInfoList []string
	for i, TxID := range args {
		fmt.Printf("查询第%d条操作记录", i)
		TxInfo, err := stub.GetState(TxID)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get state for " + TxID + "\"}"
			return shim.Error(jsonResp)
		}
		if TxInfo == nil {
			jsonResp := "{\"Error\":\"Nil content for " + TxID + "\"}"
			TxInfoList = append(TxInfoList, jsonResp)
		}
		TxInfoList = append(TxInfoList, string([]byte(TxInfo)))
	}
	TxInfoListAsByte, err := json.Marshal(TxInfoList)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(TxInfoListAsByte)
}

// #======================================================================
// # function：批量查询所有的操作日志
// # input： 查询的开始开始键值和结束键值:[startKey, endKey string] 注意：键值可以为空字符串，表示没有边界，全部查询
// # output：输出key:Value的键值对
// #======================================================================
func (t *SimpleChaincode) queryAllTxRecord(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2. ")
	}
	startKey := args[0]
	endKey := args[1]

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
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

	fmt.Printf("- queryAllTxRecords:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}
