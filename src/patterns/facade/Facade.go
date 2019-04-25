//
//  Facade.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package facade

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/core/controller"
	"github.com/puremvc/puremvc-go-multicore-framework/src/core/model"
	"github.com/puremvc/puremvc-go-multicore-framework/src/core/view"
	"github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/observer"
	"sync"
)

/**
A base Multiton `IFacade` implementation.
*/
type Facade struct {
	Key        string                 // The Multiton Key
	controller interfaces.IController // Reference to the Controller
	model      interfaces.IModel      // Reference to the Model
	view       interfaces.IView       // Reference to the View
}

var instanceMap = map[string]interfaces.IFacade{} // The Multiton Facade instanceMap.
var instanceMapMutex = sync.RWMutex{}             // instanceMapMutex for the instance

/**
  Facade Multiton Factory method

  - parameter key: multitonKey
  - parameter facadeFunc: reference that returns `IFacade`
  - returns: the Multiton instance of the `IFacade`
*/
func GetInstance(key string, facadeFunc func() interfaces.IFacade) interfaces.IFacade {
	instanceMapMutex.Lock()
	defer instanceMapMutex.Unlock()

	if instanceMap[key] == nil {
		instanceMap[key] = facadeFunc()
		instanceMap[key].InitializeFacade()
	}
	return instanceMap[key]
}

/**
  Initialize the Multiton `Facade` instance.

  Called automatically by the GetInstance. Override in your
  subclass to do any subclass specific initializations. Be
  sure to call `self.Facade.initializeFacade()`, though.
*/
func (self *Facade) InitializeFacade() {
	self.InitializeModel()
	self.InitializeController()
	self.InitializeView()
}

/**
  Initialize the `Controller`.

  Called by the `initializeFacade` method.
  Override this method in your subclass of `Facade`
  if one or both of the following are true:

  * You wish to initialize a different `IController`.
  * You have `Commands` to register with the `Controller` at startup.`.

  If you don't want to initialize a different `IController`,
  call `self.Facade.initializeController()` at the beginning of your
  method, then register `Command`s.
*/
func (self *Facade) InitializeController() {
	self.controller = controller.GetInstance(self.Key, func() interfaces.IController { return &controller.Controller{Key: self.Key} })
}

/**
  Initialize the `Model`.

  Called by the `initializeFacade` method.
  Override this method in your subclass of `Facade`
  if one or both of the following are true:

  * You wish to initialize a different `IModel`.
  * You have `Proxy`s to register with the Model that do not retrieve a reference to the Facade at construction time.`

  If you don't want to initialize a different `IModel`,
  call `self.Facade.initializeModel()` at the beginning of your
  method, then register `Proxy`s.

  Note: This method is *rarely* overridden; in practice you are more
  likely to use a `Command` to create and register `Proxy`s
  with the `Model`, since `Proxy`s with mutable data will likely
  need to send `INotification`s and thus will likely want to fetch a reference to
  the `Facade` during their construction.
*/
func (self *Facade) InitializeModel() {
	self.model = model.GetInstance(self.Key, func() interfaces.IModel { return &model.Model{Key: self.Key} })
}

/**
  Initialize the `View`.

  Called by the `initializeFacade` method.
  Override this method in your subclass of `Facade`
  if one or both of the following are true:

  * You wish to initialize a different `IView`.
  * You have `Observers` to register with the `View`

  If you don't want to initialize a different `IView`,
  call `self.Facade.initializeView()` at the beginning of your
  method, then register `IMediator` instances.

  Note: This method is *rarely* overridden; in practice you are more
  likely to use a `Command` to create and register `Mediator`s
  with the `View`, since `IMediator` instances will need to send
  `INotification`s and thus will likely want to fetch a reference
  to the `Facade` during their construction.
*/
func (self *Facade) InitializeView() {
	self.view = view.GetInstance(self.Key, func() interfaces.IView { return &view.View{Key: self.Key} })
}

/**
  Register an `ICommand` with the `Controller` by Notification name.

  - parameter notificationName: the name of the `INotification` to associate the `ICommand` with
  - parameter commandFunc: reference that returns `ICommand`
*/
func (self *Facade) RegisterCommand(notificationName string, commandFunc func() interfaces.ICommand) {
	self.controller.RegisterCommand(notificationName, commandFunc)
}

/**
  Remove a previously registered `ICommand` to `INotification` mapping from the Controller.

  - parameter notificationName: the name of the `INotification` to remove the `ICommand` mapping for
*/
func (self *Facade) RemoveCommand(notificationName string) {
	self.controller.RemoveCommand(notificationName)
}

