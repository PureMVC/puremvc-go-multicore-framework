//
//  Facade_test.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package facade

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/facade"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/mediator"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/proxy"
	"testing"
)

/**
Test the PureMVC Facade class.
*/

func TestGetInstance(t *testing.T) {
	// Test Factory Method
	var facade = facade.GetInstance("FacadeTestKey1", func() interfaces.IFacade { return &facade.Facade{Key: "FacadeTestKey1"} })

	// test assertions
	if facade == nil {
		t.Error("Expecting instance not nil")
	}
}

/**
  Tests Command registration and execution via the Facade.

  This test gets a Multiton Facade instance
  and registers the FacadeTestCommand class
  to handle 'FacadeTest' Notifcations.

  It then sends a notification using the Facade.
  Success is determined by evaluating
  a property on an object placed in the body of
  the Notification, which will be modified by the Command.
*/
func TestRegisterCommandAndSendNotification(t *testing.T) {
	// Create the Facade, register the FacadeTestCommand to
	// handle 'FacadeTest' notifications
	var facade = facade.GetInstance("FacadeTestKey2", func() interfaces.IFacade { return &facade.Facade{Key: "FacadeTestKey2"} })
	facade.RegisterCommand("FacadeTestNote", func() interfaces.ICommand { return &FacadeTestCommand{} })

	// Send notification. The Command associated with the event
	// (FacadeTestCommand) will be invoked, and will multiply
	// the vo.input value by 2 and set the result on vo.result
	var vo = FacadeTestVO{Input: 32}
	facade.SendNotification("FacadeTestNote", &vo, "")

	// test assertions
	if vo.Result != 64 {
		t.Error("Expecting vo.result == 64")
	}
}

/**
  Tests Command removal via the Facade.

  This test gets a Multiton Facade instance
  and registers the FacadeTestCommand class
  to handle 'FacadeTest' Notifcations. Then it removes the command.

  It then sends a Notification using the Facade.
  Success is determined by evaluating
  a property on an object placed in the body of
  the Notification, which will NOT be modified by the Command.
*/
func TestRegisterAndRemoveCommandAndSendNotification(t *testing.T) {
	// Create the Facade, register the FacadeTestCommand to
	// handle 'FacadeTest' events
	var facade = facade.GetInstance("FacadeTestKey3", func() interfaces.IFacade { return &facade.Facade{Key: "FacadeTestKey2"} })
	facade.RegisterCommand("FacadeTestNote", func() interfaces.ICommand { return &FacadeTestCommand{} })
	facade.RemoveCommand("FacadeTestNote")

	// Send notification. The Command associated with the event
	// (FacadeTestCommand) will NOT be invoked, and will NOT multiply
	// the vo.input value by 2
	var vo = FacadeTestVO{Input: 32}
	facade.SendNotification("FacadeTestNote", &vo, "")

	// test assertions
	if vo.Result == 64 {
		t.Error("Expecting vo.result != 64")
	}
}

