package telegraph

import "container/list"

//Broadcaster ...
type Broadcaster struct {
	listeners []chan interface{}
}

//Listen ...
func (instance *Broadcaster) Listen() <-chan interface{} {
	var newChannel = make(chan interface{})
	instance.listeners = append(instance.listeners, newChannel)
	return newChannel
}

//Notify ...
func (instance *Broadcaster) Notify(notification interface{}) {
	for _, listener := range instance.listeners {
		listener <- notification
	}
}

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
		default:
		}
	}
}

//Unsubscribe ...
func (instance LinkedBroadcaster) Unsubscribe(subscription *list.Element) {
	instance.listeners.Remove(subscription)
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
