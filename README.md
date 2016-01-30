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
