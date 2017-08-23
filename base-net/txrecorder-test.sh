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
TxRecorder
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

echo "POST Install chaincode on Org1"
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


echo "POST Install chaincode on Org2"
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

echo "POST Install chaincode on Org3"
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

echo "POST Install chaincode on Org4"
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
echo "POST instantiate chaincode on peer1 of Org1"
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

echo "POST invoke chaincode on peers of Org2"
echo "POST add TxRecorder"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/TxRecorder \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"add",
	"args":["{\"ProductID\":\"123\",\"ProductName\":\"钱包汇通第一期保理ABS\",\"ProductType\":\"信托计划\",\"BasicAssets\":\"保理车贷\",\"ProjectScale\":400000000,\"Originators\":\"qbht\",\"Investor\":[\"qbjf\",\"shyh\",\"zrj\"],\"ExpectedReturn\":\"15\",\"PaymentMethod\":\"按季付\",\"TrustInstitution\":\"zrgj\",\"DifferenceComplement\":\"amdq\",\"AssetRatingAgency\":\"zhypg\",\"AccountFirm\":\"dhhs\",\"LawOffice\":\"zlls\",\"TrustManagementFee\":10,\"AssetRatingFee\":10,\"CounselFee\":100,\"AccountancyFee\":100,\"BasicCreditorInfo\":{\"Url\":\"www.qianbao/cc/12\",\"Hashcode\":\"40b3fa8de4e01e5b37928ff03c7c6f0b\"},\"Remark\":\"无\"}"]
}')
echo "Transacton ID is $TRX_ID"
echo

echo "GET query chaincode on peer1 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/TxRecorder?peer=peer1&fcn=query&args=%5B%22123%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
echo
echo "Total execution time : $(($(date +%s)-starttime)) secs ..."
