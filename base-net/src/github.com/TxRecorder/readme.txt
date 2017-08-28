用于记录每次改写操作的详情
数据结构
{
"TxID":"",
"TxProposer":"",
"TxProductID"："",
"TxTime":"",
"TxChaincode":"",
"TxFunction":"",
"TxArguments":"",
"TxDescription":""
}

world status：
	key		  value
	TxID	  TxInfo
	注释：key这一列下面，带引号就是key的实际值，不带引号的是变量，变量是什么，key的值是什么

示例数据
["cb6f47d8dd51246f79a956af28a52522f2dd6f224c38c6d948f0204602b35329","testuser","2017-08-23T02:57:43.018Z","ClaimsPackageInfo","add","{\"ProductID\":\"123\",\"ProductName\":\"钱包汇通第一期保理ABS\",\"ProductType\":\"信托计划\",\"BasicAssets\":\"保理车贷\",\"ProjectScale\":400000000,\"Originators\":\"qbht\",\"Investor\":[\"qbjf\",\"shyh\",\"zrj\"],\"ExpectedReturn\":\"15\",\"PaymentMethod\":\"按季付\",\"TrustInstitution\":\"zrgj\",\"DifferenceComplement\":\"amdq\",\"AssetRatingAgency\":\"zhypg\",\"AccountFirm\":\"dhhs\",\"LawOffice\":\"zlls\",\"TrustManagementFee\":10,\"AssetRatingFee\":10,\"CounselFee\":100,\"AccountancyFee\":100,\"BasicCreditorInfo\":{\"Url\":\"www.qianbao/cc/12\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"},\"Remark\":\"无\"}","基础资产打包上传操作"]

调用add()增加一条操作记录，由其他链码调用：
接口测试：
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/TxRecorder \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"add",
	"args":["cb6f47d8dd51246f79a956af28a52522f2dd6f224c38c6d948f0204602b35329","zlls","123","2017-08-23T05:05:13.359Z","ClaimsPackageInfo","add","{\"ProductID\":\"123\",\"ProductName\":\"钱包汇通第一期保理ABS\",\"ProductType\":\"信托计划\",\"BasicAssets\":\"保理车贷\",\"ProjectScale\":400000000,\"Originators\":\"qbht\",\"Investor\":[\"qbjf\",\"shyh\",\"zrj\"],\"ExpectedReturn\":\"15\",\"PaymentMethod\":\"按季付\",\"TrustInstitution\":\"zrgj\",\"DifferenceComplement\":\"amdq\",\"AssetRatingAgency\":\"zhypg\",\"AccountFirm\":\"dhhs\",\"LawOffice\":\"zlls\",\"TrustManagementFee\":10,\"AssetRatingFee\":10,\"CounselFee\":100,\"AccountancyFee\":100,\"BasicCreditorInfo\":{\"Url\":\"www.qianbao/cc/12\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"},\"Remark\":\"无\"}","基础资产打包上传操作"]
}'
返回值为交易ID

调用query()查询操作记录，传入待查TxID的list:
echo "GET query chaincode TxRecorder on peer1 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/TxRecorder?peer=peer1&fcn=query&args=%5B%22cb6f47d8dd51246f79a956af28a52522f2dd6f224c38c6d948f0204602b35329%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
返回值为结果的字符串list：
["{\"TxID\":\"cb6f47d8dd51246f79a956af28a52522f2dd6f224c38c6d948f0204602b35329\",\"TxProposer\":\"zlls\",\"TxProductID\":\"123\",\"TxTime\":\"2017-08-23T05:05:13.359Z\",\"TxChaincode\":\"ClaimsPackageInfo\",\"TxFunction\":\"add\",\"TxArguments\":\"{\\\"ProductID\\\":\\\"123\\\",\\\"ProductName\\\":\\\"钱包汇通第一期保理ABS\\\",\\\"ProductType\\\":\\\"信托计划\\\",\\\"BasicAssets\\\":\\\"保理车贷\\\",\\\"ProjectScale\\\":400000000,\\\"Originators\\\":\\\"qbht\\\",\\\"Investor\\\":[\\\"qbjf\\\",\\\"shyh\\\",\\\"zrj\\\"],\\\"ExpectedReturn\\\":\\\"15\\\",\\\"PaymentMethod\\\":\\\"按季付\\\",\\\"TrustInstitution\\\":\\\"zrgj\\\",\\\"DifferenceComplement\\\":\\\"amdq\\\",\\\"AssetRatingAgency\\\":\\\"zhypg\\\",\\\"AccountFirm\\\":\\\"dhhs\\\",\\\"LawOffice\\\":\\\"zlls\\\",\\\"TrustManagementFee\\\":10,\\\"AssetRatingFee\\\":10,\\\"CounselFee\\\":100,\\\"AccountancyFee\\\":100,\\\"BasicCreditorInfo\\\":{\\\"Url\\\":\\\"www.qianbao/cc/12\\\",\\\"Hashcode\\\":\\\"40b3fa8de4e01e5b37928ff03c7c6f0b\\\"},\\\"Remark\\\":\\\"无\\\"}\",\"TxDescription\":\"基础资产打包上传操作\"}"]


调用queryAllTxRecord()查询所有操作记录
echo "GET query chaincode AllTxRecorder on peer1 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/TxRecorder?peer=peer1&fcn=queryAllTxRecord&args=%5B%22%22%2C%22%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
