//
//  IMediator.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package interfaces

/**
The interface definition for a PureMVC Mediator.

In PureMVC, `IMediator` implementors assume these responsibilities:

* Implement a common method which returns a list of all `INotification`s the `IMediator` has interest in.
* Implement a notification callback method.
* Implement methods that are called when the IMediator is registered or removed from the View.

Additionally, `IMediator`s typically:

* Act as an intermediary between one or more view components such as text boxes or list controls, maintaining references and coordinating their behavior.
* In Flash-based apps, this is often the place where event listeners are added to view components, and their handlers implemented.
* Respond to and generate `INotifications`, interacting with of the rest of the PureMVC app.

When an `IMediator` is registered with the `IView`,
the `IView` will call the `IMediator`'s
`listNotificationInterests` method. The `IMediator` will
return an `Array` of `INotification` names which
it wishes to be notified about.

The `IView` will then create an `Observer` object
encapsulating that `IMediator`'s (`handleNotification`) method
and register it as an Observer for each `INotification` name returned by
`listNotificationInterests`.
*/
type IMediator interface {
	INotifier

	// Get the `IMediator` instance name
	GetMediatorName() string

	// Get the `IMediator`'s view component.
	GetViewComponent() interface{}

	// Set the `IMediator`'s view component.
	SetViewComponent(viewComponent interface{})

	/**
	  List `INotification` interests.

	  - returns: an `Array` of the `INotification` names this `IMediator` has an interest in.
	*/
	ListNotificationInterests() []string

	/**
	  Handle an `INotification`.

	  - parameter notification: the `INotification` to be handled
	*/
	HandleNotification(notification INotification)

	/**
	  Called by the View when the Mediator is registered
	*/
	OnRegister()

	/**
	  Called by the View when the Mediator is removed
	*/
	OnRemove()
}
