package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleSNSMessage)
}

func HandleSNSMessage(ctx context.Context, snsEvent events.SNSEvent) error {
	fmt.Println("===============================================================================")
	fmt.Printf("%+v\n", snsEvent)
	fmt.Println("===============================================================================")

	for _, record := range snsEvent.Records {
		fmt.Printf("SNS Record: %s\n", record.SNS)
		fmt.Printf("SNS MessageID: %s\n", record.SNS.MessageID)
		fmt.Printf("SNS Signature: %s\n", record.SNS.Signature)
		fmt.Printf("SNS Type: %s\n", record.SNS.Type)
		fmt.Printf("SNS Timestamp: %s\n", record.SNS.Timestamp)
		fmt.Printf("SNS TopicArn: %s\n", record.SNS.TopicArn)
		fmt.Printf("SNS Subject: %s\n", record.SNS.Subject)
		fmt.Printf("SNS Message: %s\n", record.SNS.Message)

		fmt.Println("...........................................................................")

		for k, v := range record.SNS.MessageAttributes {
			fmt.Printf("SNS MessageAttributes[%s]: %v\n", k, v)
		}

		fmt.Println("...........................................................................")
	}

	return nil
}
