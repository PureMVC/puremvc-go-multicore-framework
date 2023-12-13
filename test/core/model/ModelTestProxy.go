//
//  ModelTestProxy.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package model

import "github.com/puremvc/puremvc-go-multicore-framework/src/patterns/proxy"

const MODEL_TEST_PROXY = "modelTestProxy"
const ON_REGISTER_CALLED = "onRegister Called"
const ON_REMOVE_CALLED = "onRemoveCalled"

type ModelTestProxy struct {
	proxy.Proxy
}

func (self *ModelTestProxy) OnRegister() {
	self.SetData(ON_REGISTER_CALLED)
}

func (self *ModelTestProxy) OnRemove() {
	self.SetData(ON_REMOVE_CALLED)
}
