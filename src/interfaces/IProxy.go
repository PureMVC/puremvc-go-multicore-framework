//
//  IProxy.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package interfaces

/*
IProxy The interface definition for a PureMVC Proxy.

In PureMVC, IProxy implementors assume these responsibilities:

* Implement a common method which returns the name of the Proxy.

* Provide methods for setting and getting the data object.

Additionally, IProxys typically:

* Maintain references to one or more pieces of model data.

* Provide methods for manipulating that data.

* Generate INotifications when their model data changes.

* Expose their name as a public static const called NAME, if they are not instantiated multiple times.

* Encapsulate interaction with local or remote services used to fetch and persist model data.
*/
type IProxy interface {
	INotifier

	/*
	  Get the Proxy name
	*/
	GetProxyName() string

	/*
	  Set the data object
	*/
	SetData(data interface{})

	/*
		Get the data object
	*/
	GetData() interface{}

	/*
	  Called by the Model when the Proxy is registered
	*/
	OnRegister()

	/*
	  Called by the Model when the Proxy is removed
	*/
	OnRemove()
}
