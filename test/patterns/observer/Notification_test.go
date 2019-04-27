//
//  Notification_test.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package observer

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/observer"
	"testing"
)

/*
Test the PureMVC Notification class.
*/

/*
  Tests setting and getting the name using Notification class accessor methods.
*/
func TestNameAccessors(t *testing.T) {
	var note = observer.NewNotification("TestNote", nil, "")
	if note.Name() != "TestNote" {
		t.Error("Expecting note.name == 'TestNote")
	}
}

/*
  Tests setting and getting the body using Notification class accessor methods.
*/
func TestBodyAccessors(t *testing.T) {
	var note = observer.NewNotification("TestNote", nil, "")

	note.SetBody(5)

	if note.Body() != 5 {
		t.Error("Expecting note.body == 5")
	}
}

func TestConstructor(t *testing.T) {
	// Create a new Notification using the Constructor to set the note name and body
	var note = observer.NewNotification("TestNote", 5, "TestNoteType")

	// test assertions
	if note.Name() != "TestNote" {
		t.Error("Expecting note.name == 'TestNote")
	}
	if note.Body() != 5 {
		t.Error("Expecting note.body == 5")
	}
	if note.Type() != "TestNoteType" {
		t.Error("Expecting note.type_ == 'TestNoteType'")
	}
}

/*
  Tests the toString method of the notification
*/
func TestToString(t *testing.T) {
	var note = observer.NewNotification("TestNote", "TestBody", "TestType")
	var ts string = note.String()

	if ts != "Notification name: TestNote\nBody: TestBody\nType: TestType" {
		t.Errorf("Expecting note.String() == %s", ts)
	}
}
