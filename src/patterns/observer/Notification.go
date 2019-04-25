//
//  Notification.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package observer

import "github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"

/**
A base `INotification` implementation.

PureMVC does not rely upon underlying event models such
as the one provided with Flash, and ActionScript 3 does
not have an inherent event model.

The Observer Pattern as implemented within PureMVC exists
to support event-driven communication between the
application and the actors of the MVC triad.

Notifications are not meant to be a replacement for Events
in Flex/Flash/Apollo. Generally, `IMediator` implementors
place event listeners on their view components, which they
then handle in the usual way. This may lead to the broadcast of `Notification`s to
trigger `ICommand`s or to communicate with other `IMediators`. `IProxy` and `ICommand`
instances communicate with each other and `IMediator`s
by broadcasting `INotification`s.

A key difference between Flash `Event`s and PureMVC
`Notification`s is that `Event`s follow the
'Chain of Responsibility' pattern, 'bubbling' up the display hierarchy
until some parent component handles the `Event`, while
PureMVC `Notification`s follow a 'Publish/Subscribe'
pattern. PureMVC classes need not be related to each other in a
parent/child relationship in order to communicate with one another
using `Notification`s.
*/
type Notification struct {
	name  string
	body  interface{}
	_type string
}

/**
Constructor.

- parameter name: name of the `Notification` instance. (required)
- parameter body: the `Notification` body. (optional)
- parameter type: the type of the `Notification` (optional)
*/
func NewNotification(name string, body interface{}, _type string) interfaces.INotification {
	return &Notification{name: name, body: body, _type: _type}
}

/**
Get the name of notification instance
*/
func (self *Notification) Name() string {
	return self.name
}

/**
  Get the body of notification instance
*/
func (self *Notification) Body() interface{} {
	return self.body
}

/**
  Set the body of notification instance
*/
func (self *Notification) SetBody(body interface{}) {
	self.body = body
}

/**
  Get the type of notification instance
*/
func (self *Notification) Type() string {
	return self._type
}

/**
  Set the type of notification instance
*/
func (self *Notification) SetType(t string) {
	self._type = t
}

/**
  Get the string representation of the `Notification` instance.

  - returns: the string representation of the `Notification` instance.
*/
func (self *Notification) String() string {
	msg := "Notification name: " + self.name
	msg += "\nBody: "
	if body, ok := self.body.(string); ok {
		msg += body
	} else {
		msg += "nil"
	}
	msg += "\nType: " + self._type

	return msg
}
