// ============================================================================================================================
// 本智能合约用于ABS业务流程控制
// 功能包括：业务状态检测与跳转;业务动作处理
// ============================================================================================================================

package main

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
   pb "github.com/hyperledger/fabric/protos/peer"
)
//"strconv"
//"github.com/hyperledger/fabric/common/util"
var logger = shim.NewLogger("AbsProcess")

type SimpleChaincode struct {
}
process := []string{"ProInfoUpload","AssetSaleAgreementUpload","GuaranteeAgreementUpload","TrustManageementUpload","AssetRatingInstructionUpload","AccountOpinionUpload","CounselOpinionUpload","ProductPlanInstructionUpload","InferiorAssetObtain","InferiorAssetObtainRecording","SubprimeAssetObtain","SubprimeAssetsObtainRecording","PriorityAssetObtain","PriorityAssetObtainRecording","BreakAccountRecording"}
// ============================================================================================================================
// ClaimsPackageInfo struct:
// 包含所有信息的struct
type ClaimsPackageInfoStruct struct {

	  InitClaimsPackageInfo    InitClaimsPackageInfoStruct  `json:"InitClaimsPackageInfo"`
	  SaleAgreement            SaleAgreementStruct        `json:"SaleAgreement"`
	  GuaranteeAgrement        GuaranteeAgrementStruct    `json:"GuaranteeAgrement"`
	  ProductDesignAgreement   ProductDesignAgreementStruct   `json:"ProductDesignAgreement"`
	  AssetRatingInstruction   AssetRatingInstructionStruct   `json:"AssetRatingInstruction"`
	  AccountOpinion           AccountOpinionStruct          `json:"AccountOpinion"`
	  LegalOpinion             LegalOpinionStruct   `json:"LegalOpinion"`
	  ProductPlanInstruction   ProductPlanInstructionStruct   `json:"ProductPlanInstruction"`
	  InferiorAssetSubscriptionAgreement    InferiorAssetSubscriptionAgreementStruct   `json:"InferiorAssetSubscriptionAgreement"`
	  SubprimeAssetSubscriptionAgreement    SubprimeAssetSubscriptionAgreementStruct   `json:"SubprimeAssetSubscriptionAgreement"`
	  PriorityAssetSubscriptionAgreement    PriorityAssetSubscriptionAgreementStruct   `json:"PriorityAssetSubscriptionAgreement"`
		Status    string `json:"Status"`
	}


	// 1.spv上传产品信息
	// ============================================================================================================================
	type InitClaimsPackageInfoStruct struct {
		ProductID               string `json:"ProductID"`
		ProductName             string `json:"ProductName"`
		ProductType             string `json:"ProductType"`
		BasicAssets             string `json:"BasicAssets"`
		ProjectScale            float64 `json:"ProjectScale"`
		Originators             string `json:"Originators"`
		Investor                []string `json:"Investor"`
	  ExpectedReturn          string `json:"ExpectedReturn"`
	  PaymentMethod           string `json:"PaymentMethod"`
	  TrustInstitution        string `json:"TrustInstitution"`
	  DifferenceComplement    string `json:"DifferenceComplement"`
	  AssetRatingAgency       string `json:"AssetRatingAgency"`
	  AccountFirm             string `json:"AccountFirm"`
	  LawOffice               string `json:"LawOffice"`
	  TrustManagementFee      float64 `json:"TrustManagementFee"`
	  AssetRatingFee          float64 `json:"AssetRatingFee"`
	  CounselFee              float64 `json:"CounselFee"`
	  AccountancyFee          float64 `json:"AccountancyFee"`
	  BasicCreditorInfo       BasicCreditorInfoStruct   `json:"BasicCreditorInfo"`
	  Remark                  string `json:"Remark"`
		CreatedTime          time.Time `json:"CreatedTime"`
	}
	//债权基础信息
	type BasicCreditorInfoStruct struct{
	  Url         string `json:"Url"`
	  Hashcode    string `json:"Hashcode"`
	}

	//2.发起人上传资产买卖协议
	type SaleAgreementStruct struct{
	  Url         string `json:"Url"`
	  Hashcode    string `json:"Hashcode"`
	}

	//3.差额补足人上传差额补足协议
	type GuaranteeAgrementStruct struct{
	  Url         string `json:"Url"`
	  Hashcode    string `json:"Hashcode"`
	}

	//4.SPV上传产品设计书
	type ProductDesignAgreementStruct struct{
	  Url         string `json:"Url"`
	  Hashcode    string `json:"Hashcode"`
	}

	//5.评级机构上传评级信息
	type AssetRatingInstructionStruct struct{
	  Url                       string `json:"Url"`
	  Hashcode                  string `json:"Hashcode"`
	  PriorityAssetRatio        string `json:"PriorityAssetRatio"`
	  SubprimeAssetRatio        string `json:"SubprimeAssetRatio"`
	  InferiorAssetRatio        string `json:"InferiorAssetRatio"`
	  PriorityAssetRating       string `json:"PriorityAssetRating"`
	  SubprimeAssetsRating      string `json:"SubprimeAssetsRating"`
	}

	//6.会计事务所上传审计信息
	type AccountOpinionStruct struct{
	  Url         string `json:"Url"`
	  Hashcode    string `json:"Hashcode"`
	}

	//7.律师事务所上传法律意见书
	type LegalOpinionStruct struct{
	  Url         string `json:"Url"`
	  Hashcode    string `json:"Hashcode"`
	}

	// 8.spv上传产品计划说明书
	type ProductPlanInstructionStruct struct{
	  Url         string `json:"Url"`
	  Hashcode    string `json:"Hashcode"`
	}

	//9.劣后级资产被确定认购，并上传认购协议书
	type InferiorAssetSubscriptionAgreementStruct struct{
	  Url         string `json:"Url"`
	  Hashcode    string `json:"Hashcode"`
	}

	//10.代币划账记录

	//11.次优级资产被确定认购，并上传认购协议书
	type SubprimeAssetSubscriptionAgreementStruct struct{
	  Url         string `json:"Url"`
	  Hashcode    string `json:"Hashcode"`
	}

	//12.代币划账记录

	//13.优先级资产被确定认购，并上传认购协议书
	type PriorityAssetSubscriptionAgreementStruct struct{
	  Url         string `json:"Url"`
	  Hashcode    string `json:"Hashcode"`
	}
	//14.代币划账记录

	//15.分账记录

