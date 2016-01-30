package telegraph

import (
	"container/list"
	"time"
)

//LinkedPublisher is the struct which provides the publish and subscribe behaviour
type LinkedPublisher struct {
	listeners *list.List
}

//NewLinkedPublisher returns a new instance of a LinkedPublisher
func NewLinkedPublisher() LinkedPublisher {
	return LinkedPublisher{
		listeners: list.New(),
	}
}

//Subscribe creates a new channel and adds it to the publisher.
//It returns a Subscription struct which exposes a Channel member to consume
func (instance LinkedPublisher) Subscribe() Subscription {

	var newChannel = make(chan interface{})
	element := instance.listeners.PushBack(newChannel)
	return Subscription{
		Channel: newChannel,
		element: element,
	}
}

//Publish iterates over the list of channels and sends the notification object
func (instance LinkedPublisher) Publish(notification interface{}) {
	for e := instance.listeners.Front(); e != nil; e = e.Next() {
		channel := e.Value.(chan interface{})
		select {
		case channel <- notification:
		case <-time.After(1 * time.Nanosecond):
		}
	}
}

//Unsubscribe removes the channel from the list and also closes it
//The channel can not longer be used.
func (instance LinkedPublisher) Unsubscribe(subscription *list.Element) {
	channel := instance.listeners.Remove(subscription)
	close(channel.(chan interface{}))
}

//Subscription struct is the main type used to subscibe
//The Channel member is used to receive published messages
type Subscription struct {
	Channel <-chan interface{}
	element *list.Element
}

//RemoveFrom invokes the Publisher with its hidden copy of the Element.
//This allows the Publisher to efficiently maintain its list of subscribers
//whilst maintain encapsulation
func (instance Subscription) RemoveFrom(publisher LinkedPublisher) {
	publisher.Unsubscribe(instance.element)
}
