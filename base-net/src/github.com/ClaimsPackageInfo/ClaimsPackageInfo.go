// 债券基础信息相关的链码操作

/*

*/

package main


import (
	"fmt"
	"encoding/json"
  "time"

	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("ClaimsPackageInfo")

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

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

const TxRecorderChaincodeName string = "TxRecorder"
const TxRecorderChaincodeChannel string = "mychannel"
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


func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response  {
	logger.Info("########### ClaimsPackageInfo Init ###########")
	return shim.Success(nil)
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### example_cc0 Invoke ###########")

	function, args := stub.GetFunctionAndParameters()
  if function == "proInfoUpload" {
		// Deletes an entity from its state
		return t.proInfoUpload(stub, args)
	}

	if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	}

	if function == "query" {
		// queries an entity state
		return t.query(stub, args)
	}
	if function == "update" {
		// Deletes an entity from its state
		return t.update(stub, args)
	}

	if function == "assetSaleAgreementUpload" {
		return t.assetSaleAgreementUpload(stub, args)
	}else if function == "guaranteeAgreementUpload" {
		return t.guaranteeAgreementUpload(stub, args)
	} else if function == "trustManageementUpload" {
		return t.trustManageementUpload(stub, args)
	}else if function == "assetRatingInstructionUpload" {
		return t.assetRatingInstructionUpload(stub, args)
	} else if function == "accountOpinionUpload" {
		return t.accountOpinionUpload(stub, args)
	}else if function == "counselOpinionUpload" {
		return t.counselOpinionUpload(stub, args)
	} else if function == "productPlanInstructionUpload" {
		return t.productPlanInstructionUpload(stub, args)
	}else if function == "inferiorAssetObtain" {
		return t.inferiorAssetObtain(stub, args)
	}
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

	logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0]))
}

func (t *SimpleChaincode) proInfoUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error


	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2. ")
	}
	var InitClaimsPackageInfoObj InitClaimsPackageInfoStruct
	InitClaimsPackageInfo :=args[1]
	err = json.Unmarshal([]byte(InitClaimsPackageInfo),&InitClaimsPackageInfoObj)
	if err != nil {
	fmt.Println("error:", err)
	return shim.Error(err.Error())
	 }
	ProductID := InitClaimsPackageInfoObj.ProductID
	ClaimsPackageInfo, _ := stub.GetState(ProductID)
	if ClaimsPackageInfo != nil {
		return shim.Error("the product is existed")
	}
	timestamp, _:= stub.GetTxTimestamp()
	CreatedTime := time.Unix(timestamp.Seconds, int64(timestamp.Nanos))
	InitClaimsPackageInfoObj.CreatedTime = CreatedTime
	var ClaimsPackageInfoObj ClaimsPackageInfoStruct
	ClaimsPackageInfoObj.Status = "ProInfoUpload"
	ClaimsPackageInfoObj.InitClaimsPackageInfo = InitClaimsPackageInfoObj
	jsonAsBytes,_:= json.Marshal(ClaimsPackageInfoObj)
	err = stub.PutState(ProductID,[]byte(jsonAsBytes))
	if err != nil {
		return shim.Error(err.Error())
	}
	// 现在开始记录操作
	var TxInfo [8]string
  TxInfo[0] = stub.GetTxID()                  //交易ID
	TxInfo[1] = args[0]                         //交易发起人
	TxInfo[2] = ProductID                       //产品ID
	TxInfo[3] = CreatedTime.Format("2006-01-02T15:04:05.000Z")      //交易时间
	TxInfo[4] = "ClaimsPackageInfo"             //链码名称
	TxInfo[5] = "add"                           //所调函数
	TxInfo[6] = args[1]                         //所传参数
	TxInfo[7] = "基础资产打包上传操作"             //交易描述
	functionName := "add"
	invokeArgs := util.ToChaincodeArgs(functionName,TxInfo[0],TxInfo[1],TxInfo[2],TxInfo[3],TxInfo[4],TxInfo[5],TxInfo[6],TxInfo[6])
	response := stub.InvokeChaincode(TxRecorderChaincodeName, invokeArgs, TxRecorderChaincodeChannel)
	if response.Status != shim.OK {
			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
			fmt.Printf(errStr)
			return shim.Error(errStr)
		}
	//fmt.Printf("Invoke chaincode successful. Got response %s", string(response))
  fmt.Println("success add a new TxInfo")
	//记录操作完成

	return shim.Success(nil);
}



// ============================================================================================================================
// function:发起人上传资产买卖协议（url和hash值）
// input：ProductID,UrlAndHashInfo
// ============================================================================================================================
func (t *SimpleChaincode) assetSaleAgreementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  var err error
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	ProductID := args[1]
	UrlAndHashInfo := args[2]

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
  err = stub.PutState(ProductID, []byte(ClaimsPackageInfoAsBytes))
  if err != nil{
		return shim.Error(err.Error())
	}
  fmt.Println("AssetSaleAgreementUpload done")

	// 现在开始记录操作
	var TxInfo [8]string
  TxInfo[0] = stub.GetTxID()                  //交易ID
	TxInfo[1] = args[0]                         //交易发起人
	TxInfo[2] = ProductID                       //产品ID
	TxInfo[3] = time.Now().Format("2006-01-02T15:04:05.000Z")      //交易时间
	TxInfo[4] = "ClaimsPackageInfo"             //链码名称
	TxInfo[5] = "AssetSaleAgreementUpload"     //所调函数
	TxInfo[6] = args[2]                         //所传参数
	TxInfo[7] = "发起人上传买卖协议"               //交易描述

	functionName := "add"
	invokeArgs := util.ToChaincodeArgs(functionName,TxInfo[0],TxInfo[1],TxInfo[2],TxInfo[3],TxInfo[4],TxInfo[5],TxInfo[6],TxInfo[7])
	response := stub.InvokeChaincode(TxRecorderChaincodeName, invokeArgs, TxRecorderChaincodeChannel)
	if response.Status != shim.OK {
			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
			fmt.Printf(errStr)
			return shim.Error(errStr)
		}
	//fmt.Printf("Invoke chaincode successful. Got response %s", string(response))
  fmt.Println("success add a new TxInfo")
	//记录操作完成

	return shim.Success(nil)
}

