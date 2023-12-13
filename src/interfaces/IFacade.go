//
//  IFacade.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package interfaces

/*
IFacade The interface definition for a PureMVC Facade.

The Facade Pattern suggests providing a single
class to act as a central point of communication
for a subsystem.

In PureMVC, the Facade acts as an interface between
the core MVC actors (Model, View, Controller) and
the rest of your application.
*/
type IFacade interface {
	INotifier

	/*
	  Initialize the Multiton Facade instance.

	  Called automatically by the constructor. Override in your
	  subclass to do any subclass specific initializations. Be
	  sure to call super.initializeFacade(), though.
	*/
	InitializeFacade()

	/*
	  Initialize the Controller.
	*/
	InitializeController()

	/*
		Initialize the Model.
	*/
	InitializeModel()

	/*
		Initialize the View.
	*/
	InitializeView()

	/*
	  Register an ICommand with the Controller.

	  - parameter noteName: the name of the INotification to associate the ICommand with.
	  - parameter factory: reference that returns ICommand
	*/
	RegisterCommand(notificationName string, factory func() ICommand)

	/*
	  Remove a previously registered ICommand to INotification mapping from the Controller.

	  - parameter notificationName: the name of the INotification to remove the ICommand mapping for
	*/
	RemoveCommand(notificationName string)

	/*
	  Check if a Command is registered for a given Notification

	  - parameter notificationName:
	  - returns: whether a Command is currently registered for the given notificationName.
	*/
	HasCommand(notificationName string) bool

	/*
	  Register an IProxy with the Model by name.

	  - parameter proxy: the IProxy to be registered with the Model.
	*/
	RegisterProxy(proxy IProxy)

	/*
	  Retrieve a IProxy from the Model by name.

	  - parameter proxyName: the name of the IProxy instance to be retrieved.
	  - returns: the IProxy previously regisetered by proxyName with the Model.
	*/
	RetrieveProxy(proxyName string) IProxy

	/*
	  Remove an IProxy instance from the Model by name.

	  - parameter proxyName: the IProxy to remove from the Model.
	  - returns: the IProxy that was removed from the Model
	*/
	RemoveProxy(proxyName string) IProxy

	/*
	  Check if a Proxy is registered

	  - parameter proxyName:
	  - returns: whether a Proxy is currently registered with the given proxyName.
	*/
	HasProxy(proxyName string) bool

	/*
	  Register an IMediator instance with the View.

	  - parameter mediator: a reference to the IMediator instance
	*/
	RegisterMediator(mediator IMediator)

	/*
	  Retrieve an IMediator instance from the View.

	  - parameter mediatorName: the name of the IMediator instance to retrievve
	  - returns: the IMediator previously registered with the given mediatorName.
	*/
	RetrieveMediator(mediatorName string) IMediator

	/*
	  Remove a IMediator instance from the View.

	  - parameter mediatorName: name of the IMediator instance to be removed.
	  - returns: the IMediator instance previously registered with the given mediatorName.
	*/
	RemoveMediator(mediatorName string) IMediator

	/*
	  Check if a Mediator is registered or not

	  - parameter mediatorName:
	  - returns: whether a Mediator is registered with the given mediatorName.
	*/
	HasMediator(mediatorName string) bool

	/*
		Notify Observers.

		This method is left public mostly for backward
		compatibility, and to allow you to send custom
		notification classes using the facade.

		Usually you should just call sendNotification
		and pass the parameters, never having to
		construct the notification yourself.

		- parameter notification: the INotification to have the View notify Observers of.
	*/
	NotifyObservers(notification INotification)
}
