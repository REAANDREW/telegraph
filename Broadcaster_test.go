package telegraph

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Broadcaster", func() {
	Describe("LinkedBroadcaster", func() {
		It("Accept subscriptions", func() {
			broadcaster := NewLinkedBroadcaster()
			subscription := broadcaster.Listen()
			go func() {
				broadcaster.Notify(struct{}{})
			}()
			<-subscription.Channel
		})

		It("Remove subscriptions", func() {
			broadcaster := NewLinkedBroadcaster()
			subscriptionOne := broadcaster.Listen()
			subscriptionTwo := broadcaster.Listen()
			subscriptionOne.RemoveFrom(broadcaster)
			go func() {
				broadcaster.Notify(struct{}{})
			}()
			<-subscriptionTwo.Channel
		})

		It("Non blocking publications", func() {
			broadcaster := NewLinkedBroadcaster()
			subscriptionOne := broadcaster.Listen()
			if subscriptionOne.Channel != nil {
				//Do somthing
			}
			subscriptionTwo := broadcaster.Listen()
			go func() {
				broadcaster.Notify(struct{}{})
			}()
			<-subscriptionTwo.Channel
		})
	})
})
