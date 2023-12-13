//
//  ViewTestNote.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package view

import (
	"github.com/puremvc/puremvc-go-multicore-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-multicore-framework/src/patterns/observer"
)

const ViewTestNote_NAME string = "ViewTestNote"

type ViewTestNote struct {
}

func ViewTestNoteNew(body interface{}) interfaces.INotification {
	return observer.NewNotification(ViewTestNote_NAME, body, "")
}
