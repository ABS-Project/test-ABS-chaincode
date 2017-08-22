// 债券基础信息相关的链码操作

/*

*/

package main


import (
	"fmt"
	"encoding/json"
  "time"

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

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response  {
	logger.Info("########### ClaimsPackageInfo Init ###########")
	return shim.Success(nil)
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### example_cc0 Invoke ###########")

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
	if function == "update" {
		// Deletes an entity from its state
		return t.update(stub, args)
	}

	logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0]))
}

func (t *SimpleChaincode) add(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 2. ")
	}
	var InitClaimsPackageInfoObj InitClaimsPackageInfoStruct
	InitClaimsPackageInfo :=args[0]
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
	InitClaimsPackageInfoObj.CreatedTime = time.Unix(timestamp.Seconds, int64(timestamp.Nanos))
	var ClaimsPackageInfoObj ClaimsPackageInfoStruct
	ClaimsPackageInfoObj.Status = "ProInfoUpload"
	ClaimsPackageInfoObj.InitClaimsPackageInfo = InitClaimsPackageInfoObj
	jsonAsBytes,_:= json.Marshal(ClaimsPackageInfoObj)
	err = stub.PutState(ProductID,[]byte(jsonAsBytes))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil);
}

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