// //UrlAndHashStruct
// 	type UrlAndHashStruct struct{
// 	  Url         string `json:"Url"`
// 	  Hashcode    string `json:"Hashcode"`
// 	}


// ============================================================================================================================
// TransferRecord  struct:
// 包含所有信息的struct
	type TransferRecordStruct struct {
		ProductID            string  `json:"ProductID"`
		WaterFlowNumber      string  `json:"WaterFlowNumber"`
    WaterFlowNumberTime  string  `json:"WaterFlowNumberTime"`
		FromAccount          string  `json:"FromAccount"`
		ToAccount            string  `json:"ToAccount"`
		BbMount              float64 `json:"BbMount"`
	}

// ============================================================================================================================
// Init - emtpy
// ============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### AbsProcess Init ###########")
	return shim.Success(nil)
}

// ============================================================================================================================
// Query - emtpy
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Error("Unknown supported call")
}

// ============================================================================================================================
// Invoke function is the entry point for Invocations
// ============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

  function, args := stub.GetFunctionAndParameters()
	// Handle different functions
	if function == "assetSaleAgreementUpload" {
		return t.assetSaleAgreementUpload(stub, args)
	}
	// else if function == "guaranteeAgreementUpload" {
	// 	return t.guaranteeAgreementUpload(stub, args)
	// } else if function == "trustManageementUpload" {
	// 	return t.trustManageementUpload(stub, args)
	// } else if function == "assetRatingInstructionUpload" {
	// 	return t.assetRatingInstructionUpload(stub, args)
	// } else if function == "accountOpinionUpload" {
	// 	return t.accountOpinionUpload(stub, args)
	// } else if function == "counselOpinionUpload" {
	// 	return t.counselOpinionUpload(stub, args)
	// } else if function == "productPlanInstructionUpload" {
	// 	return t.productPlanInstructionUpload(stub, args)
	// } else if function == "inferiorAssetObtain" {
	// 	return t.inferiorAssetObtain(stub, args)
	// }else if function == "inferiorAssetObtainRecording" {
	// 	return t.inferiorAssetObtainRecording(stub, args)
	// }else if function == "subprimeAssetObtain" {
	// 	return t.subprimeAssetObtain(stub, args)
	// }else if function == "subprimeAssetsObtainRecording" {
	// 	return t.subprimeAssetsObtainRecording(stub, args)
	// }else if function == "priorityAssetObtain" {
	// 	return t.priorityAssetObtain(stub, args)
	// }else if function == "priorityAssetObtainRecording" {
	// 	return t.priorityAssetObtainRecording(stub, args)
	// }else if function == "breakAccountRecording" {
	// 	return t.breakAccountRecording(stub, args)
	// }

	return shim.Error("Received unknown function invocation")
}

