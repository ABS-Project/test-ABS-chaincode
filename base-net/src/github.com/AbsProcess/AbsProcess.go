// ============================================================================================================================
// 本智能合约用于ABS业务流程控制
// 功能包括：业务状态检测与跳转;业务动作处理
// ============================================================================================================================

package main

import (
	"encoding/json"
  "errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"

	pb "github.com/hyperledger/fabric/protos/peer"

)
//"strconv"
//"github.com/hyperledger/fabric/common/util"
var logger = shim.NewLogger("AbsProcess")

type AbsProcess struct {
  }


// ============================================================================================================================
// TransferRecord  struct:
// 包含所有信息的struct
	type TransferRecordStruct struct {
		WaterFlowNumber      string  `json:"WaterFlowNumber"`
    WaterFlowNumberTime  string  `json:"WaterFlowNumberTime"`
		FromAccount          string  `json:"FromAccount"`
		ToAccount            string  `json:"ToAccount"`
		BbMount              float64 `json:"BbMount"`
	}

// ============================================================================================================================
// Init - emtpy
// ============================================================================================================================
func (t *AbsProcess) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// ============================================================================================================================
// Query - emtpy
// ============================================================================================================================
func (t *AbsProcess) Query(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Error("Unknown supported call")
}

// ============================================================================================================================
// Invoke function is the entry point for Invocations
// ============================================================================================================================
func (t *AbsProcess) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

  function, args := stub.GetFunctionAndParameters()
	// Handle different functions
	if function == "init" {
		return t.Init(stub)
	} else if function == "assetSaleAgreementUpload" {
		return t.assetSaleAgreementUpload(stub, args)
	} else if function == "guaranteeAgreementUpload" {
		return t.guaranteeAgreementUpload(stub, args)
	} else if function == "trustManageementUpload" {
		return t.trustManageementUpload(stub, args)
	} else if function == "assetRatingInstructionUpload" {
		return t.assetRatingInstructionUpload(stub, args)
	} else if function == "accountOpinionUpload" {
		return t.accountOpinionUpload(stub, args)
	} else if function == "counselOpinionUpload" {
		return t.counselOpinionUpload(stub, args)
	} else if function == "productPlanInstructionUpload" {
		return t.productPlanInstructionUpload(stub, args)
	} else if function == "inferiorAssetObtain" {
		return t.inferiorAssetObtain(stub, args)
	}else if function == "inferiorAssetObtainRecording" {
		return t.inferiorAssetObtainRecording(stub, args)
	}else if function == "subprimeAssetObtain" {
		return t.subprimeAssetObtain(stub, args)
	}else if function == "subprimeAssetsObtainRecording" {
		return t.subprimeAssetsObtainRecording(stub, args)
	}else if function == "priorityAssetObtain" {
		return t.priorityAssetObtain(stub, args)
	}else if function == "priorityAssetObtainRecording" {
		return t.priorityAssetObtainRecording(stub, args)
	}else if function == "breakAccountRecording" {
		return t.breakAccountRecording(stub, args)
	}

	return shim.Error("Received unknown function invocation")
}

