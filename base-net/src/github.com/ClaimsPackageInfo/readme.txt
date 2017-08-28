//==========================================================================================================
//流程状态表示,通过流程状态控制业务逻辑。调用合约接口时必须按照流程状态顺序调用。
//{"ProInfoUpload","AssetSaleAgreementUpload","GuaranteeAgreementUpload","TrustManageementUpload",
//"AssetRatingInstructionUpload","AccountOpinionUpload","CounselOpinionUpload","ProductPlanInstructionUpload",
//"InferiorAssetObtain","InferiorAssetObtainRecording","SubprimeAssetObtain","SubprimeAssetsObtainRecording",
//"PriorityAssetObtain","PriorityAssetObtainRecording","BreakAccountRecording"}
//==========================================================================================================

ClaimsPackageInfo数据结构定义，主要是响应第一步资产上传操作

数据结构：
	{
	  "InitClaimsPackageInfo":{
		    "ProductID":"",
				"ProductName":"",
				"ProductType":"",
				"BasicAssets":"",
				"ProjectScale":"",
				"Originators":"",
				"Investor":["",""],
			  "ExpectedReturn":"",
			  "PaymentMethod":"",
			  "TrustInstitution":"",
			  "DifferenceComplement":"",
			  "AssetRatingAgency":"",
			  "AccountFirm":"",
			  "LawOffice":"",
			  "TrustManagementFee":,
			  "AssetRatingFee":,
			  "CounselFee":,
			  "AccountancyFee":,
			  "BasicCreditorInfo":{
							"Url":"",
							"Hashcode":""
				 },
			  "Remark":""
				"CreatedTime":""
		},
	  "SaleAgreement":{
				"Url":"",
				"Hashcode":""
		},
	  "GuaranteeAgrement":{
				"Url":"",
				"Hashcode":""
		},
	  "ProductDesignAgreement":{
				"Url":"",
				"Hashcode":""
		},
	  "AssetRatingInstruction":{
				"Url":"",
				"Hashcode":"",
				""PriorityAssetRatio":"",
				"SubprimeAssetRatio":"",
				"InferiorAssetRatio":"",
				"PriorityAssetRating":"",
				"SubprimeAssetsRating":""
		},
	  "AccountOpinion":{
				"Url":"",
				"Hashcode":""
		},
	  "LegalOpinion":{
				"Url":"",
				"Hashcode":""
		},
	  "ProductPlanInstruction":{
				"Url":"",
				"Hashcode":""
		},
	  "InferiorAssetSubscriptionAgreement":{
				"Url":"",
				"Hashcode":""
		},
	  "SubprimeAssetSubscriptionAgreement":{
				"Url":"",
				"Hashcode":""
		},
	  "PriorityAssetSubscriptionAgreement":{
				"Url":"",
				"Hashcode":""
		},
		"Status":""
	}

TransferRecord  struct: 划帐信息，由代币节点进行记录。最后的分帐等同于代币节点进行转账记录。
type TransferRecordStruct struct {
	ProductID           string  `json:"ProductID"`
	WaterFlowNumber     string  `json:"WaterFlowNumber"`
	WaterFlowNumberTime string  `json:"WaterFlowNumberTime"`
	FromAccount         string  `json:"FromAccount"`
	ToAccount           string  `json:"ToAccount"`
	BbMount             float64 `json:"BbMount"`
}

world status：
	key		  value
	ProductName	  ClaimsPackageInfo
	注释：key这一列下面，带引号就是key的实际值，不带引号的是变量，变量是什么，key的值是什么

数据示例
"InitClaimsPackageInfo":

"{\"ProductID\":\"123\",\"ProductName\":\"钱包汇通第一期保理ABS\",\"ProductType\":\"信托计划\",\"BasicAssets\":\"保理车贷\",\"ProjectScale\":400000000,\"Originators\":\"qbht\",\"Investor\":[\"qbjf\",\"shyh\",\"zrj\"],\"ExpectedReturn\":\"15\",\"PaymentMethod\":\"按季付\",\"TrustInstitution\":\"zrgj\",\"DifferenceComplement\":\"amdq\",\"AssetRatingAgency\":\"zhypg\",\"AccountFirm\":\"dhhs\",\"LawOffice\":\"zlls\",\"TrustManagementFee\":10,\"AssetRatingFee\":10,\"CounselFee\":100,\"AccountancyFee\":100,\"BasicCreditorInfo\":{\"Url\":\"www.qianbao/cc/12\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"},\"Remark\":\"无\"}"


