#!/bin/bash
#
# Copyright Tongji Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

jq --version > /dev/null 2>&1
if [ $? -ne 0 ]; then
	echo "Please Install 'jq' https://stedolan.github.io/jq/ to execute this script"
	echo
	exit 1
fi
starttime=$(date +%s)

echo "POST request Enroll on Org1  ..."
echo
ORG1_TOKEN=$(curl -s -X POST \
  http://localhost:4000/users \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'username=Jim&orgName=org1')
echo $ORG1_TOKEN
ORG1_TOKEN=$(echo $ORG1_TOKEN | jq ".token" | sed "s/\"//g")
echo
echo "ORG1 token is $ORG1_TOKEN"
echo

echo
ORG2_TOKEN=$(curl -s -X POST \
  http://localhost:4000/users \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'username=Barry&orgName=org2')
echo $ORG2_TOKEN
ORG2_TOKEN=$(echo $ORG2_TOKEN | jq ".token" | sed "s/\"//g")
echo
echo "ORG2 token is $ORG2_TOKEN"
echo

ORG3_TOKEN=$(curl -s -X POST \
  http://localhost:4000/users \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'username=Irisxu&orgName=org3')
echo $ORG3_TOKEN
ORG3_TOKEN=$(echo $ORG3_TOKEN | jq ".token" | sed "s/\"//g")
echo
echo "ORG3 token is $ORG3_TOKEN"
echo

ORG4_TOKEN=$(curl -s -X POST \
  http://localhost:4000/users \
  -H "content-type: application/x-www-form-urlencoded" \
  -d 'username=xiaowang&orgName=org4')
echo $ORG4_TOKEN
ORG4_TOKEN=$(echo $ORG4_TOKEN | jq ".token" | sed "s/\"//g")
echo
echo "ORG4 token is $ORG4_TOKEN"
echo

echo
echo "POST request Create channel  ..."
echo
curl -s -X POST \
  http://localhost:4000/channels \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"channelName":"mychannel",
	"channelConfigPath":"../artifacts/channel/mychannel.tx"
}'
echo
echo
sleep 5
echo "POST request Join channel on Org1"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:7051","localhost:7056"]
}'
echo
echo

echo "POST request Join channel on Org2"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051","localhost:8056"]
}'
echo
echo

echo "POST request Join channel on Org3"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \
  -H "authorization: Bearer $ORG3_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:9051","localhost:9056"]
}'
echo
echo

echo "POST request Join channel on Org4"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \
  -H "authorization: Bearer $ORG4_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:10051","localhost:10056"]
}'
echo
echo

echo "POST Install chaincode BusinessPartnerInfo on Org1"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:7051","localhost:7056"],
	"chaincodeName":"BusinessPartnerInfo",
	"chaincodePath":"github.com/BusinessPartnerInfo",
	"chaincodeVersion":"v0"
}'
echo
echo


echo "POST Install chaincode BusinessPartnerInfo on Org2"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051","localhost:8056"],
	"chaincodeName":"BusinessPartnerInfo",
	"chaincodePath":"github.com/BusinessPartnerInfo",
	"chaincodeVersion":"v0"
}'
echo
echo

echo "POST Install chaincode BusinessPartnerInfo on Org3"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG3_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:9051","localhost:9056"],
	"chaincodeName":"BusinessPartnerInfo",
	"chaincodePath":"github.com/BusinessPartnerInfo",
	"chaincodeVersion":"v0"
}'
echo
echo

echo "POST Install chaincode BusinessPartnerInfo on Org4"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG4_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:10051","localhost:10056"],
	"chaincodeName":"BusinessPartnerInfo",
	"chaincodePath":"github.com/BusinessPartnerInfo",
	"chaincodeVersion":"v0"
}'
echo
echo "POST instantiate chaincode BusinessPartnerInfo on peer1 of Org1"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"chaincodeName":"BusinessPartnerInfo",
	"chaincodeVersion":"v0",
	"functionName":"init",
	"args":[]
}'
echo
echo

