//
//  Mediator.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package mediator

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/facade"
)

const NAME = "Mediator" // default name for the mediator

/**
  A base `IMediator` implementation.
*/
type Mediator struct {
	facade.Notifier
	Name          string      // the mediator name
	ViewComponent interface{} // The view component
}

// Get the name of the `Mediator`.
func (self *Mediator) GetMediatorName() string {
	return self.Name
}

// Get the `IMediator`'s view component.
func (self *Mediator) GetViewComponent() interface{} {
	return self.ViewComponent
}

// Set the `IMediator`'s view component.
func (self *Mediator) SetViewComponent(viewComponent interface{}) {
	self.ViewComponent = viewComponent
}

/**
  List the `INotification` names this
  `Mediator` is interested in being notified of.

  - returns: Array the list of `INotification` names
*/
func (self *Mediator) ListNotificationInterests() []string {
	return []string{}
}

/**
  Handle `INotification`s.

  Typically this will be handled in a switch statement,
  with one 'case' entry per `INotification`
  the `Mediator` is interested in.
*/
func (self *Mediator) HandleNotification(notification interfaces.INotification) {

}

/**
  Called by the View when the Mediator is registered
*/
func (self *Mediator) OnRegister() {

}

/**
  Called by the View when the Mediator is removed
*/
func (self *Mediator) OnRemove() {

}
