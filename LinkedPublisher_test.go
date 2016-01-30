package telegraph

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Something struct {
	Score int
	Info  string
}

var _ = Describe("LinkedPublisher", func() {
	It("Accept subscriptions", func() {
		broadcaster := NewLinkedPublisher()
		subscription := broadcaster.Subscribe()
		go func() {
			broadcaster.Publish(struct{}{})
		}()
		<-subscription.Channel
	})

	It("Remove subscriptions", func() {
		broadcaster := NewLinkedPublisher()
		subscriptionOne := broadcaster.Subscribe()
		subscriptionTwo := broadcaster.Subscribe()
		subscriptionOne.RemoveFrom(broadcaster)
		go func() {
			broadcaster.Publish(struct{}{})
		}()
		<-subscriptionTwo.Channel
	})

	It("Non blocking publications", func() {
		broadcaster := NewLinkedPublisher()
		subscriptionOne := broadcaster.Subscribe()
		if subscriptionOne.Channel != nil {
			//Do somthing
		}
		subscriptionTwo := broadcaster.Subscribe()
		go func() {
			broadcaster.Publish(struct{}{})
		}()
		<-subscriptionTwo.Channel
	})

	It("Publish a struct", func() {
		const expectedInfo = "BOOM"
		const expectedScore = 5

		broadcaster := NewLinkedPublisher()
		subscriptionOne := broadcaster.Subscribe()
		go func() {
			broadcaster.Publish(Something{
				Info:  expectedInfo,
				Score: expectedScore,
			})
		}()
		item := <-subscriptionOne.Channel
		something := item.(Something)
		Expect(something.Info).To(Equal(expectedInfo))
		Expect(something.Score).To(Equal(expectedScore))
	})

	It("Closes a range on unsubscribe", func() {
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
	})
})
