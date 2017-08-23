// 操作记录链码操作

/*

*/

package main
import (
	"fmt"
  "time"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/common/util"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("TxRecorder")

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}


// ============================================================================================================================
// TxInfo struct
// ============================================================================================================================
type TxInfoStruct struct{
	TxID               string     `json:"TxID"`          //交易ID
	TxProposer         string     `json:"TxProposer"`    //交易发起人
	TxTime             time.Time  `json:"TxTime"`        //交易时间
	TxChaincode        string     `json:"TxChaincode"`   //链码名称
	TxFunction         string     `json:"TxFunction"`    //所调函数
	TxArguments        string     `json:"TxArguments"`   //所传参数
	TxDescription      string     `json:"TxDescription"` //交易描述
}


func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response  {
	logger.Info("########### TxRecorder Init ###########")
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### TxRecorder Invoke ###########")

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

	logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0]))
}

func (t *SimpleChaincode) add(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 2. ")
	}

	TxID := args[0]
	TxTest, _ := stub.GetState(TxID)
	if TxTest != nil {
		return shim.Error("the Transaction is sexisted")
	}
	TxProposer := args[1]
	functionName := "addOperateLog"
	invokeArgs := util.ToChaincodeArgs(functionName,TxProposer,TxID)
	response := stub.InvokeChaincode("BusinessPartnerInfo", invokeArgs, "mychannel")
	if response.Status != shim.OK {
			errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", string(response.Payload))
			fmt.Printf(errStr)
			return shim.Error(errStr)
	}
	var TxInfoObj TxInfoStruct
	TxInfoObj.TxID = args[0]
	TxInfoObj.TxProposer = args[1]
	TxInfoObj.TxTime,_ =time.Parse("2006-01-02T15:04:05.000Z",args[2])
	TxInfoObj.TxChaincode = args[3]
	TxInfoObj.TxFunction = args[4]
	TxInfoObj.TxArguments = args[5]
	TxInfoObj.TxDescription = args[6]

	jsonAsBytes,_:= json.Marshal(TxInfoObj)
	err = stub.PutState(TxID,[]byte(jsonAsBytes))
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
	TxID := args[0]
 	TxInfo, err := stub.GetState(TxID)
 if err != nil {
 	jsonResp := "{\"Error\":\"Failed to get state for " + TxID + "\"}"
 	return shim.Error(jsonResp)
  }
 if TxInfo == nil {
 	jsonResp := "{\"Error\":\"Nil content for " + TxID + "\"}"
 	return shim.Error(jsonResp)
  }
 return shim.Success(TxInfo)
}



func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}
