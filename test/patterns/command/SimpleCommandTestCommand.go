//
//  SimpleCommandTestCommand.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package command

import "github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"

/**
A SimpleCommand subclass used by SimpleCommandTest.
*/
type SimpleCommandTestCommand struct {
}

/**
  Fabricate a result by multiplying the input by 2

  - parameter event: the `INotification` carrying the `SimpleCommandTestVO`
*/
func (command SimpleCommandTestCommand) execute(notification interfaces.INotification) {
	var vo = notification.Body().(*SimpleCommandTestVO)

	//Fabricate a result
	vo.Result = 2 * vo.Input
}
