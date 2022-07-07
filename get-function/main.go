package main

import (
	"log"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	var input = &dynamodb.GetItemInput{
		TableName: aws.String("ronit-cloud-resume"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String("visitors"),
			},
		},
	}

	queryOutput, err := svc.GetItem(input)

	type Count struct {
		ID 		 string `json:"ID"`
		Visitors string `json:"visitors"`
	}

	count := Count{}

	err = dynamodbattribute.UnmarshalMap(queryOutput.Item, &count)

	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*",
			"Access-Control-Allow-Headers": "*",
		},
		Body:       fmt.Sprintf("{ \"count\": \"%s\" }", count.Visitors),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
