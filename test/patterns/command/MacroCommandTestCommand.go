//
//  MacroCommandTestCommand.go
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

/*
MacroCommandTestCommand A MacroCommand subclass used by MacroCommandTest.
*/
type MacroCommandTestCommand struct {
	command.MacroCommand
}

/*
InitializeMacroCommand Initialize the MacroCommandTestCommand by adding
its 2 SubCommands.
*/
func (self *MacroCommandTestCommand) InitializeMacroCommand() {
	self.AddSubCommand(func() interfaces.ICommand { return &MacroCommandTestSub1Command{} })
	self.AddSubCommand(func() interfaces.ICommand { return &MacroCommandTestSub2Command{} })
}

func (self *MacroCommandTestCommand) Execute(notification interfaces.INotification) {
	self.InitializeMacroCommand()           // AddSubCommands
	self.MacroCommand.Execute(notification) // Execute SubCommands
}
