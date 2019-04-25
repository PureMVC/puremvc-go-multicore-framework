//
//  View_test.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package view

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/core/view"
	"github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/mediator"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/observer"
	"testing"
)

/**
Test the PureMVC View class.
*/

/**
  Tests the View Multiton Factory Method
*/
func TestGetInstance(t *testing.T) {
	// Test Factory Method
	var view = view.GetInstance("ViewTestKey1", func() interfaces.IView { return &view.View{Key: "ViewTestKey1"} })

	// test assertions
	if view == nil {
		t.Error("Expecting instance not nil")
	}
}

/**
  Tests registration and notification of Observers.

  An Observer is created to callback the viewTestMethod of
  this ViewTest instance. This Observer is registered with
  the View to be notified of 'ViewTestEvent' events. Such
  an event is created, and a value set on its payload. Then
  the View is told to notify interested observers of this
  Event.

  The View calls the Observer's notifyObserver method
  which calls the viewTestMethod on this instance
  of the ViewTest class. The viewTestMethod method will set
  an instance variable to the value passed in on the Event
  payload. We evaluate the instance variable to be sure
  it is the same as that passed out as the payload of the
  original 'ViewTestEvent'.
*/
func TestRegisterAndNotifyObserver(t *testing.T) {
	// Get the Multiton View instance
	var view = view.GetInstance("ViewTestKey2", func() interfaces.IView { return &view.View{Key: "ViewTestKey2"} })

	// Create observer, passing in notification method and context
	var observerTest = ObserverTest{NotifyMethod: NotifyTestMethod}
	var obs = &observer.Observer{Notify: observerTest.NotifyMethod, Context: observerTest}

	// Register Observer's interest in a particulat Notification with the View
	view.RegisterObserver(ViewTestNote_NAME, obs)

	// Create a ViewTestNote, setting
	// a body value, and tell the View to notify
	// Observers. Since the Observer is this class
	// and the notification method is viewTestMethod,
	// successful notification will result in our local
	// viewTestVar being set to the value we pass in
	// on the note body.
	var note = ViewTestNoteNew(10)
	view.NotifyObservers(note)

	// test assertions
	if ViewTestVar != 10 {
		t.Error("Expecting ViewTestVar = 10 ", ViewTestVar)
	}
}

/**
  Tests registering and retrieving a mediator with
  the View.
*/
func TestRegisterAndRetrieveMediator(t *testing.T) {
	// Get the Multiton View instance
	var view = view.GetInstance("ViewTestKey3", func() interfaces.IView { return &view.View{Key: "ViewTestKey3"} })

	// Create and register the test mediator
	var viewTestMediator interfaces.IMediator = &ViewTestMediator{Mediator: mediator.Mediator{Name: ViewTestMediator_NAME, ViewComponent: nil}}
	view.RegisterMediator(viewTestMediator)

	// Retrieve the component
	var mediator = view.RetrieveMediator(ViewTestMediator_NAME)

	// test assertions
	if mediator != viewTestMediator {
		t.Error("Expecting mediator is ViewTestMediator")
	}
}

/**
  Tests the hasMediator Method
*/
func TestHasMediator(t *testing.T) {
	// register a Mediator
	var view = view.GetInstance("ViewTestKey4", func() interfaces.IView { return &view.View{Key: "ViewTestKey4"} })

	// Create and register the test mediator
	var mediator interfaces.IMediator = &mediator.Mediator{Name: "hasMediatorTest", ViewComponent: nil}
	view.RegisterMediator(mediator)

	// assert that the view.hasMediator method returns true
	// for that mediator name
	if view.HasMediator("hasMediatorTest") != true {
		t.Error("Expecting view.HasMediator('hasMediatorTest')")
	}

	view.RemoveMediator("hasMediatorTest")

	// assert that the view.hasMediator method returns false
	// for that mediator name
	if view.HasMediator("hasMediatorTest") != false {
		t.Error("Expecting view.HasMediator('hasMediatorTest') == false")
	}
}

