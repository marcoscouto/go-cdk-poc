package components

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/jsii-runtime-go"
)

func ConfigS3(stack *awscdk.Stack) {

	awss3.NewBucket(*stack, jsii.String("bucket-test-1"), &awss3.BucketProps{
		Versioned: jsii.Bool(true),
	})

	awss3.NewBucket(*stack, jsii.String("bucket-test-2"), nil)
}
