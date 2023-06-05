package chaincode

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) IsInit(ctx contractapi.TransactionContextInterface) error {

	log.Printf("Initial Isinit run")

	// Initial Isinit run
	err := ctx.GetStub().PutState("key", []byte("key"))
	if err != nil {
		return err
	}

	return nil

}

type Kyle struct {
	Key   string
	Value string
}

func (s *SmartContract) GetData(ctx contractapi.TransactionContextInterface, args map[string]interface{}) (interface{}, error) {

	a := ctx.GetStub().GetArgs()
	b := ctx.GetStub().GetStringArgs()
	c, _ := ctx.GetStub().GetArgsSlice()
	d, e := ctx.GetStub().GetFunctionAndParameters()

	f, _ := ctx.GetClientIdentity().GetID()
	g, _ := ctx.GetClientIdentity().GetMSPID()

	h, _ := ctx.GetStub().GetCreator()

	var i interface{}
	err := json.Unmarshal(h, &i)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	key := args["key"].(string)

	// Initial Isinit run
	bytes, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SmartContract) PutData(ctx contractapi.TransactionContextInterface, args map[string]interface{}) error {

	key := args["key"].(string)
	value := args["value"].(string)

	in := Kyle{Key: key, Value: value}
	in2, err := json.Marshal(in)
	if err != nil {
		return err
	}
	// Initial Isinit run
	err = ctx.GetStub().PutState(key, in2)
	if err != nil {
		return err
	}

	return nil
}

func (s *SmartContract) Kyle(ctx contractapi.TransactionContextInterface, a string, b ...string) error {

	fmt.Println(a)
	fmt.Println(b)
	return nil
}

func (s *SmartContract) EmitEvent1(ctx contractapi.TransactionContextInterface) (string, error) {

	result := "Hi!"
	bytes, _ := json.Marshal(result)
	err := ctx.GetStub().SetEvent("Test1", bytes)
	if err != nil {
		return "fail", fmt.Errorf("failed to set event: %v", err)
	}

	return "success", nil
}

func (s *SmartContract) EmitEvent2(ctx contractapi.TransactionContextInterface) (string, error) {

	result := "Hi!"
	bytes, _ := json.Marshal(result)

	err := ctx.GetStub().SetEvent("Test2", bytes)
	if err != nil {
		return "fail", fmt.Errorf("failed to set event: %v", err)
	}

	return "success", nil
}

func (s *SmartContract) CallOtherChaincodeFunction(ctx contractapi.TransactionContextInterface, other string, funcNm string, args map[string]interface{}) (interface{}, error) {

	otherChaincodeName := other

	channelID := ctx.GetStub().GetChannelID()

	bytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	arg := [][]byte{[]byte(funcNm), bytes}
	fmt.Println(arg)

	response := ctx.GetStub().InvokeChaincode(otherChaincodeName, arg, channelID)
	if response.GetStatus() != 200 {
		return nil, fmt.Errorf("failed to call function in chaincode %s", otherChaincodeName)
	}

	return response.String(), nil
}
