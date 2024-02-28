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
		fmt.Printf("SNS Message: %s\n", record.SNS.Message)
		fmt.Printf("SNS MessageAttributes: %s\n", record.SNS.MessageAttributes)
		fmt.Printf("SNS MessageAttributes: %v\n", record.SNS.MessageAttributes)
		fmt.Printf("SNS MessageAttributes: %v\n", record.SNS.MessageAttributes["jsonMessage"])
	}

	return nil
}
