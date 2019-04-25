//
//  INotifier.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package interfaces

/**
The interface definition for a PureMVC Notifier.

`MacroCommand, Command, Mediator` and `Proxy`
all have a need to send `Notifications`.

The `INotifier` interface provides a common method called
`sendNotification` that relieves implementation code of
the necessity to actually construct `Notifications`.

The `Notifier` class, which all of the above mentioned classes
extend, also provides an initialized reference to the `Facade`
Singleton, which is required for the convienience method
for sending `Notifications`, but also eases implementation as these
classes have frequent `Facade` interactions and usually require
access to the facade anyway.
*/
type INotifier interface {
	/**
	  Send a `INotification`.

	  Convenience method to prevent having to construct new
	  notification instances in our implementation code.

	  - parameter notificationName: the name of the notification to send
	  - parameter body: the body of the notification (optional)
	  - parameter type: the type of the notification (optional)
	*/
	SendNotification(notificationName string, body interface{}, _type string)

	/**
	  Initialize this INotifier instance.

	  This is how a Notifier get to Calls to
	  sendNotification or to access the
	  facade will fail until after this method
	  has been called.

	*/
	InitializeNotifier(key string)
}
