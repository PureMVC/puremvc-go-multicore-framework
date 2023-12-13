//
//  ICommand.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package interfaces

/*
ICommand The interface definition for a PureMVC Command.
*/
type ICommand interface {
	INotifier

	/*
	  Execute the ICommand's logic to handle a given INotification.

	  - parameter note: an INotification to handle.
	*/
	Execute(notification INotification)
}
