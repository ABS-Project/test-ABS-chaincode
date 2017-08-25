BusinessPartnerInfo智能合约：每一个合作伙伴用户都在这一个合约中创建，修改，删除...

数据结构：
	{
  "UserName":"abc",
	"Organization":"律师事务所",
	"Company":"A律所",
	"Account":"23456789",
	"CreatedTime":"Mon Aug 21 2017 15:45:41 GMT+0800 (CST)",
	"OperateLog":["1","2","3","4"]
	}

world status：
	key		  value
	UserName	  BusinessPartnerInfo
	注释：key这一列下面，带引号就是key的实际值，不带引号的是变量，变量是什么，key的值是什么

"{\"UserName\":\"dhhs\",\"Organization\":\"会计事务所\",\"Company\":\"大华会计师事务所\",\"Account\":\"dhhs125\",\"CreatedTime\":\"2017-08-23T05:05:13.359Z\",\"OperateLog\":[\"1\",\"2\",\"3\",\"4\"]}"

调用addBusinessPartnerInfo()增加商业合作伙伴：
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/BusinessPartnerInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"add",
	"args":["{\"UserName\":\"zlls\",\"Organization\":\"律师事务所\",\"Company\":\"北京市中伦律师事务所\",\"Account\":\"zlls124\"}"]
}'
返回值：b75f752f5b84b6c511e471f739843151f36a9e773ce5feff0ad41acd06df825e

调用updateBusinessPartnerInfo()更新商业合作伙伴信息：
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/BusinessPartnerInfo \
  -H "authorization: Bearer $ORG2_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["localhost:8051"],
	"fcn":"updateBusinessPartnerInfo",
	"args":["{\"UserName\":\"zlls\",\"Organization\":\"律师事务所\",\"Company\":\"北京市中伦律师事务所\",\"Account\":\"zlls666\"}"]
}'
返回值：c99ceff61f543d063f4783ac5511156b744199dc4a26bc8551ade0c8aea2db1c

调用queryBusinessPartnerInfo()查询商业合作伙伴zlls信息：
echo "GET query chaincode BusinessPartnerInfo on peer1 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/BusinessPartnerInfo?peer=peer1&fcn=queryBusinessPartnerInfo&args=%5B%22zlls%22%5D" \
  -H "authorization: Bearer $ORG1_TOKEN" \
  -H "content-type: application/json"
echo
echo
返回值：{"UserName":"zlls","Organization":"律师事务所","Company":"北京市中伦律师事务所","Account":"zlls666","CreatedTime":"2017-08-25T03:33:16.431Z","OperateLog":null}

TxRecorder链码调用addOperateLog()给商业合作伙伴追加操作记录的TxID
传入参数：UserName，TxID
返回值为追加操作的TxID


其他示例数据：
"{\"UserName\":\"zlls\",\"Organization\":\"律师事务所\",\"Company\":\"北京市中伦律师事务所\",\"Account\":\"zlls124\"}"

"{\"UserName\":\"zrgj\",\"Organization\":\"信托机构\",\"Company\":\"中融国际信托有限公司\",\"Account\":\"zrgj849\"}"

"{\"UserName\":\"qbht\",\"Organization\":\"原始债权人\",\"Company\":\"钱包汇通保理公司\",\"Account\":\"qbht439\"}"

"{\"UserName\":\"qbjf\",\"Organization\":\"劣后级认购方\",\"Company\":\"钱包金服科技公司\",\"Account\":\"qbjf293\"}"

"{\"UserName\":\"shyh\",\"Organization\":\"次优级认购方\",\"Company\":\"上海银行\",\"Account\":\"shyh934\"}"

"{\"UserName\":\"zrj\",\"Organization\":\"优先级认购方\",\"Company\":\"中融金（北京）科技有限公司\",\"Account\":\"zrj243\"}"

"{\"UserName\":\"amdq\",\"Organization\":\"差额支付承诺人\",\"Company\":\"广东奥马电器股份有限公司\",\"Account\":\"amdq539\"}"

"{\"UserName\":\"zhypg\",\"Organization\":\"资产评级机构\",\"Company\":\"北京中和谊资产评估有限公司\",\"Account\":\"zhypg873\"}"