/**
  Tests registering and removing a mediator
*/
func TestRegisterAndRemoveMediator(t *testing.T) {
	// Get the Multiton View instance
	var view = view.GetInstance("ViewTestKey5", func() interfaces.IView { return &view.View{Key: "ViewTestKey5"} })

	// Create and register the test mediator
	var mediator interfaces.IMediator = &mediator.Mediator{Name: "testing", ViewComponent: nil}
	view.RegisterMediator(mediator)

	// Remove the component
	var removedMediator = view.RemoveMediator("testing")

	// assert that we have removed the appropriate mediator
	if removedMediator.GetMediatorName() != "testing" {
		t.Error("Expecting removedMediator.GetMediatorName() == 'testing'")
	}

	// assert that the mediator is no longer retrievable
	if view.RetrieveMediator("testing") != nil {
		t.Error("Expecting view.RetrieveMediator('testing') == nil")
	}
}

/**
  Tests that the View callse the onRegister and onRemove methods
*/
func TestOnRegisterAndOnRemove(t *testing.T) {
	// Get the Multiton View instance
	var view = view.GetInstance("ViewTestKey6", func() interfaces.IView { return &view.View{Key: "ViewTestKey6"} })

	// Create and register the test mediator
	var data = Data{}
	var mediator interfaces.IMediator = &ViewTestMediator4{mediator.Mediator{Name: ViewTestMediator4_NAME, ViewComponent: &data}}
	view.RegisterMediator(mediator)

	// assert that onRegsiter was called, and the mediator responded by setting our boolean
	if data.onRegisterCalled != true {
		t.Error("Expecting data.onRegisterCalled == true")
	}

	// Remove the component
	view.RemoveMediator(ViewTestMediator4_NAME)

	// assert that the mediator is no longer retrievable
	if data.onRemoveCalled != true {
		t.Error("Expecting onRemoveCalled == true")
	}
}

/**
  Tests successive regster and remove of same mediator.
*/
func TestSuccessiveRegisterAndRemoveMediator(t *testing.T) {
	// Get the Multiton View instance
	var view = view.GetInstance("ViewTestKey7", func() interfaces.IView { return &view.View{Key: "ViewTestKey7"} })
	var _mediator = &ViewTestMediator{Mediator: mediator.Mediator{Name: ViewTestMediator_NAME, ViewComponent: nil}}

	// Create and register the test mediator,
	// but not so we have a reference to it
	view.RegisterMediator(_mediator)

	// test that we can retrieve it
	if view.RetrieveMediator(ViewTestMediator_NAME) != _mediator {
		t.Error("Expecting view.RetrieveMediator(ViewTestMediatorNAME) == mediator")
	}

	// Remove the Mediator
	view.RemoveMediator(ViewTestMediator_NAME)

	// test that retrieving it now returns nil
	if view.RetrieveMediator(ViewTestMediator_NAME) != nil {
		t.Error("Expecting view.RetrieveMediator(ViewTestMediator.NAME) == nil")
	}

	// test that removing the mediator again once its gone doesn't cause crash
	if view.RetrieveMediator(ViewTestMediator_NAME) != nil {
		t.Error("Expecting view.RetrieveMediator(ViewTestMediator.NAME) == nil")
	}

	// Create and register another instance of the test mediator,
	view.RegisterMediator(&ViewTestMediator{Mediator: mediator.Mediator{Name: ViewTestMediator_NAME, ViewComponent: nil}})

	if view.RetrieveMediator(ViewTestMediator_NAME) == nil {
		t.Error("Expecting view.RetrieveMediator(ViewTestMediator_NAME) != nil")
	}

	// Remove the Mediator
	view.RemoveMediator(ViewTestMediator_NAME)

	// test that retrieving it now returns nil
	if view.RetrieveMediator(ViewTestMediator_NAME) != nil {
		t.Error("Expecting view.RetrieveMediator(ViewTestMediator_NAME) == nil")
	}
}

