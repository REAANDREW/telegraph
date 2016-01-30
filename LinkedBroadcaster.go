package telegraph

import (
	"container/list"
	"time"
)

//LinkedBroadcaster ...
type LinkedBroadcaster struct {
	listeners *list.List
}

//NewLinkedBroadcaster ...
func NewLinkedBroadcaster() LinkedBroadcaster {
	return LinkedBroadcaster{
		listeners: list.New(),
	}
}

//Listen ...
func (instance LinkedBroadcaster) Listen() Subscription {

	var newChannel = make(chan interface{})
	element := instance.listeners.PushBack(newChannel)
	return Subscription{
		Channel: newChannel,
		element: element,
	}
}

//Notify ...
func (instance LinkedBroadcaster) Notify(notification interface{}) {
	for e := instance.listeners.Front(); e != nil; e = e.Next() {
		channel := e.Value.(chan interface{})
		select {
		case channel <- notification:
		case <-time.After(1 * time.Nanosecond):
		}
	}
}

//Unsubscribe ...
func (instance LinkedBroadcaster) Unsubscribe(subscription *list.Element) {
	channel := instance.listeners.Remove(subscription)
	close(channel.(chan interface{}))
}

//Subscription ...
type Subscription struct {
	Channel <-chan interface{}
	element *list.Element
}

//RemoveFrom ...
func (instance Subscription) RemoveFrom(broadcaster LinkedBroadcaster) {
	broadcaster.Unsubscribe(instance.element)
}
