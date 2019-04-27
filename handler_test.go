package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestAbs(t *testing.T) {
	response, err := HandleRequest(events.APIGatewayProxyRequest{Body: ""})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(response)
}
