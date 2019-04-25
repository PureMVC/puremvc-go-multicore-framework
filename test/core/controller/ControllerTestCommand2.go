//
//  ControllerTestCommand2.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package controller

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/command"
)

/**
A SimpleCommand subclass used by ControllerTest.
*/
type ControllerTestCommand2 struct {
	command.SimpleCommand
}

/**
  Fabricate a result by multiplying the input by 2 and adding to the existing result

  This tests accumulation effect that would show if the command were executed more than once.

  - parameter note: the note carrying the ControllerTestVO
*/
func (controller *ControllerTestCommand2) Execute(notification interfaces.INotification) {
	var vo = notification.Body().(*ControllerTestVO)

	// Fabricate a result
	vo.Result = vo.Result + (2 * vo.Input)
}
