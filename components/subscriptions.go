package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions"
	"github.com/aws/jsii-runtime-go"
)

func ConfigSubscriptions(stack *awscdk.Stack) {

	firstTopic.AddSubscription(awssnssubscriptions.NewSqsSubscription(firstQueue, &awssnssubscriptions.SqsSubscriptionProps{
		RawMessageDelivery: jsii.Bool(true),
	}))

	secondTopic.AddSubscription(awssnssubscriptions.NewSqsSubscription(secondQueue, &awssnssubscriptions.SqsSubscriptionProps{
		RawMessageDelivery: jsii.Bool(true),
	}))

}
