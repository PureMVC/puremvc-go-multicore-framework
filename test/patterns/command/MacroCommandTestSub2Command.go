//
//  MacroCommandTestSub2Command.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package command

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/command"
)

type MacroCommandTestSub2Command struct {
	command.MacroCommand
}

/*
Execute Fabricate a result by multiplying the input by itself

- parameter event: the IEvent carrying the MacroCommandTestVO
*/
func (self *MacroCommandTestSub2Command) Execute(notification interfaces.INotification) {
	var vo = notification.Body().(*MacroCommandTestVO)

	// Fabricate a result
	vo.Result2 = vo.Input * vo.Input
}
