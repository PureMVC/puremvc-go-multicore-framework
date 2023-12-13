//
//  Notifier.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package facade

import "github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"

/*
Notifier A Base INotifier implementation.

MacroCommand, Command, Mediator and Proxy
all have a need to send Notifications.

The INotifier interface provides a common method called
sendNotification that relieves implementation code of
the necessity to actually construct Notifications.

The Notifier class, which all of the above mentioned classes
extend, provides an initialized reference to the Facade
Multiton, which is required for the convenience method
for sending Notifications, but also eases implementation as these
classes have frequent Facade interactions and usually require
access to the facade anyway.

NOTE: In the MultiCore version of the framework, there is one caveat to
notifiers, they cannot send notifications or reach the facade until they
have a valid multitonKey.

The multitonKey is set:

* on a Command when it is executed by the Controller

* on a Mediator is registered with the View

* on a Proxy is registered with the Model.
*/
type Notifier struct {
	Facade interfaces.IFacade
	Key    string // The Multiton Key for this app
}

/*
SendNotification Create and send an INotification.

Keeps us from having to construct new INotification
instances in our implementation code.

- parameter notificationName: the name of the notification to send

- parameter body: the body of the notification (optional)

- parameter type: the _type of the notification
*/
func (self *Notifier) SendNotification(notificationName string, body interface{}, _type string) {
	self.Facade.SendNotification(notificationName, body, _type)
}

/*
InitializeNotifier Initialize this INotifier instance.

This is how a Notifier gets its multitonKey.
Calls to sendNotification or to access the
facade will fail until after this method
has been called.

Mediators, Commands or Proxies may override
this method in order to send notifications
or access the Multiton Facade instance as
soon as possible. They CANNOT access the facade
in their constructors, since this method will not
yet have been called.

- parameter key: the multitonKey for this INotifier to use
*/
func (self *Notifier) InitializeNotifier(key string) {
	self.Key = key
	self.Facade = GetInstance(key, func() interfaces.IFacade { return &Facade{Key: key} })
}
