//
//  Observer_test.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package observer

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/observer"
	"testing"
)

/**
Tests PureMVC Observer class.

Since the Observer encapsulates the interested object's
callback information, there are no getters, only setters.
It is, in effect write-only memory.

Therefore, the only way to test it is to set the
notification method and context and call the notifyObserver
method.
*/

/**
  Tests observer class when initialized by accessor methods.
*/
func TestObserverAccessor(t *testing.T) {
	// Create observer
	test := &Test{}
	var obs interfaces.IObserver = &observer.Observer{Notify: nil, Context: nil}
	obs.SetNotifyContext(test)
	obs.SetNotifyMethod(test.NotifyMethod)

	// create a test event, setting a payload value and notify
	// the observer with it. since the observer is this class
	// and the notification method is observerTestMethod,
	// successful notification will result in our local
	// observerTestVar being set to the value we pass in
	// on the note body.
	var note = observer.NewNotification("ObserverTestNote", 10, "")
	obs.NotifyObserver(note)

	// test assertions
	if test.Var != 10 {
		t.Error("Expecting test.Var == 10")
	}
}

/**
  Tests observer class when initialized by constructor.
*/
func TestObserverConstructor(t *testing.T) {
	// Create observer passing in notification method and context
	var test = &Test{}
	var obs = &observer.Observer{Notify: test.NotifyMethod, Context: test}

	// create a test note, setting a body value and notify
	// the observer with it. since the observer is this class
	// and the notification method is observerTestMethod,
	// successful notification will result in our local
	// observerTestVar being set to the value we pass in
	// on the note body.
	var note = observer.NewNotification("ObserverTestNote", 5, "")
	obs.NotifyObserver(note)

	// test assertions
	if test.Var != 5 {
		t.Error("Expecting test.Var == 5")
	}
}

/**
  Tests the compareNotifyContext method of the Observer class
*/
func TestCompareNotifyContext(t *testing.T) {
	// Create observer passing in notification method and context
	var test = &Test{}
	var negTestObj = &NegTestObj{}

	var obs = &observer.Observer{Notify: test.NotifyMethod, Context: test}

	// test assertions
	if obs.CompareNotifyContext(negTestObj) != false {
		t.Error("Expecting observer.compareNotifyContext(negTestObj) == false")
	}
	if obs.CompareNotifyContext(test) != true {
		t.Error("Expecting observer.compareNotifyContext(this) == true")
	}
}

type Test struct {
	Var int
}

/**
  A function that is used as the observer notification
  method. It multiplies the input number by the
  observerTestVar value
*/
func (o *Test) NotifyMethod(note interfaces.INotification) {
	o.Var = note.Body().(int)
}

type NegTestObj struct{}