/**
  Tests registering a Mediator for 2 different notifications, removing the
  Mediator from the View, and seeing that neither notification causes the
  Mediator to be notified.
*/
func TestRemoveMediatorAndSubsequentNotify(t *testing.T) {
	// Get the Multiton View instance
	var view = view.GetInstance("ViewTestKey8", func() interfaces.IView { return &view.View{Key: "ViewTestKey8"} })

	// Create and register the test mediator to be removed.
	var data = Data{}
	view.RegisterMediator(&ViewTestMediator2{Mediator: mediator.Mediator{Name: ViewTestMediator2_NAME, ViewComponent: &data}})

	// test that notifications work
	view.NotifyObservers(observer.NewNotification(VIEWTEST_NOTE1, "", ""))
	if data.lastNotification != VIEWTEST_NOTE1 {
		t.Error("Expecting data.lastNotification == VIEWTEST_NOTE1")
	}

	view.NotifyObservers(observer.NewNotification(VIEWTEST_NOTE2, "", ""))
	if data.lastNotification != VIEWTEST_NOTE2 {
		t.Error("Expecting data.lastNotification == VIEWTEST_NOTE2")
	}

	// Remove the Mediator
	view.RemoveMediator(ViewTestMediator2_NAME)

	// test that retrieving it now returns nil
	if view.RetrieveMediator(ViewTestMediator2_NAME) != nil {
		t.Error("Expecting view.RetrieveMediator(ViewTestMediator2.NAME) == nil")
	}

	// test that notifications no longer work
	// (ViewTestMediator2 is the one that sets lastNotification
	// on this component, and ViewTestMediator)
	data.lastNotification = ""

	view.NotifyObservers(observer.NewNotification(VIEWTEST_NOTE1, "", ""))
	if data.lastNotification != "" {
		t.Error("Expecting data.lastNotification == ''")
	}

	view.NotifyObservers(observer.NewNotification(VIEWTEST_NOTE2, "", ""))
	if data.lastNotification != "" {
		t.Error("Expecting data.lastNotification == ''")
	}
}

/**
  Tests registering one of two registered Mediators and seeing
  that the remaining one still responds.
*/
func TestRemoveOneOfTwoMediatorsAndSubsequentNotify(t *testing.T) {
	// Get the Multiton View instance
	var view = view.GetInstance("ViewTestKey9", func() interfaces.IView { return &view.View{Key: "ViewTestKey9"} })

	// Create and register that responds to notifications 1 and 2
	var data = Data{}
	view.RegisterMediator(&ViewTestMediator2{Mediator: mediator.Mediator{Name: ViewTestMediator2_NAME, ViewComponent: &data}})

	// Create and register that responds to notification 3
	view.RegisterMediator(&ViewTestMediator3{Mediator: mediator.Mediator{Name: ViewTestMediator3_NAME, ViewComponent: &data}})

	// test that all notifications work
	view.NotifyObservers(observer.NewNotification(VIEWTEST_NOTE1, "", ""))
	if data.lastNotification != VIEWTEST_NOTE1 {
		t.Error("Expecting data.lastNotification == VIEWTEST_NOTE1")
	}

	view.NotifyObservers(observer.NewNotification(VIEWTEST_NOTE2, "", ""))
	if data.lastNotification != VIEWTEST_NOTE2 {
		t.Error("Expecting data.lastNotification == VIEWTEST_NOTE2")
	}

	view.NotifyObservers(observer.NewNotification(VIEWTEST_NOTE3, "", ""))
	if data.lastNotification != VIEWTEST_NOTE3 {
		t.Error("Expecting data.lastNotification == VIEWTEST_NOTE3")
	}

	// Remove the Mediator that responds to 1 and 2
	view.RemoveMediator(ViewTestMediator2_NAME)

	// test that retrieving it now returns nil
	if view.RetrieveMediator(ViewTestMediator2_NAME) != nil {
		t.Error("Expecting view.RetrieveMediator(ViewTestMediator2_NAME) == nil")
	}

	// test that notifications no longer work
	// for notifications 1 and 2, but still work for 3
	data.lastNotification = ""

	view.NotifyObservers(observer.NewNotification(VIEWTEST_NOTE1, "", ""))
	if data.lastNotification == VIEWTEST_NOTE1 {
		t.Error("Expecting data.lastNotification != VIEWTEST_NOTE1")
	}

	view.NotifyObservers(observer.NewNotification(VIEWTEST_NOTE2, "", ""))
	if data.lastNotification == VIEWTEST_NOTE2 {
		t.Error("Expecting data.lastNotification != VIEWTEST_NOTE2")
	}

	view.NotifyObservers(observer.NewNotification(VIEWTEST_NOTE3, "", ""))
	if data.lastNotification != VIEWTEST_NOTE3 {
		t.Error("Expecting data.lastNotification == VIEWTEST_NOTE3")
	}
}

