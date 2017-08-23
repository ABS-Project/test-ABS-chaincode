用于记录每次改写操作的详情
数据结构
{
"TxID":"",
"TxProposer":"",
"TxTime":,
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
["cb6f47d8dd51246f79a956af28a52522f2dd6f224c38c6d948f0204602b35329","testuser",2017-08-23T02:57:43.018Z,"ClaimsPackageInfo","add","{\"ProductID\":\"123\",\"ProductName\":\"钱包汇通第一期保理ABS\",\"ProductType\":\"信托计划\",\"BasicAssets\":\"保理车贷\",\"ProjectScale\":400000000,\"Originators\":\"qbht\",\"Investor\":[\"qbjf\",\"shyh\",\"zrj\"],\"ExpectedReturn\":\"15\",\"PaymentMethod\":\"按季付\",\"TrustInstitution\":\"zrgj\",\"DifferenceComplement\":\"amdq\",\"AssetRatingAgency\":\"zhypg\",\"AccountFirm\":\"dhhs\",\"LawOffice\":\"zlls\",\"TrustManagementFee\":10,\"AssetRatingFee\":10,\"CounselFee\":100,\"AccountancyFee\":100,\"BasicCreditorInfo\":{\"Url\":\"www.qianbao/cc/12\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"},\"Remark\":\"无\"}","基础资产打包上传操作"]
