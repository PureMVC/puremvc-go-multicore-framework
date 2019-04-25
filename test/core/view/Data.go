//
//  Data.go
//  PureMVC Go Multicore
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package view

const VIEWTEST_NOTE1 = "Notification1"
const VIEWTEST_NOTE2 = "Notification2"
const VIEWTEST_NOTE3 = "Notification3"
const VIEWTEST_NOTE4 = "Notification4"
const VIEWTEST_NOTE5 = "Notification5"
const VIEWTEST_NOTE6 = "Notification6"

type Data struct {
	lastNotification string
	onRegisterCalled bool
	onRemoveCalled   bool
	counter          int
}
