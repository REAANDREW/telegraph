[![Build Status](https://travis-ci.org/REAANDREW/telegraph.svg?branch=master)](https://travis-ci.org/REAANDREW/telegraph) [![Coverage Status](https://coveralls.io/repos/github/REAANDREW/telegraph/badge.svg?branch=master)](https://coveralls.io/github/REAANDREW/telegraph?branch=master)

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [telegraph](#telegraph)
  - [Usage](#usage)
      - [type LinkedPublisher](#type-linkedpublisher)
      - [func  NewLinkedPublisher](#func--newlinkedpublisher)
      - [func (LinkedPublisher) Publish](#func-linkedpublisher-publish)
      - [func (LinkedPublisher) Subscribe](#func-linkedpublisher-subscribe)
      - [func (LinkedPublisher) Unsubscribe](#func-linkedpublisher-unsubscribe)
      - [type Subscription](#type-subscription)
      - [func (Subscription) RemoveFrom](#func-subscription-removefrom)
  - [Examples](#examples)
    - [Subscribing](#subscribing)
    - [Unsubscribing](#unsubscribing)
    - [Non blocking publications](#non-blocking-publications)
    - [Ending a range after Unsubscribing](#ending-a-range-after-unsubscribing)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# telegraph
--
    import "github.com/reaandrew/telegraph"


## Usage

#### type LinkedPublisher

```go
type LinkedPublisher struct {
}
```

LinkedPublisher is the struct which provides the publish and subscribe behaviour

#### func  NewLinkedPublisher

```go
func NewLinkedPublisher() LinkedPublisher
```
NewLinkedPublisher returns a new instance of a LinkedPublisher

#### func (LinkedPublisher) Publish

```go
func (instance LinkedPublisher) Publish(notification interface{})
```
Publish iterates over the list of channels and sends the notification object

#### func (LinkedPublisher) Subscribe

```go
func (instance LinkedPublisher) Subscribe() Subscription
```
Subscribe creates a new channel and adds it to the publisher. It returns a
Subscription struct which exposes a Channel member to consume

#### func (LinkedPublisher) Unsubscribe

```go
func (instance LinkedPublisher) Unsubscribe(subscription *list.Element)
```
Unsubscribe removes the channel from the list and also closes it The channel can
not longer be used.

#### type Subscription

```go
type Subscription struct {
	Channel <-chan interface{}
}
```

Subscription struct is the main type used to subscibe The Channel member is used
to receive published messages

#### func (Subscription) RemoveFrom

```go
func (instance Subscription) RemoveFrom(publisher LinkedPublisher)
```
RemoveFrom invokes the Publisher with its hidden copy of the Element. This
allows the Publisher to efficiently maintain its list of subscribers whilst
maintain encapsulation

## Examples

### Subscribing

```go
broadcaster := NewLinkedPublisher()
subscription := broadcaster.Subscribe()
go func() {
  broadcaster.Publish(1)
}()
Expect(<-subscription.Channel).To(Equal(1))
```

### Unsubscribing

```go
broadcaster := NewLinkedPublisher()
subscriptionOne := broadcaster.Subscribe()
subscriptionTwo := broadcaster.Subscribe()
subscriptionOne.RemoveFrom(broadcaster)
go func() {
  broadcaster.Publish(2)
}()
Expect(<-subscriptionTwo.Channel).To(Equal(2))
```

### Non blocking publications

```go
broadcaster := NewLinkedPublisher()
subscriptionOne := broadcaster.Subscribe()
if subscriptionOne.Channel != nil {
  //Do somthing
}
subscriptionTwo := broadcaster.Subscribe()
go func() {
  broadcaster.Publish(3)
}()
Expect(<-subscriptionTwo.Channel).To(Equal(3))
```

### Ending a range after Unsubscribing

```go
broadcaster := NewLinkedPublisher()
subscription := broadcaster.Subscribe()
go func() {
  for i := 0; i < 10; i++ {
    broadcaster.Publish(i + 1)
  }
}()
for i := range subscription.Channel {
  if i.(int) == 10 {
    subscription.RemoveFrom(broadcaster)
  }
}
fmt.Println("You will see this message!")
```
