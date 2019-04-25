//
//  ViewTestMediator2.go
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

const ViewTestMediator2_NAME = "viewTestMediator2"

/**
A Mediator class used by ViewTest.
*/
type ViewTestMediator2 struct {
	mediator.Mediator
}

func (mediator *ViewTestMediator2) ListNotificationInterests() []string {
	// be sure that the mediator has some Observers created
	// in order to test removeMediator
	return []string{VIEWTEST_NOTE1, VIEWTEST_NOTE2}
}

func (mediator *ViewTestMediator2) HandleNotification(notification interfaces.INotification) {
	mediator.ViewComponent.(*Data).lastNotification = notification.Name()
}