// ============================================================================================================================
// function:发起人上传资产买卖协议（url和hash值）
// input：ProductID,ulr,hash
// ============================================================================================================================
func (t *SimpleChaincode) assetSaleAgreementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  var err error
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	ProductID := args[0]
	UrlAndHashInfo := args[1]

	// currentStatus, err := checkStatus(stub, ProductID)
	// if currentStatus != "ProInfoUpload" {
	// 	return shim.Error("status is worry! Expect: ProInfoUpload")
	// }
	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
	if err != nil {
		return shim.Error(err.Error())
	}
	ClaimsPackageInfoObj := ClaimsPackageInfoStruct{}
	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfoObj)
	if( ClaimsPackageInfoObj.Status != "ProInfoUpload" ){
		return shim.Error("Error Status!")
	}

	SaleAgreementObj := SaleAgreementStruct{}
	err = json.Unmarshal([]byte(UrlAndHashInfo),&SaleAgreementObj)
	if err != nil {
	  return shim.Error(err.Error())
	}
	ClaimsPackageInfoObj.SaleAgreement = SaleAgreementObj

	ClaimsPackageInfoObj.Status = "AssetSaleAgreementUpload"

  ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfoObj)
	if err != nil {
		return shim.Error(err.Error())
	}
  err = stub.PutState(ProductID, ClaimsPackageInfoAsBytes)
  if err != nil{
		return shim.Error(err.Error())
	}
  fmt.Println("AssetSaleAgreementUpload done")
	return shim.Success(nil)
}