echo "POST Install chaincode TxRecorder on Org1"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:7051","localhost:7056"],
	"chaincodeName":"TxRecorder",
	"chaincodePath":"github.com/TxRecorder",
	"chaincodeVersion":"v0"
}'
echo
echo


echo "POST Install chaincode TxRecorder on Org2"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051","localhost:8056"],
	"chaincodeName":"TxRecorder",
	"chaincodePath":"github.com/TxRecorder",
	"chaincodeVersion":"v0"
}'
echo
echo

echo "POST Install chaincode TxRecorder on Org3"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG3_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:9051","localhost:9056"],
	"chaincodeName":"TxRecorder",
	"chaincodePath":"github.com/TxRecorder",
	"chaincodeVersion":"v0"
}'
echo
echo

echo "POST Install chaincode TxRecorder on Org4"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG4_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:10051","localhost:10056"],
	"chaincodeName":"TxRecorder",
	"chaincodePath":"github.com/TxRecorder",
	"chaincodeVersion":"v0"
}'
echo
echo "POST instantiate chaincode TxRecorder on peer1 of Org1"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"chaincodeName":"TxRecorder",
	"chaincodeVersion":"v0",
	"functionName":"init",
	"args":[]
}'
echo
echo

echo "POST Install chaincode ClaimsPackageInfo on Org1"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:7051","localhost:7056"],
	"chaincodeName":"ClaimsPackageInfo",
	"chaincodePath":"github.com/ClaimsPackageInfo",
	"chaincodeVersion":"v0"
}'
echo
echo


echo "POST Install chaincode ClaimsPackageInfo on Org2"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051","localhost:8056"],
	"chaincodeName":"ClaimsPackageInfo",
	"chaincodePath":"github.com/ClaimsPackageInfo",
	"chaincodeVersion":"v0"
}'
echo

echo "POST Install chaincode ClaimsPackageInfo on Org3"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG3_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:9051","localhost:9056"],
	"chaincodeName":"ClaimsPackageInfo",
	"chaincodePath":"github.com/ClaimsPackageInfo",
	"chaincodeVersion":"v0"
}'
echo
echo

echo "POST Install chaincode ClaimsPackageInfo on Org4"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer $ORG4_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:10051","localhost:10056"],
	"chaincodeName":"ClaimsPackageInfo",
	"chaincodePath":"github.com/ClaimsPackageInfo",
	"chaincodeVersion":"v0"
}'
echo
echo "POST instantiate chaincode ClaimsPackageInfo on peer1 of Org1"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"chaincodeName":"ClaimsPackageInfo",
	"chaincodeVersion":"v0",
	"functionName":"init",
	"args":[]
}'
echo
###############################################################################
##########    以上为各个链码的安装与初始化过程，下面为链码测试过程   ###################
###############################################################################
echo
echo "POST invoke chaincode BusinessPartnerInfo on peers of Org2"
echo "POST add BusinessPartnerInfo"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/BusinessPartnerInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"addBusinessPartnerInfo",
	"args":["{\"UserName\":\"zlls\",\"Organization\":\"律师事务所\",\"Company\":\"北京市中伦律师事务所\",\"Account\":\"zlls124\"}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "POST invoke chaincode BusinessPartnerInfo on peers of Org2"
echo "POST update BusinessPartnerInfo"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/BusinessPartnerInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"updateBusinessPartnerInfo",
	"args":["{\"UserName\":\"zlls\",\"Organization\":\"律师事务所\",\"Company\":\"北京市中伦律师事务所\",\"Account\":\"zlls666\"}"]
}')
echo "Transacton ID is $TRX_ID"

echo "GET query chaincode BusinessPartnerInfo on peer1 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/BusinessPartnerInfo?peer=peer1&fcn=queryBusinessPartnerInfo&args=%5B%22zlls%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
echo


