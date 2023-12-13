//
//  Proxy.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package proxy

import "github.com/puremvc/puremvc-go-multicore-framework/src/patterns/facade"

const NAME = "Proxy" // default name for the proxy

/*
Proxy A base IProxy implementation.

In PureMVC, Proxy classes are used to manage parts of the
application's data model.

A Proxy might simply manage a reference to a local data object,
in which case interacting with it might involve setting and
getting of its data in synchronous fashion.

Proxy classes are also used to encapsulate the application's
interaction with remote services to save or retrieve data, in which case,
we adopt an asyncronous idiom; setting data (or calling a method) on the
Proxy and listening for a Notification to be sent
when the Proxy has retrieved the data from the service.
*/
type Proxy struct {
	facade.Notifier
	Name string      // the proxy name
	Data interface{} // the data object
}

/*
GetProxyName  Get the proxy name
*/
func (self *Proxy) GetProxyName() string {
	return self.Name
}

/*
SetData  Set the data object
*/
func (self *Proxy) SetData(data interface{}) {
	self.Data = data
}

/*
GetData  Get the data object
*/
func (self *Proxy) GetData() interface{} {
	return self.Data
}

/*
OnRegister  Called by the Model when the Proxy is registered
*/
func (self *Proxy) OnRegister() {

}

/*
OnRemove  Called by the Model when the Proxy is removed
*/
func (self *Proxy) OnRemove() {

}
