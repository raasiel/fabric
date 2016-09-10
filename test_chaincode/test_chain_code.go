package main

import (
        "errors"
        "fmt"
        "github.com/hyperledger/fabric/core/chaincode/shim"
      )

type AssetManagementChaincode struct {
}

func (t *AssetManagementChaincode) Init(stub shim.ChaincodeStubInterface,function string, args []string) ([]byte,error) {
       return nil, nil
   }


func (t *AssetManagementChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte,error) {

       stub.PutState(args[0], []byte(args[1]))
      return nil, nil
}

func (t *AssetManagementChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte,error) {

        var err error
        savedUser, err := stub.GetState(args[0])

       if err != nil {
            return nil,errors.New("User not found")
        }

      return savedUser, nil
 }


 func main() {
      err := shim.Start(new(AssetManagementChaincode))
      if err != nil {
           fmt.Printf("Error starting AssetManagementChaincode: %s",err)
      }
 }