echo "POST invoke chaincode TxRecorder on peers of Org2"
echo "POST add TxRecorder"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/TxRecorder \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"add",
	"args":["cb6f47d8dd51246f79a956af28a52522f2dd6f224c38c6d948f0204602b35329","zlls","123","2017-08-23T05:05:13.359Z","ClaimsPackageInfo","add","{\"ProductID\":\"123\",\"ProductName\":\"钱包汇通第一期保理ABS\",\"ProductType\":\"信托计划\",\"BasicAssets\":\"保理车贷\",\"ProjectScale\":400000000,\"Originators\":\"qbht\",\"Investor\":[\"qbjf\",\"shyh\",\"zrj\"],\"ExpectedReturn\":\"15\",\"PaymentMethod\":\"按季付\",\"TrustInstitution\":\"zrgj\",\"DifferenceComplement\":\"amdq\",\"AssetRatingAgency\":\"zhypg\",\"AccountFirm\":\"dhhs\",\"LawOffice\":\"zlls\",\"TrustManagementFee\":10,\"AssetRatingFee\":10,\"CounselFee\":100,\"AccountancyFee\":100,\"BasicCreditorInfo\":{\"Url\":\"www.qianbao/cc/12\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"},\"Remark\":\"无\"}","基础资产打包上传操作"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "GET query chaincode TxRecorder on peer1 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/TxRecorder?peer=peer1&fcn=query&args=%5B%22cb6f47d8dd51246f79a956af28a52522f2dd6f224c38c6d948f0204602b35329%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
echo


echo "POST invoke chaincode ClaimsPackageInfo on peers of Org2"
echo "POST add ClaimsPackageInfo"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"proInfoUpload",
	"args":["zlls","{\"ProductID\":\"123\",\"ProductName\":\"钱包汇通第一期保理ABS\",\"ProductType\":\"信托计划\",\"BasicAssets\":\"保理车贷\",\"ProjectScale\":400000000,\"Originators\":\"qbht\",\"Investor\":[\"qbjf\",\"shyh\",\"zrj\"],\"ExpectedReturn\":\"15\",\"PaymentMethod\":\"按季付\",\"TrustInstitution\":\"zrgj\",\"DifferenceComplement\":\"amdq\",\"AssetRatingAgency\":\"zhypg\",\"AccountFirm\":\"dhhs\",\"LawOffice\":\"zlls\",\"TrustManagementFee\":10,\"AssetRatingFee\":10,\"CounselFee\":100,\"AccountancyFee\":100,\"BasicCreditorInfo\":{\"Url\":\"www.qianbao/cc/12\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"},\"Remark\":\"无\"}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "POST invoke chaincode  on peers of Org2"
