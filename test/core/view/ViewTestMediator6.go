//
//  ViewTestMediator6.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package view

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/mediator"
)

const ViewTestMediator6_NAME = "ViewTestMediator6" // The Mediator base name

/*
ViewTestMediator6 A Mediator class used by ViewTest.
*/
type ViewTestMediator6 struct {
	mediator.Mediator
}

func (self *ViewTestMediator6) ListNotificationInterests() []string {
	return []string{VIEWTEST_NOTE6}
}

func (self *ViewTestMediator6) HandleNotification(notification interfaces.INotification) {
	//temp implementation until facade is developed
	self.Notifier.Facade.RemoveMediator(self.GetMediatorName())
	//var view2 = view.GetInstance("ViewTestKey11", func() interfaces.IView {
	//	return &view.View{Key: "ViewTestKey11"}
	//})
	//view2.RemoveMediator(mediator.GetMediatorName())
}

func (self *ViewTestMediator6) OnRemove() {
	self.ViewComponent.(*Data).counter++
}
