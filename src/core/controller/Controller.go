//
//  Controller.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package controller

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/core/view"
	"github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/observer"
	"sync"
)

/*
Controller A Multiton IController implementation.

In PureMVC, the Controller class follows the
'Command and Controller' strategy, and assumes these
responsibilities:

* Remembering which ICommands are intended to handle which INotifications.

* Registering itself as an IObserver with the View for each INotification that it has an ICommand mapping for.

* Creating a new instance of the proper ICommand to handle a given INotification when notified by the View.

* Calling the ICommand's execute method, passing in the INotification.

Your application must register ICommands with the
Controller.

The simplest way is to subclass Facade,
and use its initializeController method to add your
registrations.
*/
type Controller struct {
	Key             string                                // The Multiton Key for this Core
	commandMap      map[string]func() interfaces.ICommand // Mapping of Notification names to funcs that returns ICommand Class instances
	commandMapMutex sync.RWMutex                          // Mutex for commandMap
	view            interfaces.IView                      // Local reference to View
}

var instanceMap = map[string]interfaces.IController{} // The Multiton Controller instanceMap.
var instanceMapMutex sync.RWMutex                     // instanceMap Mutex

/*
GetInstance Controller Multiton Factory method.

- parameter key: multitonKey

- parameter factory reference that returns IController

- returns: the Multiton instance
*/
func GetInstance(key string, factory func() interfaces.IController) interfaces.IController {
	instanceMapMutex.Lock()
	defer instanceMapMutex.Unlock()

	if instanceMap[key] == nil {
		instanceMap[key] = factory()
		instanceMap[key].InitializeController()
	}
	return instanceMap[key]
}

/*
InitializeController Initialize the Singleton Controller instance.

Called automatically by the GetInstance.

Note that if you are using a subclass of View
in your application, you should also subclass Controller
and override the InitializeController method in the
following way:

	func (self *MyController) InitializeController() {
	  self.commandMap = map[string]func() interfaces.ICommand{}
	  self.view = MyView.GetInstance(self.Key, func() interfaces.IView { return &MyView{Key: self.Key} })
	}
*/
func (self *Controller) InitializeController() {
	self.commandMap = map[string]func() interfaces.ICommand{}
	self.view = view.GetInstance(self.Key, func() interfaces.IView { return &view.View{Key: self.Key} })
}

/*
ExecuteCommand If an ICommand has previously been registered
to handle a the given INotification, then it is executed.

- parameter note: an INotification
*/
func (self *Controller) ExecuteCommand(notification interfaces.INotification) {
	self.commandMapMutex.RLock()
	defer self.commandMapMutex.RUnlock()

	var factory = self.commandMap[notification.Name()]
	if factory == nil {
		return
	}
	commandInstance := factory()
	commandInstance.InitializeNotifier(self.Key)
	commandInstance.Execute(notification)
}

/*
RegisterCommand Register a particular ICommand class as the handler
for a particular INotification.

If an ICommand has already been registered to
handle INotifications with this name, it is no longer
used, the new ICommand is used instead.

The Observer for the new ICommand is only created if this the
first time an ICommand has been regisered for this Notification name.

- parameter notificationName: the name of the INotification

- parameter factory: reference that returns ICommand
*/
func (self *Controller) RegisterCommand(notificationName string, factory func() interfaces.ICommand) {
	self.commandMapMutex.Lock()
	defer self.commandMapMutex.Unlock()

	if self.commandMap[notificationName] == nil {
		self.view.RegisterObserver(notificationName, &observer.Observer{Notify: self.ExecuteCommand, Context: self})
	}
	self.commandMap[notificationName] = factory
}

/*
HasCommand Check if a Command is registered for a given Notification

- parameter notificationName:

- returns: whether a Command is currently registered for the given notificationName.
*/
func (self *Controller) HasCommand(notificationName string) bool {
	self.commandMapMutex.RLock()
	defer self.commandMapMutex.RUnlock()

	return self.commandMap[notificationName] != nil
}

/*
RemoveCommand Remove a previously registered ICommand to INotification mapping.

- parameter notificationName: the name of the INotification to remove the ICommand mapping for
*/
func (self *Controller) RemoveCommand(notificationName string) {
	self.commandMapMutex.Lock()
	defer self.commandMapMutex.Unlock()

	if self.commandMap[notificationName] != nil {
		self.view.RemoveObserver(notificationName, self)
		delete(self.commandMap, notificationName)
	}
}

/*
RemoveController Remove an IController instance

- parameter multitonKey: of IController instance to remove
*/
func RemoveController(key string) {
	instanceMapMutex.Lock()
	defer instanceMapMutex.Unlock()

	delete(instanceMap, key)
}
