//
//  Model.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package model

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"sync"
)

/*
A Multiton IModel implementation.

In PureMVC, the Model class provides
access to model objects (Proxies) by named lookup.

The Model assumes these responsibilities:

* Maintain a cache of IProxy instances.

* Provide methods for registering, retrieving, and removing IProxy instances.

Your application must register IProxy instances
with the Model. Typically, you use an
ICommand to create and register IProxy
instances once the Facade has initialized the Core
actors.
*/
type Model struct {
	Key           string                       // The Multiton Key for this Core
	proxyMap      map[string]interfaces.IProxy // Mapping of proxyNames to IProxy instances
	proxyMapMutex sync.RWMutex                 // Mutex for proxyMap
}

var instanceMap = map[string]interfaces.IModel{} // The Multiton Model instanceMap.
var instanceMapMutex sync.RWMutex                // instanceMapMutex for thread safety

/*
  Model Multiton Factory method.

  - parameter key: multitonKey

  - parameter modelFunc: reference that returns IModel

  - returns: the instance returned by the passed modelFunc
*/
func GetInstance(key string, modelFunc func() interfaces.IModel) interfaces.IModel {
	instanceMapMutex.Lock()
	defer instanceMapMutex.Unlock()

	if instanceMap[key] == nil {
		instanceMap[key] = modelFunc()
		instanceMap[key].InitializeModel()
	}
	return instanceMap[key]
}

/*
  Initialize the Model instance.

  Called automatically by the GetInstance, this
  is your opportunity to initialize the Multiton
  instance in your subclass without overriding the
  constructor.
*/
func (self *Model) InitializeModel() {
	self.proxyMap = map[string]interfaces.IProxy{}
}

/*
  Register an IProxy with the Model.

  - parameter proxy: an IProxy to be held by the Model.
*/
func (self *Model) RegisterProxy(proxy interfaces.IProxy) {
	self.proxyMapMutex.Lock()
	defer self.proxyMapMutex.Unlock()

	proxy.InitializeNotifier(self.Key)
	self.proxyMap[proxy.GetProxyName()] = proxy
	proxy.OnRegister()
}

/*
  Retrieve an IProxy from the Model.

  - parameter proxyName:

  - returns: the IProxy instance previously registered with the given proxyName.
*/
func (self *Model) RetrieveProxy(proxyName string) interfaces.IProxy {
	self.proxyMapMutex.RLock();
	defer self.proxyMapMutex.RUnlock()

	return self.proxyMap[proxyName]
}

/*
  Remove an IProxy from the Model.

  - parameter proxyName: name of the IProxy instance to be removed.

  - returns: the IProxy that was removed from the Model
*/
func (self *Model) RemoveProxy(proxyName string) interfaces.IProxy {
	self.proxyMapMutex.Lock()
	defer self.proxyMapMutex.Unlock()

	var proxy = self.proxyMap[proxyName]
	if proxy != nil {
		delete(self.proxyMap, proxyName)
		proxy.OnRemove()
	}
	return proxy
}

/*
  Check if a Proxy is registered

  - parameter proxyName:

  - returns: whether a Proxy is currently registered with the given proxyName.
*/
func (self *Model) HasProxy(proxyName string) bool {
	self.proxyMapMutex.RLock()
	defer self.proxyMapMutex.RUnlock()

	return self.proxyMap[proxyName] != nil
}

/*
  Remove an IModel instance

  - parameter multitonKey: of IModel instance to remove
*/
func RemoveModel(key string) {
	instanceMapMutex.Lock()
	defer instanceMapMutex.Unlock()

	delete(instanceMap, key)
}
