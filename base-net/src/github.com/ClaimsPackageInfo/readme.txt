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

world status：
	key		  value
	ProductName	  ClaimsPackageInfo
	注释：key这一列下面，带引号就是key的实际值，不带引号的是变量，变量是什么，key的值是什么

数据示例
"InitClaimsPackageInfo":

"{\"ProductID\":\"123\",\"ProductName\":\"钱包汇通第一期保理ABS\",\"ProductType\":\"信托计划\",\"BasicAssets\":\"保理车贷\",\"ProjectScale\":400000000,\"Originators\":\"qbht\",\"Investor\":[\"qbjf\",\"shyh\",\"zrj\"],\"ExpectedReturn\":\"15\",\"PaymentMethod\":\"按季付\",\"TrustInstitution\":\"zrgj\",\"DifferenceComplement\":\"amdq\",\"AssetRatingAgency\":\"zhypg\",\"AccountFirm\":\"dhhs\",\"LawOffice\":\"zlls\",\"TrustManagementFee\":10,\"AssetRatingFee\":10,\"CounselFee\":100,\"AccountancyFee\":100,\"BasicCreditorInfo\":{\"Url\":\"www.qianbao/cc/12\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"},\"Remark\":\"无\"}"

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

ABS产品信息查询调用queryClaimsPackageInfo():
如查询ProductID =123的产品信息：
echo "GET query chaincode ClaimsPackageInfo on peer1 of Org1"
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo?peer=peer1&fcn=queryClaimsPackageInfo&args=%5B%22123%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"

返回值：
{"InitClaimsPackageInfo":{"ProductID":"123","ProductName":"钱包汇通第一期保理ABS","ProductType":"信托计划","BasicAssets":"保理车贷","ProjectScale":4e+08,"Originators":"qbht","Investor":["qbjf","shyh","zrj"],"ExpectedReturn":"15","PaymentMethod":"按季付","TrustInstitution":"zrgj","DifferenceComplement":"amdq","AssetRatingAgency":"zhypg","AccountFirm":"dhhs","LawOffice":"zlls","TrustManagementFee":10,"AssetRatingFee":10,"CounselFee":100,"AccountancyFee":100,"BasicCreditorInfo":{"Url":"www.qianbao/cc/12","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"Remark":"无","CreatedTime":"2017-08-25T03:53:53.611Z"},"SaleAgreement":{"Url":"www.qianbao/cc/11","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"GuaranteeAgrement":{"Url":"www.qianbao/cc/12","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"ProductDesignAgreement":{"Url":"www.qianbao/cc/13","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"AssetRatingInstruction":{"Url":"www.qianbao/cc/15","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b","PriorityAssetRatio":"10","SubprimeAssetRatio":"20","InferiorAssetRatio":"30","PriorityAssetRating":"40","SubprimeAssetsRating":"50"},"AccountOpinion":{"Url":"www.qianbao/cc/16","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"LegalOpinion":{"Url":"www.qianbao/cc/17","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"ProductPlanInstruction":{"Url":"www.qianbao/cc/18","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"InferiorAssetSubscriptionAgreement":{"Url":"www.qianbao/cc/19","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"SubprimeAssetSubscriptionAgreement":{"Url":"www.qianbao/cc/20","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"PriorityAssetSubscriptionAgreement":{"Url":"www.qianbao/cc/21","Hashcode":"40b3fa8de4e01e5b37928ff03c7c6f0b"},"Status":"BreakAccountRecording"}