// ============================================================================================================================
// function:发起人上传资产买卖协议（url和hash值）
// input：ProductName,ulr,hash
// ============================================================================================================================
func (t *AbsProcess) assetSaleAgreementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "ProInfoUpload" {
		return shim.Error("status is worry! Expect: ProInfoUpload")
	}

  err = agreementUpload(ProductName, SaleAgreement, Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

	_, err = changeStatus(ProductName)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:差额补足人上传差额补足协议（url和hash值）
// input：ProductName,ulr,hash
// ============================================================================================================================
func (t *AbsProcess) guaranteeAgreementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "AssetSaleAgreementUpload" {
		return shim.Error("status is worry! Expect: AssetSaleAgreementUpload")
	}

  err = agreementUpload(ProductName, GuaranteeAgrement, Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

  _, err = changeStatus(ProductName)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:spv上传产品设计书（url和hash值）
// input：ProductName,ulr,hash
// ============================================================================================================================
func (t *AbsProcess) trustManageementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "GuaranteeAgreementUpload" {
		return shim.Error("status is worry! Expect: GuaranteeAgreementUpload")
	}

  err = agreementUpload(ProductName, ProductDesignAgreement, Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

	_, err = changeStatus(ProductName)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:资产评级机构上传资产评级报告（url和hash值，资产评级信息）
// input：ProductName,ulr,hash，PriorityAssetRatio，SubprimeAssetRatio，InferiorAssetRatio，PriorityAssetRating，SubprimeAssetsRating
// ============================================================================================================================
func (t *AbsProcess) assetRatingInstructionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}
	ProductName := args[0]
	Url := args[1]
	Hashcode := args[2]
	PriorityAssetRatio := args[4]
	SubprimeAssetRatio := args[5]
	InferiorAssetRatio := args[6]
	PriorityAssetRating := args[7]
	SubprimeAssetsRating := args[8]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "TrustManageementUpload" {
		return shim.Error("status is worry! Expect: TrustManageementUpload")
	}

	ClaimsPackageInfoAsBytes, err :=  stub.Getstatus(ProductName)
	if err != nil {
		shim.Error("Fail to get ClaimsPackageInfo")
	}
	ClaimsPackageInfo := ClaimsPackageInfoStruct{}

	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfo)
  ClaimsPackageInfo.AssetRatingInstruction.Url =  Url
	ClaimsPackageInfo.AssetRatingInstruction.Hashcode =  Hashcode
	ClaimsPackageInfo.AssetRatingInstruction.PriorityAssetRatio =  PriorityAssetRatio
	ClaimsPackageInfo.AssetRatingInstruction.SubprimeAssetRatio =  SubprimeAssetRatio
	ClaimsPackageInfo.AssetRatingInstruction.InferiorAssetRatio =  InferiorAssetRatio
	ClaimsPackageInfo.AssetRatingInstruction.PriorityAssetRating = PriorityAssetRating
	ClaimsPackageInfo.AssetRatingInstruction.SubprimeAssetsRating = SubprimeAssetsRating

  ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfo)
	if err != nil {
		shim.Error("Fail to marshal ClaimsPackageInfo")
	}
  err = stub.PutState(ProductName, ClaimsPackageInfoAsBytes)
  if err != nil{
		return shim.Error("Fail to marshal ClaimsPackageInfo")
	}

	_, err = changeStatus(ProductName)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:会计事务所上传会计审计报告（url和hash值）