/**
  Tests the regsitering and retrieving Model proxies via the Facade.

  Tests `registerProxy` and `retrieveProxy` in the same test.
  These methods cannot currently be tested separately
  in any meaningful way other than to show that the
  methods do not throw exception when called.
*/
func TestRegisterAndRetrieveProxy(t *testing.T) {
	// register a proxy and retrieve it.
	var facade = facade.GetInstance("FacadeTestKey4", func() interfaces.IFacade {
		return &facade.Facade{Key: "FacadeTestKey4"}
	})
	facade.RegisterProxy(&proxy.Proxy{Name: "colors", Data: []string{"red", "green", "blue"}})
	var proxy = facade.RetrieveProxy("colors").(*proxy.Proxy)

	// retrieve data from proxy
	var data = proxy.Data.([]string)

	// test assertions
	if data == nil {
		t.Error("Expecting data not nil")
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

/**
  Tests the removing Proxies via the Facade.
*/
func TestRegisterAndRemoveProxy(t *testing.T) {
	// register a proxy, remove it, then try to retrieve it
	var facade = facade.GetInstance("FacadeTestKey5", func() interfaces.IFacade {
		return &facade.Facade{Key: "FacadeTestKey5"}
	})
	var proxy interfaces.IProxy = &proxy.Proxy{Name: "sizes", Data: []string{"7", "13", "21"}}
	facade.RegisterProxy(proxy)

	// remove the proxy
	var removedProxy = facade.RemoveProxy("sizes")

	// assert that we removed the appropriate proxy
	if removedProxy.GetProxyName() != "sizes" {
		t.Error("Expecting removedProxy.GetProxyName() == 'sizes'")
	}

	// make sure we can no longer retrieve the proxy from the model
	var proxy2 = facade.RetrieveProxy("sizes")
	// test assertions
	if proxy2 != nil {
		t.Error("Expecting proxy is nil")
	}
}

/**
  Tests registering, retrieving and removing Mediators via the Facade.
*/
func TestRegisterRetrieveAndRemoveMediator(t *testing.T) {
	// register a mediator, remove it, then try to retrieve it
	var facade = facade.GetInstance("FacadeTestKey6", func() interfaces.IFacade { return &facade.Facade{Key: "FacadeTestKey6"} })
	facade.RegisterMediator(&mediator.Mediator{Name: mediator.NAME, ViewComponent: []string{}})

	// retrieve the mediator
	if facade.RetrieveMediator(mediator.NAME) == nil {
		t.Error("Expecting mediator is not nil")
	}

	// remove the mediator
	removedMediator := facade.RemoveMediator(mediator.NAME)

	// assert that we have removed the appropriate mediator
	if removedMediator.GetMediatorName() != mediator.NAME {
		t.Error("Expecting removedMediator.GetMediatorName() == Mediator.NAME")
	}

	// assert that the mediator is no longer retrievable
	if facade.RetrieveMediator(mediator.NAME) != nil {
		t.Error("Expecting facade.RetrieveMediator( Mediator.NAME ) == nil")
	}
}

/**
  Tests the hasProxy Method
*/
func TestHasProxy(t *testing.T) {
	// register a Proxy
	var facade = facade.GetInstance("FacadeTestKey7", func() interfaces.IFacade { return &facade.Facade{Key: "FacadeTestKey7"} })
	facade.RegisterProxy(&proxy.Proxy{Name: "hasProxyTest", Data: []int{1, 2, 3}})

	// assert that the model.hasProxy method returns true
	// for that proxy name
	if facade.HasProxy("hasProxyTest") != true {
		t.Error("Expecting facade.HasProxy('hasProxyTest') == true")
	}
}

/**
  Tests the hasMediator Method
*/
func TestHasMediator(t *testing.T) {
	// register a Mediator
	var facade = facade.GetInstance("FacadeTestKey8", func() interfaces.IFacade { return &facade.Facade{Key: "FacadeTestKey8"} })
	facade.RegisterMediator(&mediator.Mediator{Name: "facadeHasMediatorTest", ViewComponent: []int{}})

	// assert that the facade.hasMediator method returns true
	// for that mediator name
	if facade.HasMediator("facadeHasMediatorTest") != true {
		t.Error("Expecting facade.HasMediator('facadeHasMediatorTest') == true")
	}

	facade.RemoveMediator("facadeHasMediatorTest")

	// assert that the facade.hasMediator method returns false
	// for that mediator name
	if facade.HasMediator("facadeHasMediatorTest") != false {
		t.Error("Expecting facade.HasMediator('facadeHasMediatorTest') == false")
	}
}

/**
  Test hasCommand method.
*/
func TestHasCommand(t *testing.T) {
	// register the ControllerTestCommand to handle 'hasCommandTest' notes
	var facade = facade.GetInstance("FacadeTestKey10", func() interfaces.IFacade {
		return &facade.Facade{Key: "FacadeTestKey10"}
	})
	facade.RegisterCommand("facadeHasCommandTest", func() interfaces.ICommand { return &FacadeTestCommand{} })

	// test that hasCommand returns true for hasCommandTest notifications
	if facade.HasCommand("facadeHasCommandTest") != true {
		t.Error("Expecting facade.HasCommand('facadeHasCommandTest') == true")
	}

	// Remove the Command from the Controller
	facade.RemoveCommand("facadeHasCommandTest")

	// test that hasCommand returns false for hasCommandTest notifications
	if facade.HasCommand("facadeHasCommentTest") != false {
		t.Error("Expecting facade.HasCommand('facadeHasCommandTest') == false")
	}
}

/**
  Tests the hasCore and removeCore methods
*/
func TestHasCoreAndRemoveCore(t *testing.T) {
	// assert that the Facade.hasCore method returns false first
	if facade.HasCore("FacadeTestKey11") != false {
		t.Error("Expecting facade.HasCore('FacadeTestKey11') == false")
	}

	// register a Core
	_ = facade.GetInstance("FacadeTestKey11", func() interfaces.IFacade {
		return &facade.Facade{Key: "FacadeTestKey11"}
	})

	// assert that the Facade.hasCore method returns true now that a Core is registered
	if facade.HasCore("FacadeTestKey11") != true {
		t.Error("Expecting facade.hasCore('FacadeTestKey11') == true")
	}

	// remove the Core
	facade.RemoveCore("FacadeTestKey11")

	// assert that the Facade.hasCore method returns false now that the core has been removed.
	if facade.HasCore("FacadeTestKey11") != false {
		t.Error("Expecting facade.HasCore('FacadeTestKey11') == false")
	}
}