// ============================================================================================================================
// function:差额补足人上传差额补足协议（url和hash值）
// input：ProductID,UrlAndHashInfo
// ============================================================================================================================
func (t *SimpleChaincode) guaranteeAgreementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  var err error
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	ProductID := args[1]
	UrlAndHashInfo := args[2]

	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
	if err != nil {
		return shim.Error(err.Error())
	}
	ClaimsPackageInfoObj := ClaimsPackageInfoStruct{}
	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfoObj)
	if( ClaimsPackageInfoObj.Status != "AssetSaleAgreementUpload" ){
		return shim.Error("Error Status!")
	}

	GuaranteeAgrementObj := GuaranteeAgrementStruct{}
	err = json.Unmarshal([]byte(UrlAndHashInfo),&GuaranteeAgrementObj)
	if err != nil {
	  return shim.Error(err.Error())
	}
	ClaimsPackageInfoObj.GuaranteeAgrement = GuaranteeAgrementObj

	ClaimsPackageInfoObj.Status = "GuaranteeAgreementUpload"

  ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfoObj)
	if err != nil {
		return shim.Error(err.Error())
	}
  err = stub.PutState(ProductID, []byte(ClaimsPackageInfoAsBytes))
  if err != nil{
		return shim.Error(err.Error())
	}
  fmt.Println("GuaranteeAgreementUpload done")

	// 现在开始记录操作
	var TxInfo [8]string
  TxInfo[0] = stub.GetTxID()                  //交易ID
	TxInfo[1] = args[0]                         //交易发起人
	TxInfo[2] = ProductID                       //产品ID
	TxInfo[3] = time.Now().Format("2006-01-02T15:04:05.000Z")      //交易时间
	TxInfo[4] = "ClaimsPackageInfo"             //链码名称
	TxInfo[5] = "GuaranteeAgreementUpload"     //所调函数
	TxInfo[6] = args[2]                         //所传参数
	TxInfo[7] = "差额补足人上传差额补足协议"               //交易描述

	functionName := "add"
	invokeArgs := util.ToChaincodeArgs(functionName,TxInfo[0],TxInfo[1],TxInfo[2],TxInfo[3],TxInfo[4],TxInfo[5],TxInfo[6],TxInfo[7])
	response := stub.InvokeChaincode(TxRecorderChaincodeName, invokeArgs, TxRecorderChaincodeChannel)
	if response.Status != shim.OK {
			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
			fmt.Printf(errStr)
			return shim.Error(errStr)
		}
	//fmt.Printf("Invoke chaincode successful. Got response %s", string(response))
  fmt.Println("success add a new TxInfo")
	//记录操作完成

	return shim.Success(nil)
}


// ============================================================================================================================
// function:SPV上传产品设计书（url和hash值）
// input：ProductID,UrlAndHashInfo
// ============================================================================================================================
func (t *SimpleChaincode) trustManageementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  var err error
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	ProductID := args[1]
	UrlAndHashInfo := args[2]

	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
	if err != nil {
		return shim.Error(err.Error())
	}

	ClaimsPackageInfoObj := ClaimsPackageInfoStruct{}
	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfoObj)
	if( ClaimsPackageInfoObj.Status != "GuaranteeAgreementUpload" ){
		return shim.Error("Error Status!")
	}

	ProductDesignAgreementObj := ProductDesignAgreementStruct{}
	err = json.Unmarshal([]byte(UrlAndHashInfo),&ProductDesignAgreementObj)
	if err != nil {
	  return shim.Error(err.Error())
	}
	ClaimsPackageInfoObj.ProductDesignAgreement = ProductDesignAgreementObj

	ClaimsPackageInfoObj.Status = "TrustManageementUpload"

  ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfoObj)
	if err != nil {
		return shim.Error(err.Error())
	}
  err = stub.PutState(ProductID, []byte(ClaimsPackageInfoAsBytes))
  if err != nil{
		return shim.Error(err.Error())
	}
  fmt.Println("TrustManageementUpload done")

	// 现在开始记录操作
	var TxInfo [8]string
  TxInfo[0] = stub.GetTxID()                  //交易ID
	TxInfo[1] = args[0]                         //交易发起人
	TxInfo[2] = ProductID                       //产品ID
	TxInfo[3] = time.Now().Format("2006-01-02T15:04:05.000Z")      //交易时间
	TxInfo[4] = "ClaimsPackageInfo"             //链码名称
	TxInfo[5] = "trustManageementUpload"     //所调函数
	TxInfo[6] = args[2]                         //所传参数
	TxInfo[7] = "SPV上传产品设计书"               //交易描述

	functionName := "add"
	invokeArgs := util.ToChaincodeArgs(functionName,TxInfo[0],TxInfo[1],TxInfo[2],TxInfo[3],TxInfo[4],TxInfo[5],TxInfo[6],TxInfo[7])
	response := stub.InvokeChaincode(TxRecorderChaincodeName, invokeArgs, TxRecorderChaincodeChannel)
	if response.Status != shim.OK {
			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
			fmt.Printf(errStr)
			return shim.Error(errStr)
		}
	//fmt.Printf("Invoke chaincode successful. Got response %s", string(response))
  fmt.Println("success add a new TxInfo")
	//记录操作完成

	return shim.Success(nil)
}

