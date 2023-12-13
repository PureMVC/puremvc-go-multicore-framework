//
//  ControllerTestCommand.go
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

/*
ControllerTestCommand A SimpleCommand subclass used by ControllerTest.
*/
type ControllerTestCommand struct {
	command.SimpleCommand
}

/*
Execute Fabricate a result by multiplying the input by 2

- parameter note: the note carrying the ControllerTestVO
*/
func (self *ControllerTestCommand) Execute(notification interfaces.INotification) {
	var vo = notification.Body().(*ControllerTestVO)

	// Fabricate a result
	vo.Result = 2 * vo.Input
}