############################业务逻辑部分###########################
资产打包时产品初始信息上传调用proInfoUpload():
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"proInfoUpload",
	"args":["zlls","{\"ProductID\":\"123\",\"ProductName\":\"钱包汇通第一期保理ABS\",\"ProductType\":\"信托计划\",\"BasicAssets\":\"保理车贷\",\"ProjectScale\":400000000,\"Originators\":\"qbht\",\"Investor\":[\"qbjf\",\"shyh\",\"zrj\"],\"ExpectedReturn\":\"15\",\"PaymentMethod\":\"按季付\",\"TrustInstitution\":\"zrgj\",\"DifferenceComplement\":\"amdq\",\"AssetRatingAgency\":\"zhypg\",\"AccountFirm\":\"dhhs\",\"LawOffice\":\"zlls\",\"TrustManagementFee\":10,\"AssetRatingFee\":10,\"CounselFee\":100,\"AccountancyFee\":100,\"BasicCreditorInfo\":{\"Url\":\"www.qianbao/cc/12\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"},\"Remark\":\"无\"}"]
}'
返回值为transaction id

.......同样的 以下业务流程类同。 代码输入详情见ClaimsPackageInfo.go 测试详情见tx-cp-test.sh
assetSaleAgreementUpload（）            发起人上传资产买卖协议
guaranteeAgreementUpload（）            差额补足人上传差额补足协议
trustManageementUpload（）              SPV上传产品设计书
assetRatingInstructionUpload（）        资产评级机构上传资产评级
accountOpinionUpload（）                会计师事务所上传审计报告
counselOpinionUpload（）                律师事务所上传法律意见书
productPlanInstructionUpload（）        SPV上传产品计划说明书
inferiorAssetObtain（）                 劣后级资产购买方认购劣后级资产
inferiorAssetObtainRecording（）        代币节点记录劣后级资产购买的转账情况
subprimeAssetObtain（）                 次优级资产购买方上传次优级资产认购协议
subprimeAssetsObtainRecording（）       代币节点记录次优级资产认购方转账记录
priorityAssetObtain（）                 优先级资产购买方上传优先级资产认购协议
priorityAssetObtainRecording（）        代币节点记录优先级资产认购方的转账记录
breakAccountRecording（）               代币节点进行分帐，没调用一次记录一条分帐记录
finishBreakAccountRecording（）         完成分帐
##########################业务逻辑部分结束#############################



查询产生的分帐记录
"{\"ProductID\":\"123\",\"WaterFlowNumber\":\"33333333333\",\"WaterFlowNumberTime\":\"2017-10\",\"FromAccount\":\"ccc\",\"ToAccount\":\"cccccc\",\"BbMount\":3000.00}"

echo "GET query chaincode TransferRecord on peer1 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo?peer=peer1&fcn=queryClaimsPackageInfo&args=%5B%22RecordID03%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
echo

ABS产品信息查询调用queryClaimsPackageInfo():
如查询ProductID =123的产品信息：
echo "GET query chaincode ClaimsPackageInfo on peer1 of Org1"
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo?peer=peer1&fcn=queryClaimsPackageInfo&args=%5B%22123%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"

返回值：
{"InitClaimsPackageInfo":{"ProductID":"123","ProductName":"钱包汇通第一期保理ABS","ProductType":"信托计划","BasicAssets":"保理车贷","ProjectScale":4e+08,"Originators":"qbht","Investor":["qbjf","shyh","zrj"],"ExpectedReturn":"15","PaymentMethod":"按季付","TrustInstitution":"zrgj","DifferenceComplement":"amdq","AssetRatingAgency":"zhypg","AccountFirm":"dhhs","LawOffice":"zlls","TrustManagementFee":10,"AssetRatingFee":10,"CounselFee":100,"AccountancyFee":100,"BasicCreditorInfo":{"Url":"www.qianbao/cc/12","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"Remark":"无","CreatedTime":"2017-08-25T03:53:53.611Z"},"SaleAgreement":{"Url":"www.qianbao/cc/11","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"GuaranteeAgrement":{"Url":"www.qianbao/cc/12","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"ProductDesignAgreement":{"Url":"www.qianbao/cc/13","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"AssetRatingInstruction":{"Url":"www.qianbao/cc/15","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b","PriorityAssetRatio":"10","SubprimeAssetRatio":"20","InferiorAssetRatio":"30","PriorityAssetRating":"40","SubprimeAssetsRating":"50"},"AccountOpinion":{"Url":"www.qianbao/cc/16","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"LegalOpinion":{"Url":"www.qianbao/cc/17","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"ProductPlanInstruction":{"Url":"www.qianbao/cc/18","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"InferiorAssetSubscriptionAgreement":{"Url":"www.qianbao/cc/19","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"SubprimeAssetSubscriptionAgreement":{"Url":"www.qianbao/cc/20","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"PriorityAssetSubscriptionAgreement":{"Url":"www.qianbao/cc/21","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"Status":"BreakAccountRecording"}
