//
//  Observer.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package observer

import "github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"

/*
Observer A base IObserver implementation.

An Observer is an object that encapsulates information
about an interested object with a method that should
be called when a particular INotification is broadcast.

In PureMVC, the Observer class assumes these responsibilities:

* Encapsulate the notification (callback) method of the interested object.

* Encapsulate the notification context (this) of the interested object.

* Provide methods for setting the notification method and context.

* Provide a method for notifying the interested object.
*/
type Observer struct {
	Notify  func(notification interfaces.INotification)
	Context interface{}
}

/*
NotifyObserver  Notify the interested object.

- parameter notification: the INotification to pass to the interested object's notification method.
*/
func (self *Observer) NotifyObserver(notification interfaces.INotification) {
	self.Notify(notification)
}

/*
CompareNotifyContext  Compare an object to the notification context.

- parameter object: the object to compare
- returns: boolean indicating if the object and the notification context are the same
*/
func (self *Observer) CompareNotifyContext(object interface{}) bool {
	return object == self.Context
}

/*
SetNotifyMethod  Set the notification method.
*/
func (self *Observer) SetNotifyMethod(notifyMethod func(notification interfaces.INotification)) {
	self.Notify = notifyMethod
}

/*
SetNotifyContext  Set the notification context.
*/
func (self *Observer) SetNotifyContext(notifyContext interface{}) {
	self.Context = notifyContext
}
