package chaincode

import (
	"fmt"
	"log"
	"math/rand"
	"sync"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) IsInit(ctx contractapi.TransactionContextInterface) error {

	log.Printf("Initial Isinit run")

	// Initial Isinit run
	err := ctx.GetStub().PutState("Isinit", []byte("Isinit"))
	if err != nil {
		return err
	}

	return nil
}

func (s *SmartContract) Minter(ctx contractapi.TransactionContextInterface) error {

	num := rand.Intn(100)

	err := ctx.GetStub().PutState("test", []byte(string(num)))
	if err != nil {
		return fmt.Errorf("failed to put state: %v", err)
	}

	return nil
}

func (s *SmartContract) Concurrency(ctx contractapi.TransactionContextInterface) error {

	var wg sync.WaitGroup

	wg.Add(1)

	errChan := make(chan error, 1)

	for i := 0; i < 2; i++ {
		go func(ctx contractapi.TransactionContextInterface) {
			defer wg.Done()

			if err := getState(ctx); err != nil {
				errChan <- err
			}
		}(ctx)
	}
	wg.Wait()

	select {
	case err := <-errChan:
		fmt.Println(err)
		return err

	default:
	}

	return nil
}

func getState(ctx contractapi.TransactionContextInterface) error {

	err := ctx.GetStub().PutState("Check", []byte("Check"))
	if err != nil {
		return err
	}

	_, err = ctx.GetStub().GetState("Check")
	if err != nil {
		return err
	}

	return nil
}