// ============================================================================================================================
// function:差额补足人上传差额补足协议（url和hash值）
// input：ProductID,ulr,hash
// ============================================================================================================================
func (t *SimpleChaincode) guaranteeAgreementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductID := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductID)
	if currentStatus != "AssetSaleAgreementUpload" {
		return shim.Error("status is worry! Expect: AssetSaleAgreementUpload")
	}

  err = agreementUpload(ProductID, "GuaranteeAgrement", Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

  err = changeStatus(ProductID)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:spv上传产品设计书（url和hash值）
// input：ProductID,ulr,hash
// ============================================================================================================================
func (t *SimpleChaincode) trustManageementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductID := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductID)
	if currentStatus != "GuaranteeAgreementUpload" {
		return shim.Error("status is worry! Expect: GuaranteeAgreementUpload")
	}

  err = agreementUpload(ProductID, "ProductDesignAgreement", Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

	err = changeStatus(ProductID)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:资产评级机构上传资产评级报告（url和hash值，资产评级信息）
// input：ProductID,ulr,hash，PriorityAssetRatio，SubprimeAssetRatio，InferiorAssetRatio，PriorityAssetRating，SubprimeAssetsRating
// ============================================================================================================================
func (t *SimpleChaincode) assetRatingInstructionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}
	ProductID := args[0]
	Url := args[1]
	Hashcode := args[2]
	PriorityAssetRatio := args[4]
	SubprimeAssetRatio := args[5]
	InferiorAssetRatio := args[6]
	PriorityAssetRating := args[7]
	SubprimeAssetsRating := args[8]

	currentStatus, err := checkStatus(ProductID)
	if currentStatus != "TrustManageementUpload" {
		return shim.Error("status is worry! Expect: TrustManageementUpload")
	}

	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
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
  err = stub.PutState(ProductID, ClaimsPackageInfoAsBytes)
  if err != nil{
		return shim.Error("Fail to marshal ClaimsPackageInfo")
	}

	err = changeStatus(ProductID)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:会计事务所上传会计审计报告（url和hash值）