// ============================================================================================================================
// function:资产评级机构上传资产评级（url和hash值）
// input：ProductID,AssetRatingInstructionStruct
// ============================================================================================================================
func (t *SimpleChaincode) assetRatingInstructionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  var err error
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	ProductID := args[1]
	UrlAndHashInfo := args[2]

	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
	if err != nil {
		return shim.Error(err.Error())
	}

	ClaimsPackageInfoObj := ClaimsPackageInfoStruct{}
	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfoObj)
	if( ClaimsPackageInfoObj.Status != "TrustManageementUpload" ){
		return shim.Error("Error Status!")
	}

	AssetRatingInstructionObj := AssetRatingInstructionStruct{}
	err = json.Unmarshal([]byte(UrlAndHashInfo),&AssetRatingInstructionObj)
	if err != nil {
	  return shim.Error(err.Error())
	}
	ClaimsPackageInfoObj.AssetRatingInstruction = AssetRatingInstructionObj

	ClaimsPackageInfoObj.Status = "AssetRatingInstructionUpload"

  ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfoObj)
	if err != nil {
		return shim.Error(err.Error())
	}
  err = stub.PutState(ProductID, []byte(ClaimsPackageInfoAsBytes))
  if err != nil{
		return shim.Error(err.Error())
	}
  fmt.Println("AssetSaleAgreementUpload done")

	// 现在开始记录操作
	var TxInfo [8]string
  TxInfo[0] = stub.GetTxID()                  //交易ID
	TxInfo[1] = args[0]                         //交易发起人
	TxInfo[2] = ProductID                       //产品ID
	TxInfo[3] = time.Now().Format("2006-01-02T15:04:05.000Z")      //交易时间
	TxInfo[4] = "ClaimsPackageInfo"             //链码名称
	TxInfo[5] = "AssetSaleAgreementUpload"     //所调函数
	TxInfo[6] = args[2]                         //所传参数
	TxInfo[7] = "资产评级机构上传资产评级信息"               //交易描述

	functionName := "add"
	invokeArgs := util.ToChaincodeArgs(functionName,TxInfo[0],TxInfo[1],TxInfo[2],TxInfo[3],TxInfo[4],TxInfo[5],TxInfo[6],TxInfo[7])
	response := stub.InvokeChaincode(TxRecorderChaincodeName, invokeArgs, TxRecorderChaincodeChannel)
	if response.Status != shim.OK {
			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
			fmt.Printf(errStr)
			return shim.Error(errStr)
		}
	//fmt.Printf("Invoke chaincode successful. Got response %s", string(response))
  fmt.Println("success add a new TxInfo")
	//记录操作完成

	return shim.Success(nil)
}

// ============================================================================================================================
// function:会计师事务所上传审计报告（url和hash值）
// input：ProductID,UrlAndHashInfo
// ============================================================================================================================
func (t *SimpleChaincode) accountOpinionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  var err error
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	ProductID := args[1]
	UrlAndHashInfo := args[2]

	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
	if err != nil {
		return shim.Error(err.Error())
	}

	ClaimsPackageInfoObj := ClaimsPackageInfoStruct{}
	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfoObj)
	if( ClaimsPackageInfoObj.Status != "AssetRatingInstructionUpload" ){
		return shim.Error("Error Status!")
	}

	AccountOpinionObj := AccountOpinionStruct{}
	err = json.Unmarshal([]byte(UrlAndHashInfo),&AccountOpinionObj)
	if err != nil {
	  return shim.Error(err.Error())
	}
	ClaimsPackageInfoObj.AccountOpinion = AccountOpinionObj

	ClaimsPackageInfoObj.Status = "AccountOpinionUpload"

  ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfoObj)
	if err != nil {
		return shim.Error(err.Error())
	}
  err = stub.PutState(ProductID, []byte(ClaimsPackageInfoAsBytes))
  if err != nil{
		return shim.Error(err.Error())
	}
  fmt.Println("AccountOpinionUpload done")

	// 现在开始记录操作
	var TxInfo [8]string
  TxInfo[0] = stub.GetTxID()                  //交易ID
	TxInfo[1] = args[0]                         //交易发起人
	TxInfo[2] = ProductID                       //产品ID
	TxInfo[3] = time.Now().Format("2006-01-02T15:04:05.000Z")      //交易时间
	TxInfo[4] = "ClaimsPackageInfo"             //链码名称
	TxInfo[5] = "accountOpinionUpload"     //所调函数
	TxInfo[6] = args[2]                         //所传参数
	TxInfo[7] = "会计师事务所上传审计报告"               //交易描述

	functionName := "add"
	invokeArgs := util.ToChaincodeArgs(functionName,TxInfo[0],TxInfo[1],TxInfo[2],TxInfo[3],TxInfo[4],TxInfo[5],TxInfo[6],TxInfo[7])
	response := stub.InvokeChaincode(TxRecorderChaincodeName, invokeArgs, TxRecorderChaincodeChannel)
	if response.Status != shim.OK {
			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
			fmt.Printf(errStr)
			return shim.Error(errStr)
		}
	//fmt.Printf("Invoke chaincode successful. Got response %s", string(response))
  fmt.Println("success add a new TxInfo")
	//记录操作完成

	return shim.Success(nil)
}

