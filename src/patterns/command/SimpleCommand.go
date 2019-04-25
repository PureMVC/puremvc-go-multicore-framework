//
//  SimpleCommand.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package command

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/facade"
)

/**
A base `ICommand` implementation.

Your subclass should override the `execute`
method where your business logic will handle the `INotification`.
*/
type SimpleCommand struct {
	facade.Notifier
}

/**
  Fulfill the use-case initiated by the given `INotification`.

  In the Command Pattern, an application use-case typically
  begins with some user action, which results in an `INotification` being broadcast, which
  is handled by business logic in the `execute` method of an
  `ICommand`.

  - parameter notification: the `INotification` to handle.
*/
func (self *SimpleCommand) Execute(notification interfaces.INotification) {

}