/**
  Check if a Command is registered for a given Notification

  - parameter notificationName:
  - returns: whether a Command is currently registered for the given `notificationName`.
*/
func (self *Facade) HasCommand(notificationName string) bool {
	return self.controller.HasCommand(notificationName)
}

/**
  Register an `IProxy` with the `Model` by name.

  - parameter proxy: the `IProxy` instance to be registered with the `Model`.
*/
func (self *Facade) RegisterProxy(proxy interfaces.IProxy) {
	self.model.RegisterProxy(proxy)
}

/**
  Retrieve an `IProxy` from the `Model` by name.

  - parameter proxyName: the name of the proxy to be retrieved.
  - returns: the `IProxy` instance previously registered with the given `proxyName`.
*/
func (self *Facade) RetrieveProxy(proxyName string) interfaces.IProxy {
	return self.model.RetrieveProxy(proxyName)
}

/**
  Remove an `IProxy` from the `Model` by name.

  - parameter proxyName: the `IProxy` to remove from the `Model`.
  - returns: the `IProxy` that was removed from the `Model`
*/
func (self *Facade) RemoveProxy(proxyName string) interfaces.IProxy {
	return self.model.RemoveProxy(proxyName)
}

/**
  Check if a Proxy is registered

  - parameter proxyName:
  - returns: whether a Proxy is currently registered with the given `proxyName`.
*/
func (self *Facade) HasProxy(proxyName string) bool {
	return self.model.HasProxy(proxyName)
}

/**
  Register a `IMediator` with the `View`.

  - parameter mediator: a reference to the `IMediator`
*/
func (self *Facade) RegisterMediator(mediator interfaces.IMediator) {
	self.view.RegisterMediator(mediator)
}

/**
  Retrieve an `IMediator` from the `View`.

  - parameter mediatorName:
  - returns: the `IMediator` previously registered with the given `mediatorName`.
*/
func (self *Facade) RetrieveMediator(mediatorName string) interfaces.IMediator {
	return self.view.RetrieveMediator(mediatorName)
}

/**
  Remove an `IMediator` from the `View`.

  - parameter mediatorName: name of the `IMediator` to be removed.
  - returns: the `IMediator` that was removed from the `View`
*/
func (self *Facade) RemoveMediator(mediatorName string) interfaces.IMediator {
	return self.view.RemoveMediator(mediatorName)
}

/**
  Check if a Mediator is registered or not

  - parameter mediatorName:
  - returns: whether a Mediator is registered with the given `mediatorName`.
*/
func (self *Facade) HasMediator(mediatorName string) bool {
	return self.view.HasMediator(mediatorName)
}

/**
  Create and send an `INotification`.

  Keeps us from having to construct new notification
  instances in our implementation code.

  - parameter notificationName: the name of the notiification to send
  - parameter body: the body of the notification (optional)
  - parameter _type: the type of the notification (optional)
*/
func (self *Facade) SendNotification(notificationName string, body interface{}, _type string) {
	self.NotifyObservers(observer.NewNotification(notificationName, body, _type))
}

/**
  Notify `Observer`s.

  This method is left public mostly for backward
  compatibility, and to allow you to send custom
  notification classes using the facade.

  Usually you should just call sendNotification
  and pass the parameters, never having to
  construct the notification yourself.

  - parameter notification: the `INotification` to have the `View` notify `Observers` of.
*/
func (self *Facade) NotifyObservers(notification interfaces.INotification) {
	self.view.NotifyObservers(notification)
}

/**
  Set the Multiton key for this facade instance.

  Not called directly, but instead from the
  constructor when getInstance is invoked.
  It is necessary to be public in order to
  implement INotifier.
*/
func (self *Facade) InitializeNotifier(key string) {
	self.Key = key
}

/**
  Check if a Core is registered or not

  - parameter key: the multiton key for the Core in question
  - returns: whether a Core is registered with the given `key`.
*/
func HasCore(key string) bool {
	instanceMapMutex.RLock()
	defer instanceMapMutex.RUnlock()

	return instanceMap[key] != nil
}

/**
  Remove a Core.

  Remove the Model, View, Controller and Facade
  instances for the given key.

  - parameter key: multitonKey of the Core to remove
*/
func RemoveCore(key string) {
	instanceMapMutex.Lock()
	defer instanceMapMutex.Unlock()

	model.RemoveModel(key)
	view.RemoveView(key)
	controller.RemoveController(key)
	delete(instanceMap, key)
}
