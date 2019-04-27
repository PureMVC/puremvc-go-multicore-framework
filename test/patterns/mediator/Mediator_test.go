//
//  Mediator_test.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package mediator

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/mediator"
	"testing"
)

/*
Test the PureMVC Mediator class.
*/

/*
  Tests getting the name using Mediator class accessor method.
*/
func TestNameAccessor(t *testing.T) {
	var m interfaces.IMediator = &mediator.Mediator{Name: mediator.NAME}
	if m.GetMediatorName() != mediator.NAME {
		t.Error("Expecting m.GetMediatorName() == Mediator.NAME")
	}
}

/*
  Tests getting the name using Mediator class accessor method.
*/
func TestViewAccessor(t *testing.T) {
	var view interface{} = new(interface{})

	var m interfaces.IMediator = &mediator.Mediator{Name: mediator.NAME, ViewComponent: view}

	if m.GetViewComponent() == nil {
		t.Error("Expecting m.GetViewComponent() not nil")
	}
}