// input：ProductID,ulr,hash
// ============================================================================================================================
func (t *SimpleChaincode) accountOpinionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductID := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductID)
	if currentStatus != "AssetRatingInstructionUpload" {
		return shim.Error("status is worry! Expect: AssetRatingInstructionUpload")
	}

  err = agreementUpload(ProductID, "AccountOpinion", Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

	err = changeStatus(ProductID)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:律师事务所上传法律意见书（url和hash值）
// input：ProductID,ulr,hash
// ============================================================================================================================
func (t *SimpleChaincode) counselOpinionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductID := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductID)
	if currentStatus != "AccountOpinionUpload" {
		return shim.Error("status is worry! Expect: AccountOpinionUpload")
	}

  err = agreementUpload(ProductID, "LegalOpinion", Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

	err = changeStatus(ProductID)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function：spv上传产品计划说明书（url和hash值）
// input：ProductID,ulr,hash
// ============================================================================================================================
func (t *SimpleChaincode) productPlanInstructionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductID := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductID)
	if currentStatus != "CounselOpinionUpload" {
		return shim.Error("status is worry! Expect: CounselOpinionUpload")
	}

  err = agreementUpload(ProductID, "ProductPlanInstruction", Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

	err = changeStatus(ProductID)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:劣后级资产购买方上传资产买卖协议（url和hash值）
// input：ProductID,ulr,hash
// ============================================================================================================================
func (t *SimpleChaincode) inferiorAssetObtain(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductID := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductID)
	if currentStatus != "ProductPlanInstructionUpload" {
		return shim.Error("status is worry! Expect: ProductPlanInstructionUpload")
	}

  err = agreementUpload(ProductID, "InferiorAssetSubscriptionAgreement", Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

	err = changeStatus(ProductID)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:代币节点记录转账:劣后级认购
// input：  key: RecordID
//        value: ProductID,WaterFlowNumber, WaterFlowNumberTime, FromAccount, ToAccount string, BbMount float64
// ============================================================================================================================
func (t *SimpleChaincode) inferiorAssetObtainRecording(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}
	RecordID := args[0]
	ProductID := args[1]
  WaterFlowNumber := args[2]
	WaterFlowNumberTime := args[3]
	FromAccount := args[4]
	ToAccount := args[5]
	BbMount := args[6]

  currentStatus, err := checkStatus(ProductID)
	if currentStatus != "InferiorAssetObtain" {
		return shim.Error("status is worry! Expect: InferiorAssetObtain")
	}

	err = addTransfeRecord(RecordID, ProductID, WaterFlowNumber, WaterFlowNumberTime, FromAccount, ToAccount, BbMount)
  if err != nil{
		return shim.Error("err")
	}

	err = changeStatus(ProductID)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:次优级资产购买方（url和hash值）
// input：ProductID,ulr,hash
// ============================================================================================================================
func (t *SimpleChaincode) subprimeAssetObtain(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductID := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductID)
	if currentStatus != "InferiorAssetObtainRecording" {
		return shim.Error("status is worry! Expect: InferiorAssetObtainRecording")
	}

  err = agreementUpload(ProductID, "SubprimeAssetSubscriptionAgreement", Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

	err = changeStatus(ProductID)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:代币节点记录转账：次优级认购
// input：
// ============================================================================================================================
func (t *SimpleChaincode) subprimeAssetsObtainRecording(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}
	RecordID := args[0]
	ProductID := args[1]
  WaterFlowNumber := args[2]
	WaterFlowNumberTime := args[3]
	FromAccount := args[4]
	ToAccount := args[5]
	BbMount := args[6]

  currentStatus, err := checkStatus(ProductID)
	if currentStatus != "SubprimeAssetObtain" {
		return shim.Error("status is worry! Expect: SubprimeAssetObtain")
	}

	err = addTransfeRecord(RecordID, ProductID, WaterFlowNumber, WaterFlowNumberTime, FromAccount, ToAccount, BbMount)
  if err != nil{
		return shim.Error("err")
	}

	err = changeStatus(ProductID)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:优先级资产购买方（url和hash值）
// input：ProductID,ulr,hash
// ============================================================================================================================
func (t *SimpleChaincode) priorityAssetObtain(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductID := args[0]
	Url := args[1]
	Hashcode := args[2]

	currentStatus, err := checkStatus(ProductID)
	if currentStatus != "SubprimeAssetsObtainRecording" {
		return shim.Error("status is worry! Expect: SubprimeAssetsObtainRecording")
	}

  err = agreementUpload(ProductID, "PriorityAssetSubscriptionAgreement", Url, Hashcode)
	if err != nil {
		return shim.Error("Fail to upload agreement")
	}

	err = changeStatus(ProductID)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:代币节点记录转账：优先级认购
// input：ProductID
// ============================================================================================================================
func (t *SimpleChaincode) priorityAssetObtainRecording(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}
	RecordID := args[0]
	ProductID := args[1]
  WaterFlowNumber := args[2]
	WaterFlowNumberTime := args[3]
	FromAccount := args[4]
	ToAccount := args[5]
	BbMount := args[6]

  currentStatus, err := checkStatus(ProductID)
	if currentStatus != "PriorityAssetObtain" {
		return shim.Error("status is worry! Expect: PriorityAssetObtain")
	}

	err = addTransfeRecord(RecordID, ProductID, WaterFlowNumber, WaterFlowNumberTime, FromAccount, ToAccount, BbMount)
  if err != nil{
		return shim.Error("err")
	}

	err = changeStatus(ProductID)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
// function:代币节点记录转账：分帐
// input：ProductID
// ============================================================================================================================
func (t *SimpleChaincode) breakAccountRecording(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}
	RecordID := args[0]
	ProductID := args[1]
  WaterFlowNumber := args[2]
	WaterFlowNumberTime := args[3]
	FromAccount := args[4]
	ToAccount := args[5]
	BbMount := args[6]

  currentStatus, err := checkStatus(ProductID)
	if currentStatus != "PriorityAssetObtainRecording" {
		return shim.Error("status is worry! Expect: PriorityAssetObtainRecording")
	}

	err = addTransfeRecord(RecordID, ProductID, WaterFlowNumber, WaterFlowNumberTime, FromAccount, ToAccount, BbMount)
  if err != nil{
		return shim.Error("err")
	}

	err = changeStatus(ProductID)
	if err != nil {
		return shim.Error("Fail to change status")
	}

	return shim.Success(nil)
}


============================================================================================================================
function:检查当前的业务状态
input：ProductID
============================================================================================================================
func checkStatus(stub shim.ChaincodeStubInterface, ProductID string) ( currentStatus string, err error ){

	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
	if err != nil {
		return "checkStaus has error:", err
	}
	ClaimsPackageInfo := ClaimsPackageInfoStruct{}
	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfo)

	return ClaimsPackageInfo.Status, nil
}

// ============================================================================================================================
// function:根据当前的业务状态转换到下一个状态
// input：ProductID
// ============================================================================================================================
func changeStatus(stub shim.ChaincodeStubInterface, ProductID string) ( err error ) {

	process := []string{"ProInfoUpload","AssetSaleAgreementUpload","GuaranteeAgreementUpload","TrustManageementUpload","AssetRatingInstructionUpload","AccountOpinionUpload","CounselOpinionUpload","ProductPlanInstructionUpload","InferiorAssetObtain","InferiorAssetObtainRecording","SubprimeAssetObtain","SubprimeAssetsObtainRecording","PriorityAssetObtain","PriorityAssetObtainRecording","BreakAccountRecording"}
	currentStatus, err := checkStatus(stub, ProductID)
	var i int
	var status string
	for i, status = range process {
    if status == currentStatus {
        fmt.Printf("found \"%s\" at process[%d]\n", status, i)
        break
    }
  }
	if i == len(process){
		err := errors.New("Alreay last status!")
		return err
	}
	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
	if err != nil {
		return err
	}
	ClaimsPackageInfo := ClaimsPackageInfoStruct{}

	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfo)
  ClaimsPackageInfo.Status = process[i+1]

  ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfo)
	if err != nil {
		return err
	}
  stub.PutState(ProductID, ClaimsPackageInfoAsBytes)

	return nil
}

