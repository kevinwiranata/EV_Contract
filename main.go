/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

func main() {
	evContract := new(EVContract)
	evContract.Info.Version = "0.0.1"
	evContract.Info.Description = "My Smart Contract"
	evContract.Info.License = new(metadata.LicenseMetadata)
	evContract.Info.License.Name = "Apache-2.0"
	evContract.Info.Contact = new(metadata.ContactMetadata)
	evContract.Info.Contact.Name = "John Doe"

	chaincode, err := contractapi.NewChaincode(evContract)
	chaincode.Info.Title = "EV_Test chaincode"
	chaincode.Info.Version = "0.0.1"

	if err != nil {
		panic("Could not create chaincode from EvContract." + err.Error())
	}

	err = chaincode.Start()

	if err != nil {
		panic("Failed to start chaincode. " + err.Error())
	}
}
