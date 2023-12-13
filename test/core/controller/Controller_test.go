//
//  Controller_test.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package controller

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/core/controller"
	"github.com/puremvc/puremvc-go-multicore-framework/src/core/view"
	"github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/observer"
	"testing"
)

/*
Test the PureMVC Controller class.
*/

/*
Tests the Controller Multiton Factory Method
*/
func TestGetInstance(t *testing.T) {
	var c = controller.GetInstance("ControllerTestKey1", func() interfaces.IController { return &controller.Controller{Key: "ControllerTestKey1"} })

	if c == nil {
		t.Error("Expecting instance not nil")
	}
}

/*
Tests Command registration and execution.

This test gets a Multiton Controller instance
and registers the ControllerTestCommand class
to handle 'ControllerTest' Notifications.

It then constructs such a Notification and tells the
Controller to execute the associated Command.
Success is determined by evaluating a property
on an object passed to the Command, which will
be modified when the Command executes.
*/
func TestRegisterAndExecuteCommand(t *testing.T) {
	// Create the controller, register the ControllerTestCommand to handle 'ControllerTest' notes
	var c = controller.GetInstance("ControllerTestKey2", func() interfaces.IController { return &controller.Controller{Key: "ControllerTestKey2"} })
	c.RegisterCommand("ControllerTest", func() interfaces.ICommand { return &ControllerTestCommand{} })

	// Create a 'ControllerTest' note
	var vo = ControllerTestVO{Input: 12}
	var note interfaces.INotification = observer.NewNotification("ControllerTest", &vo, "")

	// Tell the controller to execute the Command associated with the note
	// the ControllerTestCommand invoked will multiply the vo.input value
	// by 2 and set the result on vo.result
	c.ExecuteCommand(note)

	// test assertions
	if vo.Result != 24 {
		t.Error("Expecting vo.Result == 24")
	}
}

/*
Tests Command registration and removal.

Tests that once a Command is registered and verified
working, it can be removed from the Controller.
*/
func TestRegisterAndRemoveCommand(t *testing.T) {
	// Create the controller, register the ControllerTestCommand to handle 'ControllerTest' notes
	var c = controller.GetInstance("ControllerTestKey3", func() interfaces.IController { return &controller.Controller{Key: "ControllerTestKey3"} })
	c.RegisterCommand("ControllerRemoveTest", func() interfaces.ICommand { return &ControllerTestCommand{} })

	// Create a 'ControllerTest' note
	var vo ControllerTestVO = ControllerTestVO{Input: 12}
	var note interfaces.INotification = observer.NewNotification("ControllerRemoveTest", &vo, "")

	// Tell the controller to execute the Command associated with the note
	// the ControllerTestCommand invoked will multiply the vo.input value
	// by 2 and set the result on vo.result
	c.ExecuteCommand(note)

	// test assertions
	if vo.Result != 24 {
		t.Error("Expecting vo.result == 24")
	}

	// Reset result
	vo.Result = 0

	// Remove the Command from the Controller
	c.RemoveCommand("ControllerRemoveTest")

	// Tell the controller to execute the Command associated with the
	// note. This time, it should not be registered, and our vo result
	// will not change
	c.ExecuteCommand(note)

	// test assertions
	if vo.Result != 0 {
		t.Error("Expecting vo.result == 0")
	}
}

/*
Test hasCommand method.
*/
func TestHasCommand(t *testing.T) {
	// register the ControllerTestCommand to handle 'hasCommandTest' notes
	var c = controller.GetInstance("ControllerTestKey4", func() interfaces.IController { return &controller.Controller{Key: "ControllerTestKey4"} })
	c.RegisterCommand("hasCommandTest", func() interfaces.ICommand { return &ControllerTestCommand{} })

	// test that hasCommand returns true for hasCommandTest notifications
	if c.HasCommand("hasCommandTest") == false {
		t.Error("Expecting controller.HasCommand('hasCommandTest') == true")
	}

	// Remove the Command from the Controller
	c.RemoveCommand("hasCommandTest")

	// test that hasCommand returns false for hasCommandTest notifications
	if c.HasCommand("hasCommandTest") == true {
		t.Error("Expecting controller.HasCommand('hasCommandTest') == false")
	}
}

/*
Tests Removing and Reregistering a Command

Tests that when a Command is re-registered that it isn't fired twice.
This involves, minimally, registration with the controller but
notification via the View, rather than direct execution of
the Controller's executeCommand method as is done above in
testRegisterAndRemove.
*/
func TestReregisterAndExecuteCommand(t *testing.T) {
	// Fetch the controller, register the ControllerTestCommand2 to handle 'ControllerTest2' notes
	var c = controller.GetInstance("ControllerTestKey5", func() interfaces.IController { return &controller.Controller{Key: "ControllerTestKey5"} })
	c.RegisterCommand("ControllerTest2", func() interfaces.ICommand { return &ControllerTestCommand2{} })

	// Remove the Command from the Controller
	c.RemoveCommand("ControllerTest2")

	// Re-register the Command with the Controller
	c.RegisterCommand("ControllerTest2", func() interfaces.ICommand { return &ControllerTestCommand2{} })

	// Create a 'ControllerTest2' note
	var vo *ControllerTestVO = &ControllerTestVO{Input: 12}
	var note interfaces.INotification = observer.NewNotification("ControllerTest2", vo, "")

	// retrieve a reference to the View from the same core.
	var v interfaces.IView = view.GetInstance("ControllerTestKey5", func() interfaces.IView { return &view.View{Key: "ControllerTestKey5"} })

	// send the notification
	v.NotifyObservers(note)

	// test assertions
	// if the command is executed once the value will be 24
	if vo.Result != 24 {
		t.Error("Expecting vo.result == 24")
	}

	// Prove that accumulation works in the VO by sending the notification again
	v.NotifyObservers(note)

	// if the command is executed twice the value will be 48
	if vo.Result != 48 {
		t.Error("Expecting vo.result == 48")
	}
}
