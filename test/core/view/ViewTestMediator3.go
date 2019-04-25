//
//  ViewTestMediator3.go
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

const ViewTestMediator3_NAME = "viewTestMediator3"

/**
A Mediator class used by ViewTest.
*/
type ViewTestMediator3 struct {
	mediator.Mediator
}

// be sure that the mediator has some Observers created
// in order to test removeMediator
func (mediator *ViewTestMediator3) ListNotificationInterests() []string {
	return []string{VIEWTEST_NOTE3}
}

func (mediator *ViewTestMediator3) HandleNotification(notification interfaces.INotification) {
	mediator.ViewComponent.(*Data).lastNotification = notification.Name()
}
