//
//  Copyright Tongji University. All Rights Reserved.
//  用于操作记录的添加和查询
//  SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("BusinessPartnerInfo")


type SimpleChaincode struct {
}

// ============================================================================================================================
// BusinessPartnerInfo struct
// ============================================================================================================================
type BusinessPartnerInfoStruct struct {
	UserName     string    `json:"UserName"`
	Organization string    `json:"Organization"`
	Company      string    `json:"Company"`
	Account      string    `json:"Account"`
	CreatedTime  time.Time `json:"CreatedTime"`
	OperateLog   []string  `json:"OperateLog"`
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### BusinessPartnerInfo Init ###########")
	return shim.Success(nil)

}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### BusinessPartnerInfo Invoke ###########")

	function, args := stub.GetFunctionAndParameters()
	if function == "addBusinessPartnerInfo" {
		// 增加商业合作伙伴
		return t.addBusinessPartnerInfo(stub, args)
	}
	if function == "addOperateLog" {
		// 追加商业合作伙伴的操作TxID
		return t.addOperateLog(stub, args)
	}
	if function == "deleteBusinessPartnerInfo" {
		// 删除商业合作伙伴
		return t.deleteBusinessPartnerInfo(stub, args)
	}

	if function == "queryBusinessPartnerInfo" {
		//根据用户名查询商业合作伙伴信息
		return t.queryBusinessPartnerInfo(stub, args)
	}
	if function == "updateBusinessPartnerInfo" {
		//更新商业合作伙伴信息，创建日期和操作日志不可更改
		return t.updateBusinessPartnerInfo(stub, args)
	}

	logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0]))
}

// 增加商业合作伙伴
// 一个参数，字符串化的用户信息json
func (t *SimpleChaincode) addBusinessPartnerInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2. ")
	}
	var PartnerInfoObj BusinessPartnerInfoStruct
	PartnerInfo := args[0]
	err = json.Unmarshal([]byte(PartnerInfo), &PartnerInfoObj)
	if err != nil {
		fmt.Println("error:", err)
		return shim.Error(err.Error())
	}
	UserName := PartnerInfoObj.UserName
	UserTest, _ := stub.GetState(UserName)
	if UserTest != nil {
		return shim.Error("the user is existed")
	}
	timestamp, _ := stub.GetTxTimestamp()
	PartnerInfoObj.CreatedTime = time.Unix(timestamp.Seconds, int64(timestamp.Nanos))
	var OperateLog []string
	// TxID := stub.GetTxID()
	// OperateLog =append(OperateLog,TxID)
	PartnerInfoObj.OperateLog = OperateLog
	jsonAsBytes, _ := json.Marshal(PartnerInfoObj)
	err = stub.PutState(UserName, []byte(jsonAsBytes))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

// 追加商业合作伙伴的操作TxID
// 两个参数：UserName，TxID
func (t *SimpleChaincode) addOperateLog(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2. ")
	}
	var PartnerInfoObj BusinessPartnerInfoStruct
	UserName := args[0]
	NewOperateLog := args[1]
	OldUserInfo, _ := stub.GetState(UserName)
	if OldUserInfo == nil {
		return shim.Error("the user is not exist!!")
	}
	err = json.Unmarshal([]byte(OldUserInfo), &PartnerInfoObj)
	if err != nil {
		return shim.Error(err.Error())
	}
	PartnerInfoObj.OperateLog = append(PartnerInfoObj.OperateLog, NewOperateLog)
	jsonAsBytes, _ := json.Marshal(PartnerInfoObj)
	err = stub.PutState(UserName, []byte(jsonAsBytes))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

// 删除商业合作伙伴
// 一个参数：UserName
func (t *SimpleChaincode) deleteBusinessPartnerInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 ")
	}
	UserName := args[0]
	UserInfo, err := stub.GetState(UserName)
	//test if the user has been existed
	if err != nil {
		return shim.Error("The user never been exited")
	}
	if UserInfo == nil {
		return shim.Error("The user`s information is empty!")
	}
	err = stub.DelState(UserName) //remove the key from chaincode state
	if err != nil {
		return shim.Error("Failed to delete the user. ")
	}
	return shim.Success(nil)

}

// 根据用户名查询商业合作伙伴信息
// 一个参数：UserName
func (t *SimpleChaincode) queryBusinessPartnerInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}
	UserName := args[0]
	UserInfo, err := stub.GetState(UserName)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + UserName + "\"}"
		return shim.Error(jsonResp)
	}
	if UserInfo == nil {
		jsonResp := "{\"Error\":\"Nil content for " + UserName + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(UserInfo)

}

// 更新商业合作伙伴信息，创建日期和操作日志不可更改
//  一个参数，新的字符串化的用户信息json
func (t *SimpleChaincode) updateBusinessPartnerInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2. ")
	}
	var NewPartnerInfoObj BusinessPartnerInfoStruct
	var OldPartnerInfoObj BusinessPartnerInfoStruct
	NewPartnerInfo := args[0]
	err = json.Unmarshal([]byte(NewPartnerInfo), &NewPartnerInfoObj)
	if err != nil {
		fmt.Println("error:", err)
		return shim.Error(err.Error())
	}
	UserName := NewPartnerInfoObj.UserName
	OldUserInfo, _ := stub.GetState(UserName)
	if OldUserInfo == nil {
		return shim.Error("the user is not exist!!")
	}
	err = json.Unmarshal([]byte(OldUserInfo), &OldPartnerInfoObj)
	NewPartnerInfoObj.CreatedTime = OldPartnerInfoObj.CreatedTime
	NewPartnerInfoObj.OperateLog = OldPartnerInfoObj.OperateLog
	jsonAsBytes, _ := json.Marshal(NewPartnerInfoObj)
	err = stub.PutState(UserName, []byte(jsonAsBytes))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}
