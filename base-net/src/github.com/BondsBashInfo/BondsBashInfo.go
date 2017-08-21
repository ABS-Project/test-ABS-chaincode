// 债券基础信息相关的链码操作

/*

*/

package main


import (
	"fmt"
	"strconv"
  "time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("BondsBashInfo")

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// BondsBashInfo struct
// ============================================================================================================================
type UserStruct struct {

}


func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response  {
	logger.Info("########### BondsBashInfo Init ###########")
	return shim.Success(nil)


}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### example_cc0 Invoke ###########")

	function, args := stub.GetFunctionAndParameters()
  if function == "addBonds" {
		// Deletes an entity from its state
		return t.addBonds(stub, args)
	}

	if function == "deleteBonds" {
		// Deletes an entity from its state
		return t.deleteBonds(stub, args)
	}

	if function == "queryBonds" {
		// queries an entity state
		return t.queryBonds(stub, args)
	}
	if function == "updateBonds" {
		// Deletes an entity from its state
		return t.updateBonds(stub, args)
	}

	logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0]))
}

func (t *SimpleChaincode) addBonds(stub shim.ChaincodeStubInterface, args []string) pb.Response {


        return shim.Success(nil);
}

// Deletes an entity from state
func (t *SimpleChaincode) deleteBonds(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}


	return shim.Success(nil)
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) queryBonds(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}
}

func (t *SimpleChaincode) updateBonds(stub shim.ChaincodeStubInterface, args []string) pb.Response {


        return shim.Success(nil);
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}