// ============================================================================================================================
// function:律师事务所上传法律意见书（url和hash值）
// input：ProductID,UrlAndHashInfo
// ============================================================================================================================
func (t *SimpleChaincode) counselOpinionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  var err error
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	ProductID := args[1]
	UrlAndHashInfo := args[2]

	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
	if err != nil {
		return shim.Error(err.Error())
	}

	ClaimsPackageInfoObj := ClaimsPackageInfoStruct{}
	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfoObj)
	if( ClaimsPackageInfoObj.Status != "AccountOpinionUpload" ){
		return shim.Error("Error Status!")
	}

	LegalOpinionObj :=  LegalOpinionStruct{}
	err = json.Unmarshal([]byte(UrlAndHashInfo),&LegalOpinionObj)
	if err != nil {
	  return shim.Error(err.Error())
	}
	ClaimsPackageInfoObj.LegalOpinion = LegalOpinionObj

	ClaimsPackageInfoObj.Status = "CounselOpinionUpload"

  ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfoObj)
	if err != nil {
		return shim.Error(err.Error())
	}
  err = stub.PutState(ProductID, []byte(ClaimsPackageInfoAsBytes))
  if err != nil{
		return shim.Error(err.Error())
	}
  fmt.Println("CounselOpinionUpload done")

	// 现在开始记录操作
	var TxInfo [8]string
  TxInfo[0] = stub.GetTxID()                  //交易ID
	TxInfo[1] = args[0]                         //交易发起人
	TxInfo[2] = ProductID                       //产品ID
	TxInfo[3] = time.Now().Format("2006-01-02T15:04:05.000Z")      //交易时间
	TxInfo[4] = "ClaimsPackageInfo"             //链码名称
	TxInfo[5] = "counselOpinionUpload"     //所调函数
	TxInfo[6] = args[2]                         //所传参数
	TxInfo[7] = "律师事务所上传法律意见书"               //交易描述

	functionName := "add"
	invokeArgs := util.ToChaincodeArgs(functionName,TxInfo[0],TxInfo[1],TxInfo[2],TxInfo[3],TxInfo[4],TxInfo[5],TxInfo[6],TxInfo[7])
	response := stub.InvokeChaincode(TxRecorderChaincodeName, invokeArgs, TxRecorderChaincodeChannel)
	if response.Status != shim.OK {
			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
			fmt.Printf(errStr)
			return shim.Error(errStr)
		}
	//fmt.Printf("Invoke chaincode successful. Got response %s", string(response))
  fmt.Println("success add a new TxInfo")
	//记录操作完成

	return shim.Success(nil)
}

// ============================================================================================================================
// function:SPV上传产品计划说明书（url和hash值）
// input：ProductID,UrlAndHashInfo
// ============================================================================================================================
func (t *SimpleChaincode) productPlanInstructionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  var err error
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	ProductID := args[1]
	UrlAndHashInfo := args[2]

	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
	if err != nil {
		return shim.Error(err.Error())
	}

	ClaimsPackageInfoObj := ClaimsPackageInfoStruct{}
	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfoObj)
	if( ClaimsPackageInfoObj.Status != "CounselOpinionUpload" ){
		return shim.Error("Error Status!")
	}

	ProductPlanInstructionObj := ProductPlanInstructionStruct{}
	err = json.Unmarshal([]byte(UrlAndHashInfo),&ProductPlanInstructionObj)
	if err != nil {
	  return shim.Error(err.Error())
	}
	ClaimsPackageInfoObj.ProductPlanInstruction = ProductPlanInstructionObj

	ClaimsPackageInfoObj.Status = "ProductPlanInstructionUpload"

  ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfoObj)
	if err != nil {
		return shim.Error(err.Error())
	}
  err = stub.PutState(ProductID, []byte(ClaimsPackageInfoAsBytes))
  if err != nil{
		return shim.Error(err.Error())
	}
  fmt.Println("ProductPlanInstructionUpload done")

	// 现在开始记录操作
	var TxInfo [8]string
  TxInfo[0] = stub.GetTxID()                  //交易ID
	TxInfo[1] = args[0]                         //交易发起人
	TxInfo[2] = ProductID                       //产品ID
	TxInfo[3] = time.Now().Format("2006-01-02T15:04:05.000Z")      //交易时间
	TxInfo[4] = "ClaimsPackageInfo"             //链码名称
	TxInfo[5] = "productPlanInstructionUpload"     //所调函数
	TxInfo[6] = args[2]                         //所传参数
	TxInfo[7] = "SPV上传产品计划说明书"               //交易描述

	functionName := "add"
	invokeArgs := util.ToChaincodeArgs(functionName,TxInfo[0],TxInfo[1],TxInfo[2],TxInfo[3],TxInfo[4],TxInfo[5],TxInfo[6],TxInfo[7])
	response := stub.InvokeChaincode(TxRecorderChaincodeName, invokeArgs, TxRecorderChaincodeChannel)
	if response.Status != shim.OK {
			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
			fmt.Printf(errStr)
			return shim.Error(errStr)
		}
	//fmt.Printf("Invoke chaincode successful. Got response %s", string(response))
  fmt.Println("success add a new TxInfo")
	//记录操作完成

	return shim.Success(nil)
}