/**
  Tests registering the same mediator twice.
  A subsequent notification should only illicit
  one response. Also, since reregistration
  was causing 2 observers to be created, ensure
  that after removal of the mediator there will
  be no further response.
*/
func TestMediatorReregistration(t *testing.T) {
	// Get the Multiton View instance
	var view = view.GetInstance("ViewTestKey10", func() interfaces.IView { return &view.View{Key: "ViewTestKey10"} })

	// Create and register that responds to notification 5
	var data = Data{}
	view.RegisterMediator(&ViewTestMediator5{Mediator: mediator.Mediator{Name: ViewTestMediator5_NAME, ViewComponent: &data}})

	// try to register another instance of that mediator (uses the same NAME constant).
	view.RegisterMediator(&ViewTestMediator5{Mediator: mediator.Mediator{Name: ViewTestMediator5_NAME, ViewComponent: &data}})

	// test that the counter is only incremented once (mediator 5's response)
	view.NotifyObservers(observer.NewNotification(VIEWTEST_NOTE5, "", ""))
	if data.counter != 1 {
		t.Error("Expecting counter == 1")
	}

	// Remove the Mediator
	view.RemoveMediator(ViewTestMediator5_NAME)

	// test that retrieving it now returns nil
	if view.RetrieveMediator(ViewTestMediator5_NAME) != nil {
		t.Error("Expecting view.RetrieveMediator(ViewTestMediator5_NAME) == nil")
	}

	// test that the counter is no longer incremented
	data.counter = 0
	view.NotifyObservers(observer.NewNotification(VIEWTEST_NOTE5, "", ""))
	if data.counter != 0 {
		t.Error("Expecting counter == 0")
	}
}

/**

  Tests the ability for the observer list to
  be modified during the process of notification,
  and all observers be properly notified. This
  happens most often when multiple Mediators
  respond to the same notification by removing
  themselves.
*/
func TestModifyObserverListDuringNotification(t *testing.T) {
	// Get the Multiton View instance
	var view = view.GetInstance("ViewTestKey11", func() interfaces.IView { return &view.View{Key: "ViewTestKey11"} })

	// Create and register several mediator instances that respond to notification 6
	// by removing themselves, which will cause the observer list for that notification
	// to change. versions prior to MultiCore Version 2.0.5 will see every other mediator
	// fails to be notified.
	var data = Data{}
	view.RegisterMediator(&ViewTestMediator6{Mediator: mediator.Mediator{Name: ViewTestMediator6_NAME + "/1", ViewComponent: &data}})
	view.RegisterMediator(&ViewTestMediator6{Mediator: mediator.Mediator{Name: ViewTestMediator6_NAME + "/2", ViewComponent: &data}})
	view.RegisterMediator(&ViewTestMediator6{Mediator: mediator.Mediator{Name: ViewTestMediator6_NAME + "/3", ViewComponent: &data}})
	view.RegisterMediator(&ViewTestMediator6{Mediator: mediator.Mediator{Name: ViewTestMediator6_NAME + "/4", ViewComponent: &data}})
	view.RegisterMediator(&ViewTestMediator6{Mediator: mediator.Mediator{Name: ViewTestMediator6_NAME + "/5", ViewComponent: &data}})
	view.RegisterMediator(&ViewTestMediator6{Mediator: mediator.Mediator{Name: ViewTestMediator6_NAME + "/6", ViewComponent: &data}})
	view.RegisterMediator(&ViewTestMediator6{Mediator: mediator.Mediator{Name: ViewTestMediator6_NAME + "/7", ViewComponent: &data}})
	view.RegisterMediator(&ViewTestMediator6{Mediator: mediator.Mediator{Name: ViewTestMediator6_NAME + "/8", ViewComponent: &data}})

	// send the notification. each of the above mediators will respond by removing
	// themselves and incrementing the counter by 1. This should leave us with a
	// count of 8, since 8 mediators will respond.
	view.NotifyObservers(observer.NewNotification(VIEWTEST_NOTE6, "", ""))

	// verify the count is correct
	if data.counter != 8 {
		t.Error("Expecting counter == 8")
	}

	// clear the counter
	data.counter = 0
	view.NotifyObservers(observer.NewNotification(VIEWTEST_NOTE6, "", ""))

	// verify the count is 0
	if data.counter != 0 {
		t.Error("Expecting counter == 0")
	}
}