echo "POST assetSaleAgreementUpload"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"assetSaleAgreementUpload",
	"args":["zlls","123","{\"Url\":\"www.qianbao/cc/11\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "GET query chaincode AllTxRecorder on peer1 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/TxRecorder?peer=peer1&fcn=queryAllTxRecord&args=%5B%22%22%2C%22%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
echo

echo "POST invoke chaincode  on peers of Org2"
echo "POST guaranteeAgreementUpload"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"guaranteeAgreementUpload",
	"args":["zlls","123","{\"Url\":\"www.qianbao/cc/12\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "POST invoke chaincode  on peers of Org2"
echo "POST trustManageementUpload"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"trustManageementUpload",
	"args":["zlls","123","{\"Url\":\"www.qianbao/cc/13\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "POST invoke chaincode  on peers of Org2"
echo "POST assetRatingInstructionUpload"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"assetRatingInstructionUpload",
	"args":["zlls","123","{\"Url\":\"www.qianbao/cc/15\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\",\"PriorityAssetRatio\":\"10\",\"SubprimeAssetRatio\":\"20\",\"InferiorAssetRatio\":\"30\",\"PriorityAssetRating\":\"40\",\"SubprimeAssetsRating\":\"50\"}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "POST invoke chaincode  on peers of Org2"
echo "POST accountOpinionUpload"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"accountOpinionUpload",
	"args":["zlls","123","{\"Url\":\"www.qianbao/cc/16\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "POST invoke chaincode  on peers of Org2"
echo "POST counselOpinionUpload"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"counselOpinionUpload",
	"args":["zlls","123","{\"Url\":\"www.qianbao/cc/17\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "POST invoke chaincode  on peers of Org2"
echo "POST productPlanInstructionUpload"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"productPlanInstructionUpload",
	"args":["zlls","123","{\"Url\":\"www.qianbao/cc/18\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "POST invoke chaincode  on peers of Org2"
echo "POST inferiorAssetObtain"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"inferiorAssetObtain",
	"args":["zlls","123","{\"Url\":\"www.qianbao/cc/19\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "POST invoke chaincode  on peers of Org2"
echo "POST inferiorAssetObtainRecording"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"inferiorAssetObtainRecording",
	"args":["zlls","123","RecordID01","{\"ProductID\":\"123\",\"WaterFlowNumber\":\"11111111\",\"WaterFlowNumberTime\":\"2017-08\",\"FromAccount\":\"aaa\",\"ToAccount\":\"aaaaaaa\",\"BbMount\":1000.00}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "POST invoke chaincode  on peers of Org2"
echo "POST subprimeAssetObtain"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"subprimeAssetObtain",
	"args":["zlls","123","{\"Url\":\"www.qianbao/cc/20\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "POST invoke chaincode  on peers of Org2"
echo "POST subprimeAssetsObtainRecording"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"subprimeAssetsObtainRecording",
	"args":["zlls","123","RecordID02","{\"ProductID\":\"123\",\"WaterFlowNumber\":\"22222222\",\"WaterFlowNumberTime\":\"2017-9\",\"FromAccount\":\"bbb\",\"ToAccount\":\"bbbbbb\",\"BbMount\":2000.00}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "POST invoke chaincode  on peers of Org2"
echo "POST priorityAssetObtain"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"priorityAssetObtain",
	"args":["zlls","123","{\"Url\":\"www.qianbao/cc/21\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "POST invoke chaincode  on peers of Org2"
echo "POST priorityAssetObtainRecording"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"priorityAssetObtainRecording",
	"args":["zlls","123","RecordID03","{\"ProductID\":\"123\",\"WaterFlowNumber\":\"33333333333\",\"WaterFlowNumberTime\":\"2017-10\",\"FromAccount\":\"ccc\",\"ToAccount\":\"cccccc\",\"BbMount\":3000.00}"]
}')
echo "Transacton ID is $TRX_ID"
echo

#post请求可以排序，但是查询结果没有返回客户端
# echo "POST invoke chaincode  on peers of Org2"
# echo "POST queryTransferRecord"
# echo
# TRX_ID=$(curl -s -X POST \
#   http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
#   -H "authorization: Bearer $ORG2_TOKEN" \
#   -H "content-type: application/json" \
#   -d '{
# 	"peers": ["localhost:8051"],
# 	"fcn":"queryTransferRecord",
# 	"args":["RecordID03"]
# }')
# echo "Transacton ID is $TRX_ID"
# echo

echo "GET query chaincode TransferRecord on peer1 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo?peer=peer1&fcn=queryClaimsPackageInfo&args=%5B%22RecordID03%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
echo

echo "POST invoke chaincode on peers of Org2"
echo "POST breakAccountRecording"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"breakAccountRecording",
	"args":["zlls","123","RecordID04","{\"ProductID\":\"123\",\"WaterFlowNumber\":\"444444444\",\"WaterFlowNumberTime\":\"2017-11\",\"FromAccount\":\"ddd\",\"ToAccount\":\"ddddddd\",\"BbMount\":4000.00}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "GET query chaincode TransferRecord on peer1 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo?peer=peer1&fcn=queryClaimsPackageInfo&args=%5B%22RecordID04%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
echo

echo "POST invoke chaincode  on peers of Org2"
echo "POST finishBreakAccountRecording"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"finishBreakAccountRecording",
	"args":["zlls","123"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "GET query chaincode ClaimsPackageInfo on peer1 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/ClaimsPackageInfo?peer=peer1&fcn=queryClaimsPackageInfo&args=%5B%22123%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
echo

echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