// ============================================================================================================================
// function:劣后级资产购买方认购劣后级资产（url和hash值）
// input：ProductID,UrlAndHashInfo
// ============================================================================================================================
func (t *SimpleChaincode) inferiorAssetObtain(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  var err error
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	ProductID := args[1]
	UrlAndHashInfo := args[2]

	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
	if err != nil {
		return shim.Error(err.Error())
	}

	ClaimsPackageInfoObj := ClaimsPackageInfoStruct{}
	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfoObj)
	if( ClaimsPackageInfoObj.Status != "ProductPlanInstructionUpload" ){
		return shim.Error("Error Status!")
	}

	InferiorAssetSubscriptionAgreementObj := InferiorAssetSubscriptionAgreementStruct{}
	err = json.Unmarshal([]byte(UrlAndHashInfo),&InferiorAssetSubscriptionAgreementObj)
	if err != nil {
	  return shim.Error(err.Error())
	}
	ClaimsPackageInfoObj.InferiorAssetSubscriptionAgreement = InferiorAssetSubscriptionAgreementObj

	ClaimsPackageInfoObj.Status = "InferiorAssetObtain"

  ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfoObj)
	if err != nil {
		return shim.Error(err.Error())
	}
  err = stub.PutState(ProductID, []byte(ClaimsPackageInfoAsBytes))
  if err != nil{
		return shim.Error(err.Error())
	}
  fmt.Println("InferiorAssetObtain done")

	// 现在开始记录操作
	var TxInfo [8]string
  TxInfo[0] = stub.GetTxID()                  //交易ID
	TxInfo[1] = args[0]                         //交易发起人
	TxInfo[2] = ProductID                       //产品ID
	TxInfo[3] = time.Now().Format("2006-01-02T15:04:05.000Z")      //交易时间
	TxInfo[4] = "ClaimsPackageInfo"             //链码名称
	TxInfo[5] = "inferiorAssetObtain"     //所调函数
	TxInfo[6] = args[2]                         //所传参数
	TxInfo[7] = "劣后级资产购买方认购劣后级资产"               //交易描述

	functionName := "add"
	invokeArgs := util.ToChaincodeArgs(functionName,TxInfo[0],TxInfo[1],TxInfo[2],TxInfo[3],TxInfo[4],TxInfo[5],TxInfo[6],TxInfo[7])
	response := stub.InvokeChaincode(TxRecorderChaincodeName, invokeArgs, TxRecorderChaincodeChannel)
	if response.Status != shim.OK {
			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
			fmt.Printf(errStr)
			return shim.Error(errStr)
		}
	//fmt.Printf("Invoke chaincode successful. Got response %s", string(response))
  fmt.Println("success add a new TxInfo")
	//记录操作完成

	return shim.Success(nil)
}

