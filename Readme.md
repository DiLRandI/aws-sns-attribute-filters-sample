# AWS-SNS-ATTRIBUTE-FILTERS-SAMPLE

This is a sample project to demonstrate how to use AWS SNS attribute filters to filter messages based on message attributes.

## Prerequisites

- GO 1.21
- NPM

## How to run

- Clone the repository
- Run `npm install` to install the dependencies
- Run `go mod tidy` to install the go dependencies
- Run `make deploy` to deploy the stack [By default stack is deployed to `ap-southeast-1`]

## How to test

- Navigate to AWS Lambda console
- Select the lambda function `publisher`
- Click on `Test` and use `Hello-World` as the test event [Repeat this for several times]
  - The publisher will publish random message but with `myIntField` will have value either `123` or `456`
- Navigate AWS Cloudwatch console and check the log groups for
  - `-sns-subscriber` logs which you can see the messages received by the subscriber are only `456` related configuration can be 
    [identify here](https://github.com/DiLRandI/aws-sns-attribute-filters-sample/blob/5346bc9ad9e7e6c89c4ff3c434516de08e8a0455/serverless.yml#L48-L49)
    
  - `-sqs-subscriber` logs which you can see the messages received by the subscriber are only `123` related configuration can be 
    [identify here](https://github.com/DiLRandI/aws-sns-attribute-filters-sample/blob/5346bc9ad9e7e6c89c4ff3c434516de08e8a0455/serverless.yml#L70-L72)
    
- The above configuration will be visible on the AWS Console SNS Topic under `MyTestTopic` -> `Subscription filter policy` section