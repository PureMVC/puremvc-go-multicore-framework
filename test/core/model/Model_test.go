//
//  Model_test.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package model

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/core/model"
	"github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/proxy"
	"testing"
)

/*
Test the PureMVC Model class.
*/

func TestGetInstance(t *testing.T) {
	// Test Factory Method
	var m = model.GetInstance("ModelTestKey1", func() interfaces.IModel { return &model.Model{Key: "ModelTestKey1"} })

	// test assertions
	if m == nil {
		t.Error("Expecting instance not nil")
	}
}

/*
Tests the proxy registration and retrieval methods.

Tests registerProxy and retrieveProxy in the same test.
These methods cannot currently be tested separately
in any meaningful way other than to show that the
methods do not throw exception when called.
*/
func TestRegisterAndRetrieveProxy(t *testing.T) {
	// register a proxy and retrieve it.
	var m = model.GetInstance("ModelTestKey2", func() interfaces.IModel { return &model.Model{Key: "ModelTestKey2"} })
	m.RegisterProxy(&proxy.Proxy{Name: "colors", Data: []string{"red", "green", "blue"}})
	var p = m.RetrieveProxy("colors")
	var data = p.GetData().([]string)

	// test assertions
	if data == nil {
		t.Error("Expecting data not nill")
	}
	if len(data) != 3 {
		t.Error("Expecting len(data) == 3")
	}
	if data[0] != "red" {
		t.Error("Expecting data[0] == 'red'")
	}
	if data[1] != "green" {
		t.Error("Expecting data[1] == 'green'")
	}
	if data[2] != "blue" {
		t.Error("Expecting data[2] == 'blue'")
	}
}

/*
Tests the proxy removal method.
*/
func TestRegisterAndRemoveProxy(t *testing.T) {
	// register a proxy, remove it, then try to retrieve it
	var m = model.GetInstance("ModelTestKey3", func() interfaces.IModel { return &model.Model{Key: "ModelTestKey3"} })
	var p interfaces.IProxy = &proxy.Proxy{Name: "sizes", Data: []string{"7", "13", "21"}}
	m.RegisterProxy(p)

	// remove the proxy
	var removedProxy = m.RemoveProxy("sizes")

	// assert that we removed the appropriate proxy
	if removedProxy.GetProxyName() != "sizes" {
		t.Error("Expecting removedProxy.GetProxyName() == 'sizes'")
	}

	// ensure that the proxy is no longer retrievable from the model
	var nilProxy = m.RetrieveProxy("sizes")

	// test assertions
	if nilProxy != nil {
		t.Error("Expecting proxy is nil")
	}
}

/*
Tests the hasProxy Method
*/
func TestHasProxy(t *testing.T) {
	// register a proxy
	var m = model.GetInstance("ModelTestKey4", func() interfaces.IModel { return &model.Model{Key: "ModelTestKey4"} })
	var p interfaces.IProxy = &proxy.Proxy{Name: "aces", Data: []string{"clubs", "spades", "hearts", "diamonds"}}
	m.RegisterProxy(p)

	// assert that the model.hasProxy method returns true
	// for that proxy name
	if m.HasProxy("aces") != true {
		t.Error("Expecting model.HasProxy('aces') == true")
	}

	// remove the proxy
	m.RemoveProxy("aces")

	// assert that the model.hasProxy method returns false
	// for that proxy name
	if m.HasProxy("aces") != false {
		t.Error("Expecting model.HasProxy('aces') == false")
	}
}

/*
Tests that the Model calls the onRegister and onRemove methods
*/
func TestOnRegisterAndOnRemove(t *testing.T) {
	// Get a Multiton View instance
	var m = model.GetInstance("ModelTestKey5", func() interfaces.IModel { return &model.Model{Key: "ModelTestKey5"} })

	// Create and register the test mediator
	var p interfaces.IProxy = &ModelTestProxy{proxy.Proxy{Name: MODEL_TEST_PROXY}}
	m.RegisterProxy(p)

	// assert that onRegsiter was called, and the proxy responded by setting its data accordingly
	if p.GetData() != ON_REGISTER_CALLED {
		t.Error("Expecting proxy.GetData() == ON_REGISTER_CALLED")
	}

	// Remove the component
	m.RemoveProxy(MODEL_TEST_PROXY)

	// assert that onRemove was called, and the proxy responded by setting its data accordingly
	if p.GetData() != ON_REMOVE_CALLED {
		t.Error("Expecting proxy.GetData() == ON_REMOVE_CALLED")
	}
}