// // ============================================================================================================================
// // function:发起人上传资产买卖协议（url和hash值）
// // input：ProductID,UrlAndHashInfo
// // ============================================================================================================================
// func (t *SimpleChaincode) assetSaleAgreementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//   var err error
// 	if len(args) != 3 {
// 		return shim.Error("Incorrect number of arguments. Expecting 3")
// 	}
//
// 	ProductID := args[1]
// 	UrlAndHashInfo := args[2]
//
// 	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	// timestamp, _:= stub.GetTxTimestamp()
// 	// CreatedTime := time.Unix(timestamp.Seconds, int64(timestamp.Nanos))
// 	ClaimsPackageInfoObj := ClaimsPackageInfoStruct{}
// 	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfoObj)
// 	if( ClaimsPackageInfoObj.Status != "ProInfoUpload" ){
// 		return shim.Error("Error Status!")
// 	}
//
// 	SaleAgreementObj := SaleAgreementStruct{}
// 	err = json.Unmarshal([]byte(UrlAndHashInfo),&SaleAgreementObj)
// 	if err != nil {
// 	  return shim.Error(err.Error())
// 	}
// 	ClaimsPackageInfoObj.SaleAgreement = SaleAgreementObj
//
// 	ClaimsPackageInfoObj.Status = "AssetSaleAgreementUpload"
//
//   ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfoObj)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
//   err = stub.PutState(ProductID, []byte(ClaimsPackageInfoAsBytes))
//   if err != nil{
// 		return shim.Error(err.Error())
// 	}
//   fmt.Println("AssetSaleAgreementUpload done")
//
// 	// 现在开始记录操作
// 	var TxInfo [8]string
//   TxInfo[0] = stub.GetTxID()                  //交易ID
// 	TxInfo[1] = args[0]                         //交易发起人
// 	TxInfo[2] = ProductID                       //产品ID
// 	TxInfo[3] = time.Now().Format("2006-01-02T15:04:05.000Z")      //交易时间
// 	TxInfo[4] = "ClaimsPackageInfo"             //链码名称
// 	TxInfo[5] = "AssetSaleAgreementUpload"     //所调函数
// 	TxInfo[6] = args[2]                         //所传参数
// 	TxInfo[7] = "发起人上传买卖协议"               //交易描述
//
// 	functionName := "add"
// 	invokeArgs := util.ToChaincodeArgs(functionName,TxInfo[0],TxInfo[1],TxInfo[2],TxInfo[3],TxInfo[4],TxInfo[5],TxInfo[6],TxInfo[7])
// 	response := stub.InvokeChaincode(TxRecorderChaincodeName, invokeArgs, TxRecorderChaincodeChannel)
// 	if response.Status != shim.OK {
// 			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
// 			fmt.Printf(errStr)
// 			return shim.Error(errStr)
// 		}
// 	//fmt.Printf("Invoke chaincode successful. Got response %s", string(response))
//   fmt.Println("success add a new TxInfo")
// 	//记录操作完成
//
// 	return shim.Success(nil)
// }
//
// // ============================================================================================================================
// // function:发起人上传资产买卖协议（url和hash值）
// // input：ProductID,UrlAndHashInfo
// // ============================================================================================================================
// func (t *SimpleChaincode) assetSaleAgreementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//   var err error
// 	if len(args) != 3 {
// 		return shim.Error("Incorrect number of arguments. Expecting 3")
// 	}
//
// 	ProductID := args[1]
// 	UrlAndHashInfo := args[2]
//
// 	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	// timestamp, _:= stub.GetTxTimestamp()
// 	// CreatedTime := time.Unix(timestamp.Seconds, int64(timestamp.Nanos))
// 	ClaimsPackageInfoObj := ClaimsPackageInfoStruct{}
// 	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfoObj)
// 	if( ClaimsPackageInfoObj.Status != "ProInfoUpload" ){
// 		return shim.Error("Error Status!")
// 	}
//
// 	SaleAgreementObj := SaleAgreementStruct{}
// 	err = json.Unmarshal([]byte(UrlAndHashInfo),&SaleAgreementObj)
// 	if err != nil {
// 	  return shim.Error(err.Error())
// 	}
// 	ClaimsPackageInfoObj.SaleAgreement = SaleAgreementObj
//
// 	ClaimsPackageInfoObj.Status = "AssetSaleAgreementUpload"
//
//   ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfoObj)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
//   err = stub.PutState(ProductID, []byte(ClaimsPackageInfoAsBytes))
//   if err != nil{
// 		return shim.Error(err.Error())
// 	}
//   fmt.Println("AssetSaleAgreementUpload done")
//
// 	// 现在开始记录操作
// 	var TxInfo [8]string
//   TxInfo[0] = stub.GetTxID()                  //交易ID
// 	TxInfo[1] = args[0]                         //交易发起人
// 	TxInfo[2] = ProductID                       //产品ID
// 	TxInfo[3] = time.Now().Format("2006-01-02T15:04:05.000Z")      //交易时间
// 	TxInfo[4] = "ClaimsPackageInfo"             //链码名称
// 	TxInfo[5] = "AssetSaleAgreementUpload"     //所调函数
// 	TxInfo[6] = args[2]                         //所传参数
// 	TxInfo[7] = "发起人上传买卖协议"               //交易描述
//
// 	functionName := "add"
// 	invokeArgs := util.ToChaincodeArgs(functionName,TxInfo[0],TxInfo[1],TxInfo[2],TxInfo[3],TxInfo[4],TxInfo[5],TxInfo[6],TxInfo[7])
// 	response := stub.InvokeChaincode(TxRecorderChaincodeName, invokeArgs, TxRecorderChaincodeChannel)
// 	if response.Status != shim.OK {
// 			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
// 			fmt.Printf(errStr)
// 			return shim.Error(errStr)
// 		}
// 	//fmt.Printf("Invoke chaincode successful. Got response %s", string(response))
//   fmt.Println("success add a new TxInfo")
// 	//记录操作完成
//
// 	return shim.Success(nil)
// }
//
// // ============================================================================================================================
// // function:发起人上传资产买卖协议（url和hash值）
// // input：ProductID,UrlAndHashInfo
// // ============================================================================================================================
// func (t *SimpleChaincode) assetSaleAgreementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//   var err error
// 	if len(args) != 3 {
// 		return shim.Error("Incorrect number of arguments. Expecting 3")
// 	}
//
// 	ProductID := args[1]
// 	UrlAndHashInfo := args[2]
//
// 	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	// timestamp, _:= stub.GetTxTimestamp()
// 	// CreatedTime := time.Unix(timestamp.Seconds, int64(timestamp.Nanos))
// 	ClaimsPackageInfoObj := ClaimsPackageInfoStruct{}
// 	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfoObj)
// 	if( ClaimsPackageInfoObj.Status != "ProInfoUpload" ){
// 		return shim.Error("Error Status!")
// 	}
//
// 	SaleAgreementObj := SaleAgreementStruct{}
// 	err = json.Unmarshal([]byte(UrlAndHashInfo),&SaleAgreementObj)
// 	if err != nil {
// 	  return shim.Error(err.Error())
// 	}
// 	ClaimsPackageInfoObj.SaleAgreement = SaleAgreementObj
//
// 	ClaimsPackageInfoObj.Status = "AssetSaleAgreementUpload"
//
//   ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfoObj)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
//   err = stub.PutState(ProductID, []byte(ClaimsPackageInfoAsBytes))
//   if err != nil{
// 		return shim.Error(err.Error())
// 	}
//   fmt.Println("AssetSaleAgreementUpload done")
//
// 	// 现在开始记录操作
// 	var TxInfo [8]string
//   TxInfo[0] = stub.GetTxID()                  //交易ID
// 	TxInfo[1] = args[0]                         //交易发起人
// 	TxInfo[2] = ProductID                       //产品ID
// 	TxInfo[3] = time.Now().Format("2006-01-02T15:04:05.000Z")      //交易时间
// 	TxInfo[4] = "ClaimsPackageInfo"             //链码名称
// 	TxInfo[5] = "AssetSaleAgreementUpload"     //所调函数
// 	TxInfo[6] = args[2]                         //所传参数
// 	TxInfo[7] = "发起人上传买卖协议"               //交易描述
//
// 	functionName := "add"
// 	invokeArgs := util.ToChaincodeArgs(functionName,TxInfo[0],TxInfo[1],TxInfo[2],TxInfo[3],TxInfo[4],TxInfo[5],TxInfo[6],TxInfo[7])
// 	response := stub.InvokeChaincode(TxRecorderChaincodeName, invokeArgs, TxRecorderChaincodeChannel)
// 	if response.Status != shim.OK {
// 			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
// 			fmt.Printf(errStr)
// 			return shim.Error(errStr)
// 		}
// 	//fmt.Printf("Invoke chaincode successful. Got response %s", string(response))
//   fmt.Println("success add a new TxInfo")
// 	//记录操作完成
//
// 	return shim.Success(nil)
// }
//
// // ============================================================================================================================
// // function:发起人上传资产买卖协议（url和hash值）
// // input：ProductID,UrlAndHashInfo
// // ============================================================================================================================
// func (t *SimpleChaincode) assetSaleAgreementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//   var err error
// 	if len(args) != 3 {
// 		return shim.Error("Incorrect number of arguments. Expecting 3")
// 	}
//
// 	ProductID := args[1]
// 	UrlAndHashInfo := args[2]
//
// 	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	// timestamp, _:= stub.GetTxTimestamp()
// 	// CreatedTime := time.Unix(timestamp.Seconds, int64(timestamp.Nanos))
// 	ClaimsPackageInfoObj := ClaimsPackageInfoStruct{}
// 	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfoObj)
// 	if( ClaimsPackageInfoObj.Status != "ProInfoUpload" ){
// 		return shim.Error("Error Status!")
// 	}
//
// 	SaleAgreementObj := SaleAgreementStruct{}
// 	err = json.Unmarshal([]byte(UrlAndHashInfo),&SaleAgreementObj)
// 	if err != nil {
// 	  return shim.Error(err.Error())
// 	}
// 	ClaimsPackageInfoObj.SaleAgreement = SaleAgreementObj
//
// 	ClaimsPackageInfoObj.Status = "AssetSaleAgreementUpload"
//
//   ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfoObj)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
//   err = stub.PutState(ProductID, []byte(ClaimsPackageInfoAsBytes))
//   if err != nil{
// 		return shim.Error(err.Error())
// 	}
//   fmt.Println("AssetSaleAgreementUpload done")
//
// 	// 现在开始记录操作
// 	var TxInfo [8]string
//   TxInfo[0] = stub.GetTxID()                  //交易ID
// 	TxInfo[1] = args[0]                         //交易发起人
// 	TxInfo[2] = ProductID                       //产品ID
// 	TxInfo[3] = time.Now().Format("2006-01-02T15:04:05.000Z")      //交易时间
// 	TxInfo[4] = "ClaimsPackageInfo"             //链码名称
// 	TxInfo[5] = "AssetSaleAgreementUpload"     //所调函数
// 	TxInfo[6] = args[2]                         //所传参数
// 	TxInfo[7] = "发起人上传买卖协议"               //交易描述
//
// 	functionName := "add"
// 	invokeArgs := util.ToChaincodeArgs(functionName,TxInfo[0],TxInfo[1],TxInfo[2],TxInfo[3],TxInfo[4],TxInfo[5],TxInfo[6],TxInfo[7])
// 	response := stub.InvokeChaincode(TxRecorderChaincodeName, invokeArgs, TxRecorderChaincodeChannel)
// 	if response.Status != shim.OK {
// 			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
// 			fmt.Printf(errStr)
// 			return shim.Error(errStr)
// 		}
// 	//fmt.Printf("Invoke chaincode successful. Got response %s", string(response))
//   fmt.Println("success add a new TxInfo")
// 	//记录操作完成
//
// 	return shim.Success(nil)
// }
//
// // ============================================================================================================================
// // function:发起人上传资产买卖协议（url和hash值）
// // input：ProductID,UrlAndHashInfo
// // ============================================================================================================================
// func (t *SimpleChaincode) assetSaleAgreementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//   var err error
// 	if len(args) != 3 {
// 		return shim.Error("Incorrect number of arguments. Expecting 3")
// 	}
//
// 	ProductID := args[1]
// 	UrlAndHashInfo := args[2]
//
// 	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	// timestamp, _:= stub.GetTxTimestamp()
// 	// CreatedTime := time.Unix(timestamp.Seconds, int64(timestamp.Nanos))
// 	ClaimsPackageInfoObj := ClaimsPackageInfoStruct{}
// 	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfoObj)
// 	if( ClaimsPackageInfoObj.Status != "ProInfoUpload" ){
// 		return shim.Error("Error Status!")
// 	}
//
// 	SaleAgreementObj := SaleAgreementStruct{}
// 	err = json.Unmarshal([]byte(UrlAndHashInfo),&SaleAgreementObj)
// 	if err != nil {
// 	  return shim.Error(err.Error())
// 	}
// 	ClaimsPackageInfoObj.SaleAgreement = SaleAgreementObj
//
// 	ClaimsPackageInfoObj.Status = "AssetSaleAgreementUpload"
//
//   ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfoObj)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
//   err = stub.PutState(ProductID, []byte(ClaimsPackageInfoAsBytes))
//   if err != nil{
// 		return shim.Error(err.Error())
// 	}
//   fmt.Println("AssetSaleAgreementUpload done")
//
// 	// 现在开始记录操作
// 	var TxInfo [8]string
//   TxInfo[0] = stub.GetTxID()                  //交易ID
// 	TxInfo[1] = args[0]                         //交易发起人
// 	TxInfo[2] = ProductID                       //产品ID
// 	TxInfo[3] = time.Now().Format("2006-01-02T15:04:05.000Z")      //交易时间
// 	TxInfo[4] = "ClaimsPackageInfo"             //链码名称
// 	TxInfo[5] = "AssetSaleAgreementUpload"     //所调函数
// 	TxInfo[6] = args[2]                         //所传参数
// 	TxInfo[7] = "发起人上传买卖协议"               //交易描述
//
// 	functionName := "add"
// 	invokeArgs := util.ToChaincodeArgs(functionName,TxInfo[0],TxInfo[1],TxInfo[2],TxInfo[3],TxInfo[4],TxInfo[5],TxInfo[6],TxInfo[7])
// 	response := stub.InvokeChaincode(TxRecorderChaincodeName, invokeArgs, TxRecorderChaincodeChannel)
// 	if response.Status != shim.OK {
// 			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
// 			fmt.Printf(errStr)
// 			return shim.Error(errStr)
// 		}
// 	//fmt.Printf("Invoke chaincode successful. Got response %s", string(response))
//   fmt.Println("success add a new TxInfo")
// 	//记录操作完成
//
// 	return shim.Success(nil)
// }
//
// // ============================================================================================================================
// // function:发起人上传资产买卖协议（url和hash值）
// // input：ProductID,UrlAndHashInfo
// // ============================================================================================================================
// func (t *SimpleChaincode) assetSaleAgreementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//   var err error
// 	if len(args) != 3 {
// 		return shim.Error("Incorrect number of arguments. Expecting 3")
// 	}
//
// 	ProductID := args[1]
// 	UrlAndHashInfo := args[2]
//
// 	ClaimsPackageInfoAsBytes, err :=  stub.GetState(ProductID)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	// timestamp, _:= stub.GetTxTimestamp()
// 	// CreatedTime := time.Unix(timestamp.Seconds, int64(timestamp.Nanos))
// 	ClaimsPackageInfoObj := ClaimsPackageInfoStruct{}
// 	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfoObj)
// 	if( ClaimsPackageInfoObj.Status != "ProInfoUpload" ){
// 		return shim.Error("Error Status!")
// 	}
//
// 	SaleAgreementObj := SaleAgreementStruct{}
// 	err = json.Unmarshal([]byte(UrlAndHashInfo),&SaleAgreementObj)
// 	if err != nil {
// 	  return shim.Error(err.Error())
// 	}
// 	ClaimsPackageInfoObj.SaleAgreement = SaleAgreementObj
//
// 	ClaimsPackageInfoObj.Status = "AssetSaleAgreementUpload"
//
//   ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfoObj)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
//   err = stub.PutState(ProductID, []byte(ClaimsPackageInfoAsBytes))
//   if err != nil{
// 		return shim.Error(err.Error())
// 	}
//   fmt.Println("AssetSaleAgreementUpload done")
//
// 	// 现在开始记录操作
// 	var TxInfo [8]string
//   TxInfo[0] = stub.GetTxID()                  //交易ID
// 	TxInfo[1] = args[0]                         //交易发起人
// 	TxInfo[2] = ProductID                       //产品ID
// 	TxInfo[3] = time.Now().Format("2006-01-02T15:04:05.000Z")      //交易时间
// 	TxInfo[4] = "ClaimsPackageInfo"             //链码名称
// 	TxInfo[5] = "AssetSaleAgreementUpload"     //所调函数
// 	TxInfo[6] = args[2]                         //所传参数
// 	TxInfo[7] = "发起人上传买卖协议"               //交易描述
//
// 	functionName := "add"
// 	invokeArgs := util.ToChaincodeArgs(functionName,TxInfo[0],TxInfo[1],TxInfo[2],TxInfo[3],TxInfo[4],TxInfo[5],TxInfo[6],TxInfo[7])
// 	response := stub.InvokeChaincode(TxRecorderChaincodeName, invokeArgs, TxRecorderChaincodeChannel)
// 	if response.Status != shim.OK {
// 			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
// 			fmt.Printf(errStr)
// 			return shim.Error(errStr)
// 		}
// 	//fmt.Printf("Invoke chaincode successful. Got response %s", string(response))
//   fmt.Println("success add a new TxInfo")
// 	//记录操作完成
//
// 	return shim.Success(nil)
// }

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}


	return shim.Success(nil)
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	 }
	 ProductID := args[0]
 	ClaimsPackageInfo, err := stub.GetState(ProductID)
 if err != nil {
 	jsonResp := "{\"Error\":\"Failed to get state for " + ProductID + "\"}"
 	return shim.Error(jsonResp)
  }
 if ClaimsPackageInfo == nil {
 	jsonResp := "{\"Error\":\"Nil content for " + ProductID + "\"}"
 	return shim.Error(jsonResp)
  }
 return shim.Success(ClaimsPackageInfo)
}

func (t *SimpleChaincode) update(stub shim.ChaincodeStubInterface, args []string) pb.Response {


        return shim.Success(nil);
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}
