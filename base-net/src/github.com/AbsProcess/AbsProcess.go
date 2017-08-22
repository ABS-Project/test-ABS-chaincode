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
	} else if function == "add" { //add a new job
		return t.Add(stub, args)
	} else if function == "delete" { //deletes an job from its state
		return t.Delete(stub, args)
	} else if function == "edit" { //change the infor of the job
		return t.Edit(stub, args)
	} else if function == "addTX" { //add a new TX
		return t.AddTX(stub, args)
	} else if function == "addTotalApplied" { //add 1 when a student applied the job
		return t.AddTotalApplied(stub, args)
	} else if function == "addTotalWaitCheck" { //add 1 when auto check not passed
		return t.AddTotalWaitCheck(stub, args)
	} else if function == "addTotalHired" { //add 1 when auto check passed or agency check passed
		return t.AddTotalHired(stub, args)
	} else if function == "addTotalSettled" { //add 1 when auto settle passed or agency settle passed
		return t.AddTotalSettled(stub, args)
	}

	return shim.Error("Received unknown function invocation")
}

// ============================================================================================================================
// function:发起人上传资产买卖协议（url和hash值）
// input：ProductName,ulr,hash
// ============================================================================================================================
func (t *SimpleChaincode) assetSaleAgreementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName = arg[0]
	Url = arg[1]
	Hashcode = arg[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "ProInfoUpload" {
		return shim.Error("status is worry! Expect: ProInfoUpload")
	}

  _, err = agreementUpload(ProductName, SaleAgreement, Url, Hashcode)
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
func (t *SimpleChaincode) guaranteeAgreementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName = arg[0]
	Url = arg[1]
	Hashcode = arg[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "AssetSaleAgreementUpload" {
		return shim.Error("status is worry! Expect: AssetSaleAgreementUpload")
	}

  _, err = agreementUpload(ProductName, GuaranteeAgrement, Url, Hashcode)
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
func (t *SimpleChaincode) trustManageementUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName = arg[0]
	Url = arg[1]
	Hashcode = arg[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "GuaranteeAgreementUpload" {
		return shim.Error("status is worry! Expect: GuaranteeAgreementUpload")
	}

  _, err = agreementUpload(ProductName, ProductDesignAgreement, Url, Hashcode)
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
func (t *SimpleChaincode) assetRatingInstructionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}
	ProductName = arg[0]
	Url = arg[1]
	Hashcode = arg[2]
	PriorityAssetRatio = arg[4]
	SubprimeAssetRatio = arg[5]
	InferiorAssetRatio = arg[6]
	PriorityAssetRating = arg[7]
	SubprimeAssetsRating = arg[8]

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
func (t *SimpleChaincode) accountOpinionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName = arg[0]
	Url = arg[1]
	Hashcode = arg[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "AssetRatingInstructionUpload" {
		return shim.Error("status is worry! Expect: AssetRatingInstructionUpload")
	}

  _, err = agreementUpload(ProductName, AccountOpinion, Url, Hashcode)
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
func (t *SimpleChaincode) counselOpinionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName = arg[0]
	Url = arg[1]
	Hashcode = arg[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "AccountOpinionUpload" {
		return shim.Error("status is worry! Expect: AccountOpinionUpload")
	}

  _, err = agreementUpload(ProductName, LegalOpinion, Url, Hashcode)
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
func (t *SimpleChaincode) productPlanInstructionUpload(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName = arg[0]
	Url = arg[1]
	Hashcode = arg[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "CounselOpinionUpload" {
		return shim.Error("status is worry! Expect: CounselOpinionUpload")
	}

  _, err = agreementUpload(ProductName, ProductPlanInstruction, Url, Hashcode)
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
func (t *SimpleChaincode) inferiorAssetObtain(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName = arg[0]
	Url = arg[1]
	Hashcode = arg[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "ProductPlanInstructionUpload" {
		return shim.Error("status is worry! Expect: ProductPlanInstructionUpload")
	}

  _, err = agreementUpload(ProductName, InferiorAssetSubscriptionAgreement, Url, Hashcode)
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
// function:代币节点记录转账
// input：ProductName,
// ============================================================================================================================
func (t *SimpleChaincode) InferiorAssetObtainRecording(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}


	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "ProductPlanInstructionUpload" {
		return shim.Error("status is worry! Expect: ProductPlanInstructionUpload")
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
func (t *SimpleChaincode) subprimeAssetObtain(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName = arg[0]
	Url = arg[1]
	Hashcode = arg[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "InferiorAssetObtainRecording" {
		return shim.Error("status is worry! Expect: InferiorAssetObtainRecording")
	}

  _, err = agreementUpload(ProductName, SubprimeAssetSubscriptionAgreement, Url, Hashcode)
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
// function:代币节点记录转账
// input：
// ============================================================================================================================
func (t *SimpleChaincode) SubprimeAssetsObtainRecording(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	currentStatus, err := checkStatus(ProductName)

	changeStatus(ProductName)
}

// ============================================================================================================================
// function:优先级资产购买方（url和hash值）
// input：ProductName,ulr,hash
// ============================================================================================================================
func (t *SimpleChaincode) PriorityAssetObtain(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	ProductName = arg[0]
	Url = arg[1]
	Hashcode = arg[2]

	currentStatus, err := checkStatus(ProductName)
	if currentStatus != "SubprimeAssetsObtainRecording" {
		return shim.Error("status is worry! Expect: SubprimeAssetsObtainRecording")
	}

  _, err = agreementUpload(ProductName, PriorityAssetSubscriptionAgreement, Url, Hashcode)
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
// function:代币节点记录转账
// input：ProductName
// ============================================================================================================================
func (t *SimpleChaincode) PriorityAssetObtainRecording(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	currentStatus, err := checkStatus(ProductName)

	changeStatus(ProductName)
}

// ============================================================================================================================
// function:代币节点记录转账
// input：ProductName
// ============================================================================================================================
func (t *SimpleChaincode) BreakAccountRecording(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	currentStatus, err := checkStatus(ProductName)

	changeStatus(ProductName)
}


// ============================================================================================================================
// function:检查当前的业务状态
// input：ProductName
// ============================================================================================================================
func checkStatus(stub shim.ChaincodeStubInterface , ProductName string) (string, err){
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
func changeStatus(stub shim.ChaincodeStubInterface , ProductName string) (string, err) {
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

func agreementUpload(stub shim.ChaincodeStubInterface , ProductName string, AgreementName string, Url string, Hashcode string) (string, err) {
	ClaimsPackageInfoAsBytes, err :=  stub.Getstatus(ProductName)
	if err != nil {
		return nil, err
	}
	ClaimsPackageInfo := ClaimsPackageInfoStruct{}

	json.Unmarshal(ClaimsPackageInfoAsBytes, &ClaimsPackageInfo)
  ClaimsPackageInfo.AgreementName.Url =  Url
	ClaimsPackageInfo.AgreementName.Hashcode =  Hashcode

  ClaimsPackageInfoAsBytes, err = json.Marshal(ClaimsPackageInfo)
	if err != nil {
		return nil, err
	}
  err = stub.PutState(ProductName, ClaimsPackageInfoAsBytes)
  if err != nil{
		return nil, err
	}

	return "upload agreement successfully", nil
}
