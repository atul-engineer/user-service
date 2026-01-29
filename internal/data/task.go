package data

import (
	"fmt"
	"time"
)

type UserCreateEvent struct {
	Id int
	Name     string
	Email    string
	IsActive bool
}

type UserTask struct {
	Event chan UserCreateEvent
	usrSvc *UserService
}

func NewUserTask() *UserTask {
	return &UserTask{
		Event: make(chan UserCreateEvent, 10),
	}
}

func (t *UserTask) WriteEvent(uv UserCreateEvent) {
	t.Event <- uv
}

func (t *UserTask) ProcessEvent() {
	fmt.Println("Processing user create event")
	for e := range t.Event {
		time.Sleep(4 * time.Second)
		err := t.usrSvc.UpdateUserStatus(e.IsActive, e.Id)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("user status updated ID: ", e.Id)
	}
}