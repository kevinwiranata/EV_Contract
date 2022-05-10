/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"testing"

// 	"github.com/hyperledger/fabric-contract-api-go/contractapi"
// 	"github.com/hyperledger/fabric-chaincode-go/shim"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// const getStateError = "world state get error"

// type MockStub struct {
// 	shim.ChaincodeStubInterface
// 	mock.Mock
// }

// func (ms *MockStub) GetState(key string) ([]byte, error) {
// 	args := ms.Called(key)

// 	return args.Get(0).([]byte), args.Error(1)
// }

// func (ms *MockStub) PutState(key string, value []byte) error {
// 	args := ms.Called(key, value)

// 	return args.Error(0)
// }

// func (ms *MockStub) DelState(key string) error {
// 	args := ms.Called(key)

// 	return args.Error(0)
// }

// type MockContext struct {
// 	contractapi.TransactionContextInterface
// 	mock.Mock
// }

// func (mc *MockContext) GetStub() shim.ChaincodeStubInterface {
// 	args := mc.Called()

// 	return args.Get(0).(*MockStub)
// }

// func configureStub() (*MockContext, *MockStub) {
// 	var nilBytes []byte

// 	testEv := new(Ev)
// 	testEv.Value = "set value"
// 	evBytes, _ := json.Marshal(testEv)

// 	ms := new(MockStub)
// 	ms.On("GetState", "statebad").Return(nilBytes, errors.New(getStateError))
// 	ms.On("GetState", "missingkey").Return(nilBytes, nil)
// 	ms.On("GetState", "existingkey").Return([]byte("some value"), nil)
// 	ms.On("GetState", "evkey").Return(evBytes, nil)
// 	ms.On("PutState", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).Return(nil)
// 	ms.On("DelState", mock.AnythingOfType("string")).Return(nil)

// 	mc := new(MockContext)
// 	mc.On("GetStub").Return(ms)

// 	return mc, ms
// }

// func TestEvExists(t *testing.T) {
// 	var exists bool
// 	var err error

// 	ctx, _ := configureStub()
// 	c := new(EvContract)

// 	exists, err = c.EvExists(ctx, "statebad")
// 	assert.EqualError(t, err, getStateError)
// 	assert.False(t, exists, "should return false on error")

// 	exists, err = c.EvExists(ctx, "missingkey")
// 	assert.Nil(t, err, "should not return error when can read from world state but no value for key")
// 	assert.False(t, exists, "should return false when no value for key in world state")

// 	exists, err = c.EvExists(ctx, "existingkey")
// 	assert.Nil(t, err, "should not return error when can read from world state and value exists for key")
// 	assert.True(t, exists, "should return true when value for key in world state")
// }

// func TestCreateEv(t *testing.T) {
// 	var err error

// 	ctx, stub := configureStub()
// 	c := new(EvContract)

// 	err = c.CreateEv(ctx, "statebad", "some value")
// 	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

// 	err = c.CreateEv(ctx, "existingkey", "some value")
// 	assert.EqualError(t, err, "The asset existingkey already exists", "should error when exists returns true")

// 	err = c.CreateEv(ctx, "missingkey", "some value")
// 	stub.AssertCalled(t, "PutState", "missingkey", []byte("{\"value\":\"some value\"}"))
// }

// func TestReadEv(t *testing.T) {
// 	var ev *Ev
// 	var err error

// 	ctx, _ := configureStub()
// 	c := new(EvContract)

// 	ev, err = c.ReadEv(ctx, "statebad")
// 	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when reading")
// 	assert.Nil(t, ev, "should not return Ev when exists errors when reading")

// 	ev, err = c.ReadEv(ctx, "missingkey")
// 	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when reading")
// 	assert.Nil(t, ev, "should not return Ev when key does not exist in world state when reading")

// 	ev, err = c.ReadEv(ctx, "existingkey")
// 	assert.EqualError(t, err, "Could not unmarshal world state data to type Ev", "should error when data in key is not Ev")
// 	assert.Nil(t, ev, "should not return Ev when data in key is not of type Ev")

// 	ev, err = c.ReadEv(ctx, "evkey")
// 	expectedEv := new(Ev)
// 	expectedEv.Value = "set value"
// 	assert.Nil(t, err, "should not return error when Ev exists in world state when reading")
// 	assert.Equal(t, expectedEv, ev, "should return deserialized Ev from world state")
// }

// func TestUpdateEv(t *testing.T) {
// 	var err error

// 	ctx, stub := configureStub()
// 	c := new(EvContract)

// 	err = c.UpdateEv(ctx, "statebad", "new value")
// 	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when updating")

// 	err = c.UpdateEv(ctx, "missingkey", "new value")
// 	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when updating")

// 	err = c.UpdateEv(ctx, "evkey", "new value")
// 	expectedEv := new(Ev)
// 	expectedEv.Value = "new value"
// 	expectedEvBytes, _ := json.Marshal(expectedEv)
// 	assert.Nil(t, err, "should not return error when Ev exists in world state when updating")
// 	stub.AssertCalled(t, "PutState", "evkey", expectedEvBytes)
// }

// func TestDeleteEv(t *testing.T) {
// 	var err error

// 	ctx, stub := configureStub()
// 	c := new(EvContract)

// 	err = c.DeleteEv(ctx, "statebad")
// 	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

// 	err = c.DeleteEv(ctx, "missingkey")
// 	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when deleting")

// 	err = c.DeleteEv(ctx, "evkey")
// 	assert.Nil(t, err, "should not return error when Ev exists in world state when deleting")
// 	stub.AssertCalled(t, "DelState", "evkey")
// }
