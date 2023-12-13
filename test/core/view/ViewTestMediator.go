//
//  ViewTestMediator.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package view

import "github.com/puremvc/puremvc-go-multicore-framework/src/patterns/mediator"

const ViewTestMediator_NAME = "ViewTestMediator"

/*
ViewTestMediator A Mediator class used by ViewTest.
*/
type ViewTestMediator struct {
	mediator.Mediator
}

func (self *ViewTestMediator) ListNotificationInterests() []string {
	// be sure that the mediator has some Observers created
	// in order to test removeMediator
	return []string{"ABC", "DEF", "GHI"}
}
