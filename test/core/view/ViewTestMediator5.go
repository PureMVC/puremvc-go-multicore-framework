//
//  ViewTestMediator5.go
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

const ViewTestMediator5_NAME = "viewTestMediator5"

/*
A Mediator class used by ViewTest.
*/
type ViewTestMediator5 struct {
	mediator.Mediator
}

func (mediator *ViewTestMediator5) ListNotificationInterests() []string {
	return []string{VIEWTEST_NOTE5}
}

func (mediator *ViewTestMediator5) HandleNotification(notification interfaces.INotification) {
	mediator.ViewComponent.(*Data).counter++
}
