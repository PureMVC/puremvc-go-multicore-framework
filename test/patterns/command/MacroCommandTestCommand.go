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
A MacroCommand subclass used by MacroCommandTest.
*/
type MacroCommandTestCommand struct {
	command.MacroCommand
}

/*
  Initialize the MacroCommandTestCommand by adding
  its 2 SubCommands.
*/
func (command *MacroCommandTestCommand) InitializeMacroCommand() {
	command.AddSubCommand(func() interfaces.ICommand { return &MacroCommandTestSub1Command{} })
	command.AddSubCommand(func() interfaces.ICommand { return &MacroCommandTestSub2Command{} })
}

func (command *MacroCommandTestCommand) Execute(notification interfaces.INotification) {
	command.InitializeMacroCommand()           // AddSubCommands
	command.MacroCommand.Execute(notification) // Execute SubCommands
}