func agreementUpload(stub shim.ChaincodeStubInterface, ProductID string, AgreementName string, UrlAndHashInfo string) (err error) {

	UrlAndHashInfoObj := UrlAndHashStruct{}
	err = json.Unmarshal([]byte(UrlAndHashInfo),&UrlAndHashInfoObj)
	if err != nil {
	  return err
	}

  ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
	if err != nil {
		return err
	}
	ClaimsPackageInfoObj := ClaimsPackageInfoStruct{}
	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfoObj)

	ClaimsPackageInfoObj.SaleAgreement = UrlAndHashInfoObj
	// AgreementNameStruct := "ClaimsPackageInfoObj." + AgreementName
  // AgreementNameStruct = UrlAndHashInfoObj


  ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfoObj)
	if err != nil {
		return err
	}
  err = stub.PutState(ProductID, ClaimsPackageInfoAsBytes)
  if err != nil{
		return err
	}

	return nil
}

func addTransfeRecord(stub shim.ChaincodeStubInterface, RecordID, ProductID, WaterFlowNumber, WaterFlowNumberTime, FromAccount, ToAccount , BbMount string) (err error) {

	RecordTest, _:=  stub.GetState(RecordID)
	if RecordTest != nil {
		return errors.New("the record is existed")
	}
	TransfeRecordInfo := TransferRecordStruct{}

	TransfeRecordInfo.ProductID =  ProductID
  TransfeRecordInfo.WaterFlowNumber =  WaterFlowNumber
	TransfeRecordInfo.WaterFlowNumberTime =  WaterFlowNumberTime
  TransfeRecordInfo.FromAccount = FromAccount
	TransfeRecordInfo.ToAccount = ToAccount
	BbMountAsFloat, err := strconv.ParseFloat(BbMount, 64)
	if err != nil{
		return err
	}
	TransfeRecordInfo.BbMount = BbMountAsFloat

  TransfeRecordAsBytes, err = json.Marshal(TransfeRecordInfo)
	if err != nil {
		return err
	}
  err = stub.PutState(RecordID, []byte(TransfeRecordAsBytes))
  if err != nil{
		return err
	}

	return err
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}
