//
//  SimpleCommand_test.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package command

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/observer"
	"testing"
)

/**
  Test the PureMVC SimpleCommand class.
*/

/**
  Tests the `execute` method of a `SimpleCommand`.

  This test creates a new `Notification`, adding a
  `SimpleCommandTestVO` as the body.
  It then creates a `SimpleCommandTestCommand` and invokes
  its `execute` method, passing in the note.

  Success is determined by evaluating a property on the
  object that was passed on the Notification body, which will
  be modified by the SimpleCommand.
*/
func TestSimpleCommandExecute(t *testing.T) {
	// Create the VO
	vo := SimpleCommandTestVO{Input: 5}

	// Create the Notification (note)
	var note = observer.NewNotification("SimpleCommandTestNote", &vo, "")

	// Create the SimpleCommand
	var command = SimpleCommandTestCommand{}

	// Execute the SimpleCommand
	command.execute(note)

	// test assertions
	if vo.Result != 10 {
		t.Error("Expecting vo.result == 10")
	}
}