// input：ProductName,ulr,hash
// ============================================================================================================================
func (t *AbsProcess) accountOpinionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "AssetRatingInstructionUpload" {
		return shim.Error("status is worry! Expect: AssetRatingInstructionUpload")
	}

  err = agreementUpload(ProductName, AccountOpinion, Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

	_, err = changeStatus(ProductName)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:律师事务所上传法律意见书（url和hash值）
// input：ProductName,ulr,hash
// ============================================================================================================================
func (t *AbsProcess) counselOpinionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "AccountOpinionUpload" {
		return shim.Error("status is worry! Expect: AccountOpinionUpload")
	}

  err = agreementUpload(ProductName, LegalOpinion, Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

	_, err = changeStatus(ProductName)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function：spv上传产品计划说明书（url和hash值）
// input：ProductName,ulr,hash
// ============================================================================================================================
func (t *AbsProcess) productPlanInstructionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "CounselOpinionUpload" {
		return shim.Error("status is worry! Expect: CounselOpinionUpload")
	}

  err = agreementUpload(ProductName, ProductPlanInstruction, Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

	_, err = changeStatus(ProductName)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:劣后级资产购买方上传资产买卖协议（url和hash值）
// input：ProductName,ulr,hash
// ============================================================================================================================
func (t *AbsProcess) inferiorAssetObtain(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "ProductPlanInstructionUpload" {
		return shim.Error("status is worry! Expect: ProductPlanInstructionUpload")
	}

  err = agreementUpload(ProductName, InferiorAssetSubscriptionAgreement, Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

	_, err = changeStatus(ProductName)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:代币节点记录转账:劣后级认购
// input：ProductName,WaterFlowNumber, WaterFlowNumberTime, FromAccount, ToAccount string, BbMount float64
// ============================================================================================================================
func (t *AbsProcess) inferiorAssetObtainRecording(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}
	ProductName := args[0]
  WaterFlowNumber := args[1]
	WaterFlowNumberTime := args[2]
	FromAccount := args[3]
	ToAccount := args[4]
	BbMount := args[5]

  currentStatus, err := checkStatus(ProductName)
	if currentStatus != "InferiorAssetObtain" {
		return shim.Error("status is worry! Expect: InferiorAssetObtain")
	}

	err = addTransfeRecord(ProductName, WaterFlowNumber, WaterFlowNumberTime, FromAccount, ToAccount, BbMount)
  if err != nil{
		return shim.Error(err)
	}

	_, err = changeStatus(ProductName)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:次优级资产购买方（url和hash值）
// input：ProductName,ulr,hash
// ============================================================================================================================
func (t *AbsProcess) subprimeAssetObtain(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "InferiorAssetObtainRecording" {
		return shim.Error("status is worry! Expect: InferiorAssetObtainRecording")
	}

  err = agreementUpload(ProductName, SubprimeAssetSubscriptionAgreement, Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

	_, err = changeStatus(ProductName)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:代币节点记录转账：次优级认购
// input：
// ============================================================================================================================
func (t *AbsProcess) subprimeAssetsObtainRecording(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}
	ProductName := args[0]
  WaterFlowNumber := args[1]
	WaterFlowNumberTime := args[2]
	FromAccount := args[3]
	ToAccount := args[4]
	BbMount := args[5]

  currentStatus, err := checkStatus(ProductName)
	if currentStatus != "SubprimeAssetObtain" {
		return shim.Error("status is worry! Expect: SubprimeAssetObtain")
	}

	err = addTransfeRecord(ProductName, WaterFlowNumber, WaterFlowNumberTime, FromAccount, ToAccount, BbMount)
  if err != nil{
		return shim.Error(err)
	}

	_, err = changeStatus(ProductName)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:优先级资产购买方（url和hash值）
// input：ProductName,ulr,hash
// ============================================================================================================================
func (t *AbsProcess) priorityAssetObtain(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "SubprimeAssetsObtainRecording" {
		return shim.Error("status is worry! Expect: SubprimeAssetsObtainRecording")
	}

  err = agreementUpload(ProductName, PriorityAssetSubscriptionAgreement, Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

	_, err = changeStatus(ProductName)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:代币节点记录转账：优先级认购
// input：ProductName
// ============================================================================================================================
func (t *AbsProcess) priorityAssetObtainRecording(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}
	ProductName := args[0]
  WaterFlowNumber := args[1]
	WaterFlowNumberTime := args[2]
	FromAccount := args[3]
	ToAccount := args[4]
	BbMount := args[5]

  currentStatus, err := checkStatus(ProductName)
	if currentStatus != "PriorityAssetObtain" {
		return shim.Error("status is worry! Expect: PriorityAssetObtain")
	}

	err = addTransfeRecord(ProductName, WaterFlowNumber, WaterFlowNumberTime, FromAccount, ToAccount, BbMount)
  if err != nil{
		return shim.Error(err)
	}

	_, err = changeStatus(ProductName)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:代币节点记录转账：分帐
// input：ProductName
// ============================================================================================================================
func (t *AbsProcess) breakAccountRecording(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}
	ProductName := args[0]
  WaterFlowNumber := args[1]
	WaterFlowNumberTime := args[2]
	FromAccount := args[3]
	ToAccount := args[4]
	BbMount := args[5]

  currentStatus, err := checkStatus(ProductName)
	if currentStatus != "PriorityAssetObtainRecording" {
		return shim.Error("status is worry! Expect: PriorityAssetObtainRecording")
	}

	err = addTransfeRecord(ProductName, WaterFlowNumber, WaterFlowNumberTime, FromAccount, ToAccount, BbMount)
  if err != nil{
		return shim.Error(err)
	}

	_, err = changeStatus(ProductName)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}


// ============================================================================================================================
// function:检查当前的业务状态
// input：ProductName
// ============================================================================================================================
func checkStatus(ProductName string) (string, err error){

	ClaimsPackageInfoAsBytes, err :=  stub.Getstatus(ProductName)
	if err != nil {
		return nil, err
	}
	ClaimsPackageInfo := ClaimsPackageInfoStruct{}
	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfo)

	return ClaimsPackageInfo.Status, nil
}

// ============================================================================================================================
// function:根据当前的业务状态转换到下一个状态
// input：ProductName
// ============================================================================================================================
func (t *AbsProcess) changeStatus(ProductName string) (string, err error) {

	process := []string{"ProInfoUpload","AssetSaleAgreementUpload","GuaranteeAgreementUpload","TrustManageementUpload","AssetRatingInstructionUpload","AccountOpinionUpload","CounselOpinionUpload","ProductPlanInstructionUpload","InferiorAssetObtain","InferiorAssetObtainRecording","SubprimeAssetObtain","SubprimeAssetsObtainRecording","PriorityAssetObtain","PriorityAssetObtainRecording","BreakAccountRecording"}
	currentStatus, err := checkStatus(ProductName)
	for i, status := range process {
    if status == currentStatus {
        fmt.Printf("found \"%s\" at process[%d]\n", status, i)
        break
    }
  }
	if i == len(process){
		err := errors.New("Alreay last status!")
		return nil, err
	}
	ClaimsPackageInfoAsBytes, err :=  stub.Getstatus(ProductName)
	if err != nil {
		return nil, err
	}
	ClaimsPackageInfo := ClaimsPackageInfoStruct{}

	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfo)
  ClaimsPackageInfo.Status = process[i+1]

  ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfo)
	if err != nil {
		return nil, err
	}
  stub.PutState(ProductName, ClaimsPackageInfoAsBytes)

	return ClaimsPackageInfo.Status, nil
}

func (t *AbsProcess) agreementUpload(ProductName string, AgreementName string, Url string, Hashcode string) (err error) {

	ClaimsPackageInfoAsBytes, err :=  stub.Getstatus(ProductName)
	if err != nil {
		return err
	}
	ClaimsPackageInfo := ClaimsPackageInfoStruct{}

	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfo)
  ClaimsPackageInfo.AgreementName.Url =  Url
	ClaimsPackageInfo.AgreementName.Hashcode =  Hashcode

  ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfo)
	if err != nil {
		return err
	}
  err = stub.PutState(ProductName, ClaimsPackageInfoAsBytes)
  if err != nil{
		return err
	}

	return nil
}

func (t *AbsProcess) addTransfeRecord(ProductName, WaterFlowNumber, WaterFlowNumberTime, FromAccount, ToAccount string, BbMount float64) (err error) {

	TransfeRecordAsBytes, err :=  stub.Getstatus(ProductName)
	if err != nil {
		return err
	}
	TransfeRecordInfo := TransferRecordStruct{}

	json.Unmarshal(TransfeRecordAsBytes, &TransfeRecordInfo)
  TransfeRecordInfo.WaterFlowNumber =  Url
	TransfeRecordInfo.WaterFlowNumberTime =  Hashcode
  TransfeRecordInfo.FromAccount = FromAccount
	TransfeRecordInfo.ToAccount = FromAccount
	TransfeRecordInfo.BbMount = FromAccount

  TransfeRecordAsBytes, err = json.Marshal(TransfeRecordInfo)
	if err != nil {
		return err
	}
  err = stub.PutState(ProductName, TransfeRecordAsBytes)
  if err != nil{
		return err
	}

	return err
}
